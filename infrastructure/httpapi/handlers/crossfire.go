package handlers

import (
	"fmt"
	"strings"

	"github.com/crypto-com/chain-indexing/appinterface/cosmosapp"

	"github.com/valyala/fasthttp"

	crossfire "github.com/crypto-com/chain-indexing/appinterface/projection/crossfire"
	crossfire_constants "github.com/crypto-com/chain-indexing/appinterface/projection/crossfire/constants"
	crossfire_validator_view "github.com/crypto-com/chain-indexing/appinterface/projection/crossfire/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
	"github.com/crypto-com/chain-indexing/internal/utctime"
)

type Crossfire struct {
	logger applogger.Logger

	validatorAddressPrefix string
	consNodeAddressPrefix  string

	cosmosAppClient              cosmosapp.Client
	crossfireValidatorsView      *crossfire_validator_view.CrossfireValidators
	crossfireValidatorsStatsView *crossfire_validator_view.CrossfireValidatorsStats
	httpClient                   *crossfire.HTTPClient
}

func NewCrossfire(
	logger applogger.Logger,
	validatorAddressPrefix string,
	consNodeAddressPrefix string,
	cosmosAppClient cosmosapp.Client,
	rdbHandle *rdb.Handle,
	participantsListURL string,
) *Crossfire {
	return &Crossfire{
		logger.WithFields(applogger.LogFields{
			"module": "CrossfireHTTPHander",
		}),

		validatorAddressPrefix,
		consNodeAddressPrefix,

		cosmosAppClient,
		crossfire_validator_view.NewCrossfireValidators(rdbHandle),
		crossfire_validator_view.NewCrossfireValidatorsStats(rdbHandle),
		crossfire.NewHTTPClient(participantsListURL),
	}
}

// ListAllCrossfireValidators Fetch all validators and combine them
func (handler *Crossfire) ListAllCrossfireValidators(ctx *fasthttp.RequestCtx) {
	var err error

	// Fetch Remote List
	participantList, errFetching := handler.httpClient.Participants()
	if errFetching != nil {
		handler.logger.Errorf("error listing validators: %v", errFetching)
		httpapi.InternalServerError(ctx)
	}

	// Fetch DB Crossfire Validators List
	crossfireValidatorsDB, err := handler.crossfireValidatorsView.List()
	if err != nil {
		handler.logger.Errorf("error listing validators: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	var finalList = []CrossfireValidatorDetails{}

	for _, participant := range *participantList {
		for _, dbValidator := range crossfireValidatorsDB {

			// For each matching participant create a return object
			if dbValidator.OperatorAddress == participant.OperatorAddress && dbValidator.InitialDelegatorAddress == participant.PrimaryAddress {

				validatorStats, err := handler.getValidatorStats(participant.OperatorAddress, dbValidator)
				if err != nil {
					continue
				}
				validatorDetail := CrossfireValidatorDetails{
					operatorAddress:            participant.OperatorAddress,
					initialDelegatorAddress:    dbValidator.InitialDelegatorAddress,
					moniker:                    dbValidator.Moniker,
					taskSetup:                  dbValidator.TaskPhase1NodeSetup,
					taskKeepActive:             dbValidator.TaskPhase2KeepNodeActive,
					taskProposalVote:           dbValidator.TaskPhase2ProposalVote,
					taskNetworkUpgrade:         dbValidator.TaskPhase2NetworkUpgrade,
					rankPhase1n2CommitmentRank: dbValidator.RankTaskPhase1n2CommitmentCount,
					rankPhase3CommitmentRank:   dbValidator.RankTaskPhase3CommitmentCount,
					rankTxSentRank:             dbValidator.RankTaskHighestTxSent,
					stats:                      validatorStats,
				}
				finalList = append(finalList, validatorDetail)
			}
		}
	}

	httpapi.Success(ctx, finalList)
}

func (handler *Crossfire) getValidatorStats(operatorAddress string, dbValidator crossfire_validator_view.CrossfireValidatorRow) (*ValidatorStats, error) {
	// Fetch Validator Stats
	key := fmt.Sprintf("%s%s", "%", operatorAddress)
	validatorStatRows, err := handler.crossfireValidatorsStatsView.FindAllLike(key)
	if err != nil {
		return nil, fmt.Errorf("[error-crossfire] Fetching Validator Stats, %v", err)
	}

	var finalValidatorStats *ValidatorStats
	for _, validatorStatRow := range validatorStatRows {
		*finalValidatorStats.joinedAtBlockHeight = dbValidator.JoinedAtBlockHeight
		*finalValidatorStats.joinedAtTimestamp = dbValidator.JoinedAtBlockTime

		switch strings.Split(validatorStatRow.Key, crossfire_constants.DB_KEY_SEPARATOR)[0] {
		case crossfire_constants.TOTAL_TX_SENT_PREFIX:
			*finalValidatorStats.totalTxSent = validatorStatRow.Value
		case crossfire_constants.PHASE_1_TX_SENT_PREFIX:
			*finalValidatorStats.txSentPhase1 = validatorStatRow.Value
		case crossfire_constants.PHASE_2_TX_SENT_PREFIX:
			*finalValidatorStats.txSentPhase2 = validatorStatRow.Value
		case crossfire_constants.PHASE_3_TX_SENT_PREFIX:
			*finalValidatorStats.txSentPhase3 = validatorStatRow.Value
		case crossfire_constants.PHASE_3_COMMIT_PREFIX:
			*finalValidatorStats.commitCountPhase3 = validatorStatRow.Value
		case crossfire_constants.PHASE_2_COMMIT_PREFIX:
			*finalValidatorStats.commitCountPhase2 = validatorStatRow.Value
		case crossfire_constants.PHASE_1N2_COMMIT_PREFIX:
			*finalValidatorStats.commitCountPhase1n2 = validatorStatRow.Value
		default:
			break
		}
	}
	return finalValidatorStats, nil
}

// ValidatorStats Validator statistics
type ValidatorStats struct {
	totalTxSent         *int64
	txSentPhase1        *int64
	txSentPhase2        *int64
	txSentPhase3        *int64
	commitCountPhase1n2 *int64
	commitCountPhase2   *int64
	commitCountPhase3   *int64
	joinedAtBlockHeight *int64
	joinedAtTimestamp   *utctime.UTCTime
}

// CrossfireValidatorDetails response object
type CrossfireValidatorDetails struct {
	operatorAddress            string
	initialDelegatorAddress    string
	moniker                    string
	taskSetup                  crossfire_constants.TaskStatus
	taskKeepActive             crossfire_constants.TaskStatus
	taskProposalVote           crossfire_constants.TaskStatus
	taskNetworkUpgrade         crossfire_constants.TaskStatus
	rankPhase1n2CommitmentRank int64
	rankPhase3CommitmentRank   int64
	rankTxSentRank             int64
	stats                      *ValidatorStats
}
