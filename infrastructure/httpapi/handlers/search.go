package handlers

import (
	"errors"
	"strings"

	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/external/tmcosmosutils"

	account_transaction_view "github.com/crypto-com/chain-indexing/projection/account_transaction/view"

	validator_view "github.com/crypto-com/chain-indexing/projection/validator/view"

	"github.com/valyala/fasthttp"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
	block_view "github.com/crypto-com/chain-indexing/projection/block/view"
	transaction_view "github.com/crypto-com/chain-indexing/projection/transaction/view"
)

type Search struct {
	logger applogger.Logger

	blocksView                   *block_view.Blocks
	transactionsView             transaction_view.BlockTransactions
	validatorsView               validator_view.Validators
	accountTransactionsTotalView *account_transaction_view.AccountTransactionsTotal
}

func NewSearch(logger applogger.Logger, rdbHandle *rdb.Handle) *Search {
	return &Search{
		logger.WithFields(applogger.LogFields{
			"module": "SearchHandler",
		}),

		block_view.NewBlocks(rdbHandle),
		transaction_view.NewTransactionsView(rdbHandle),
		validator_view.NewValidatorsView(rdbHandle),
		account_transaction_view.NewAccountTransactionsTotal(rdbHandle),
	}
}

func (search *Search) Search(ctx *fasthttp.RequestCtx) {
	keyword := string(ctx.QueryArgs().Peek("keyword"))

	var results SearchResults

	blocks, err := search.blocksView.Search(keyword)
	if err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			blocks = []block_view.Block{}
		} else {
			search.logger.Errorf("error searching block: %v", err)
			httpapi.InternalServerError(ctx)
			return
		}
	}

	transactions, err := search.transactionsView.Search(keyword)
	if err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			transactions = []transaction_view.TransactionRow{}
		} else {
			search.logger.Errorf("error searching transaction: %v", err)
			httpapi.InternalServerError(ctx)
			return
		}
	}

	validators, err := search.validatorsView.Search(keyword)
	if err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			validators = []validator_view.ValidatorRow{}
		} else {
			search.logger.Errorf("error searching validator: %v", err)
			httpapi.InternalServerError(ctx)
			return
		}
	}

	results.Accounts = make([]string, 0)
	if tmcosmosutils.IsValidCosmosAddress(keyword) {
		isAccountExist, err := search.accountTransactionsTotalView.Search(keyword)
		if err != nil {
			search.logger.Errorf("error searching account: %v", err)
			httpapi.InternalServerError(ctx)
			return
		}
		if isAccountExist {
			results.Accounts = []string{keyword}
		}
	}

	// If keyword contains a "/", then it is a {account}/{memo} combination
	strs := strings.SplitN(keyword, "/", 2)
	if len(strs) == 2 {
		address := strs[0]

		if tmcosmosutils.IsValidCosmosAddress(address) {
			isAccountExist, err := search.accountTransactionsTotalView.Search(keyword)
			if err != nil {
				search.logger.Errorf("error searching account: %v", err)
				httpapi.InternalServerError(ctx)
				return
			}
			if isAccountExist {
				results.Accounts = []string{keyword}
			}
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
	Accounts     []string                          `json:"accounts"`
}
