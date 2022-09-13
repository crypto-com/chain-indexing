package cosmosapp

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/crypto-com/chain-indexing/usecase/coin"
	"github.com/crypto-com/chain-indexing/usecase/model"
	jsoniter "github.com/json-iterator/go"

	cosmosapp_interface "github.com/crypto-com/chain-indexing/appinterface/cosmosapp"
)

var _ cosmosapp_interface.Client = &HTTPClient{}

const ERR_CODE_ACCOUNT_NOT_FOUND = 2
const ERR_CODE_ACCOUNT_NO_DELEGATION = 5

type HTTPClient struct {
	httpClient *http.Client
	rpcUrl     string

	bondingDenom string
}

// NewHTTPClient returns a new HTTPClient for tendermint request
func NewHTTPClient(rpcUrl string, bondingDenom string) *HTTPClient {
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
	}

	return &HTTPClient{
		httpClient,
		strings.TrimSuffix(rpcUrl, "/"),

		bondingDenom,
	}
}

func NewInsecureHTTPClient(rpcUrl string, bondingDenom string) *HTTPClient {
	// nolint:gosec
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{
		Timeout:   30 * time.Second,
		Transport: transport,
	}

	return &HTTPClient{
		httpClient,
		strings.TrimSuffix(rpcUrl, "/"),

		bondingDenom,
	}
}

