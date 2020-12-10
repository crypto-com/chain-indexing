# Events Documentation
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
* `name` : Name for the event  
* `version` : Numerical version maintained for the event  
* `height` : Block height of the event  
* `uuid` : Unique ID that is assigned on event creation  

Associated functions with the structure:  
* `Name()` : returns name of the event as `string`  
* `Version()` : returns the event version as `int`  
* `Height()` : returns the height of the block containing the event  
* `UUID()` : returns the unique event ID as `string`  

Before going further let's first understand [MsgBase](./README.md#MsgBase).

## MsgBase
It extends the `Base` object by enriching it with other blockchain properties. `MsgBase` consists of a few properties pertaining to a blockchain event. Refer to below structure:

```go
type MsgBase struct {
	event.Base

	MsgName   string `json:"msgName"`
	MsgTxHash string `json:"txHash"`
	MsgIndex  int    `json:"msgIndex"`
}
```
where  
* `event.Base` : `Base` properties as explained [here](./README.md#Understanding-an-EVENT)  
* `msgName` : `msg*` of the Crypto.com blockchain such as `MsgSend` or `MsgCreateValidator` etc.  
* `txHash` : Blockchain TxID for the transaction containing the event  
* `msgIndex` : Corresponding index of the `Msg*` inside the `tx.Body.Messages` list  

In addition to the above there are different methods available which can be described as below:  
* `MsgType()` : returns `msgName` as `string`  
* `TxHash()` : returns `txHash` as `string`  
* `TxSuccess()` : returns `true` if the transaction is successful or else `false`  
* `eventName(msgName string, txSuccess bool)` : returns the concatenated eventName as `string`  


## Categories
- [Bank](./bank)
- Staking
- Distribution
- Genesis
- Governance
- Validator
- Slashing