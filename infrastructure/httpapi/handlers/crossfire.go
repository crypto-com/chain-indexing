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

	var dbParticipantsMap = make(map[string]int)
	for idx, dbParticipant := range crossfireValidatorsDB {
		dbParticipantsMap[dbParticipant.OperatorAddress] = idx
	}

	var finalList = make([]CrossfireValidatorDetails, 0)

	for _, participantRemote := range *participantList {

		if idx, ok := dbParticipantsMap[participantRemote.OperatorAddress]; !ok {
			defaultValidator := getDefaultCrossfireValidatorDetailsResponse(participantRemote)
			finalList = append(finalList, defaultValidator)
		} else {
			validatorStats, err := handler.getValidatorStats(participantRemote.OperatorAddress, crossfireValidatorsDB[idx])
			if err != nil {
				continue
			}
			validatorDetail := CrossfireValidatorDetails{
				OperatorAddress:            participantRemote.OperatorAddress,
				InitialDelegatorAddress:    crossfireValidatorsDB[idx].InitialDelegatorAddress,
				Moniker:                    crossfireValidatorsDB[idx].Moniker,
				TaskSetup:                  crossfireValidatorsDB[idx].TaskPhase1NodeSetup,
				TaskKeepActive:             crossfireValidatorsDB[idx].TaskPhase2KeepNodeActive,
				TaskProposalVote:           validatorStats.TaskVote,
				TaskNetworkUpgrade:         crossfireValidatorsDB[idx].TaskPhase2NetworkUpgrade,
				RankPhase1n2CommitmentRank: crossfireValidatorsDB[idx].RankTaskPhase1n2CommitmentCount,
				RankPhase3CommitmentRank:   crossfireValidatorsDB[idx].RankTaskPhase3CommitmentCount,
				RankTxSentRank:             crossfireValidatorsDB[idx].RankTaskHighestTxSent,
				Stats:                      validatorStats,
			}

			finalList = append(finalList, validatorDetail)
		}
	}

	httpapi.Success(ctx, finalList)
}

func (handler *Crossfire) getValidatorStats(operatorAddress string, dbValidator crossfire_validator_view.CrossfireValidatorRow) (ValidatorStats, error) {
	// Fetch Validator Stats
	key := fmt.Sprintf("%s%s", "%", operatorAddress)
	validatorStatRows, err := handler.crossfireValidatorsStatsView.FindAllLike(key)
	if err != nil {
		return ValidatorStats{}, fmt.Errorf("[error-crossfire] Fetching Validator Stats, %v", err)
	}

	var finalValidatorStats ValidatorStats
	for _, validatorStatRow := range validatorStatRows {
		finalValidatorStats.JoinedAtBlockHeight = dbValidator.JoinedAtBlockHeight
		finalValidatorStats.JoinedAtTimestamp = dbValidator.JoinedAtBlockTime
		finalValidatorStats.TaskVote = crossfire_constants.INCOMPLETED

		switch strings.Split(validatorStatRow.Key, crossfire_constants.DB_KEY_SEPARATOR)[0] {
		case crossfire_constants.TOTAL_TX_SENT_PREFIX:
			finalValidatorStats.TotalTxSent = validatorStatRow.Value
		case crossfire_constants.PHASE_1_TX_SENT_PREFIX:
			finalValidatorStats.TxSentPhase1 = validatorStatRow.Value
		case crossfire_constants.PHASE_2_TX_SENT_PREFIX:
			finalValidatorStats.TxSentPhase2 = validatorStatRow.Value
		case crossfire_constants.PHASE_3_TX_SENT_PREFIX:
			finalValidatorStats.TxSentPhase3 = validatorStatRow.Value
		case crossfire_constants.PHASE_3_COMMIT_PREFIX:
			finalValidatorStats.CommitCountPhase3 = validatorStatRow.Value
		case crossfire_constants.PHASE_2_COMMIT_PREFIX:
			finalValidatorStats.CommitCountPhase2 = validatorStatRow.Value
		case crossfire_constants.PHASE_1N2_COMMIT_PREFIX:
			finalValidatorStats.CommitCountPhase1n2 = validatorStatRow.Value
		case crossfire_constants.VOTED_PROPOSAL_ID:
			finalValidatorStats.TaskVote = crossfire_constants.COMPLETED
		default:
			break
		}
	}
	return finalValidatorStats, nil
}

func getDefaultCrossfireValidatorDetailsResponse(remoteParticipant crossfire.Participant) CrossfireValidatorDetails {

	return CrossfireValidatorDetails{
		OperatorAddress:            remoteParticipant.OperatorAddress,
		InitialDelegatorAddress:    remoteParticipant.PrimaryAddress,
		Moniker:                    remoteParticipant.Moniker,
		TaskSetup:                  crossfire_constants.INCOMPLETED,
		TaskKeepActive:             crossfire_constants.INCOMPLETED,
		TaskProposalVote:           crossfire_constants.INCOMPLETED,
		TaskNetworkUpgrade:         crossfire_constants.INCOMPLETED,
		RankPhase1n2CommitmentRank: 0,
		RankPhase3CommitmentRank:   0,
		RankTxSentRank:             0,
		Stats: ValidatorStats{
			TotalTxSent:         0,
			TxSentPhase1:        0,
			TxSentPhase2:        0,
			TxSentPhase3:        0,
			CommitCountPhase1n2: 0,
			CommitCountPhase2:   0,
			CommitCountPhase3:   0,
			JoinedAtBlockHeight: 0,
			JoinedAtTimestamp:   utctime.FromUnixNano(0),
		},
	}
}

// ValidatorStats Validator statistics
type ValidatorStats struct {
	TotalTxSent         int64           `json:"totalTxSent"`
	TxSentPhase1        int64           `json:"txSentPhase1"`
	TxSentPhase2        int64           `json:"txSentPhase2"`
	TxSentPhase3        int64           `json:"txSentPhase3"`
	CommitCountPhase1n2 int64           `json:"commitCountPhase1n2"`
	CommitCountPhase2   int64           `json:"commitCountPhase2"`
	CommitCountPhase3   int64           `json:"commitCountPhase3"`
	JoinedAtBlockHeight int64           `json:"joinedAtBlockHeight"`
	JoinedAtTimestamp   utctime.UTCTime `json:"joinedAtTimestamp"`
	TaskVote            crossfire_constants.TaskStatus
}

// CrossfireValidatorDetails response object
type CrossfireValidatorDetails struct {
	OperatorAddress            string                         `json:"operatorAddress"`
	InitialDelegatorAddress    string                         `json:"initialDelegatorAddress"`
	Moniker                    string                         `json:"moniker"`
	TaskSetup                  crossfire_constants.TaskStatus `json:"taskSetup"`
	TaskKeepActive             crossfire_constants.TaskStatus `json:"taskKeepActive"`
	TaskProposalVote           crossfire_constants.TaskStatus `json:"taskProposalVote"`
	TaskNetworkUpgrade         crossfire_constants.TaskStatus `json:"taskNetworkUpgrade"`
	RankPhase1n2CommitmentRank int64                          `json:"rankPhase1n2CommitmentRank"`
	RankPhase3CommitmentRank   int64                          `json:"rankPhase3CommitmentRank"`
	RankTxSentRank             int64                          `json:"rankTxSentRank"`
	Stats                      ValidatorStats                 `json:"stats"`
}
