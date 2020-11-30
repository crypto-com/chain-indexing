package handlers

import (
	validator_view "github.com/crypto-com/chainindex/appinterface/projection/validator/view"

	"github.com/valyala/fasthttp"

	block_view "github.com/crypto-com/chainindex/appinterface/projection/block/view"
	transaction_view "github.com/crypto-com/chainindex/appinterface/projection/transasaction/view"
	"github.com/crypto-com/chainindex/appinterface/rdb"
	"github.com/crypto-com/chainindex/infrastructure/httpapi"
	applogger "github.com/crypto-com/chainindex/internal/logger"
)

type Search struct {
	logger applogger.Logger

	blocksView       *block_view.Blocks
	transactionsView *transaction_view.BlockTransactions
	validatorsView   *validator_view.Validators
}

func NewSearch(logger applogger.Logger, rdbHandle *rdb.Handle) *Search {
	return &Search{
		logger.WithFields(applogger.LogFields{
			"module": "SearchHandler",
		}),

		block_view.NewBlocks(rdbHandle),
		transaction_view.NewTransactions(rdbHandle),
		validator_view.NewValidators(rdbHandle),
	}
}

func (search *Search) Search(ctx *fasthttp.RequestCtx) {
	keyword := string(ctx.QueryArgs().Peek("keyword"))

	var results SearchResults

	blocks, err := search.blocksView.Search(keyword)
	if err != nil {
		if err == rdb.ErrNoRows {
			blocks = []block_view.Block{}
		} else {
			search.logger.Errorf("error searching block: %v", err)
			httpapi.InternalServerError(ctx)
			return
		}
	}

	transactions, err := search.transactionsView.Search(keyword)
	if err != nil {
		if err == rdb.ErrNoRows {
			transactions = []transaction_view.TransactionRow{}
		} else {
			search.logger.Errorf("error searching transaction: %v", err)
			httpapi.InternalServerError(ctx)
			return
		}
	}

	validators, err := search.validatorsView.Search(keyword)
	if err != nil {
		if err == rdb.ErrNoRows {
			validators = []validator_view.ValidatorRow{}
		} else {
			search.logger.Errorf("error searching validator: %v", err)
			httpapi.InternalServerError(ctx)
			return
		}
	}

	results.Blocks = blocks
	results.Transactions = transactions
	results.Validators = validators

	httpapi.Success(ctx, results)
}

type SearchResults struct {
	Blocks       []block_view.Block                `json:"blocks"`
	Transactions []transaction_view.TransactionRow `json:"transactions"`
	Validators   []validator_view.ValidatorRow     `json:"validators"`
}
