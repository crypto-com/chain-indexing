# Change Log
 
The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).
 
## [0.0.1] - 2021-10-27
 
### Added
   
### Changed
 
- `ParseMsgRecvPacket()`

  When `fungible_token_packet` not in TxsResult log, we do not write a customized error message to `MsgIBCRecvPacket.Params.PacketAck` anymore. 
  
  Previously, `IBCChannel` projection is using `MsgIBCRecvPacket.Params.PacketAck.MaybeError` to check MsgRecvPacket is success or not.
  
  Now `IBCChannel` projection is using `MsgIBCRecvPacket.Params.MaybeFungibleTokenPacketData.Success` to check MsgRecvPacket is success or not. The customized error message is removed.
 
### Fixed