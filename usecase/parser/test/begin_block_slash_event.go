package usecase_parser_test

var BEGIN_BLOCK_SLASH_EVENT_BLOCK_RESULTS_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "156320",
    "txs_results": [],
    "begin_block_events": [
      {
        "type": "slash",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxNnNhOWNmbmV2bGwwcmVjd2E1aDdzZW1xZnB0emRxdXI3dnFybDQ=",
            "index": true
          },
          {
            "key": "cG93ZXI=",
            "value": "MTY1NDM3ODA=",
            "index": true
          },
          {
            "key": "cmVhc29u",
            "value": "ZG91YmxlX3NpZ24=",
            "index": true
          }
        ]
      },
      {
        "type": "slash",
        "attributes": [
          {
            "key": "amFpbGVk",
            "value": "Y3JvY25jbGNvbnMxNnNhOWNmbmV2bGwwcmVjd2E1aDdzZW1xZnB0emRxdXI3dnFybDQ=",
            "index": true
          }
        ]
      }
    ],
    "end_block_events": [
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "Y3JvMXR5Z21zM3hoaHMzeXY0ODdwaHgzZHc0YTk1am43dDdsZXFhOG5q",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "Y3JvMWZsNDh2c25tc2R6Y3Y4NXE1ZDJxNHo1YWpkaGE4eXUzZHFwazl4",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTU3MTY1OTEwNjIyMTFiYXNldGNybw==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "Y3JvMWZsNDh2c25tc2R6Y3Y4NXE1ZDJxNHo1YWpkaGE4eXUzZHFwazl4",
            "index": true
          }
        ]
      }
    ],
    "validator_updates": [
      {
        "pub_key": {
          "Sum": {
            "type": "tendermint.crypto.PublicKey_Ed25519",
            "value": {
              "ed25519": "6e3k6KCzT+VXMyM34i6LrKZT3YM3oByTsXj0OtLApg8="
            }
          }
        }
      }
    ],
    "consensus_param_updates": {
      "block": {
        "max_bytes": "22020096",
        "max_gas": "-1"
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
