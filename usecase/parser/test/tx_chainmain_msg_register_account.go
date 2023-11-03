package usecase_parser_test

const TX_CHAINMAIN_MSG_REGISTER_ACCOUNT_BLOCK_RESP = `
{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "block_id": {
      "hash": "B69E8227BFBA171B5D1729A62989E6039AB5A0ACCE178A1241506EE936F6AC9C",
      "parts": {
        "total": 1,
        "hash": "E1598B165E4448A7F0932230EE5D1AFC3D73594318350A3CE5130A2E5A3EFA14"
      }
    },
    "block": {
      "header": {
        "version": {
          "block": "11"
        },
        "chain_id": "testnet-croeseid-5",
        "height": "67461",
        "time": "2022-11-18T09:45:50.24590599Z",
        "last_block_id": {
          "hash": "25FE8EE277297834956B98C943E4CB8012D881F5B67C35AB3015D9AB5292DCEC",
          "parts": {
            "total": 1,
            "hash": "1355F96F8E709F191CE40872BCE4D4A66567B406266183B4D333A4D96E226482"
          }
        },
        "last_commit_hash": "4CBEDDF3082D725B8CF2AFC7B329A6545F0F2368E1BA320DAB7B02B1CEDED96C",
        "data_hash": "D89BA4817545EA69DC395747D64D56326548EDF0D069D0C88BC88EBC57503576",
        "validators_hash": "30876D8A0FA3F977F37F7A96B1D0747232E03328378354F885996E265C161ADC",
        "next_validators_hash": "30876D8A0FA3F977F37F7A96B1D0747232E03328378354F885996E265C161ADC",
        "consensus_hash": "372B4AE845086C837EFEF79A189B085B1FD6610C53F3BEB17EE0E27B347C06DE",
        "app_hash": "A7F503198899968B69D744A676115BF751D8488B11C3C89B4D1FA068EDB4E1AF",
        "last_results_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "proposer_address": "C9000A9E3F0093E25ED2E2484B8899C30ABC7334"
      },
      "data": {
        "txs": [
          "CmoKaAooL2NoYWlubWFpbi5pY2FhdXRoLnYxLk1zZ1JlZ2lzdGVyQWNjb3VudBI8Cit0Y3JvMW5wN3p0Y2ZleWNxd2hqMG5yOGh4ZnUwbGZqejI3dGVscXg1M3JhEg1jb25uZWN0aW9uLTE4EmsKUQpGCh8vY29zbW9zLmNyeXB0by5zZWNwMjU2azEuUHViS2V5EiMKIQOnMc0FKyjIWGBK3hD5JiMyD0L7aLCL8g7zCXRnNN5x1BIECgIIARirARIWChAKCGJhc2V0Y3JvEgQ1MDAwEP7hCRpA1xZfFFDtIYo1d6qsEcGJbMnVG0PLl4k4UU1/15ro+XdTSN1Ye2Jgu8NJdUWzWrCL87kCG9TtaiEovv+cd14sgA=="
        ]
      },
      "evidence": {
        "evidence": []
      },
      "last_commit": {
        "height": "67460",
        "round": 0,
        "block_id": {
          "hash": "25FE8EE277297834956B98C943E4CB8012D881F5B67C35AB3015D9AB5292DCEC",
          "parts": {
            "total": 1,
            "hash": "1355F96F8E709F191CE40872BCE4D4A66567B406266183B4D333A4D96E226482"
          }
        },
        "signatures": [
          {
            "block_id_flag": 2,
            "validator_address": "47FFED8892045D94FBD2DD59B7CD8E633E0593D8",
            "timestamp": "2022-11-18T09:45:50.144690726Z",
            "signature": "WZhxHwiTFPXpSrrrnyld4Ek0U9nZOy8V0F2NYniEauw0gnkD2e6eVcE/K7IEOCbEtXszCs+QdOf5kx9Ic5JhDA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "A0704795B61B37C87190D97672C14E9E9ACA408F",
            "timestamp": "2022-11-18T09:45:50.24590599Z",
            "signature": "YfV7pR7oVNkCaYJp5i7NuHwvRfqYKhB8YwMlhhx1sKM5Gt9iQiFJrNUsdGc5xferyRgTt7hJfnjVZ+d5AR32Dg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "C9000A9E3F0093E25ED2E2484B8899C30ABC7334",
            "timestamp": "2022-11-18T09:45:50.246109367Z",
            "signature": "zVdvhhbY7cJjzOb0PNnRxdde8LL8MHvuvaWGMCYEj9dM9k7pxcDNlv7IUFRmMXcwcGMlh/7ww3h6fC6gVTUtAA=="
          }
        ]
      }
    }
  }
}
`

