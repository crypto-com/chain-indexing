package handlers

import (
	"errors"

	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
	nft_view "github.com/crypto-com/chain-indexing/projection/nft/view"
	"github.com/valyala/fasthttp"
)

type NFTs struct {
	logger applogger.Logger

	denomsView         *nft_view.Denoms
	tokensView         *nft_view.Tokens
	tokenTransfersView *nft_view.TokenTransfers
}

func NewNFTs(
	logger applogger.Logger, rdbHandle *rdb.Handle,
) *NFTs {
	return &NFTs{
		logger,

		nft_view.NewDenoms(rdbHandle),
		nft_view.NewTokens(rdbHandle),
		nft_view.NewTokenTransfers(rdbHandle),
	}
}

func (handler *NFTs) ListDenoms(ctx *fasthttp.RequestCtx) {
	pagination, paginationErr := httpapi.ParsePagination(ctx)
	if paginationErr != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	createdAtOrder := view.ORDER_ASC
	queryArgs := ctx.QueryArgs()
	if queryArgs.Has("order") {
		if string(queryArgs.Peek("order")) == "createdAt.desc" {
			createdAtOrder = view.ORDER_DESC
		}
	}

	filter := nft_view.DenomListFilter{
		MaybeCreator: nil,
	}
	if queryArgs.Has("filter.creator") {
		creator := string(queryArgs.Peek("filter.creator"))
		filter.MaybeCreator = &creator
	}

	denoms, paginationResult, err := handler.denomsView.List(filter, nft_view.DenomListOrder{
		CreatedAt: createdAtOrder,
	}, pagination)
	if err != nil {
		handler.logger.Errorf("error listing NFT denoms: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, denoms, paginationResult)
}

func (handler *NFTs) FindDenomById(ctx *fasthttp.RequestCtx) {
	idParam, _ := ctx.UserValue("denomId").(string)
	denom, err := handler.denomsView.FindById(idParam)
	if err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			httpapi.NotFound(ctx)
			return
		}
		handler.logger.Errorf("error finding denom by id: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.Success(ctx, denom)
}

func (handler *NFTs) ListTokensByDenomId(ctx *fasthttp.RequestCtx) {
	pagination, paginationErr := httpapi.ParsePagination(ctx)
	if paginationErr != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	mintedAtOrder := view.ORDER_ASC
	queryArgs := ctx.QueryArgs()
	if queryArgs.Has("order") {
		if string(queryArgs.Peek("order")) == "mintedAt.desc" {
			mintedAtOrder = view.ORDER_DESC
		}
	}

	filter := nft_view.TokenListFilter{
		MaybeDenomId: nil,
		MaybeDrop:    nil,
		MaybeMinter:  nil,
		MaybeOwner:   nil,
	}
	if queryArgs.Has("filter.denomId") {
		denomId := string(queryArgs.Peek("filter.denomId"))
		filter.MaybeDenomId = &denomId
	}
	if queryArgs.Has("filter.drop") {
		drop := string(queryArgs.Peek("filter.drop"))
		filter.MaybeDrop = &drop
	}
	if queryArgs.Has("filter.minter") {
		minter := string(queryArgs.Peek("filter.minter"))
		filter.MaybeMinter = &minter
	}
	if queryArgs.Has("filter.owner") {
		owner := string(queryArgs.Peek("filter.owner"))
		filter.MaybeOwner = &owner
	}

	denoms, paginationResult, err := handler.tokensView.List(filter, nft_view.TokenListOrder{
		MintedAt: mintedAtOrder,
	}, pagination)
	if err != nil {
		handler.logger.Errorf("error listing NFT tokens: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, denoms, paginationResult)
}

func (handler *NFTs) FindTokenById(ctx *fasthttp.RequestCtx) {
	denomIdParam, _ := ctx.UserValue("denomId").(string)
	tokenIdParam, _ := ctx.UserValue("tokenId").(string)
	token, err := handler.tokensView.FindById(denomIdParam, tokenIdParam)
	if err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			httpapi.NotFound(ctx)
			return
		}
		handler.logger.Errorf("error finding token by id: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.Success(ctx, token)
}

func (handler *NFTs) ListDrops(ctx *fasthttp.RequestCtx) {
	pagination, paginationErr := httpapi.ParsePagination(ctx)
	if paginationErr != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	drops, paginationResult, err := handler.tokensView.ListDrops(pagination)
	if err != nil {
		handler.logger.Errorf("error listing NFT drops: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, drops, paginationResult)
}

func (handler *NFTs) ListTokensByDrop(ctx *fasthttp.RequestCtx) {
	dropParam, _ := ctx.UserValue("drop").(string)

	pagination, paginationErr := httpapi.ParsePagination(ctx)
	if paginationErr != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	mintedAtOrder := view.ORDER_ASC
	queryArgs := ctx.QueryArgs()
	if queryArgs.Has("order") {
		if string(queryArgs.Peek("order")) == "mintedAt.desc" {
			mintedAtOrder = view.ORDER_DESC
		}
	}

	filter := nft_view.TokenListFilter{
		MaybeDenomId: nil,
		MaybeDrop:    &dropParam,
		MaybeMinter:  nil,
		MaybeOwner:   nil,
	}

	denoms, paginationResult, err := handler.tokensView.List(filter, nft_view.TokenListOrder{
		MintedAt: mintedAtOrder,
	}, pagination)
	if err != nil {
		handler.logger.Errorf("error listing NFT tokens by drop: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, denoms, paginationResult)
}

func (handler *NFTs) ListTokensByAccount(ctx *fasthttp.RequestCtx) {
	accountParam, _ := ctx.UserValue("account").(string)

	pagination, paginationErr := httpapi.ParsePagination(ctx)
	if paginationErr != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	mintedAtOrder := view.ORDER_ASC
	queryArgs := ctx.QueryArgs()
	if queryArgs.Has("order") {
		if string(queryArgs.Peek("order")) == "mintedAt.desc" {
			mintedAtOrder = view.ORDER_DESC
		}
	}

	filter := nft_view.TokenListFilter{
		MaybeDenomId: nil,
		MaybeDrop:    nil,
		MaybeMinter:  nil,
		MaybeOwner:   &accountParam,
	}

	denoms, paginationResult, err := handler.tokensView.List(filter, nft_view.TokenListOrder{
		MintedAt: mintedAtOrder,
	}, pagination)
	if err != nil {
		handler.logger.Errorf("error listing NFT tokens by account: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, denoms, paginationResult)
}

func (handler *NFTs) ListTransfersByToken(ctx *fasthttp.RequestCtx) {
	denomIdParam, _ := ctx.UserValue("denomId").(string)
	tokenIdParam, _ := ctx.UserValue("tokenId").(string)

	pagination, paginationErr := httpapi.ParsePagination(ctx)
	if paginationErr != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	idOrder := view.ORDER_ASC
	queryArgs := ctx.QueryArgs()
	if queryArgs.Has("order") {
		if string(queryArgs.Peek("order")) == "transferredAt.desc" {
			idOrder = view.ORDER_DESC
		}
	}

	filter := nft_view.TokenTransferListFilter{
		MaybeDenomId:     &denomIdParam,
		MaybeTokenId:     &tokenIdParam,
		MaybeDrop:        nil,
		MaybeBlockHeight: nil,
		MaybeSender:      nil,
		MaybeRecipient:   nil,
		MaybeAccount:     nil,
	}

	transfers, paginationResult, err := handler.tokenTransfersView.List(filter, nft_view.TokenTransferListOrder{
		Id: idOrder,
	}, pagination)
	if err != nil {
		handler.logger.Errorf("error listing NFT token transfers by token: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, transfers, paginationResult)
}

func (handler *NFTs) ListTransfersByAccount(ctx *fasthttp.RequestCtx) {
	accountParam, _ := ctx.UserValue("account").(string)

	pagination, paginationErr := httpapi.ParsePagination(ctx)
	if paginationErr != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	idOrder := view.ORDER_ASC
	queryArgs := ctx.QueryArgs()
	if queryArgs.Has("order") {
		if string(queryArgs.Peek("order")) == "transferredAt.desc" {
			idOrder = view.ORDER_DESC
		}
	}

	filter := nft_view.TokenTransferListFilter{
		MaybeDenomId:     nil,
		MaybeTokenId:     nil,
		MaybeDrop:        nil,
		MaybeBlockHeight: nil,
		MaybeSender:      nil,
		MaybeRecipient:   nil,
		MaybeAccount:     &accountParam,
	}

	transfers, paginationResult, err := handler.tokenTransfersView.List(filter, nft_view.TokenTransferListOrder{
		Id: idOrder,
	}, pagination)
	if err != nil {
		handler.logger.Errorf("error listing NFT token transfers by account: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, transfers, paginationResult)
}

func (handler *NFTs) ListTransfers(ctx *fasthttp.RequestCtx) {
	pagination, paginationErr := httpapi.ParsePagination(ctx)
	if paginationErr != nil {
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

	filter := nft_view.TokenTransferListFilter{
		MaybeDenomId:     nil,
		MaybeTokenId:     nil,
		MaybeDrop:        nil,
		MaybeBlockHeight: nil,
		MaybeSender:      nil,
		MaybeRecipient:   nil,
		MaybeAccount:     nil,
	}
	if queryArgs.Has("filter.denomId") {
		denomId := string(queryArgs.Peek("filter.denomId"))
		filter.MaybeDenomId = &denomId
	}
	if queryArgs.Has("filter.tokenId") {
		tokenId := string(queryArgs.Peek("filter.tokenId"))
		filter.MaybeTokenId = &tokenId
	}
	if queryArgs.Has("filter.drop") {
		drop := string(queryArgs.Peek("filter.drop"))
		filter.MaybeDrop = &drop
	}
	if queryArgs.Has("filter.blockHeight") {
		blockHeight := string(queryArgs.Peek("filter.blockHeight"))
		filter.MaybeBlockHeight = &blockHeight
	}
	if queryArgs.Has("filter.sender") {
		sender := string(queryArgs.Peek("filter.sender"))
		filter.MaybeSender = &sender
	}
	if queryArgs.Has("filter.recipient") {
		recipient := string(queryArgs.Peek("filter.recipient"))
		filter.MaybeRecipient = &recipient
	}
	if queryArgs.Has("filter.account") {
		account := string(queryArgs.Peek("filter.account"))
		filter.MaybeAccount = &account
	}

	transfers, paginationResult, err := handler.tokenTransfersView.List(filter, nft_view.TokenTransferListOrder{
		Id: idOrder,
	}, pagination)
	if err != nil {
		handler.logger.Errorf("error listing NFT token transfers by account: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, transfers, paginationResult)
}
