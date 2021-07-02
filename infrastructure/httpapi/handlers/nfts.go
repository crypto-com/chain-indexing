package handlers

import (
	"errors"
	"strings"

	"github.com/crypto-com/chain-indexing/internal/primptr"

	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
	nft_view "github.com/crypto-com/chain-indexing/projection/nft/view"
	"github.com/valyala/fasthttp"
)

type NFTs struct {
	logger applogger.Logger

	denomsView   *nft_view.Denoms
	tokensView   *nft_view.Tokens
	messagesView *nft_view.Messages
}

func NewNFTs(
	logger applogger.Logger, rdbHandle *rdb.Handle,
) *NFTs {
	return &NFTs{
		logger,

		nft_view.NewDenoms(rdbHandle),
		nft_view.NewTokens(rdbHandle),
		nft_view.NewMessages(rdbHandle),
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

func (handler *NFTs) FindDenomByName(ctx *fasthttp.RequestCtx) {
	nameParam, _ := ctx.UserValue("denomName").(string)
	denom, err := handler.denomsView.FindByName(nameParam)
	if err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			httpapi.NotFound(ctx)
			return
		}
		handler.logger.Errorf("error finding denom by name: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.Success(ctx, denom)
}

func (handler *NFTs) ListTokens(ctx *fasthttp.RequestCtx) {
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

func (handler *NFTs) ListTokensByDenomId(ctx *fasthttp.RequestCtx) {
	denomId, _ := ctx.UserValue("denomId").(string)

	pagination, paginationErr := httpapi.ParsePagination(ctx)
	if paginationErr != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	mintedAtOrder := view.ORDER_ASC
	var lastEditedAtOrder view.ORDER
	queryArgs := ctx.QueryArgs()
	if queryArgs.Has("order") {
		orderArg := string(queryArgs.Peek("order"))
		if orderArg == "mintedAt.desc" {
			mintedAtOrder = view.ORDER_DESC
		} else if orderArg == "lastEditedAt" {
			lastEditedAtOrder = view.ORDER_ASC
		} else if orderArg == "lastEditedAt.desc" {
			lastEditedAtOrder = view.ORDER_DESC
		}
	}

	filter := nft_view.TokenListFilter{
		MaybeDenomId: &denomId,
		MaybeDrop:    nil,
		MaybeMinter:  nil,
		MaybeOwner:   nil,
	}
	order := nft_view.TokenListOrder{
		MintedAt:     mintedAtOrder,
		LastEditedAt: lastEditedAtOrder,
	}

	denoms, paginationResult, err := handler.tokensView.List(filter, order, pagination)
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
	var lastEditedAtOrder view.ORDER
	queryArgs := ctx.QueryArgs()
	if queryArgs.Has("order") {
		orderArg := string(queryArgs.Peek("order"))
		if orderArg == "mintedAt.desc" {
			mintedAtOrder = view.ORDER_DESC
		} else if orderArg == "lastEditedAt" {
			lastEditedAtOrder = view.ORDER_ASC
		} else if orderArg == "lastEditedAt.desc" {
			lastEditedAtOrder = view.ORDER_DESC
		}
	}

	filter := nft_view.TokenListFilter{
		MaybeDenomId: nil,
		MaybeDrop:    &dropParam,
		MaybeMinter:  nil,
		MaybeOwner:   nil,
	}
	order := nft_view.TokenListOrder{
		MintedAt:     mintedAtOrder,
		LastEditedAt: lastEditedAtOrder,
	}

	denoms, paginationResult, err := handler.tokensView.List(filter, order, pagination)
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
	var lastEditedAtOrder view.ORDER
	queryArgs := ctx.QueryArgs()
	if queryArgs.Has("order") {
		orderArg := string(queryArgs.Peek("order"))
		if orderArg == "mintedAt.desc" {
			mintedAtOrder = view.ORDER_DESC
		} else if orderArg == "lastEditedAt" {
			lastEditedAtOrder = view.ORDER_ASC
		} else if orderArg == "lastEditedAt.desc" {
			lastEditedAtOrder = view.ORDER_DESC
		}
	}

	filter := nft_view.TokenListFilter{
		MaybeDenomId: nil,
		MaybeDrop:    nil,
		MaybeMinter:  nil,
		MaybeOwner:   &accountParam,
	}
	order := nft_view.TokenListOrder{
		MintedAt:     mintedAtOrder,
		LastEditedAt: lastEditedAtOrder,
	}

	denoms, paginationResult, err := handler.tokensView.List(filter, order, pagination)
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

	filter := nft_view.MessagesListFilter{
		MaybeDenomId:  &denomIdParam,
		MaybeTokenId:  &tokenIdParam,
		MaybeDrop:     nil,
		MaybeMsgTypes: []string{"MsgTransferNFT"},
	}

	transfers, paginationResult, err := handler.messagesView.List(filter, nft_view.MessagesListOrder{
		Id: idOrder,
	}, pagination)
	if err != nil {
		handler.logger.Errorf("error listing NFT token transfers by token: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, transfers, paginationResult)
}

func (handler *NFTs) ListMessagesByDenom(ctx *fasthttp.RequestCtx) {
	var err error

	pagination, err := httpapi.ParsePagination(ctx)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	denomId := ctx.UserValue("denomId").(string)
	filter := nft_view.MessagesListFilter{
		MaybeDenomId:  &denomId,
		MaybeTokenId:  nil,
		MaybeDrop:     nil,
		MaybeMsgTypes: nil,
	}
	queryArgs := ctx.QueryArgs()
	if queryArgs.Has("filter.msgType") {
		filter.MaybeMsgTypes = strings.Split(string(queryArgs.Peek("filter.msgType")), ",")
	}

	idOrder := view.ORDER_ASC
	if queryArgs.Has("order") {
		if string(queryArgs.Peek("order")) == "height.desc" {
			idOrder = view.ORDER_DESC
		}
	}

	blocks, paginationResult, err := handler.messagesView.List(
		filter, nft_view.MessagesListOrder{Id: idOrder}, pagination,
	)
	if err != nil {
		handler.logger.Errorf("error listing NFT messages: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, blocks, paginationResult)
}

func (handler *NFTs) ListMessagesByToken(ctx *fasthttp.RequestCtx) {
	var err error

	pagination, err := httpapi.ParsePagination(ctx)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	denomId := ctx.UserValue("denomId").(string)
	tokenId := ctx.UserValue("tokenId").(string)
	filter := nft_view.MessagesListFilter{
		MaybeDenomId:  &denomId,
		MaybeTokenId:  &tokenId,
		MaybeDrop:     nil,
		MaybeMsgTypes: nil,
	}
	queryArgs := ctx.QueryArgs()
	if queryArgs.Has("filter.msgType") {
		filter.MaybeMsgTypes = strings.Split(string(queryArgs.Peek("filter.msgType")), ",")
	}

	idOrder := view.ORDER_ASC
	if queryArgs.Has("order") {
		if string(queryArgs.Peek("order")) == "height.desc" {
			idOrder = view.ORDER_DESC
		}
	}

	blocks, paginationResult, err := handler.messagesView.List(
		filter, nft_view.MessagesListOrder{Id: idOrder}, pagination,
	)
	if err != nil {
		handler.logger.Errorf("error listing NFT messages: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, blocks, paginationResult)
}

func (handler *NFTs) ListMessages(ctx *fasthttp.RequestCtx) {
	var err error

	pagination, err := httpapi.ParsePagination(ctx)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	filter := nft_view.MessagesListFilter{
		MaybeDenomId:  nil,
		MaybeTokenId:  nil,
		MaybeDrop:     nil,
		MaybeMsgTypes: nil,
	}
	queryArgs := ctx.QueryArgs()
	if queryArgs.Has("filter.denomId") {
		filter.MaybeDenomId = primptr.String(string(queryArgs.Peek("filter.denomId")))
	}
	if queryArgs.Has("filter.tokenId") {
		filter.MaybeTokenId = primptr.String(string(queryArgs.Peek("filter.tokenId")))
	}
	if queryArgs.Has("filter.drop") {
		filter.MaybeDrop = primptr.String(string(queryArgs.Peek("filter.drop")))
	}
	if queryArgs.Has("filter.msgType") {
		filter.MaybeMsgTypes = strings.Split(string(queryArgs.Peek("filter.msgType")), ",")
	}

	idOrder := view.ORDER_ASC
	if queryArgs.Has("order") {
		if string(queryArgs.Peek("order")) == "height.desc" {
			idOrder = view.ORDER_DESC
		}
	}

	blocks, paginationResult, err := handler.messagesView.List(
		filter, nft_view.MessagesListOrder{Id: idOrder}, pagination,
	)
	if err != nil {
		handler.logger.Errorf("error listing NFT messages: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, blocks, paginationResult)
}
