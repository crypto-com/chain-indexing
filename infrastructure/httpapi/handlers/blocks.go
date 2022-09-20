package handlers

import (
	"errors"
	"strconv"

	applogger "github.com/crypto-com/chain-indexing/external/logger"
	validator_view "github.com/crypto-com/chain-indexing/projection/validator/view"
	"github.com/valyala/fasthttp"

	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
	block_view "github.com/crypto-com/chain-indexing/projection/block/view"
	blockrawevent_view "github.com/crypto-com/chain-indexing/projection/block_raw_event/view"
	blockevent_view "github.com/crypto-com/chain-indexing/projection/blockevent/view"
	raw_transaction_view "github.com/crypto-com/chain-indexing/projection/raw_transaction/view"
	transaction_view "github.com/crypto-com/chain-indexing/projection/transaction/view"
)

type Blocks struct {
	logger applogger.Logger

	blocksView                    *block_view.Blocks
	transactionsView              transaction_view.BlockTransactions
	blockEventsView               *blockevent_view.BlockEvents
	validatorBlockCommitmentsView *validator_view.ValidatorBlockCommitments
	rawTransactionsView           raw_transaction_view.BlockRawTransactions
	blockRawEventView             *blockrawevent_view.BlockRawEvents
}

func NewBlocks(logger applogger.Logger, rdbHandle *rdb.Handle) *Blocks {
	return &Blocks{
		logger.WithFields(applogger.LogFields{
			"module": "BlocksHandler",
		}),

		block_view.NewBlocks(rdbHandle),
		transaction_view.NewTransactionsView(rdbHandle),
		blockevent_view.NewBlockEvents(rdbHandle),
		validator_view.NewValidatorBlockCommitments(rdbHandle),
		raw_transaction_view.NewRawTransactionsView(rdbHandle),
		blockrawevent_view.NewBlockRawEvents(rdbHandle),
	}
}

func (handler *Blocks) FindBy(ctx *fasthttp.RequestCtx) {
	heightOrHashParam, heightOrHashParamOk := URLValueGuard(ctx, handler.logger, "height-or-hash")
	if !heightOrHashParamOk {
		return
	}

	height, err := strconv.ParseInt(heightOrHashParam, 10, 64)
	var identity block_view.BlockIdentity
	if err == nil {
		identity.MaybeHeight = &height
	} else {
		identity.MaybeHash = &heightOrHashParam
	}
	block, err := handler.blocksView.FindBy(&identity)
	if err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			httpapi.NotFound(ctx)
			return
		}
		handler.logger.Errorf("error finding block by height or hash: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.Success(ctx, block)
}