func (client *HTTPClient) Account(accountAddress string) (*cosmosapp_interface.Account, error) {
	rawRespBody, err := client.request(
		fmt.Sprintf("%s/%s", client.getUrl("auth", "accounts"), accountAddress), "",
	)
	if err != nil {
		return nil, err
	}
	defer rawRespBody.Close()

	var accountResp AccountResp
	if err := jsoniter.NewDecoder(rawRespBody).Decode(&accountResp); err != nil {
		return nil, err
	}
	rawAccount := accountResp.Account

	var account cosmosapp_interface.Account
	switch rawAccount.Type {
	case cosmosapp_interface.ACCOUNT_BASE:
		account = cosmosapp_interface.Account{
			Type:          rawAccount.Type,
			Address:       *rawAccount.MaybeAddress,
			MaybePubkey:   rawAccount.MaybePubKey,
			AccountNumber: *rawAccount.MaybeAccountNumber,
			Sequence:      *rawAccount.MaybeSequence,

			MaybeModuleAccount:            nil,
			MaybeDelayedVestingAccount:    nil,
			MaybeContinuousVestingAccount: nil,
			MaybePeriodicVestingAccount:   nil,
		}
	case cosmosapp_interface.ACCOUNT_MODULE:
		account = cosmosapp_interface.Account{
			Type:          rawAccount.Type,
			Address:       rawAccount.MaybeBaseAccount.Address,
			MaybePubkey:   nil,
			AccountNumber: rawAccount.MaybeBaseAccount.AccountNumber,
			Sequence:      rawAccount.MaybeBaseAccount.Sequence,
			MaybeModuleAccount: &cosmosapp_interface.ModuleAccount{
				Name:        *rawAccount.MaybeName,
				Permissions: rawAccount.MaybePermissions,
			},

			MaybeDelayedVestingAccount:    nil,
			MaybeContinuousVestingAccount: nil,
			MaybePeriodicVestingAccount:   nil,
		}
	case "/cosmos.vesting.v1beta1.DelayedVestingAccount":
		account = cosmosapp_interface.Account{
			Type:          rawAccount.Type,
			Address:       rawAccount.MaybeBaseVestingAccount.BaseAccount.Address,
			MaybePubkey:   rawAccount.MaybeBaseVestingAccount.BaseAccount.MaybePubKey,
			AccountNumber: rawAccount.MaybeBaseVestingAccount.BaseAccount.AccountNumber,
			Sequence:      rawAccount.MaybeBaseVestingAccount.BaseAccount.Sequence,

			MaybeModuleAccount: nil,
			MaybeDelayedVestingAccount: &cosmosapp_interface.DelayedVestingAccount{
				OriginalVesting:  rawAccount.MaybeBaseVestingAccount.OriginalVesting,
				DelegatedFree:    rawAccount.MaybeBaseVestingAccount.DelegatedFree,
				DelegatedVesting: rawAccount.MaybeBaseVestingAccount.DelegatedVesting,
				EndTime:          rawAccount.MaybeBaseVestingAccount.EndTime,
			},
			MaybeContinuousVestingAccount: nil,
			MaybePeriodicVestingAccount:   nil,
		}

	case "/cosmos.vesting.v1beta1.ContinuousVestingAccount":
		account = cosmosapp_interface.Account{
			Type:          rawAccount.Type,
			Address:       rawAccount.MaybeBaseVestingAccount.BaseAccount.Address,
			MaybePubkey:   rawAccount.MaybeBaseVestingAccount.BaseAccount.MaybePubKey,
			AccountNumber: rawAccount.MaybeBaseVestingAccount.BaseAccount.AccountNumber,
			Sequence:      rawAccount.MaybeBaseVestingAccount.BaseAccount.Sequence,

			MaybeModuleAccount:         nil,
			MaybeDelayedVestingAccount: nil,
			MaybeContinuousVestingAccount: &cosmosapp_interface.ContinuousVestingAccount{
				OriginalVesting:  rawAccount.MaybeBaseVestingAccount.OriginalVesting,
				DelegatedFree:    rawAccount.MaybeBaseVestingAccount.DelegatedFree,
				DelegatedVesting: rawAccount.MaybeBaseVestingAccount.DelegatedVesting,
				StartTime:        *rawAccount.MaybeStartTime,
				EndTime:          rawAccount.MaybeBaseVestingAccount.EndTime,
			},
			MaybePeriodicVestingAccount: nil,
		}
	case cosmosapp_interface.ACCOUNT_VESTING_PERIODIC:
		account = cosmosapp_interface.Account{
			Type:          rawAccount.Type,
			Address:       rawAccount.MaybeBaseVestingAccount.BaseAccount.Address,
			MaybePubkey:   rawAccount.MaybeBaseVestingAccount.BaseAccount.MaybePubKey,
			AccountNumber: rawAccount.MaybeBaseVestingAccount.BaseAccount.AccountNumber,
			Sequence:      rawAccount.MaybeBaseVestingAccount.BaseAccount.Sequence,

			MaybeModuleAccount:            nil,
			MaybeDelayedVestingAccount:    nil,
			MaybeContinuousVestingAccount: nil,
			MaybePeriodicVestingAccount: &cosmosapp_interface.PeriodicVestingAccount{
				OriginalVesting:  rawAccount.MaybeBaseVestingAccount.OriginalVesting,
				DelegatedFree:    rawAccount.MaybeBaseVestingAccount.DelegatedFree,
				DelegatedVesting: rawAccount.MaybeBaseVestingAccount.DelegatedVesting,
				StartTime:        *rawAccount.MaybeStartTime,
				EndTime:          rawAccount.MaybeBaseVestingAccount.EndTime,
				VestingPeriods:   rawAccount.MaybeVestingPeriods,
			},
		}
	default:
		return nil, fmt.Errorf("unrecognized account type: %s", rawAccount.Type)
	}

	return &account, nil
}

func (client *HTTPClient) Balances(accountAddress string) (coin.Coins, error) {
	resp := &BankBalancesResp{
		Pagination: Pagination{
			MaybeNextKey: nil,
			Total:        "",
		},
	}
	balances := coin.NewEmptyCoins()
	for {
		queryUrl := fmt.Sprintf("%s/%s", client.getUrl("bank", "balances"), accountAddress)
		if resp.Pagination.MaybeNextKey != nil {
			queryUrl = fmt.Sprintf(
				"%s?pagination.key=%s",
				queryUrl, url.QueryEscape(*resp.Pagination.MaybeNextKey),
			)
		}

		rawRespBody, err := client.request(queryUrl)
		if err != nil {
			return nil, err
		}
		defer rawRespBody.Close()

		if err := jsoniter.NewDecoder(rawRespBody).Decode(&resp); err != nil {
			return nil, err
		}
		for _, balanceKVPair := range resp.BankBalanceResponses {
			balance, coinErr := coin.NewCoinFromString(balanceKVPair.Denom, balanceKVPair.Amount)
			if coinErr != nil {
				return nil, coinErr
			}
			balances = balances.Add(balance)
		}

		if resp.Pagination.MaybeNextKey == nil {
			break
		}
	}

	return balances, nil
}

