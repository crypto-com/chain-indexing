package handlers

import (
	"errors"
	"math/big"
	"time"

	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbparambase/types"

	"github.com/crypto-com/chain-indexing/appinterface/cosmosapp"
	param_view "github.com/crypto-com/chain-indexing/appinterface/projection/rdbparambase/view"
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
	proposal_view "github.com/crypto-com/chain-indexing/projection/proposal/view"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	"github.com/valyala/fasthttp"
)

type Proposals struct {
	logger applogger.Logger

	cosmosClient       cosmosapp.Client
	proposalsView      *proposal_view.Proposals
	votesView          *proposal_view.Votes
	depositorsView     *proposal_view.Depositors
	proposalParamsView *param_view.Params

	totalBonded              coin.Coin
	totalBondedLastUpdatedAt time.Time
}

func NewProposals(logger applogger.Logger, rdbHandle *rdb.Handle, cosmosClient cosmosapp.Client) *Proposals {
	return &Proposals{
		logger,

		cosmosClient,
		proposal_view.NewProposals(rdbHandle),
		proposal_view.NewVotes(rdbHandle),
		proposal_view.NewDepositors(rdbHandle),
		param_view.NewParams(rdbHandle, proposal_view.PARAMS_TABLE_NAME),

		coin.Coin{},
		time.Unix(int64(0), int64(0)),
	}
}

func (handler *Proposals) FindById(ctx *fasthttp.RequestCtx) {
	idParam, _ := ctx.UserValue("id").(string)
	proposal, err := handler.proposalsView.FindById(idParam)
	if err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			httpapi.NotFound(ctx)
			return
		}
		handler.logger.Errorf("error finding proposal by id: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	tally, queryTallyErr := handler.cosmosClient.ProposalTally(idParam)
	if queryTallyErr != nil {
		if !errors.Is(queryTallyErr, cosmosapp.ErrProposalNotFound) {
			handler.logger.Errorf("error retrieving proposal tally: %v", queryTallyErr)
			httpapi.InternalServerError(ctx)
			return
		}
		tally = cosmosapp.Tally{
			Yes:        "0",
			Abstain:    "0",
			No:         "0",
			NoWithVeto: "0",
		}
	}

	if handler.totalBondedLastUpdatedAt.Add(1 * time.Hour).Before(time.Now()) {
		var queryTotalBondedErr error
		handler.totalBonded, queryTotalBondedErr = handler.cosmosClient.TotalBondedBalance()
		if queryTotalBondedErr != nil {
			handler.logger.Errorf("error retrieving total bonded balance: %v", queryTallyErr)
			httpapi.InternalServerError(ctx)
			return
		}

		handler.totalBondedLastUpdatedAt = time.Now()
	}
	quorumStr, queryQuorumErr := handler.proposalParamsView.FindBy(types.ParamAccessor{
		Module: "gov",
		Key:    "quorum",
	})
	if queryQuorumErr != nil {
		handler.logger.Errorf("error retrieving gov quorum param: %v", queryQuorumErr)
		httpapi.InternalServerError(ctx)
		return
	}

	quorum, parseQuorumOk := new(big.Float).SetString(quorumStr)
	if !parseQuorumOk {
		handler.logger.Errorf("error parsing gov quorum param to big.Float: %v", parseQuorumOk)
		httpapi.InternalServerError(ctx)
		return
	}
	totalBonded, parseTotalBondedOk := new(big.Float).SetString(handler.totalBonded.Amount.String())
	if !parseTotalBondedOk {
		handler.logger.Errorf("error parsing total bonded balance to big.Float: %v", parseTotalBondedOk)
		httpapi.InternalServerError(ctx)
		return
	}
	requiredVotingPower := new(big.Float).Mul(totalBonded, quorum)

	totalVotedPower := new(big.Int).SetInt64(int64(0))
	for _, votedPowerStr := range []string{
		tally.Yes,
		tally.No,
		tally.NoWithVeto,
		tally.Abstain,
	} {
		votedPower, parseVotedPowerOk := new(big.Int).SetString(votedPowerStr, 10)
		if !parseVotedPowerOk {
			handler.logger.Error("error parsing voted power")
			httpapi.InternalServerError(ctx)
			return
		}

		totalVotedPower = new(big.Int).Add(totalVotedPower, votedPower)
	}

	proposalDetails := ProposalDetails{
		proposal,

		requiredVotingPower.Text('f', 0),
		totalVotedPower.Text(10),
		ProposalVotedPowerResult{
			Yes:        tally.Yes,
			Abstain:    tally.Abstain,
			No:         tally.No,
			NoWithVeto: tally.NoWithVeto,
		},
	}

	httpapi.Success(ctx, proposalDetails)
}