const TX_CHAINMAIN_MSG_REGISTER_ACCOUNT_BLOCK_RESULTS_RESP = `
{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "67461",
    "txs_results": [
      {
        "code": 0,
        "data": "EjIKMC9jaGFpbm1haW4uaWNhYXV0aC52MS5Nc2dSZWdpc3RlckFjY291bnRSZXNwb25zZQ==",
        "log": "[{\"msg_index\":0,\"events\":[{\"type\":\"channel_open_init\",\"attributes\":[{\"key\":\"port_id\",\"value\":\"icacontroller-tcro1np7ztcfeycqwhj0nr8hxfu0lfjz27telqx53ra\"},{\"key\":\"channel_id\",\"value\":\"channel-48\"},{\"key\":\"counterparty_port_id\",\"value\":\"icahost\"},{\"key\":\"counterparty_channel_id\"},{\"key\":\"connection_id\",\"value\":\"connection-18\"},{\"key\":\"version\",\"value\":\"{\\\"fee_version\\\":\\\"ics29-1\\\",\\\"app_version\\\":\\\"{\\\\\\\"version\\\\\\\":\\\\\\\"ics27-1\\\\\\\",\\\\\\\"controller_connection_id\\\\\\\":\\\\\\\"connection-18\\\\\\\",\\\\\\\"host_connection_id\\\\\\\":\\\\\\\"connection-1\\\\\\\",\\\\\\\"address\\\\\\\":\\\\\\\"\\\\\\\",\\\\\\\"encoding\\\\\\\":\\\\\\\"proto3\\\\\\\",\\\\\\\"tx_type\\\\\\\":\\\\\\\"sdk_multi_msg\\\\\\\"}\\\"}\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"register_account\"},{\"key\":\"module\",\"value\":\"ibc_channel\"}]}]}]",
        "info": "",
        "gas_wanted": "159998",
        "gas_used": "158041",
        "events": [
          {
            "type": "coin_spent",
            "attributes": [
              {
                "key": "c3BlbmRlcg==",
                "value": "dGNybzFucDd6dGNmZXljcXdoajBucjhoeGZ1MGxmanoyN3RlbHF4NTNyYQ==",
                "index": true
              },
              {
                "key": "YW1vdW50",
                "value": "NTAwMGJhc2V0Y3Jv",
                "index": true
              }
            ]
          },
          {
            "type": "coin_received",
            "attributes": [
              {
                "key": "cmVjZWl2ZXI=",
                "value": "dGNybzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHhoemFoYQ==",
                "index": true
              },
              {
                "key": "YW1vdW50",
                "value": "NTAwMGJhc2V0Y3Jv",
                "index": true
              }
            ]
          },
          {
            "type": "transfer",
            "attributes": [
              {
                "key": "cmVjaXBpZW50",
                "value": "dGNybzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHhoemFoYQ==",
                "index": true
              },
              {
                "key": "c2VuZGVy",
                "value": "dGNybzFucDd6dGNmZXljcXdoajBucjhoeGZ1MGxmanoyN3RlbHF4NTNyYQ==",
                "index": true
              },
              {
                "key": "YW1vdW50",
                "value": "NTAwMGJhc2V0Y3Jv",
                "index": true
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "c2VuZGVy",
                "value": "dGNybzFucDd6dGNmZXljcXdoajBucjhoeGZ1MGxmanoyN3RlbHF4NTNyYQ==",
                "index": true
              }
            ]
          },
          {
            "type": "tx",
            "attributes": [
              {
                "key": "ZmVl",
                "value": "NTAwMGJhc2V0Y3Jv",
                "index": true
              },
              {
                "key": "ZmVlX3BheWVy",
                "value": "dGNybzFucDd6dGNmZXljcXdoajBucjhoeGZ1MGxmanoyN3RlbHF4NTNyYQ==",
                "index": true
              }
            ]
          },
          {
            "type": "tx",
            "attributes": [
              {
                "key": "YWNjX3NlcQ==",
                "value": "dGNybzFucDd6dGNmZXljcXdoajBucjhoeGZ1MGxmanoyN3RlbHF4NTNyYS8xNzE=",
                "index": true
              }
            ]
          },
          {
            "type": "tx",
            "attributes": [
              {
                "key": "c2lnbmF0dXJl",
                "value": "MXhaZkZGRHRJWW8xZDZxc0VjR0piTW5WRzBQTGw0azRVVTEvMTVybytYZFRTTjFZZTJKZ3U4TkpkVVd6V3JDTDg3a0NHOVR0YWlFb3Z2K2NkMTRzZ0E9PQ==",
                "index": true
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "YWN0aW9u",
                "value": "cmVnaXN0ZXJfYWNjb3VudA==",
                "index": true
              }
            ]
          },
          {
            "type": "channel_open_init",
            "attributes": [
              {
                "key": "cG9ydF9pZA==",
                "value": "aWNhY29udHJvbGxlci10Y3JvMW5wN3p0Y2ZleWNxd2hqMG5yOGh4ZnUwbGZqejI3dGVscXg1M3Jh",
                "index": true
              },
              {
                "key": "Y2hhbm5lbF9pZA==",
                "value": "Y2hhbm5lbC00OA==",
                "index": true
              },
              {
                "key": "Y291bnRlcnBhcnR5X3BvcnRfaWQ=",
                "value": "aWNhaG9zdA==",
                "index": true
              },
              {
                "key": "Y291bnRlcnBhcnR5X2NoYW5uZWxfaWQ=",
                "value": null,
                "index": true
              },
              {
                "key": "Y29ubmVjdGlvbl9pZA==",
                "value": "Y29ubmVjdGlvbi0xOA==",
                "index": true
              },
              {
                "key": "dmVyc2lvbg==",
                "value": "eyJmZWVfdmVyc2lvbiI6ImljczI5LTEiLCJhcHBfdmVyc2lvbiI6IntcInZlcnNpb25cIjpcImljczI3LTFcIixcImNvbnRyb2xsZXJfY29ubmVjdGlvbl9pZFwiOlwiY29ubmVjdGlvbi0xOFwiLFwiaG9zdF9jb25uZWN0aW9uX2lkXCI6XCJjb25uZWN0aW9uLTFcIixcImFkZHJlc3NcIjpcIlwiLFwiZW5jb2RpbmdcIjpcInByb3RvM1wiLFwidHhfdHlwZVwiOlwic2RrX211bHRpX21zZ1wifSJ9",
                "index": true
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "bW9kdWxl",
                "value": "aWJjX2NoYW5uZWw=",
                "index": true
              }
            ]
          }
        ],
        "codespace": ""
      }
    ],
    "begin_block_events": [
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "dGNybzFtM2gzMHdsdnNmOGxscnV4dHB1a2R2c3kwa20ya3VtODdseDltcQ==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "ODk2NjU2MzU4NWJhc2V0Y3Jv",
            "index": true
          }
        ]
      },
      {
        "type": "coinbase",
        "attributes": [
          {
            "key": "bWludGVy",
            "value": "dGNybzFtM2gzMHdsdnNmOGxscnV4dHB1a2R2c3kwa20ya3VtODdseDltcQ==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "ODk2NjU2MzU4NWJhc2V0Y3Jv",
            "index": true
          }
        ]
      },
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "dGNybzFtM2gzMHdsdnNmOGxscnV4dHB1a2R2c3kwa20ya3VtODdseDltcQ==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "ODk2NjU2MzU4NWJhc2V0Y3Jv",
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "dGNybzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHhoemFoYQ==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "ODk2NjU2MzU4NWJhc2V0Y3Jv",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "dGNybzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHhoemFoYQ==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "dGNybzFtM2gzMHdsdnNmOGxscnV4dHB1a2R2c3kwa20ya3VtODdseDltcQ==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "ODk2NjU2MzU4NWJhc2V0Y3Jv",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "dGNybzFtM2gzMHdsdnNmOGxscnV4dHB1a2R2c3kwa20ya3VtODdseDltcQ==",
            "index": true
          }
        ]
      },
      {
        "type": "mint",
        "attributes": [
          {
            "key": "Ym9uZGVkX3JhdGlv",
            "value": "MC4wMDA1MzE4MDEwMTc1MzA2MjU=",
            "index": true
          },
          {
            "key": "aW5mbGF0aW9u",
            "value": "MC4wMjAwNjQwMTc2MDY5NDI3MDM=",
            "index": true
          },
          {
            "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
            "value": "NTY1OTI2NDU0MDM2NzQ3MTEuNDcxMzMzMjUyNTkwNjM5OTk3",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "ODk2NjU2MzU4NQ==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "dGNybzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHhoemFoYQ==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "ODk2NjU2MzU4NWJhc2V0Y3Jv",
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "dGNybzFqdjY1czNncnFmNnY2amwzZHA0dDZjOXQ5cms5OWNkODMzOXA0bA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "ODk2NjU2MzU4NWJhc2V0Y3Jv",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "dGNybzFqdjY1czNncnFmNnY2amwzZHA0dDZjOXQ5cms5OWNkODMzOXA0bA==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "dGNybzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHhoemFoYQ==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "ODk2NjU2MzU4NWJhc2V0Y3Jv",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "dGNybzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHhoemFoYQ==",
            "index": true
          }
        ]
      },
      {
        "type": "proposer_reward",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NDQ4MzI4MTc5LjI1MDAwMDAwMDAwMDAwMDAwMGJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxN3h6ZDloNDAwN2o0djhmNGs3dm41N2Y3ank4d2c4bWo5N2o1aHU=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NDQ4MzI4MTcuOTI1MDAwMDAwMDAwMDAwMDAwYmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxN3h6ZDloNDAwN2o0djhmNGs3dm41N2Y3ank4d2c4bWo5N2o1aHU=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NDQ4MzI4MTc5LjI1MDAwMDAwMDAwMDAwMDAwMGJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxN3h6ZDloNDAwN2o0djhmNGs3dm41N2Y3ank4d2c4bWo5N2o1aHU=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MjgzOTQxMTgwLjE5MTY2NjY2NjM4MjcyNTQ4NmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxMHdycjIwNmQ2eWRsZW13bnJrcHJzeWphejZkemZqczYwZWV3N3A=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MjgzOTQxMTgwMS45MTY2NjY2NjM4MjcyNTQ4NjRiYXNldGNybw==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxMHdycjIwNmQ2eWRsZW13bnJrcHJzeWphejZkemZqczYwZWV3N3A=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MjgzOTQxMTgwLjE5MTY2NjY2NjM4MjcyNTQ4NmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxN3h6ZDloNDAwN2o0djhmNGs3dm41N2Y3ank4d2c4bWo5N2o1aHU=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MjgzOTQxMTgwMS45MTY2NjY2NjM4MjcyNTQ4NjRiYXNldGNybw==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxN3h6ZDloNDAwN2o0djhmNGs3dm41N2Y3ank4d2c4bWo5N2o1aHU=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MjgzOTQxMTgwLjE5MTY2NjY2NjM4MjcyNTQ4NmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxcmdtcXNjbnpzM2FlNjM2OHI1ZXljejJtNnQ4MmhnYWZmaDQzcXQ=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MjgzOTQxMTgwMS45MTY2NjY2NjM4MjcyNTQ4NjRiYXNldGNybw==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxcmdtcXNjbnpzM2FlNjM2OHI1ZXljejJtNnQ4MmhnYWZmaDQzcXQ=",
            "index": true
          }
        ]
      }
    ],
    "end_block_events": null,
    "validator_updates": null,
    "consensus_param_updates": {
      "block": {
        "max_bytes": "500000",
        "max_gas": "81500000"
      },
      "evidence": {
        "max_age_num_blocks": "403200",
        "max_age_duration": "2419200000000000",
        "max_bytes": "150000"
      },
      "validator": {
        "pub_key_types": [
          "ed25519"
        ]
      }
    }
  }
}
`