func (client *HTTPClient) BondedBalance(accountAddress string) (coin.Coins, error) {
	resp := &DelegationsResp{
		MaybePagination: &Pagination{
			MaybeNextKey: nil,
			Total:        "",
		},
	}
	balance := coin.NewEmptyCoins()
	for {
		queryUrl := fmt.Sprintf("%s/%s", client.getUrl("staking", "delegations"), accountAddress)
		if resp.MaybePagination.MaybeNextKey != nil {
			queryUrl = fmt.Sprintf(
				"%s?pagination.key=%s",
				queryUrl, url.QueryEscape(*resp.MaybePagination.MaybeNextKey),
			)
		}

		rawRespBody, statusCode, err := client.rawRequest(queryUrl)
		if err != nil {
			return nil, err
		}
		defer rawRespBody.Close()

		if decodeErr := jsoniter.NewDecoder(rawRespBody).Decode(&resp); decodeErr != nil {
			return nil, decodeErr
		}
		if resp.MaybeCode != nil {
			if *resp.MaybeCode == ERR_CODE_ACCOUNT_NOT_FOUND {
				return nil, cosmosapp_interface.ErrAccountNotFound
			} else if *resp.MaybeCode == ERR_CODE_ACCOUNT_NO_DELEGATION {
				return nil, cosmosapp_interface.ErrAccountNoDelegation
			}
		}
		if statusCode != 200 {
			return nil, fmt.Errorf("error requesting Cosmos %s endpoint: status code %d", queryUrl, statusCode)
		}
		for _, delegation := range resp.MaybeDelegationResponses {
			delegatedCoin, coinErr := coin.NewCoinFromString(delegation.Balance.Denom, delegation.Balance.Amount)
			if coinErr != nil {
				return nil, fmt.Errorf("error parsing Coin from delegation balance: %v", coinErr)
			}
			balance = balance.Add(delegatedCoin)
		}

		if resp.MaybePagination.MaybeNextKey == nil {
			break
		}
	}

	return balance, nil
}

func (client *HTTPClient) RedelegatingBalance(accountAddress string) (coin.Coins, error) {
	resp := &UnbondingResp{
		Pagination: Pagination{
			MaybeNextKey: nil,
			Total:        "",
		},
	}
	balance := coin.NewEmptyCoins()
	for {
		queryUrl := fmt.Sprintf(
			"%s/%s/redelegations", client.getUrl("staking", "delegators"), accountAddress,
		)
		if resp.Pagination.MaybeNextKey != nil {
			queryUrl = fmt.Sprintf(
				"%s?pagination.key=%s",
				queryUrl, url.QueryEscape(*resp.Pagination.MaybeNextKey),
			)
		}

		rawRespBody, err := client.request(queryUrl)
		if err != nil {
			return nil, err
		}
		defer rawRespBody.Close()

		if err := jsoniter.NewDecoder(rawRespBody).Decode(&resp); err != nil {
			return nil, err
		}
		for _, unbonding := range resp.UnbondingResponses {
			for _, entry := range unbonding.Entries {
				unbondingCoin, coinErr := coin.NewCoinFromString(client.bondingDenom, entry.Balance)
				if coinErr != nil {
					return nil, fmt.Errorf("error parsing Coin from unbonding balance: %v", coinErr)
				}
				balance = balance.Add(unbondingCoin)
			}
		}

		if resp.Pagination.MaybeNextKey == nil {
			break
		}
	}

	return balance, nil
}

