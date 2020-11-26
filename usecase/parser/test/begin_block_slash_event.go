package usecase_parser_test

var BEGIN_BLOCK_SLASH_EVENT_BLOCK_RESULTS_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "28",
    "txs_results": null,
    "begin_block_events": [
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxNTQ4ZjVoeWRkZGcwZWE0c2RneHNlN3Q3ajRqbjg0enA2eTk1NGU=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "OQ==",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "Mjg=",
            "index": true
          }
        ]
      },
      {
        "type": "slash",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxNTQ4ZjVoeWRkZGcwZWE0c2RneHNlN3Q3ajRqbjg0enA2eTk1NGU=",
            "index": true
          },
          {
            "key": "cG93ZXI=",
            "value": "MzEy",
            "index": true
          },
          {
            "key": "cmVhc29u",
            "value": "bWlzc2luZ19zaWduYXR1cmU=",
            "index": true
          },
          {
            "key": "amFpbGVk",
            "value": "Y3JvY25jbGNvbnMxNTQ4ZjVoeWRkZGcwZWE0c2RneHNlN3Q3ajRqbjg0enA2eTk1NGU=",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxbmZ0ZzJuOWd6anIybDdsZW1jc2hrMHY4d2RtdXV6cThudWxudnM=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "OQ==",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "Mjg=",
            "index": true
          }
        ]
      },
      {
        "type": "slash",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxbmZ0ZzJuOWd6anIybDdsZW1jc2hrMHY4d2RtdXV6cThudWxudnM=",
            "index": true
          },
          {
            "key": "cG93ZXI=",
            "value": "MjEy",
            "index": true
          },
          {
            "key": "cmVhc29u",
            "value": "bWlzc2luZ19zaWduYXR1cmU=",
            "index": true
          },
          {
            "key": "amFpbGVk",
            "value": "Y3JvY25jbGNvbnMxbmZ0ZzJuOWd6anIybDdsZW1jc2hrMHY4d2RtdXV6cThudWxudnM=",
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
            "value": "NTE5NDUxMzU2YmFzZWNybw==",
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
              "ed25519": "jZmiyA+S/yVqVuN2Px/9OqB/xgMPaj4mPdHpUOg/Kj0="
            }
          }
        }
      },
      {
        "pub_key": {
          "Sum": {
            "type": "tendermint.crypto.PublicKey_Ed25519",
            "value": {
              "ed25519": "LNa+qkaUeJ97z/uLAKv1YTLMspaGxSkQyipkAmtwivo="
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
