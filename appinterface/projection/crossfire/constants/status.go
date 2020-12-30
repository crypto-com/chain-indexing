package constants

const DO_NOT_MODIFY = "[do-not-modify]"

const BONDED NodeStatus = "Bonded"
const JAILED NodeStatus = "Jailed"
const UNBONDED NodeStatus = "Unbonded"
const UNBONDING NodeStatus = "Unbonding"

type NodeStatus = string

const INCOMPLETED TaskStatus = "Incompleted"
const COMPLETED TaskStatus = "Completed"
const MISSED TaskStatus = "Missed"

type TaskStatus = string

<<<<<<< HEAD
const VOTE_OPTION_NO = "VOTE_OPTION_NO"
const VOTE_OPTION_NO_WITH_VETO = "VOTE_OPTION_NO_WITH_VETO"

const VOTE_OPTION_YES = "VOTE_OPTION_YES"
const VOTE_OPTION_ABSTAIN = "VOTE_OPTION_ABSTAIN"
const VOTE_OPTION_UNSPECIFIED = "VOTE_OPTION_UNSPECIFIED"

const DO_NOT_MODIFY = "[do-not-modify]"
=======
const PHASE1_COMMIT_COUNT ChainStatsKey = "phase1_commit_count"
const PHASE2_COMMIT_COUNT ChainStatsKey = "phase2_commit_count"
const PHASE3_COMMIT_COUNT ChainStatsKey = "phase3_commit_count"
type ChainStatsKey = string
>>>>>>> feat: update crossfireChainStatsView for storing integers
