package types

type Chain = string

type Status = string

type BridgeType = string

type Direction = string

const (
	CHAIN_CRYPTO_ORG_CHAIN Chain = "Crypto.org-Chain"
	CHAIN_CRONOS           Chain = "Cronos"
	CHAIN_ETHEREUM         Chain = "Ethereum"
)

const (
	STATUS_PENDING                Status = "Pending"
	STATUS_CANCELLED              Status = "Cancelled"
	STATUS_COUNTERPARTY_CONFIRMED Status = "CounterpartyConfirmed"
	STATUS_FAILED_ON_CHAIN        Status = "FailedOnChain"
	STATUS_COUNTERPARTY_REJECTED  Status = "CounterpartyRejected"
	// Packet already relayed, do nothing (no-op)
	// https://github.com/cosmos/ibc-go/blob/760d15a3a55397678abe311b7f65203b2e8437d6/modules/core/04-channel/keeper/packet.go#L239
	STATUS_NO_OPERATION Status = "NoOperation"
)

const (
	BRIDGE_TYPE_GRAVITY BridgeType = "Gravity"
	BRIDGE_TYPE_IBC     BridgeType = "IBC"
)

const (
	DIRECTION_OUTGOING Direction = "Outgoing"
	DIRECTION_INCOMING Direction = "Incoming"
	DIRECTION_RESPONSE Direction = "Response"
)
