package usecase_parser_test

const TX_MSG_TIER_CLAIM_TIER_REWARDS_TXS_RESP = `{
  "tx": {
    "body": {
      "messages": [
        {
          "@type": "/chainmain.tieredrewards.v1.MsgClaimTierRewards",
          "owner": "cro1dulwqgcdpemn8c34sjd92fxepz5p0sqpeevw84",
          "position_ids": ["1", "2", "3"]
        }
      ],
      "memo": "",
      "timeout_height": "0",
      "extension_options": [],
      "non_critical_extension_options": []
    },
    "auth_info": {
      "signer_infos": [],
      "fee": { "amount": [], "gas_limit": "200000", "payer": "", "granter": "" }
    },
    "signatures": []
  },
  "tx_response": {
    "height": "100000",
    "txhash": "D0E1F2A3B4C5D6E7F8A9B0C1D2E3F4A5B6C7D8E9F0A1B2C3D4E5F6A7B8C9D0E1",
    "codespace": "",
    "code": 0,
    "data": "",
    "raw_log": "[]",
    "logs": [],
    "info": "",
    "gas_wanted": "200000",
    "gas_used": "100000",
    "tx": {
      "@type": "/cosmos.tx.v1beta1.Tx",
      "body": {
        "messages": [
          {
            "@type": "/chainmain.tieredrewards.v1.MsgClaimTierRewards",
            "owner": "cro1dulwqgcdpemn8c34sjd92fxepz5p0sqpeevw84",
            "position_ids": ["1", "2", "3"]
          }
        ],
        "memo": "",
        "timeout_height": "0",
        "extension_options": [],
        "non_critical_extension_options": []
      },
      "auth_info": {
        "signer_infos": [],
        "fee": { "amount": [], "gas_limit": "200000", "payer": "", "granter": "" }
      },
      "signatures": []
    },
    "timestamp": "2024-06-15T10:00:00Z",
    "events": []
  }
}`

const TX_MSG_TIER_CLAIM_TIER_REWARDS_BLOCK_RESULTS_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "100000",
    "txs_results": [
      {
        "code": 0,
        "data": "",
        "log": "[]",
        "info": "",
        "gas_wanted": "200000",
        "gas_used": "100000",
        "events": [
          {
            "type": "message",
            "attributes": [
              {
                "key": "YWN0aW9u",
                "value": "dGllcg==",
                "index": false
              }
            ]
          },
          {
            "type": "chainmain.tieredrewards.v1.EventTierRewardsClaimed",
            "attributes": [
              {
                "key": "bXNnX2luZGV4",
                "value": "MA==",
                "index": false
              },
              {
                "key": "b3duZXI=",
                "value": "Y3JvMWR1bHdxZ2NkcGVtbjhjMzRzamQ5MmZ4ZXB6NXAwc3FwZWV2dzg0",
                "index": false
              },
              {
                "key": "cG9zaXRpb25faWRz",
                "value": "MSwyLDM=",
                "index": false
              },
              {
                "key": "YmFzZV9yZXdhcmRz",
                "value": "MTAwMGJhc2Vjcm8=",
                "index": false
              },
              {
                "key": "Ym9udXNfcmV3YXJkcw==",
                "value": "NTAwYmFzZWNybw==",
                "index": false
              }
            ]
          }
        ],
        "codespace": ""
      }
    ],
    "begin_block_events": [],
    "end_block_events": [],
    "validator_updates": null,
    "consensus_param_updates": {
      "block": {
        "max_bytes": "1048576",
        "max_gas": "81500000"
      },
      "evidence": {
        "max_age_num_blocks": "100000",
        "max_age_duration": "172800000000000",
        "max_bytes": "1048576"
      },
      "validator": {
        "pub_key_types": [
          "ed25519"
        ]
      }
    }
  }
}`
