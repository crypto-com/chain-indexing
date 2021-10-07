package view

type Status string

type BridgeType string

type Direction string

const (
	STATUS_PENDING                Status = "Pending"
	STATUS_CANCELLED              Status = "Cancelled"
	STATUS_COUNTERPARTY_CONFIRMED Status = "CounterpartyConfirmed"
	STATUS_FAILED                 Status = "Failed"
	STATUS_COUNTERPARTY_REJECTED  Status = "CounterpartyRejected"
)

const (
	BRIDGE_TYPE_GRAVITY BridgeType = "Gravity"
	BRIDGE_TYPE_IBC     BridgeType = "IBC"
)

const (
	DIRECTION_OUTGOING Direction = "Outgoing"
	DIRECTION_INCOMING Direction = "Incoming"
)
