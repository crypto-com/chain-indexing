package usecase_parser_test

const TX_MSG_TIER_BLOCK_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "block_id": {
      "hash": "A1B2C3D4E5F6A1B2C3D4E5F6A1B2C3D4E5F6A1B2C3D4E5F6A1B2C3D4E5F6A1B2",
      "parts": {
        "total": 1,
        "hash": "B2C3D4E5F6A1B2C3D4E5F6A1B2C3D4E5F6A1B2C3D4E5F6A1B2C3D4E5F6A1B2C3"
      }
    },
    "block": {
      "header": {
        "version": {
          "block": "11"
        },
        "chain_id": "crypto-org-chain-mainnet-1",
        "height": "100000",
        "time": "2024-06-15T10:00:00.000000Z",
        "last_block_id": {
          "hash": "C3D4E5F6A1B2C3D4E5F6A1B2C3D4E5F6A1B2C3D4E5F6A1B2C3D4E5F6A1B2C3D4",
          "parts": {
            "total": 1,
            "hash": "D4E5F6A1B2C3D4E5F6A1B2C3D4E5F6A1B2C3D4E5F6A1B2C3D4E5F6A1B2C3D4E5"
          }
        },
        "last_commit_hash": "E5F6A1B2C3D4E5F6A1B2C3D4E5F6A1B2C3D4E5F6A1B2C3D4E5F6A1B2C3D4E5F6",
        "data_hash": "F6A1B2C3D4E5F6A1B2C3D4E5F6A1B2C3D4E5F6A1B2C3D4E5F6A1B2C3D4E5F6A1",
        "validators_hash": "5F3936EBC0AC11C73A0A361ACA707A08B7F1229C57371D48D4E515C135237C9D",
        "next_validators_hash": "5F3936EBC0AC11C73A0A361ACA707A08B7F1229C57371D48D4E515C135237C9D",
        "consensus_hash": "252FE7CF36DD1BB85DAFC47A08961DF0CFD8C027DEFA5E01E958BE121599DB9D",
        "app_hash": "9794111060A1EB15780166CB2ED8565B7D371FE2624162A14926630FB6C7722E",
        "last_results_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "proposer_address": "521A28F60F5B1D0879260EB7E0CF775FFF0C9A57"
      },
      "data": {
        "txs": [
          "dGVzdA=="
        ]
      },
      "evidence": {
        "evidence": []
      },
      "last_commit": {
        "height": "99999",
        "round": 0,
        "block_id": {
          "hash": "C3D4E5F6A1B2C3D4E5F6A1B2C3D4E5F6A1B2C3D4E5F6A1B2C3D4E5F6A1B2C3D4",
          "parts": {
            "total": 1,
            "hash": "D4E5F6A1B2C3D4E5F6A1B2C3D4E5F6A1B2C3D4E5F6A1B2C3D4E5F6A1B2C3D4E5"
          }
        },
        "signatures": [
          {
            "block_id_flag": 2,
            "validator_address": "521A28F60F5B1D0879260EB7E0CF775FFF0C9A57",
            "timestamp": "2024-06-15T10:00:00.000000Z",
            "signature": "dGVzdHNpZw=="
          }
        ]
      }
    }
  }
}`

const TX_MSG_TIER_BLOCK_RESULTS_RESP = `{
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