const TX_CHAINMAIN_MSG_REGISTER_ACCOUNT_TXS_RESP = `{
  "tx": {
    "body": {
      "messages": [
        {
          "@type": "/chainmain.icaauth.v1.MsgRegisterAccount",
          "owner": "tcro1np7ztcfeycqwhj0nr8hxfu0lfjz27telqx53ra",
          "connection_id": "connection-18",
          "version": ""
        }
      ],
      "memo": "",
      "timeout_height": "0",
      "extension_options": [
      ],
      "non_critical_extension_options": [
      ]
    },
    "auth_info": {
      "signer_infos": [
        {
          "public_key": {
            "@type": "/cosmos.crypto.secp256k1.PubKey",
            "key": "A6cxzQUrKMhYYEreEPkmIzIPQvtosIvyDvMJdGc03nHU"
          },
          "mode_info": {
            "single": {
              "mode": "SIGN_MODE_DIRECT"
            }
          },
          "sequence": "171"
        }
      ],
      "fee": {
        "amount": [
          {
            "denom": "basetcro",
            "amount": "5000"
          }
        ],
        "gas_limit": "159998",
        "payer": "",
        "granter": ""
      },
      "tip": null
    },
    "signatures": [
      "1xZfFFDtIYo1d6qsEcGJbMnVG0PLl4k4UU1/15ro+XdTSN1Ye2Jgu8NJdUWzWrCL87kCG9TtaiEovv+cd14sgA=="
    ]
  },
  "tx_response": {
    "height": "67461",
    "txhash": "5FC4ECBEB0FFD3809F62052CADA161E542BAA1198EE94CCD0DF197EE4A885F35",
    "codespace": "",
    "code": 0,
    "data": "12320A302F636861696E6D61696E2E696361617574682E76312E4D736752656769737465724163636F756E74526573706F6E7365",
    "raw_log": "[{\"msg_index\":0,\"events\":[{\"type\":\"channel_open_init\",\"attributes\":[{\"key\":\"port_id\",\"value\":\"icacontroller-tcro1np7ztcfeycqwhj0nr8hxfu0lfjz27telqx53ra\"},{\"key\":\"channel_id\",\"value\":\"channel-48\"},{\"key\":\"counterparty_port_id\",\"value\":\"icahost\"},{\"key\":\"counterparty_channel_id\"},{\"key\":\"connection_id\",\"value\":\"connection-18\"},{\"key\":\"version\",\"value\":\"{\\\"fee_version\\\":\\\"ics29-1\\\",\\\"app_version\\\":\\\"{\\\\\\\"version\\\\\\\":\\\\\\\"ics27-1\\\\\\\",\\\\\\\"controller_connection_id\\\\\\\":\\\\\\\"connection-18\\\\\\\",\\\\\\\"host_connection_id\\\\\\\":\\\\\\\"connection-1\\\\\\\",\\\\\\\"address\\\\\\\":\\\\\\\"\\\\\\\",\\\\\\\"encoding\\\\\\\":\\\\\\\"proto3\\\\\\\",\\\\\\\"tx_type\\\\\\\":\\\\\\\"sdk_multi_msg\\\\\\\"}\\\"}\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"register_account\"},{\"key\":\"module\",\"value\":\"ibc_channel\"}]}]}]",
    "logs": [
      {
        "msg_index": 0,
        "log": "",
        "events": [
          {
            "type": "channel_open_init",
            "attributes": [
              {
                "key": "port_id",
                "value": "icacontroller-tcro1np7ztcfeycqwhj0nr8hxfu0lfjz27telqx53ra"
              },
              {
                "key": "channel_id",
                "value": "channel-48"
              },
              {
                "key": "counterparty_port_id",
                "value": "icahost"
              },
              {
                "key": "counterparty_channel_id",
                "value": ""
              },
              {
                "key": "connection_id",
                "value": "connection-18"
              },
              {
                "key": "version",
                "value": "{\"fee_version\":\"ics29-1\",\"app_version\":\"{\\\"version\\\":\\\"ics27-1\\\",\\\"controller_connection_id\\\":\\\"connection-18\\\",\\\"host_connection_id\\\":\\\"connection-1\\\",\\\"address\\\":\\\"\\\",\\\"encoding\\\":\\\"proto3\\\",\\\"tx_type\\\":\\\"sdk_multi_msg\\\"}\"}"
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "action",
                "value": "register_account"
              },
              {
                "key": "module",
                "value": "ibc_channel"
              }
            ]
          }
        ]
      }
    ],
    "info": "",
    "gas_wanted": "159998",
    "gas_used": "158041",
    "tx": {
      "@type": "/cosmos.tx.v1beta1.Tx",
      "body": {
        "messages": [
          {
            "@type": "/chainmain.icaauth.v1.MsgRegisterAccount",
            "owner": "tcro1np7ztcfeycqwhj0nr8hxfu0lfjz27telqx53ra",
            "connectionId": "connection-18",
            "version": ""
          }
        ],
        "memo": "",
        "timeout_height": "0",
        "extension_options": [
        ],
        "non_critical_extension_options": [
        ]
      },
      "auth_info": {
        "signer_infos": [
          {
            "public_key": {
              "@type": "/cosmos.crypto.secp256k1.PubKey",
              "key": "A6cxzQUrKMhYYEreEPkmIzIPQvtosIvyDvMJdGc03nHU"
            },
            "mode_info": {
              "single": {
                "mode": "SIGN_MODE_DIRECT"
              }
            },
            "sequence": "171"
          }
        ],
        "fee": {
          "amount": [
            {
              "denom": "basetcro",
              "amount": "5000"
            }
          ],
          "gas_limit": "159998",
          "payer": "",
          "granter": ""
        },
        "tip": null
      },
      "signatures": [
        "1xZfFFDtIYo1d6qsEcGJbMnVG0PLl4k4UU1/15ro+XdTSN1Ye2Jgu8NJdUWzWrCL87kCG9TtaiEovv+cd14sgA=="
      ]
    },
    "timestamp": "2022-11-18T09:45:50Z",
    "events": [
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "dGNybzFucDd6dGNmZXljcXdoajBucjhoeGZ1MGxmanoyN3RlbHF4NTNyYQ==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "NTAwMGJhc2V0Y3Jv",
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "dGNybzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHhoemFoYQ==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "NTAwMGJhc2V0Y3Jv",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "dGNybzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHhoemFoYQ==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "dGNybzFucDd6dGNmZXljcXdoajBucjhoeGZ1MGxmanoyN3RlbHF4NTNyYQ==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "NTAwMGJhc2V0Y3Jv",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "dGNybzFucDd6dGNmZXljcXdoajBucjhoeGZ1MGxmanoyN3RlbHF4NTNyYQ==",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "ZmVl",
            "value": "NTAwMGJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "ZmVlX3BheWVy",
            "value": "dGNybzFucDd6dGNmZXljcXdoajBucjhoeGZ1MGxmanoyN3RlbHF4NTNyYQ==",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "YWNjX3NlcQ==",
            "value": "dGNybzFucDd6dGNmZXljcXdoajBucjhoeGZ1MGxmanoyN3RlbHF4NTNyYS8xNzE=",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "c2lnbmF0dXJl",
            "value": "MXhaZkZGRHRJWW8xZDZxc0VjR0piTW5WRzBQTGw0azRVVTEvMTVybytYZFRTTjFZZTJKZ3U4TkpkVVd6V3JDTDg3a0NHOVR0YWlFb3Z2K2NkMTRzZ0E9PQ==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "YWN0aW9u",
            "value": "cmVnaXN0ZXJfYWNjb3VudA==",
            "index": true
          }
        ]
      },
      {
        "type": "channel_open_init",
        "attributes": [
          {
            "key": "cG9ydF9pZA==",
            "value": "aWNhY29udHJvbGxlci10Y3JvMW5wN3p0Y2ZleWNxd2hqMG5yOGh4ZnUwbGZqejI3dGVscXg1M3Jh",
            "index": true
          },
          {
            "key": "Y2hhbm5lbF9pZA==",
            "value": "Y2hhbm5lbC00OA==",
            "index": true
          },
          {
            "key": "Y291bnRlcnBhcnR5X3BvcnRfaWQ=",
            "value": "aWNhaG9zdA==",
            "index": true
          },
          {
            "key": "Y291bnRlcnBhcnR5X2NoYW5uZWxfaWQ=",
            "value": null,
            "index": true
          },
          {
            "key": "Y29ubmVjdGlvbl9pZA==",
            "value": "Y29ubmVjdGlvbi0xOA==",
            "index": true
          },
          {
            "key": "dmVyc2lvbg==",
            "value": "eyJmZWVfdmVyc2lvbiI6ImljczI5LTEiLCJhcHBfdmVyc2lvbiI6IntcInZlcnNpb25cIjpcImljczI3LTFcIixcImNvbnRyb2xsZXJfY29ubmVjdGlvbl9pZFwiOlwiY29ubmVjdGlvbi0xOFwiLFwiaG9zdF9jb25uZWN0aW9uX2lkXCI6XCJjb25uZWN0aW9uLTFcIixcImFkZHJlc3NcIjpcIlwiLFwiZW5jb2RpbmdcIjpcInByb3RvM1wiLFwidHhfdHlwZVwiOlwic2RrX211bHRpX21zZ1wifSJ9",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "bW9kdWxl",
            "value": "aWJjX2NoYW5uZWw=",
            "index": true
          }
        ]
      }
    ]
  }
}`