func (handler *Blocks) List(ctx *fasthttp.RequestCtx) {
	pagination, err := httpapi.ParsePagination(ctx)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	heightOrder := view.ORDER_ASC
	queryArgs := ctx.QueryArgs()
	if queryArgs.Has("order") {
		if string(queryArgs.Peek("order")) == "height.desc" {
			heightOrder = view.ORDER_DESC
		}
	}

	blocks, paginationResult, err := handler.blocksView.List(block_view.BlocksListOrder{
		Height: heightOrder,
	}, pagination)
	if err != nil {
		handler.logger.Errorf("error listing blocks: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, blocks, paginationResult)
}

func (handler *Blocks) ListTransactionsByHeight(ctx *fasthttp.RequestCtx) {
	pagination, err := httpapi.ParsePagination(ctx)
	if err != nil {
		httpapi.BadRequest(ctx, err)
		return
	}

	blockHeightParam, blockHeightParamOk := URLValueGuard(ctx, handler.logger, "height")
	if !blockHeightParamOk {
		return
	}
	blockHeight, err := strconv.ParseInt(blockHeightParam, 10, 64)
	if err != nil {
		httpapi.BadRequest(ctx, errors.New("invalid block height"))
		return
	}

	heightOrder := view.ORDER_ASC
	queryArgs := ctx.QueryArgs()
	if queryArgs.Has("order") {
		if string(queryArgs.Peek("order")) == "height.desc" {
			heightOrder = view.ORDER_DESC
		}
	}

	blocks, paginationResult, err := handler.transactionsView.List(transaction_view.TransactionsListFilter{
		MaybeBlockHeight: &blockHeight,
	}, transaction_view.TransactionsListOrder{
		Height: heightOrder,
	}, pagination)
	if err != nil {
		handler.logger.Errorf("error listing transactions: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, blocks, paginationResult)
}

func (handler *Blocks) ListEventsByHeight(ctx *fasthttp.RequestCtx) {
	pagination, err := httpapi.ParsePagination(ctx)
	if err != nil {
		httpapi.BadRequest(ctx, err)
		return
	}

	heightOrder := view.ORDER_ASC
	queryArgs := ctx.QueryArgs()
	if queryArgs.Has("order") {
		if string(queryArgs.Peek("order")) == "height.desc" {
			heightOrder = view.ORDER_DESC
		}
	}

	blockHeightParam, blockHeightParamOk := URLValueGuard(ctx, handler.logger, "height")
	if !blockHeightParamOk {
		return
	}
	blockHeight, err := strconv.ParseInt(blockHeightParam, 10, 64)
	if err != nil {
		httpapi.BadRequest(ctx, errors.New("invalid block height"))
		return
	}

	blocks, paginationResult, err := handler.blockEventsView.List(blockevent_view.BlockEventsListFilter{
		MaybeBlockHeight: &blockHeight,
	}, blockevent_view.BlockEventsListOrder{
		Height: heightOrder,
	}, pagination)
	if err != nil {
		handler.logger.Errorf("error listing events: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, blocks, paginationResult)
}

func (handler *Blocks) ListCommitmentsByHeight(ctx *fasthttp.RequestCtx) {
	pagination, err := httpapi.ParsePagination(ctx)
	if err != nil {
		httpapi.BadRequest(ctx, err)
		return
	}

	blockHeightParam, blockHeightParamOk := URLValueGuard(ctx, handler.logger, "height")
	if !blockHeightParamOk {
		return
	}
	blockHeight, err := strconv.ParseInt(blockHeightParam, 10, 64)
	if err != nil {
		httpapi.BadRequest(ctx, errors.New("invalid block height"))
		return
	}

	blocks, paginationResult, err := handler.validatorBlockCommitmentsView.List(
		validator_view.ValidatorBlockCommitmentsListFilter{
			MaybeBlockHeight: &blockHeight,
		}, pagination)
	if err != nil {
		handler.logger.Errorf("error listing block commitments: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, blocks, paginationResult)
}

func (handler *Blocks) ListCommitmentsByConsensusNodeAddress(ctx *fasthttp.RequestCtx) {
	pagination, err := httpapi.ParsePagination(ctx)
	if err != nil {
		httpapi.BadRequest(ctx, err)
		return
	}

	addressParam := ctx.UserValue("address")
	if addressParam == nil {
		httpapi.BadRequest(ctx, errors.New("missing consensus node address"))
		return
	}
	address := addressParam.(string)

	blocks, paginationResult, err := handler.validatorBlockCommitmentsView.List(
		validator_view.ValidatorBlockCommitmentsListFilter{
			MaybeConsensusNodeAddress: &address,
		}, pagination)
	if err != nil {
		handler.logger.Errorf("error listing block commitments: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, blocks, paginationResult)
}

func (handler *Blocks) ListRawTransactionsByHeight(ctx *fasthttp.RequestCtx) {
	pagination, err := httpapi.ParsePagination(ctx)
	if err != nil {
		httpapi.BadRequest(ctx, err)
		return
	}

	blockHeightParam, blockHeightParamOk := URLValueGuard(ctx, handler.logger, "height")
	if !blockHeightParamOk {
		return
	}
	blockHeight, err := strconv.ParseInt(blockHeightParam, 10, 64)
	if err != nil {
		httpapi.BadRequest(ctx, errors.New("invalid block height"))
		return
	}

	heightOrder := view.ORDER_ASC
	queryArgs := ctx.QueryArgs()
	if queryArgs.Has("order") {
		if string(queryArgs.Peek("order")) == "height.desc" {
			heightOrder = view.ORDER_DESC
		}
	}

	blocks, paginationResult, err := handler.rawTransactionsView.List(raw_transaction_view.RawTransactionsListFilter{
		MaybeBlockHeight: &blockHeight,
	}, raw_transaction_view.RawTransactionsListOrder{
		Height: heightOrder,
	}, pagination)
	if err != nil {
		handler.logger.Errorf("error listing transactions: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, blocks, paginationResult)
}

func (handler *Blocks) ListBlockRawEventsByHeight(ctx *fasthttp.RequestCtx) {
	pagination, err := httpapi.ParsePagination(ctx)
	if err != nil {
		httpapi.BadRequest(ctx, err)
		return
	}

	blockHeightParam, blockHeightParamOk := URLValueGuard(ctx, handler.logger, "height")
	if !blockHeightParamOk {
		return
	}
	blockHeight, err := strconv.ParseInt(blockHeightParam, 10, 64)
	if err != nil {
		httpapi.BadRequest(ctx, errors.New("invalid block height"))
		return
	}

	heightOrder := view.ORDER_ASC
	queryArgs := ctx.QueryArgs()
	if queryArgs.Has("order") {
		if string(queryArgs.Peek("order")) == "height.desc" {
			heightOrder = view.ORDER_DESC
		}
	}

	blocks, paginationResult, err := handler.blockRawEventView.List(blockrawevent_view.BlockRawEventsListFilter{
		MaybeBlockHeight: &blockHeight,
	}, blockrawevent_view.BlockRawEventsListOrder{
		Height: heightOrder,
	}, pagination)
	if err != nil {
		handler.logger.Errorf("error listing transactions: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, blocks, paginationResult)
}