func (handler *Proposals) List(ctx *fasthttp.RequestCtx) {
	var err error

	pagination, err := httpapi.ParsePagination(ctx)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	idOrder := view.ORDER_ASC
	queryArgs := ctx.QueryArgs()
	if queryArgs.Has("order") {
		if string(queryArgs.Peek("order")) == "id.desc" {
			idOrder = view.ORDER_DESC
		}
	}

	filter := proposal_view.ProposalListFilter{
		MaybeStatus: nil,
	}
	if queryArgs.Has("filter.status") {
		status := string(queryArgs.Peek("filter.status"))
		filter.MaybeStatus = &status
	}

	proposals, paginationResult, err := handler.proposalsView.List(filter, proposal_view.ProposalListOrder{
		Id: idOrder,
	}, pagination)
	if err != nil {
		handler.logger.Errorf("error listing proposals: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, proposals, paginationResult)
}

func (handler *Proposals) ListVotesById(ctx *fasthttp.RequestCtx) {
	idParam, _ := ctx.UserValue("id").(string)

	pagination, paginationError := httpapi.ParsePagination(ctx)
	if paginationError != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	voteAtOrder := view.ORDER_ASC
	queryArgs := ctx.QueryArgs()
	if queryArgs.Has("order") {
		if string(queryArgs.Peek("order")) == "voteAt.desc" {
			voteAtOrder = view.ORDER_DESC
		}
	}

	votes, paginationResult, err := handler.votesView.ListByProposalId(idParam, proposal_view.VoteListOrder{
		VoteAtBlockHeight: voteAtOrder,
	}, pagination)
	if err != nil {
		handler.logger.Errorf("error listing proposal votes: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, votes, paginationResult)
}

func (handler *Proposals) ListDepositorsById(ctx *fasthttp.RequestCtx) {
	idParam, _ := ctx.UserValue("id").(string)

	pagination, paginationError := httpapi.ParsePagination(ctx)
	if paginationError != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	depositAtOrder := view.ORDER_ASC
	queryArgs := ctx.QueryArgs()
	if queryArgs.Has("order") {
		if string(queryArgs.Peek("order")) == "depositAt.desc" {
			depositAtOrder = view.ORDER_DESC
		}
	}

	depositors, paginationResult, err := handler.depositorsView.ListByProposalId(idParam, proposal_view.DepositorListOrder{
		DepositAtBlockHeight: depositAtOrder,
	}, pagination)
	if err != nil {
		handler.logger.Errorf("error listing proposal votes: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, depositors, paginationResult)
}

type ProposalDetails struct {
	*proposal_view.ProposalWithMonikerRow

	RequiredVotingPower string                   `json:"requiredVotingPower"`
	TotalVotedPower     string                   `json:"totalVotedPower"`
	VotedPowerResult    ProposalVotedPowerResult `json:"votedPowerResult"`
}

type ProposalVotedPowerResult struct {
	Yes        string `json:"yes"`
	Abstain    string `json:"abstain"`
	No         string `json:"no"`
	NoWithVeto string `json:"noWithVeto"`
}
