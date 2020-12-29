package constants

const BONDED NodeStatus = "Bonded"
const JAILED NodeStatus = "Jailed"
const UNBONDED NodeStatus = "Unbonded"
const UNBONDING NodeStatus = "Unbonding"
type NodeStatus = string

const INCOMPLETED TaskStatus = "Incompleted"
const COMPLETED TaskStatus = "Completed"
const MISSED TaskStatus = "Missed"
type TaskStatus = string

const DO_NOT_MODIFY = "[do-not-modify]"
