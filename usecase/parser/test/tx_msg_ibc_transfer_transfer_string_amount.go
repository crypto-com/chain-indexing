package usecase_parser_test

const TX_MSG_TRANSFER_STRING_AMOUNT_BLOCK_RESP = `
{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "block_id": {
      "hash": "73E6D0F8C45BC19281F7B6EB03A0C002B7632608411D4BE9020DCF3BE93CD52B",
      "parts": {
        "total": 1,
        "hash": "23B2107DF8D2B288BE796C09120B196F56F4453877C44D73AA317DA294F13A13"
      }
    },
    "block": {
      "header": {
        "version": {
          "block": "11"
        },
        "chain_id": "cronostestnet_338-3",
        "height": "140108",
        "time": "2021-10-22T09:52:49.796664095Z",
        "last_block_id": {
          "hash": "6C11D7F1C9C1EC1B3EB2A6658D6A3BE674589103F4183FCE83F7BAF478678FC3",
          "parts": {
            "total": 1,
            "hash": "914CC0A465D36B98787A9F8ED2A3B0E266692FF409D796DF44474BCCADB91F0D"
          }
        },
        "last_commit_hash": "05C9130A63DB4432DBD0CF5624EC5860C4F96525583F5782E636BF339F4F2B97",
        "data_hash": "E5B4A3528B4C71801A4AAC40EC56A90D8EE3C09E7321BC32E0B92E08138F76EC",
        "validators_hash": "415F9F8A48C405E536379D8C327C2176EA651AD1D06BC3B525D508F51645255D",
        "next_validators_hash": "415F9F8A48C405E536379D8C327C2176EA651AD1D06BC3B525D508F51645255D",
        "consensus_hash": "252FE7CF36DD1BB85DAFC47A08961DF0CFD8C027DEFA5E01E958BE121599DB9D",
        "app_hash": "E51E4EF0216A75CE87EA2AA1825A58BEE95EADC817776A230D4B6C91EC5290FF",
        "last_results_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "proposer_address": "A3F7397F81D4CC82D566633F9728DFB50A35D965"
      },
      "data": {
        "txs": [
          "CtQBCtEBCikvaWJjLmFwcGxpY2F0aW9ucy50cmFuc2Zlci52MS5Nc2dUcmFuc2ZlchKjAQoIdHJhbnNmZXISCWNoYW5uZWwtMBogCghiYXNldGNybxIUMjg4MzY4MzYyMzY4MjgzOTg5MDAiK3RjcmMxdjc2cjd1NHV5cjNld2RrczhjcW11dzdjYTRsZWp2Yzg5cHhoZXYqK3Rjcm8xNTU4cmZsOHBucGs5MHJtcnl2eTNlcDAzeXNscTRhcjBsa3Zkbm4yBggEELDTPzib/tWvtoGU2BYSggEKWQpPCigvZXRoZXJtaW50LmNyeXB0by52MS5ldGhzZWNwMjU2azEuUHViS2V5EiMKIQNSRQMj3TAYrtQpkK6s3oV6E8+j7H14YvoDU1ufmGR6XxIECgIIARgFEiUKHwoIYmFzZXRjcm8SEzEwMDAwMDAwMDAwMDAwMDAwMDAQwJoMGkE2SUpCNobSr65c2e1CKOQzCFW9vnJHJAs+dMRzTFSVdx1aTBD7+0tZBBtMxHqzJG0uWKvDOXebgs9ESzSefuj6AA=="
        ]
      },
      "evidence": {
        "evidence": []
      },
      "last_commit": {
        "height": "140107",
        "round": 0,
        "block_id": {
          "hash": "6C11D7F1C9C1EC1B3EB2A6658D6A3BE674589103F4183FCE83F7BAF478678FC3",
          "parts": {
            "total": 1,
            "hash": "914CC0A465D36B98787A9F8ED2A3B0E266692FF409D796DF44474BCCADB91F0D"
          }
        },
        "signatures": [
          {
            "block_id_flag": 2,
            "validator_address": "0421821D46C46F86F1EAD79EEB31CFA88A5578CB",
            "timestamp": "2021-10-22T09:52:49.796664095Z",
            "signature": "tYNjx6qDqBaWDfwwtYIhnBTfejRCrqJZCZ2X781OJhawf2GNivbp1qoEtQoYVuFsf0RWPKmfezulizmoeRH+Bg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "A3F7397F81D4CC82D566633F9728DFB50A35D965",
            "timestamp": "2021-10-22T09:52:49.664512645Z",
            "signature": "6OjYS6xkRixaUidKAYZCDp7W+h9enqeQbbWuKM9245qGeSk1B7w8Rlk3arm9mRmT0IVtJkscOb+9QIsuA7P1CQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "1473CC22B535A206B956CB47C65C272EE5324927",
            "timestamp": "2021-10-22T09:52:49.991604313Z",
            "signature": "d/Xzrkshs9eSEZ1xWt/rCxUGd+zC0De0po33N5EG7VibR3uoK0cePjOo0GQ8vBpbeKbO8WRpGjiQv1NefSTNCw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "1E036C0FE3472F33A9FFC9F3ED14B56DE85BE389",
            "timestamp": "2021-10-22T09:52:49.988441807Z",
            "signature": "lylJlKnEsN8ntdoFDxhmpTJjl5mq4xtFL4cTGMiRZ5gl9zbfWnPItMgMiw0tC/9B2HEQA2kKf2oB8o9/fo6mCg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "258CEFCAD9A77B8CEB7F2BAAE9A44B4EC2A9B5BF",
            "timestamp": "2021-10-22T09:52:50.123182527Z",
            "signature": "p33e7VJJrVpFUd6NM+rPSqLr+HCtBi9srnEzRXQI74unSnRrTlp7mObQ1MffimDrFJrRREZ/3Pih5N2YS1l0BQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "559AB20018F2E53A89533C390DCA2CEA6721D869",
            "timestamp": "2021-10-22T09:52:50.049928334Z",
            "signature": "0HUyjuSy10Ys1AidL9tQjeOmiWZT6//dgZXePO1S/rZukLoSXx6XvZvqEtJtT8nhFDXjdWqL+TUSrl/lWpe2Bg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "5600A9A2FF66520305EC37F1F698FEE236FACF98",
            "timestamp": "2021-10-22T09:52:50.026410974Z",
            "signature": "A44P6yFSpWdteO15GEznOkZWPoQPDdeoDDQ43xpSIEPHIEoCTwzWWdZjSxkww+EyIjr7NRNImGX1q0u1kJbxDw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "72A52EBF69F6670C05A3EB2BA831ADBF629B8948",
            "timestamp": "2021-10-22T09:52:50.070870959Z",
            "signature": "38twPd/vOnkUqVYLLKXHT8yeQ0v6VFstSEoU2pFv5bsJyLEKpPWnJ4nBMccw3nxWZUt+eY9Z/HNECT9LdHh6AQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "739109BF2993601BEF624283C9A572DF3175CBA8",
            "timestamp": "2021-10-22T09:52:50.003908949Z",
            "signature": "GvQUBghxqs/n0MYup2vRS3zqQHJiMXZhJaEfP9VEzF3iZsO4yY/QydtUTqOOkl5SSaM4LjDwlMqiRaRiSJpXBg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "A6C944F5ECC995C358572B7FF1B769B6DB60000D",
            "timestamp": "2021-10-22T09:52:50.041091345Z",
            "signature": "06JhUmOLAWwu0uV3EWoNmno22CO+jLAcWGrW47IqaKDmCj1WFdftOhk+m8s+qs12cQEF6NjbNZIP9PwS2VDqAg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "A72AA9230A83D4D6D18AD07EC03B5168DEA702AE",
            "timestamp": "2021-10-22T09:52:50.041066368Z",
            "signature": "8K7n1/FKbDhIo1DX3GremVAL0ZcFLN5sd7GcxqGMRlCEAugpKriZgjfgzRSBm7ORQrJ12L+KAV3d76RvbJ0QDA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "A7E012CB6FFABA17C6C98600F3F7402DC1A06C8C",
            "timestamp": "2021-10-22T09:52:50.036825393Z",
            "signature": "t9UE6BsofMidbPVDUHIsBSSWNGkSYisaw/TIZPg1jAq74v5bx30I95PERZe3ZrxXS/4tHaNLviQIk5jh3veSAg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "AD582CC7756A0C2F3803840997BAF10075D6D005",
            "timestamp": "2021-10-22T09:52:50.032562981Z",
            "signature": "5fS2Ph8uJQ7sH3O5nZsyCOXkXfpPH/gBTSYraqUn/HX9SwGW/SqDl1qdMM1GjTYn0yMjinjqlipxEHwVlLCrAA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "B6581676389C2D34EC256DA777F3D9D743AD10AB",
            "timestamp": "2021-10-22T09:52:50.166166888Z",
            "signature": "I2F8kF1hDN3B1ysxYamKbJJA+kT5mjSwRlvlA1yuFoHQKXaBrJIXAdU2fmmqW+OS3dXBZMqMGQ5NDIj9GfFHDQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "B8735B4BBA3AD413B7DBB3A74ABC75A5907BE7A5",
            "timestamp": "2021-10-22T09:52:50.073235918Z",
            "signature": "OMKnQ17kG+cfPDIN8mLYwl3L+ug0rJpy2dDeR5pcc9AgmYlaB7HEeTBY8RwbEiJ9/5r3QbkC82RHomSf7Kc2DQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "BE937B566A3EBBD72E6553F3D216E2869D7FD1DD",
            "timestamp": "2021-10-22T09:52:50.172061485Z",
            "signature": "DDJwvT9Dguu19AcxKvP5YBQj/W95cw/T1yPHg6JGhSdgyNHeSrZjNoiFZ7jQdaKhE18YlZC25B5jP7JW5dWHAg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "D03C1628B633072EF67E42141110CCD6AA420456",
            "timestamp": "2021-10-22T09:52:49.980316908Z",
            "signature": "dvNHPU0FIyIssjvSJVjDblVcozMYWjGVpT11AM2dsDhMPJicqT6EQ1J+mT+s16xhpTltSgWaAgeqQATDGLv6Cg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "09D283BB0AC4B6A8BA05F6600E018E1D4DD25C12",
            "timestamp": "2021-10-22T09:52:50.0809697Z",
            "signature": "W13UdcIH4sV9mUo2cITw8QNOijSYtwiiTBLK9Fbgw0ujlVWU+bIphbjVOxWHlV35cm5bPR4mKHwUA3tB7+UUDQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "B359B56836AF6117E780045985D45F988065B746",
            "timestamp": "2021-10-22T09:52:49.9789385Z",
            "signature": "11PIwHP7stgO50swZuQ7xgeVKSJqr/QCT3FtYu0BdlLUE3Q1UT9mOR48SzYWWZi8UPRofP/6HCx3F2BsQzqJDg=="
          }
        ]
      }
    }
  }
}
`

