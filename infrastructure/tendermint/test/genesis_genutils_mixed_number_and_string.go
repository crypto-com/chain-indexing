package infrastructure_tendermint_test

// This genesis sample has two gentxs. The first one has number in its numeric fields (e.g.
// timeout_height, sequence). The second one use string in its numeric fields. Both are valid
// genesis file
const GENESIS_MIXED_NUMBER_AND_STRING_JSON = `
{
	"jsonrpc": "2.0",
	"id": -1,
	"result": {
		"genesis": {
			"genesis_time": "2021-01-01T00:00:00.000000Z",
			"chain_id": "devnet",
			"app_hash": "",
			"app_state": {
			"auth": {
				"accounts": [],
				"params": {
				"max_memo_characters": "256",
				"sig_verify_cost_ed25519": "590",
				"sig_verify_cost_secp256k1": "1000",
				"tx_sig_limit": "7",
				"tx_size_cost_per_byte": "10"
				}
			},
			"bank": {
				"balances": [],
				"denom_metadata": [
					{
						"base": "basecro",
						"denom_units": [
						{
							"aliases": [
								"carson"
							],
							"denom": "basecro",
							"exponent": 0
						},
						{
							"denom": "cro",
							"exponent": 8
						}
						],
						"description": "The native token of Crypto.org app.",
						"display": "cro"
					}
				],
				"params": {
					"default_send_enabled": true,
					"send_enabled": []
				},
				"supply": []
			},
			"capability": {
				"index": "1",
				"owners": []
			},
			"chainmain": {},
			"distribution": {
				"delegator_starting_infos": [],
				"delegator_withdraw_infos": [],
				"fee_pool": {
					"community_pool": []
				},
				"outstanding_rewards": [],
				"params": {
					"base_proposer_reward": "0.010000000000000000",
					"bonus_proposer_reward": "0.040000000000000000",
					"community_tax": "0.000000000000000000",
					"withdraw_addr_enabled": true
				},
				"previous_proposer": "",
				"validator_accumulated_commissions": [],
				"validator_current_rewards": [],
				"validator_historical_rewards": [],
				"validator_slash_events": []
			},
			"evidence": {
				"evidence": []
			},
			"genutil": {
				"gen_txs": [
					{
						"body": {
							"messages": [
								{
									"@type": "/cosmos.staking.v1beta1.MsgCreateValidator",
									"description": {
										"moniker": "test-01",
										"identity": "",
										"website": "https://crypto.org",
										"security_contact": "chain-security@crypto.org",
										"details": "Test 01"
									},
									"commission": {
										"rate": "1.000000000000000000",
										"max_rate": "1.000000000000000000",
										"max_change_rate": "0.100000000000000000"
									},
									"min_self_delegation": "100000000",
									"delegator_address": "cro16wry2347tzptx79kkk6v5rrecvrmf5zkmtsw72",
									"validator_address": "crocncl16wry2347tzptx79kkk6v5rrecvrmf5zkcxn8uk",
									"pubkey": {
										"@type": "/cosmos.crypto.ed25519.PubKey",
										"key": "JRS9Ldx0f1xmtTeN/Wi1q7rQdqNzZne/68wXMz2Vk+Y="
									},
									"value": {
										"denom": "basecro",
										"amount": "1000"
									}
								}
							],
							"memo": "Test 01",
							"timeout_height": 0,
							"extension_options": [],
							"non_critical_extension_options": []
						},
						"auth_info": {
							"signer_infos": [
								{
									"public_key": {
										"@type": "/cosmos.crypto.secp256k1.PubKey",
										"key": "AvCj8A1kBC89VrONuCDA7wgbU5LYYl1HRvn5wwr42Os9"
									},
									"mode_info": {
										"single": {
										"mode": "SIGN_MODE_DIRECT"
										}
									},
									"sequence": 0
								}
							],
							"fee": {
								"amount": [],
								"gas_limit": 200000,
								"payer": "",
								"granter": ""
							}
						},
						"signatures": [
							"S0lDdl1jGtPglQpTuYMp1VYMXny78CUR27WUHvE9+vpvcT4W7ZYAYWZC1bqp04Dvlsn6fCzcRAljekUuveBfbg=="
						]
					},
					{
						"body": {
						  "messages": [
							{
							  "@type": "/cosmos.staking.v1beta1.MsgCreateValidator",
							  "description": {
								"moniker": "Test 02",
								"identity": "",
								"website": "",
								"security_contact": "Test 02",
								"details": ""
							  },
							  "commission": {
								"rate": "0.100000000000000000",
								"max_rate": "0.200000000000000000",
								"max_change_rate": "0.010000000000000000"
							  },
							  "min_self_delegation": "1",
							  "delegator_address": "cro1sseazswz3zyn70a79mjwtwnw5ccdxw2cydeetx",
							  "validator_address": "crocncl1sseazswz3zyn70a79mjwtwnw5ccdxw2c8q6sf6",
							  "pubkey": {
								"@type": "/cosmos.crypto.ed25519.PubKey",
								"key": "fd0vscPhXhbmqvNkt1TWu7jritG31+NFgmTkGYaMVc4="
							  },
							  "value": {
								"denom": "basecro",
								"amount": "50000000000000"
							  }
							}
						  ],
						  "memo": "",
						  "timeout_height": "0",
						  "extension_options": [],
						  "non_critical_extension_options": []
						},
						"auth_info": {
						  "signer_infos": [],
						  "fee": {
							"amount": [
							  {
								"denom": "basecro",
								"amount": "80000"
							  }
							],
							"gas_limit": "800000",
							"payer": "",
							"granter": ""
						  }
						},
						"signatures": []
					}
				]
			},
			"gov": {
				"deposit_params": {
				"max_deposit_period": "1209600000000000ns",
				"min_deposit": [
					{
					"denom": "basecro",
					"amount": "2000000000000"
					}
				]
				},
				"deposits": [],
				"proposals": [],
				"starting_proposal_id": "1",
				"tally_params": {
				"quorum": "0.334000000000000000",
				"threshold": "0.500000000000000000",
				"veto_threshold": "0.334000000000000000"
				},
				"votes": [],
				"voting_params": {
				"voting_period": "1209600000000000ns"
				}
			},
			"ibc": {
				"channel_genesis": {
				"ack_sequences": [],
				"acknowledgements": [],
				"channels": [],
				"commitments": [],
				"next_channel_sequence": "0",
				"receipts": [],
				"recv_sequences": [],
				"send_sequences": []
				},
				"client_genesis": {
				"clients": [],
				"clients_consensus": [],
				"clients_metadata": [],
				"create_localhost": false,
				"next_client_sequence": "0",
				"params": {
					"allowed_clients": [
					"06-solomachine",
					"07-tendermint"
					]
				}
				},
				"connection_genesis": {
				"client_connection_paths": [],
				"connections": [],
				"next_connection_sequence": "0"
				}
			},
			"mint": {
				"minter": {
				"annual_provisions": "0.000000000000000000",
				"inflation": "0.020000000000000000"
				},
				"params": {
				"blocks_per_year": "6311520",
				"goal_bonded": "0.300000000000000000",
				"inflation_max": "0.025000000000000000",
				"inflation_min": "0.012000000000000000",
				"inflation_rate_change": "0.006000000000000000",
				"mint_denom": "basecro"
				}
			},
			"params": null,
			"slashing": {
				"missed_blocks": [],
				"params": {
				"downtime_jail_duration": "86400s",
				"min_signed_per_window": "0.500000000000000000",
				"signed_blocks_window": "5000",
				"slash_fraction_double_sign": "0.050000000000000000",
				"slash_fraction_downtime": "0.000000000000000000"
				},
				"signing_infos": []
			},
			"staking": {
				"delegations": [],
				"exported": false,
				"last_total_power": "0",
				"last_validator_powers": [],
				"params": {
				"bond_denom": "basecro",
				"historical_entries": 10000,
				"max_entries": 7,
				"max_validators": 100,
				"unbonding_time": "2419200000000000ns"
				},
				"redelegations": [],
				"unbonding_delegations": [],
				"validators": []
			},
			"supply": {},
			"transfer": {
				"denom_traces": [],
				"params": {
				"receive_enabled": false,
				"send_enabled": false
				},
				"port_id": "transfer"
			},
			"upgrade": {},
			"vesting": {}
			},
			"consensus_params": {
			"block": {
				"max_bytes": "500000",
				"max_gas": "81500000",
				"time_iota_ms": "1000"
			},
			"evidence": {
				"max_age_duration": "2419200000000000",
				"max_age_num_blocks": "403200",
				"max_bytes": "150000"
			},
			"validator": {
				"pub_key_types": [
				"ed25519"
				]
			},
			"version": {}
			},
			"initial_height": "1"
		}
	}
}`
