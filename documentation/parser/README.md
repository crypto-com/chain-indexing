# Possible signer in each msg type

| Msg Type                          | Possible Signer                                | Code Reference from Cosmos SDK                                                                                 |
| --------------------------------- | ---------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `MsgSend`                         | `msg.FromAddress`                              | [x/bank/types/msgs.go](https://github.com/cosmos/cosmos-sdk/blob/main/x/bank/types/msgs.go#L56)                |
| `MsgMultiSend`                    | `msg.Inputs`                                   | [x/bank/types/msgs.go](https://github.com/cosmos/cosmos-sdk/blob/main/x/bank/types/msgs.go#L95)                |
| `MsgSetWithdrawAddress`           | `msg.DelegatorAddress`                         | [x/distribution/types/msg.go](https://github.com/cosmos/cosmos-sdk/blob/main/x/distribution/types/msg.go#L31)  |
| `MsgWithdrawDelegatorReward`      | `msg.DelegatorAddress`                         | [x/distribution/types/msg.go](https://github.com/cosmos/cosmos-sdk/blob/main/x/distribution/types/msg.go#L65)  |
| `MsgWithdrawValidatorCommission`  | `msg.ValidatorAddress`                         | [x/distribution/types/msg.go](https://github.com/cosmos/cosmos-sdk/blob/main/x/distribution/types/msg.go#L97)  |
| `MsgFundCommunityPool`            | `msg.Depositor`                                | [x/distribution/types/msg.go](https://github.com/cosmos/cosmos-sdk/blob/main/x/distribution/types/msg.go#L133) |
| `MsgSubmitProposal`               | `m.Proposer`                                   | [x/gov/types/v1/msgs.go](https://github.com/cosmos/cosmos-sdk/blob/main/x/gov/types/v1/msgs.go#L90)            |
| `MsgVote`                         | `msg.Voter`                                    | [x/gov/types/v1/msgs.go](https://github.com/cosmos/cosmos-sdk/blob/main/x/gov/types/v1/msgs.go#L171)           |
| `MsgDeposit`                      | `msg.Depositor`                                | [x/gov/types/v1/msgs.go](https://github.com/cosmos/cosmos-sdk/blob/main/x/gov/types/v1/msgs.go#L135)           |
| `MsgDelegate`                     | `msg.DelegatorAddress`                         | [x/staking/types/msg.go](https://github.com/cosmos/cosmos-sdk/blob/main/x/staking/types/msg.go#L213)           |
| `MsgUndelegate`                   | `msg.DelegatorAddress`                         | [x/staking/types/msg.go](https://github.com/cosmos/cosmos-sdk/blob/main/x/staking/types/msg.go#L313)           |
| `MsgBeginRedelegate`              | `msg.DelegatorAddress`                         | [x/staking/types/msg.go](https://github.com/cosmos/cosmos-sdk/blob/main/x/staking/types/msg.go#L263)           |
| `MsgCreateValidator`              | `msg.DelegatorAddress`, `msg.ValidatorAddress` | [x/staking/types/msg.go](https://github.com/cosmos/cosmos-sdk/blob/main/x/staking/types/msg.go#L66)            |
| `MsgEditValidator`                | `msg.ValidatorAddress`                         | [x/staking/types/msg.go](https://github.com/cosmos/cosmos-sdk/blob/main/x/staking/types/msg.go#L159)           |
| `MsgUnjail`                       | `msg.ValidatorAddr`                            | [x/slashing/types/msg.go](https://github.com/cosmos/cosmos-sdk/blob/main/x/slashing/types/msg.go#L29)          |
| `MsgNFTIssueDenom`                | `m.Sender`                                     | [x/nft/msgs.go](https://github.com/cosmos/cosmos-sdk/blob/main/x/nft/msgs.go#L38)                              |
| `MsgMintNFT"`                     | `m.sender`                                     | [x/nft/msgs.go](https://github.com/cosmos/cosmos-sdk/blob/main/x/nft/msgs.go#L38)                              |
| `MsgTransferNFT`                  | `m.Sender `                                    | [x/nft/msgs.go](https://github.com/cosmos/cosmos-sdk/blob/main/x/nft/msgs.go#L38)                              |
| `MsgEditNFT"`                     | `m.Sender`                                     | [x/nft/msgs.go](https://github.com/cosmos/cosmos-sdk/blob/main/x/nft/msgs.go#L38)                              |
| `MsgBurnNFT"`                     | `m.Sender`                                     | [x/nft/msgs.go](https://github.com/cosmos/cosmos-sdk/blob/main/x/nft/msgs.go#L38)                              |
| `MsgGrant`                        | `msg.Granter`                                  | [x/authz/msgs.go](https://github.com/cosmos/cosmos-sdk/blob/main/x/authz/msgs.go#L46)                          |
| `MsgRevoke`                       | `msg.Granter`                                  | [x/authz/msgs.go](https://github.com/cosmos/cosmos-sdk/blob/main/x/authz/msgs.go#L46)                          |
| `MsgExec`                         | `msg.Grantee`                                  | [x/authz/msgs.go](https://github.com/cosmos/cosmos-sdk/blob/main/x/authz/msgs.go#L46)                          |
| `MsgGrantAllowance`               | `msg.Granter`                                  | [x/feegrant/msgs.go](https://github.com/cosmos/cosmos-sdk/blob/main/x/feegrant/msgs.go#L58)                    |
| `MsgRevokeAllowance`              | `msg.Granter`                                  | [x/feegrant/msgs.go](https://github.com/cosmos/cosmos-sdk/blob/main/x/feegrant/msgs.go#L118)                   |
| `MsgCreateVestingAccount`         | `msg.FromAddress`                              | [x/auth/vesting/types/msgs.go](https://github.com/cosmos/cosmos-sdk/blob/main/x/auth/vesting/types/msgs.go#L74 |          