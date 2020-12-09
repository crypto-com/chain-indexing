# Events Detailing Documentation
This documentation explains different events supported on the chain-indexing service.

## Understanding an EVENT
 An event is a basic detailed representation of an state change transaction on Crypto.com blockchain. An event will always have the `Base` properties which are:
```go
type Base struct {
	EventName    string `json:"name"`
	EventVersion int    `json:"version"`
	BlockHeight  int64  `json:"height"`
	EventUUID    string `json:"uuid"`
}
```
where 
`name` : Name for the event
`version` : Numerical version maintained for the event
`height` : Block height of the event
`uuid` : Unique ID that is assigned on event creation

## MsgBase
`MsgBase` consists of a few properties pertaining to a blockchain event. Refer to below structure for more details.
```go
type MsgBase struct {
	event.Base

	MsgName   string `json:"msgName"`
	MsgTxHash string `json:"txHash"`
	MsgIndex  int    `json:"msgIndex"`
}
```
where 
`event.Base` : All the properties of `Base` structure explained [here](./README.md#Understanding-an-EVENT)
`msgName` : `msg*` of the Crypto.com blockchain such as `MsgSend` or `MsgCreateValidator` etc.
`txHash` : Blockchain TxID for the transaction containing the event
`msgIndex` : Message Index on the block created

## Categories
- [Bank](./bank)
- Staking
- Distribution
- Genesis
- Governance
- Validator
- Slashing