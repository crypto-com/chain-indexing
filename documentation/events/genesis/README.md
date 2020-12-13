# Genesis Module Event List
  - [event::GENESIS_CREATED](#event_genesis_created)

## event::GENESIS_CREATED  
More about genesis [here](https://chain.crypto.com/docs/chain-details/genesis_file.html#genesis)  

*Name* : GenesisCreated

*Type* : [Base](../README.md#Understanding_an_EVENT)

*Structure* : 

| Key                                                  | Type            | Value                                                                                                                       |
| ---------------------------------------------------- | --------------- | --------------------------------------------------------------------------------------------------------------------------- |
| `genesis_time`                                       | *string*        | Time of Genesis. [Tendermint timestamps](https://github.com/tendermint/spec/blob/master/spec/core/data_structures.md#time). |
| `chain_id`                                           | *string*        | Tendermint chain ID. Value: `"testnet-croeseid-1"`                                                                          |
| `initial_height`                                     | *string*        | Initial blockchain height                                                                                                   |
| `app_hash`                                           | *string*        | Application hash                                                                                                            |
| `consensus_params`                                   | *object*        | Consensus parameters wrapper                                                                                                |
| `consensus_params.block`                             | *object*        | Consensus parameters Block wrapper                                                                                          |
| `consensus_params.block.max_bytes`                   | *string*        | Maximum bytes for a block                                                                                                   |
| `consensus_params.block.max_gas`                     | *string*        | Maximum gas for a block                                                                                                     |
| `consensus_params.block.time_iota_ms`                | *string*        | Time iota for a block in milliseconds                                                                                       |
| `consensus_params.evidence`                          | *object*        | Consensus parameters Evidence wrapper                                                                                       |
| `consensus_params.evidence.max_age_num_blocks`       | *string*        | Maximum age in number of blocks                                                                                             |
| `consensus_params.evidence.max_age_duration`         | *string*        | Maximum Duration of the evidence                                                                                            |
| `consensus_params.validator`                         | *object*        | Consensus parameters Validator wrapper                                                                                      |
| `consensus_params.validator.pub_key_types`           | *array(string)* | Array of public key types supported                                                                                         |
| `consensus_params.version`                           | *object*        | Empty value. Value: `{}`                                                                                                    |
| `app_state`                                          | *object*        | Application state wrapper                                                                                                   |
| `app_state.auth`                                     | *object*        | Auth Application state wrapper                                                                                              |
| `app_state.auth.params`                              | *object*        | Auth parameters wrapper                                                                                                     |
| `app_state.auth.params.max_memo_characters`          | *string*        | Maximum memo characters allowed                                                                                             |
| `app_state.auth.params.tx_sig_limit`                 | *string*        | Transaction signature limit                                                                                                 |
| `app_state.auth.params.tx_size_cost_per_byte`        | *string*        | Transaction cost per byte                                                                                                   |
| `app_state.auth.params.sig_verify_cost_ed25519`      | *string*        | Verify cost for `ed25519` curve                                                                                             |
| `app_state.auth.params.sig_verify_cost_secp256k1`    | *string*        | Verify cost for `secp256k1` curve                                                                                           |
| `app_state.auth.accounts`                            | *array(object)* | Array of auth accounts                                                                                                      |
| `_._.accounts.@type`                                 | *string*        | Cosmos SDK Type URL                                                                                                         |
| `_._.accounts.address`                               | *string*        | *`(Optional)`* Auth account blockchain address                                                                              |
| `_._.accounts.pub_key`                               | *object*        | Golang type: `interface{}`. Value : `null`                                                                                  |
| `_._.accounts.account_number`                        | *string*        | *`(Optional)`* Node account number for this account                                                                         |
| `_._.accounts.sequence`                              | *string*        | *`(Optional)`* Number of sends from this account                                                                            |
| `app_state.auth.accounts.base_vesting_account`       | *string*        | *`(Optional)`* Base Vesting account wrapper                                                                                 |
| `_._._.base_vesting_account.base_account`            | *string*        | *`(Optional)`* Base Vesting account wrapper                                                                                 |
| `_._._.base_vesting_account.original_vesting`        | *array(object)* | *`(Optional)`* Base Vesting account wrapper                                                                                 |
| `_._._.base_vesting_account.original_vesting.denom`  | *array(object)* | CRO amount denomination label                                                                                               |
| `_._._.base_vesting_account.original_vesting.amount` | *array(object)* | Vesting CRO amount in base unit                                                                                             |
| `_._._.base_vesting_account.delegated_free`          | *array(object)* | Golang type: *`[]interface{}`*                                                                                              |
| `_._._.base_vesting_account.delegated_vesting`       | *array(object)* | Golang type: *`[]interface{}`*                                                                                              |
| `_._._.base_vesting_account.end_time`                | *string*        | Designated end time for an auth account. (`seconds`)                                                                        |
| `app_state.chainmain`                                | *object*        | Chain main application state                                                                                                |
| `app_state.bank`                                     | *object*        | Bank module. `T.B.D`                                                                                                        |
| `app_state.distribution`                             | *object*        | Distribution module. `T.B.D`                                                                                                |
| `app_state.evidence`                                 | *object*        | Evidence module. `T.B.D`                                                                                                    |
| `app_state.genutil`                                  | *object*        | Genutil module. `T.B.D`                                                                                                     |
| `app_state.gov`                                      | *object*        | Governance module. `T.B.D`                                                                                                  |
| `app_state.mint`                                     | *object*        | Minting module. `T.B.D`                                                                                                     |
| `app_state.params`                                   | *object*        | Golang type: *`interface{}`*                                                                                                |
| `app_state.slashing`                                 | *object*        | Slashing module. `T.B.D`                                                                                                    |
| `app_state.staking`                                  | *object*        | Staking module. `T.B.D`                                                                                                     |
| `app_state.upgrade`                                  | *object*        | Upgrade module. `T.B.D`                                                                                                     |
| `name`                                               | *string*        | Specific Event Name. Value: `GenesisCreated`                                                                                |
| `version`                                            | *int*           | Event Version. Value: `1`                                                                                                   |
| `height`                                             | *int64*         | Height of the block containing the transaction                                                                              |
| `uuid`                                               | *string*        | Unique ID that is assigned on event creation                                                                                |

*Example* :  
```json
{
    "name": "GenesisCreated",
    "uuid": "2f0859f8-f44e-4af3-a0b2-cfdc3e56fff3",
    "height": 0,
    "Genesis": {
        "app_hash": "",
        "chain_id": "testnet-croeseid-1",
        "app_state": {
            "gov": {
                "votes": [],
                "deposits": [],
                "proposals": [],
                "tally_params": {
                    "quorum": "0.334000000000000000",
                    "threshold": "0.500000000000000000",
                    "veto_threshold": "0.334000000000000000"
                },
                "voting_params": {
                    "voting_period": "43200s"
                },
                "deposit_params": {
                    "min_deposit": [
                        {
                            "denom": "basetcro",
                            "amount": "10000000"
                        }
                    ],
                    "max_deposit_period": "43200s"
                },
                "starting_proposal_id": "1"
            },
            "auth": {
                "params": {
                    "tx_sig_limit": "7",
                    "max_memo_characters": "256",
                    "tx_size_cost_per_byte": "10",
                    "sig_verify_cost_ed25519": "590",
                    "sig_verify_cost_secp256k1": "1000"
                },
                "accounts": [
                    {
                        "@type": "/cosmos.vesting.v1beta1.DelayedVestingAccount",
                        "pub_key": null,
                        "base_vesting_account": {
                            "end_time": "1603788958",
                            "base_account": {
                                "address": "tcro1e3mpg4kkz9j5h4r28fl74mzmggmjw5e9rece0k",
                                "pub_key": null,
                                "sequence": "0",
                                "account_number": "0"
                            },
                            "delegated_free": [],
                            "original_vesting": [
                                {
                                    "denom": "basetcro",
                                    "amount": "2000000000000000000"
                                }
                            ],
                            "delegated_vesting": []
                        }
                    }
                ]
            },
            "bank": {
                "params": {
                    "send_enabled": [],
                    "default_send_enabled": true
                },
                "supply": [],
                "balances": [
                    {
                        "coins": [
                            {
                                "denom": "basetcro",
                                "amount": "20000000000000"
                            }
                        ],
                        "address": "tcro1fja5nsxz7gsqw4zccuuy8r7pjnjmc7dsdjun5p"
                    }
                ],
                "denom_metadata": [
                    {
                        "base": "basetcro",
                        "display": "tcro",
                        "denom_units": [
                            {
                                "denom": "basetcro",
                                "aliases": [
                                    "carson"
                                ],
                                "exponent": 0
                            },
                            {
                                "denom": "tcro",
                                "aliases": [],
                                "exponent": 8
                            }
                        ],
                        "description": "The native token of Crypto.com app."
                    }
                ]
            },
            "mint": {
                "minter": {
                    "inflation": "0.013000000000000000",
                    "annual_provisions": "0.000000000000000000"
                },
                "params": {
                    "mint_denom": "basetcro",
                    "goal_bonded": "0.670000000000000000",
                    "inflation_max": "0.020000000000000000",
                    "inflation_min": "0.007000000000000000",
                    "blocks_per_year": "6311520",
                    "inflation_rate_change": "0.013000000000000000"
                }
            },
            "params": null,
            "genutil": {
                "gen_txs": [
                    {
                        "body": {
                            "memo": "6989460472eb5f65faf4fdfff910f1bea800b10e@10.10.5.7:26656",
                            "messages": [
                                {
                                    "@type": "/cosmos.staking.v1beta1.MsgCreateValidator",
                                    "value": {
                                        "denom": "basetcro",
                                        "amount": "10000000000000"
                                    },
                                    "pubkey": "tcrocnclconspub1zcjduepq5xp88wqmrhkg3xuyl6vcey3d93kw6cdglkmq4ley3ysvjfx90jnqlvaxpc",
                                    "commission": {
                                        "rate": "0.100000000000000000",
                                        "max_rate": "0.200000000000000000",
                                        "max_change_rate": "0.010000000000000000"
                                    },
                                    "description": {
                                        "details": "",
                                        "moniker": "jotasdfsdro",
                                        "website": "",
                                        "identity": "",
                                        "security_contact": ""
                                    },
                                    "delegator_address": "tcro16kqr009ptgken6qsxnzfnyjfsq6q97g3fxwppq",
                                    "validator_address": "tcrocncl16kqr009ptgken6qsxnzfnyjfsq6q97g3uedcer",
                                    "min_self_delegation": "1"
                                }
                            ],
                            "timeout_height": "0",
                            "extension_options": [],
                            "non_critical_extension_options": []
                        },
                        "auth_info": {
                            "fee": {
                                "payer": "",
                                "amount": [],
                                "granter": "",
                                "gas_limit": "200000"
                            },
                            "signer_infos": [
                                {
                                    "sequence": "0",
                                    "mode_info": {
                                        "single": {
                                            "mode": "SIGN_MODE_DIRECT"
                                        }
                                    },
                                    "public_key": {
                                        "key": "AqdoWLPCuN3Jf1FSrsvadkq0Q2Eo0wCZA7aVC",
                                        "@type": "/cosmos.crypto.secp256k1.PubKey"
                                    }
                                }
                            ]
                        },
                        "signatures": [
                            "Pk0xhx2h8bFzGu1RYqWbahjoyHkoQWXkuQNefIc3RBwOXY2lNDz3Y0NDBOAxtV6QhXwNtSpbe2riqFrzX+RZ+g=="
                        ]
                    }
                ]
            },
            "staking": {
                "params": {
                    "bond_denom": "basetcro",
                    "max_entries": 7,
                    "max_validators": 150,
                    "unbonding_time": "600s",
                    "historical_entries": 100
                },
                "exported": false,
                "validators": [],
                "delegations": [],
                "redelegations": [],
                "last_total_power": "0",
                "last_validator_powers": [],
                "unbonding_delegations": []
            },
            "upgrade": {},
            "evidence": {
                "evidence": []
            },
            "slashing": {
                "params": {
                    "signed_blocks_window": "2000",
                    "min_signed_per_window": "0.500000000000000000",
                    "downtime_jail_duration": "3600s",
                    "slash_fraction_downtime": "0.001",
                    "slash_fraction_double_sign": "0.050000000000000000"
                },
                "missed_blocks": [],
                "signing_infos": []
            },
            "chainmain": {},
            "distribution": {
                "params": {
                    "community_tax": "0",
                    "base_proposer_reward": "0.010000000000000000",
                    "bonus_proposer_reward": "0.040000000000000000",
                    "withdraw_addr_enabled": true
                },
                "fee_pool": {
                    "community_pool": []
                },
                "previous_proposer": "",
                "outstanding_rewards": [],
                "validator_slash_events": [],
                "delegator_starting_infos": [],
                "delegator_withdraw_infos": [],
                "validator_current_rewards": [],
                "validator_historical_rewards": [],
                "validator_accumulated_commissions": []
            }
        },
        "genesis_time": "2020-10-13T08:55:58.046949Z",
        "initial_height": "1",
        "consensus_params": {
            "block": {
                "max_gas": "-1",
                "max_bytes": "22020096",
                "time_iota_ms": "1000"
            },
            "version": {},
            "evidence": {
                "max_age_duration": "172800000000000",
                "max_age_num_blocks": "100000"
            },
            "validator": {
                "pub_key_types": [
                    "ed25519"
                ]
            }
        }
    },
    "version": 1
}
```  