func (client *HTTPClient) UnbondingBalance(accountAddress string) (coin.Coins, error) {
	resp := &UnbondingResp{
		Pagination: Pagination{
			MaybeNextKey: nil,
			Total:        "",
		},
	}
	balance := coin.NewEmptyCoins()
	for {
		queryUrl := fmt.Sprintf(
			"%s/%s/unbonding_delegations",
			client.getUrl("staking", "delegators"), accountAddress,
		)
		if resp.Pagination.MaybeNextKey != nil {
			queryUrl = fmt.Sprintf(
				"%s?pagination.key=%s",
				queryUrl, url.QueryEscape(*resp.Pagination.MaybeNextKey),
			)
		}

		rawRespBody, err := client.request(queryUrl)
		if err != nil {
			return nil, err
		}
		defer rawRespBody.Close()

		if err := jsoniter.NewDecoder(rawRespBody).Decode(&resp); err != nil {
			return nil, err
		}
		for _, unbonding := range resp.UnbondingResponses {
			for _, entry := range unbonding.Entries {
				unbondingCoin, coinErr := coin.NewCoinFromString(client.bondingDenom, entry.Balance)
				if coinErr != nil {
					return nil, fmt.Errorf("error parsing Coin from unbonding balance: %v", coinErr)
				}
				balance = balance.Add(unbondingCoin)
			}
		}

		if resp.Pagination.MaybeNextKey == nil {
			break
		}
	}

	return balance, nil
}

func (client *HTTPClient) TotalRewards(accountAddress string) (coin.DecCoins, error) {
	rawRespBody, err := client.request(
		fmt.Sprintf(
			"%s/%s/rewards",
			client.getUrl("distribution", "delegators"),
			accountAddress,
		), "",
	)
	if err != nil {
		return nil, err
	}
	defer rawRespBody.Close()

	var delegatorRewardResp DelegatorRewardResp
	if err := jsoniter.NewDecoder(rawRespBody).Decode(&delegatorRewardResp); err != nil {
		return nil, err
	}

	rewards := coin.NewEmptyDecCoins()
	for _, total := range delegatorRewardResp.Total {
		rewardCoin, coinErr := coin.NewDecCoinFromString(total.Denom, total.Amount)
		if coinErr != nil {
			return nil, fmt.Errorf("error parsing Coin from total reward: %v", coinErr)
		}
		rewards = rewards.Add(rewardCoin)
	}
	return rewards, nil
}

func (client *HTTPClient) Validator(validatorAddress string) (*cosmosapp_interface.Validator, error) {
	rawRespBody, err := client.request(
		fmt.Sprintf(
			"%s/%s",
			client.getUrl("staking", "validators"), validatorAddress,
		), "",
	)
	if err != nil {
		return nil, err
	}
	defer rawRespBody.Close()

	var validatorResp ValidatorResp
	if err := jsoniter.NewDecoder(rawRespBody).Decode(&validatorResp); err != nil {
		return nil, err
	}

	return &validatorResp.Validator, nil
}

func (client *HTTPClient) Commission(validatorAddress string) (coin.DecCoins, error) {
	rawRespBody, err := client.request(
		fmt.Sprintf("%s/%s/commission",
			client.getUrl("distribution", "validators"), validatorAddress,
		), "",
	)
	if err != nil {
		return nil, err
	}
	defer rawRespBody.Close()

	var commissionResp ValidatorCommissionResp
	if err := jsoniter.NewDecoder(rawRespBody).Decode(&commissionResp); err != nil {
		return nil, err
	}

	totalCommission := coin.NewEmptyDecCoins()
	for _, commission := range commissionResp.Commissions.Commission {
		commissionCoin, coinErr := coin.NewDecCoinFromString(commission.Denom, commission.Amount)
		if coinErr != nil {
			return nil, fmt.Errorf("error parsing Coin from commission: %v", coinErr)
		}
		totalCommission = totalCommission.Add(commissionCoin)
	}
	return totalCommission, nil
}

