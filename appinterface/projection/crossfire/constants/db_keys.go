package constants

import "fmt"

const VOTED_PROPOSAL_ID string = "voted_proposal_id"
const NETWORK_UPGRADE string = "network_upgrade"
const DB_KEY_SEPARATOR string = ":"
const TASK_PHASE_2_PROPOSAL_VOTE_COLUMN_NAME = "task_phase_2_proposal_vote"

// Tx Sent DB Keys
const PHASE_1_TX_SENT_PREFIX = "phase_1_tx_sent"
const PHASE_2_TX_SENT_PREFIX = "phase_2_tx_sent"
const PHASE_3_TX_SENT_PREFIX = "phase_3_tx_sent"
const TOTAL_TX_SENT_PREFIX = "total_tx_sent"

// JACKPOT Tx Sent DB Keys
const JACKPOT_1_TX_SENT_PREFIX string = "jackpot_1_tx_sent"
const JACKPOT_2_TX_SENT_PREFIX string = "jackpot_2_tx_sent"
const JACKPOT_3_TX_SENT_PREFIX string = "jackpot_3_tx_sent"
const JACKPOT_4_TX_SENT_PREFIX string = "jackpot_4_tx_sent"

const TYPE_URL_PUBKEY = "/cosmos.crypto.secp256k1.PubKey"

const PHASE_1_BLOCK_COUNT ChainStatsKey = "phase1_block_count"
const PHASE_2_BLOCK_COUNT ChainStatsKey = "phase2_block_count"
const PHASE_3_BLOCK_COUNT ChainStatsKey = "phase3_block_count"

type ChainStatsKey = string

const PHASE_1N2_COMMIT_PREFIX CommitPhaseKey = "phase_1n2_commit_count"
const PHASE_2_COMMIT_PREFIX CommitPhaseKey = "phase_2_commit_count"
const PHASE_3_COMMIT_PREFIX CommitPhaseKey = "phase_3_commit_count"

type CommitPhaseKey = string

func ValidatorCommitmentKey(operatorAddress string, phasePrefix CommitPhaseKey) ChainStatsKey {
	return fmt.Sprintf("%s%s%s", phasePrefix, DB_KEY_SEPARATOR, operatorAddress)
}

func NETWORK_UPGRADE_TARGET_TIMESTAMP_KEY() string {
	return fmt.Sprintf("%s%s%s", NETWORK_UPGRADE, DB_KEY_SEPARATOR, "timestamp")
}

func NETWORK_UPGRADE_TARGET_BLOCKHEIGHT_KEY() string {
	return fmt.Sprintf("%s%s%s", NETWORK_UPGRADE, DB_KEY_SEPARATOR, "blockheight")
}