const TX_MSG_TRANSFER_STRING_AMOUNT_BLOCK_RESULTS_RESP = `
{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "140108",
    "txs_results": [
      {
        "code": 111222,
        "data": null,
        "log": "panic message redacted to hide potentially sensitive system info: panic",
        "info": "",
        "gas_wanted": "200000",
        "gas_used": "100212",
        "events": [],
        "codespace": "undefined"
      }
    ],
    "begin_block_events": [
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "dGNyYzFtM2gzMHdsdnNmOGxscnV4dHB1a2R2c3kwa20ya3VtODM2NTI0MA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "dGNyYzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bGZqc2plag==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "dGNyYzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bGZqc2plag==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "dGNyYzFtM2gzMHdsdnNmOGxscnV4dHB1a2R2c3kwa20ya3VtODM2NTI0MA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "dGNyYzFtM2gzMHdsdnNmOGxscnV4dHB1a2R2c3kwa20ya3VtODM2NTI0MA==",
            "index": true
          }
        ]
      },
      {
        "type": "mint",
        "attributes": [
          {
            "key": "Ym9uZGVkX3JhdGlv",
            "value": "MC4wMDAwMDAwMDAxMDc3OTk5OTg=",
            "index": true
          },
          {
            "key": "aW5mbGF0aW9u",
            "value": "MC4wMDAwMDAwMDAwMDAwMDAwMDA=",
            "index": true
          },
          {
            "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
            "value": "MC4wMDAwMDAwMDAwMDAwMDAwMDA=",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MA==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "dGNyYzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bGZqc2plag==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "dGNyYzFqdjY1czNncnFmNnY2amwzZHA0dDZjOXQ5cms5OWNkODc1aHdtcw==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "dGNyYzFqdjY1czNncnFmNnY2amwzZHA0dDZjOXQ5cms5OWNkODc1aHdtcw==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "dGNyYzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bGZqc2plag==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "dGNyYzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bGZqc2plag==",
            "index": true
          }
        ]
      },
      {
        "type": "proposer_reward",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxdjc2cjd1NHV5cjNld2RrczhjcW11dzdjYTRsZWp2Yzh1bmFnMm0=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxdjc2cjd1NHV5cjNld2RrczhjcW11dzdjYTRsZWp2Yzh1bmFnMm0=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxdjc2cjd1NHV5cjNld2RrczhjcW11dzdjYTRsZWp2Yzh1bmFnMm0=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxdjc2cjd1NHV5cjNld2RrczhjcW11dzdjYTRsZWp2Yzh1bmFnMm0=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxdjc2cjd1NHV5cjNld2RrczhjcW11dzdjYTRsZWp2Yzh1bmFnMm0=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxczllM2Y1eTR0dDRma3ozanlqODY1cWF1ZDJjcWhzNjZ5bm1qNGY=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxczllM2Y1eTR0dDRma3ozanlqODY1cWF1ZDJjcWhzNjZ5bm1qNGY=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxc2h5OHY0ZnQ3OG16dnNzZDkwYWw1Y2x6MzNydTZhZ3JsYW04Z2U=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxc2h5OHY0ZnQ3OG16dnNzZDkwYWw1Y2x6MzNydTZhZ3JsYW04Z2U=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxMDhnODBoZmcwcWt6d3U1bm4wN3Bkcnhqc2R4ZmthY2g5c2N0cjQ=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxMDhnODBoZmcwcWt6d3U1bm4wN3Bkcnhqc2R4ZmthY2g5c2N0cjQ=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxNnRueDY4MnBqYXIycnE4MjQ5OHJwemZ1cjhxNWg3bXdweXZnc3U=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxNnRueDY4MnBqYXIycnE4MjQ5OHJwemZ1cjhxNWg3bXdweXZnc3U=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxbGxkamhqbm4zMmU4dmVrN2N4ZTlnMDVuZjhqNzR5MHhsejVnMjM=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxbGxkamhqbm4zMmU4dmVrN2N4ZTlnMDVuZjhqNzR5MHhsejVnMjM=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxbXRjbjI1MDVrMzdtbHp0eXdmOGVnOHNwdjBrcG5zcWFwYWZ6c3o=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxbXRjbjI1MDVrMzdtbHp0eXdmOGVnOHNwdjBrcG5zcWFwYWZ6c3o=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxejI3bG4zM3FsZG45NnVhMmxkZDRkdnk3cDhlMjlwN3dsbXhzdGw=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxejI3bG4zM3FsZG45NnVhMmxkZDRkdnk3cDhlMjlwN3dsbXhzdGw=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxNXE4Zmx2bTVxZnp0dDV6dHY5bXkyaHQyeGwyeGFra3g5c2F4eDc=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxNXE4Zmx2bTVxZnp0dDV6dHY5bXkyaHQyeGwyeGFra3g5c2F4eDc=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxdW1zaGY4dDZjY3NkZHpua2cwOXNheHlycDIzNXF6c2Z2MDBlNWs=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxdW1zaGY4dDZjY3NkZHpua2cwOXNheHlycDIzNXF6c2Z2MDBlNWs=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxM2F3dm0zNnVsZTl0NXNkank4YXZ5cndkd2t2bHZtaDNzcHE2eDU=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxM2F3dm0zNnVsZTl0NXNkank4YXZ5cndkd2t2bHZtaDNzcHE2eDU=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxcTBnY3EwMjVqeHM4ZnJzNnZhZGc4cnphNjUwNzI0MzU2eWFzYXU=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxcTBnY3EwMjVqeHM4ZnJzNnZhZGc4cnphNjUwNzI0MzU2eWFzYXU=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxM2s2ejYyY2c2MHhmcW10MGFqbGU2dTRmemc0dWU0MmZncWM3bno=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxM2s2ejYyY2c2MHhmcW10MGFqbGU2dTRmemc0dWU0MmZncWM3bno=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxeHlodjJnYzZ1dnR1am0zNWN0eTh0Z2czeDBudjdyanMwZWd3NG4=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxeHlodjJnYzZ1dnR1am0zNWN0eTh0Z2czeDBudjdyanMwZWd3NG4=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxa2Q3bGRoZHh6cjM4ZWt5bnFranJwM3A5bHhwZGN6eHhzeWFqYXE=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxa2Q3bGRoZHh6cjM4ZWt5bnFranJwM3A5bHhwZGN6eHhzeWFqYXE=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxZ3l3eGZkcjRuZXRzODdjenhscGNhenBsYTVxY2tlZmpkczNjeHU=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxZ3l3eGZkcjRuZXRzODdjenhscGNhenBsYTVxY2tlZmpkczNjeHU=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxNXV0dTJjMjN2Mmg0d3RkZXV3bXI0eXZjaHFseGphZnY4dTgzcnI=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxNXV0dTJjMjN2Mmg0d3RkZXV3bXI0eXZjaHFseGphZnY4dTgzcnI=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxaDBqY2cwM3hqZnE0bjl3NmwybHh1dHA4bDU1bm1jejl3enkwMHY=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxaDBqY2cwM3hqZnE0bjl3NmwybHh1dHA4bDU1bm1jejl3enkwMHY=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxOWE2cjc0ZHZmeGp5dmp6ZjNwZzl5M3k1cmhrNnJkczJwaDM5OHk=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": null,
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyY3ZhbG9wZXIxOWE2cjc0ZHZmeGp5dmp6ZjNwZzl5M3k1cmhrNnJkczJwaDM5OHk=",
            "index": true
          }
        ]
      }
    ],
    "end_block_events": [
      {
        "type": "block_bloom",
        "attributes": [
          {
            "key": "Ymxvb20=",
            "value": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==",
            "index": true
          }
        ]
      }
    ],
    "validator_updates": null,
    "consensus_param_updates": {
      "block": {
        "max_bytes": "1048576",
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

const TX_MSG_TRANSFER_STRING_AMOUNT_TXS_RESP = `{
  "tx": {
    "body": {
      "messages": [
        {
          "@type": "/ibc.applications.transfer.v1.MsgTransfer",
          "source_port": "transfer",
          "source_channel": "channel-0",
          "token": {
            "denom": "basetcro",
            "amount": "28836836236828398900"
          },
          "sender": "tcrc1v76r7u4uyr3ewdks8cqmuw7ca4lejvc89pxhev",
          "receiver": "tcro1558rfl8pnpk90rmryvy3ep03yslq4ar0lkvdnn",
          "timeout_height": {
            "revision_number": "4",
            "revision_height": "1042864"
          },
          "timeout_timestamp": "1634894674620940059"
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
            "@type": "/ethermint.crypto.v1.ethsecp256k1.PubKey",
            "key": "A1JFAyPdMBiu1CmQrqzehXoTz6PsfXhi+gNTW5+YZHpf"
          },
          "mode_info": {
            "single": {
              "mode": "SIGN_MODE_DIRECT"
            }
          },
          "sequence": "5"
        }
      ],
      "fee": {
        "amount": [
          {
            "denom": "basetcro",
            "amount": "1000000000000000000"
          }
        ],
        "gas_limit": "200000",
        "payer": "",
        "granter": ""
      }
    },
    "signatures": [
      "NklKQjaG0q+uXNntQijkMwhVvb5yRyQLPnTEc0xUlXcdWkwQ+/tLWQQbTMR6syRtLlirwzl3m4LPREs0nn7o+gA="
    ]
  },
  "tx_response": {
    "height": "140108",
    "txhash": "D924F6E1A16ACDFFBF0B5BFDECC8E010E8F8D746B379FFC63D477C472B4128B7",
    "codespace": "undefined",
    "code": 111222,
    "data": "",
    "raw_log": "panic message redacted to hide potentially sensitive system info: panic",
    "logs": [
    ],
    "info": "",
    "gas_wanted": "200000",
    "gas_used": "100212",
    "tx": {
      "@type": "/cosmos.tx.v1beta1.Tx",
      "body": {
        "messages": [
          {
            "@type": "/ibc.applications.transfer.v1.MsgTransfer",
            "source_port": "transfer",
            "source_channel": "channel-0",
            "token": {
              "denom": "basetcro",
              "amount": "28836836236828398900"
            },
            "sender": "tcrc1v76r7u4uyr3ewdks8cqmuw7ca4lejvc89pxhev",
            "receiver": "tcro1558rfl8pnpk90rmryvy3ep03yslq4ar0lkvdnn",
            "timeout_height": {
              "revision_number": "4",
              "revision_height": "1042864"
            },
            "timeout_timestamp": "1634894674620940059"
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
              "@type": "/ethermint.crypto.v1.ethsecp256k1.PubKey",
              "key": "A1JFAyPdMBiu1CmQrqzehXoTz6PsfXhi+gNTW5+YZHpf"
            },
            "mode_info": {
              "single": {
                "mode": "SIGN_MODE_DIRECT"
              }
            },
            "sequence": "5"
          }
        ],
        "fee": {
          "amount": [
            {
              "denom": "basetcro",
              "amount": "1000000000000000000"
            }
          ],
          "gas_limit": "200000",
          "payer": "",
          "granter": ""
        }
      },
      "signatures": [
        "NklKQjaG0q+uXNntQijkMwhVvb5yRyQLPnTEc0xUlXcdWkwQ+/tLWQQbTMR6syRtLlirwzl3m4LPREs0nn7o+gA="
      ]
    },
    "timestamp": "2021-10-22T09:52:49Z",
    "events": [
    ]
  }
}`