func (client *HTTPClient) Delegation(
	delegator string, validator string,
) (*cosmosapp_interface.DelegationResponse, error) {
	resp := &DelegationsResp{
		MaybePagination: &Pagination{
			MaybeNextKey: nil,
			Total:        "",
		},
	}
	for {
		queryUrl := fmt.Sprintf("%s/%s", client.getUrl("staking", "delegations"), delegator)
		if resp.MaybePagination.MaybeNextKey != nil {
			queryUrl = fmt.Sprintf(
				"%s?pagination.key=%s",
				queryUrl, url.QueryEscape(*resp.MaybePagination.MaybeNextKey),
			)
		}

		rawRespBody, statusCode, err := client.rawRequest(queryUrl)
		if err != nil {
			return nil, err
		}
		defer rawRespBody.Close()

		if decodeErr := jsoniter.NewDecoder(rawRespBody).Decode(&resp); decodeErr != nil {
			return nil, decodeErr
		}
		if resp.MaybeCode != nil && *resp.MaybeCode == ERR_CODE_ACCOUNT_NOT_FOUND {
			return nil, cosmosapp_interface.ErrAccountNotFound
		}
		if statusCode != 200 {
			return nil, fmt.Errorf("error requesting Cosmos %s endpoint: status code %d", queryUrl, statusCode)
		}

		for _, delegation := range resp.MaybeDelegationResponses {
			if delegation.Delegation.DelegatorAddress == delegator &&
				delegation.Delegation.ValidatorAddress == validator {
				return &delegation, nil
			}
		}

		if resp.MaybePagination.MaybeNextKey == nil {
			break
		}
	}

	return nil, nil
}

func (client *HTTPClient) AnnualProvisions() (coin.DecCoin, error) {
	rawRespBody, err := client.request(client.getUrl("mint", "annual_provisions"))
	if err != nil {
		return coin.DecCoin{}, err
	}
	defer rawRespBody.Close()

	var annualProvisionsResp AnnualProvisionsResp
	if err := jsoniter.NewDecoder(rawRespBody).Decode(&annualProvisionsResp); err != nil {
		return coin.DecCoin{}, err
	}

	annualProvisions, coinErr := coin.NewDecCoinFromString(client.bondingDenom, annualProvisionsResp.AnnualProvisions)
	if coinErr != nil {
		return coin.DecCoin{}, fmt.Errorf("error parsing coin from annual provision: %v", annualProvisions)
	}

	return annualProvisions, nil
}

func (client *HTTPClient) TotalBondedBalance() (coin.Coin, error) {
	resp := &ValidatorsResp{
		MaybePagination: &Pagination{
			MaybeNextKey: nil,
			Total:        "",
		},
	}

	totalBondedBalance, newCoinErr := coin.NewCoin(client.bondingDenom, coin.ZeroInt())
	if newCoinErr != nil {
		return coin.Coin{}, fmt.Errorf("error when creating new coin: %v", newCoinErr)
	}
	for {
		queryUrl := client.getUrl("staking", "validators")
		if resp.MaybePagination.MaybeNextKey != nil {
			queryUrl = fmt.Sprintf(
				"%s?pagination.key=%s",
				queryUrl, url.QueryEscape(*resp.MaybePagination.MaybeNextKey),
			)
		}

		rawRespBody, statusCode, err := client.rawRequest(queryUrl)
		if err != nil {
			return coin.Coin{}, err
		}
		defer rawRespBody.Close()

		if decodeErr := jsoniter.NewDecoder(rawRespBody).Decode(&resp); decodeErr != nil {
			return coin.Coin{}, decodeErr
		}
		if resp.MaybeCode != nil && *resp.MaybeCode == ERR_CODE_ACCOUNT_NOT_FOUND {
			return coin.Coin{}, cosmosapp_interface.ErrAccountNotFound
		}
		if statusCode != 200 {
			return coin.Coin{}, fmt.Errorf("error requesting Cosmos %s endpoint: status code %d", queryUrl, statusCode)
		}

		for _, validator := range resp.MaybeValidatorResponse {
			bondedCoin, coinErr := coin.NewCoinFromString(client.bondingDenom, validator.Tokens)
			if coinErr != nil {
				if coinErr != nil {
					return coin.Coin{}, fmt.Errorf("error parsing Coin from validator tokens: %v", coinErr)
				}
			}
			totalBondedBalance = totalBondedBalance.Add(bondedCoin)
		}

		if resp.MaybePagination.MaybeNextKey == nil {
			break
		}
	}

	return totalBondedBalance, nil
}

func (client *HTTPClient) Proposals() ([]cosmosapp_interface.Proposal, error) {
	resp := &ProposalsResp{
		MaybePagination: &Pagination{
			MaybeNextKey: nil,
			Total:        "",
		},
	}

	proposals := make([]cosmosapp_interface.Proposal, 0)
	for {
		queryUrl := client.getUrl("gov", "proposals")
		if resp.MaybePagination.MaybeNextKey != nil {
			queryUrl = fmt.Sprintf(
				"%s?pagination.key=%s",
				queryUrl, url.QueryEscape(*resp.MaybePagination.MaybeNextKey),
			)
		}

		rawRespBody, statusCode, err := client.rawRequest(queryUrl)
		if err != nil {
			return nil, err
		}
		defer rawRespBody.Close()

		if decodeErr := jsoniter.NewDecoder(rawRespBody).Decode(&resp); decodeErr != nil {
			return nil, decodeErr
		}
		if statusCode != 200 {
			return nil, fmt.Errorf("error requesting Cosmos %s endpoint: status code %d", queryUrl, statusCode)
		}

		proposals = append(proposals, resp.MaybeProposalsResponse...)

		if resp.MaybePagination.MaybeNextKey == nil {
			break
		}
	}

	return proposals, nil
}

func (client *HTTPClient) ProposalById(id string) (cosmosapp_interface.Proposal, error) {
	method := fmt.Sprintf(
		"%s/%s",
		client.getUrl("gov", "proposals"), id,
	)
	rawRespBody, statusCode, err := client.rawRequest(
		method, "",
	)
	if err != nil {
		return cosmosapp_interface.Proposal{}, err
	}
	if statusCode == 404 {
		return cosmosapp_interface.Proposal{}, cosmosapp_interface.ErrProposalNotFound
	}
	if statusCode != 200 {
		rawRespBody.Close()
		return cosmosapp_interface.Proposal{}, fmt.Errorf("error requesting Cosmos %s endpoint: %d", method, statusCode)
	}
	defer rawRespBody.Close()

	var proposalResp ProposalResp
	if err := jsoniter.NewDecoder(rawRespBody).Decode(&proposalResp); err != nil {
		return cosmosapp_interface.Proposal{}, err
	}

	return proposalResp.Proposal, nil
}

func (client *HTTPClient) ProposalTally(id string) (cosmosapp_interface.Tally, error) {
	method := fmt.Sprintf(
		"%s/%s/tally",
		client.getUrl("gov", "proposals"), id,
	)
	rawRespBody, statusCode, err := client.rawRequest(
		method, "",
	)
	if err != nil {
		return cosmosapp_interface.Tally{}, err
	}
	if statusCode == 404 {
		return cosmosapp_interface.Tally{}, cosmosapp_interface.ErrProposalNotFound
	}
	if statusCode != 200 {
		rawRespBody.Close()
		return cosmosapp_interface.Tally{}, fmt.Errorf("error requesting Cosmos %s endpoint: %d", method, statusCode)
	}
	defer rawRespBody.Close()

	var tallyResp TallyResp
	if err := jsoniter.NewDecoder(rawRespBody).Decode(&tallyResp); err != nil {
		return cosmosapp_interface.Tally{}, err
	}

	return tallyResp.Tally, nil
}

func (client *HTTPClient) Tx(hash string) (*model.Tx, error) {
	rawRespBody, err := client.request(
		fmt.Sprintf(
			"%s/%s",
			client.getUrl("tx", "txs"),
			hash,
		), "",
	)
	if err != nil {
		return nil, err
	}
	defer rawRespBody.Close()

	tx, err := ParseTxsResp(rawRespBody)
	if err != nil {
		return nil, fmt.Errorf("error parsing Tx(%s): %v", hash, err)
	}

	return tx, nil
}

func ParseTxsResp(rawRespReader io.Reader) (*model.Tx, error) {
	var txsResp TxsResp
	if err := jsoniter.NewDecoder(rawRespReader).Decode(&txsResp); err != nil {
		return nil, err
	}

	height, err := strconv.ParseInt(txsResp.TxResponse.Height, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing txsResp.TxResponse.Height to int64 param: %v", err)
	}

	gasWanted, err := strconv.ParseInt(txsResp.TxResponse.GasWanted, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing txsResp.TxResponse.GasWanted to int64 param: %v", err)
	}

	gasUsed, err := strconv.ParseInt(txsResp.TxResponse.GasUsed, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing txsResp.TxResponse.GasUsed to int64 param: %v", err)
	}

	var tx = &model.Tx{
		Tx: txsResp.Tx,
		TxResponse: model.TxResponse{
			Height:    height,
			TxHash:    txsResp.TxResponse.TxHash,
			Codespace: txsResp.TxResponse.Codespace,
			Code:      txsResp.TxResponse.Code,
			Data:      txsResp.TxResponse.Data,
			RawLog:    txsResp.TxResponse.RawLog,
			Info:      txsResp.TxResponse.Info,
			GasWanted: gasWanted,
			GasUsed:   gasUsed,
			Timestamp: txsResp.TxResponse.Timestamp,
			Events:    txsResp.TxResponse.Events,
		},
	}

	if txsResp.TxResponse.Logs != nil {
		logs := make([]model.TxResponseLog, 0)
		for _, log := range txsResp.TxResponse.Logs {
			parsedLog := model.TxResponseLog{
				MsgIndex: log.MsgIndex,
				Log:      log.Log,
				Events:   log.Events,
			}
			logs = append(logs, parsedLog)
		}
		tx.TxResponse.Logs = logs
	}

	tx.TxResponse.Tx = model.TxResponseTx{
		Type:     txsResp.TxResponse.Tx.Type,
		CosmosTx: txsResp.TxResponse.Tx.CosmosTx,
	}

	return tx, nil
}

func (client *HTTPClient) getUrl(module string, method string) string {
	return fmt.Sprintf("cosmos/%s/v1beta1/%s", module, method)
}

// request construct tendermint getUrl and issues an HTTP request
// returns the success http Body
func (client *HTTPClient) request(method string, queryString ...string) (io.ReadCloser, error) {
	var err error

	queryUrl := client.rpcUrl + "/" + method
	if len(queryString) > 0 {
		queryUrl += "?" + queryString[0]
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, queryUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request with context: %v", err)
	}
	rawResp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error requesting Cosmos %s endpoint: %v", queryUrl, err)
	}

	if rawResp.StatusCode != 200 {
		rawResp.Body.Close()
		return nil, fmt.Errorf("error requesting Cosmos %s endpoint: %s", method, rawResp.Status)
	}

	return rawResp.Body, nil
}

// rawRequest construct tendermint getUrl and issues an HTTP request
// returns the http Body with any status code
func (client *HTTPClient) rawRequest(method string, queryString ...string) (io.ReadCloser, int, error) {
	var err error

	queryUrl := client.rpcUrl + "/" + method
	if len(queryString) > 0 {
		queryUrl += "?" + queryString[0]
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, queryUrl, nil)
	if err != nil {
		return nil, 0, fmt.Errorf("error creating HTTP request with context: %v", err)
	}
	// nolint:bodyclose
	rawResp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("error requesting Cosmos %s endpoint: %v", queryUrl, err)
	}

	return rawResp.Body, rawResp.StatusCode, nil
}

type Pagination struct {
	MaybeNextKey *string `json:"next_key"`
	Total        string  `json:"total"`
}
