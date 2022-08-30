package usecase_parser_test

const TX_MSG_REVOKE_BLOCK_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "block_id": {
      "hash": "7F4811841ADAE208C36D673D1AFFC06B2EC6D9E49B1C05842FDE0E634C4DE54F",
      "parts": {
        "total": 1,
        "hash": "0667AAE19BF56EB8C860FBFDD19548798B25022A26E2EB5A55CCBEF43B940A58"
      }
    },
    "block": {
      "header": {
        "version": {
          "block": "11"
        },
        "chain_id": "testnet-croeseid-4",
        "height": "123731",
        "time": "2021-08-26T20:13:47.674429514Z",
        "last_block_id": {
          "hash": "3D440C3AC4FC58D74EF5D01817045C09EECD54E0E69D7F5DD1F7102F9D6DBBEF",
          "parts": {
            "total": 1,
            "hash": "43556F6DCCB648C1F0EA11E93CA1F4ED633CF99288F01F296BE5500E67D75608"
          }
        },
        "last_commit_hash": "A7E1EC02E9BB3D5390849B7471D5D47A68C5334EC9ADD6C66E5BD8BAA14E857D",
        "data_hash": "666D2A40A64E09F41FB0E57BF9B4C737EC32322CEF6DFC83473469BDC6C698AE",
        "validators_hash": "FA93FE6B0DE3DD8DA45C80B4D3FB659A9EE0EA8BE2557656125DDB5FE3BF22E3",
        "next_validators_hash": "FA93FE6B0DE3DD8DA45C80B4D3FB659A9EE0EA8BE2557656125DDB5FE3BF22E3",
        "consensus_hash": "372B4AE845086C837EFEF79A189B085B1FD6610C53F3BEB17EE0E27B347C06DE",
        "app_hash": "EEE2D6238DBD8D2D82E3B1D052B676CF68B6383C781B816DAFCFE5C341E4A7EE",
        "last_results_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "proposer_address": "F9A4582896E5F2692BEDDD7C251755E33D2DE013"
      },
      "data": {
        "txs": [
          "Cp4BCpsBCh8vY29zbW9zLmF1dGh6LnYxYmV0YTEuTXNnUmV2b2tlEngKK3Rjcm8xdnVyZmhxZjBqMmpnZnBqYWhsamE2ZzZ1cTJ0czJyNjBzd20yZDkSK3Rjcm8xNXpoNXRuN3hqZGVjdTR6amNsc21sbmxodDVlYWQybXg4NGdhdTIaHC9jb3Ntb3MuYmFuay52MWJldGExLk1zZ1NlbmQSawpQCkYKHy9jb3Ntb3MuY3J5cHRvLnNlY3AyNTZrMS5QdWJLZXkSIwohAxcZ9W91aKz8X9eZOQTjv4JiFcLP3JjDIc/+go8FR13mEgQKAggBGAMSFwoRCghiYXNldGNybxIFNTAwMDAQwJoMGkBDVmq2Aw1EKkzB4AF8egZomqyZf/cMRch+y+Fkuwa+kTlxqq/j8wV/Pg6Hd6kpMGKETpOyy1X8iSzJJx7gLwad"
        ]
      },
      "evidence": {
        "evidence": []
      },
      "last_commit": {
        "height": "123730",
        "round": 0,
        "block_id": {
          "hash": "3D440C3AC4FC58D74EF5D01817045C09EECD54E0E69D7F5DD1F7102F9D6DBBEF",
          "parts": {
            "total": 1,
            "hash": "43556F6DCCB648C1F0EA11E93CA1F4ED633CF99288F01F296BE5500E67D75608"
          }
        },
        "signatures": [
          {
            "block_id_flag": 2,
            "validator_address": "F9A4582896E5F2692BEDDD7C251755E33D2DE013",
            "timestamp": "2021-08-26T20:13:47.679796094Z",
            "signature": "rAs1G2E4WJA6v+tUb0y7ice2nscVzzZcp9geey4Wf/OB8TPxY9mMzor6DQLxYUMyH0uWFMNDBkaOSLpbliJ8BQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "71DA178AE52B2D0654102A86166CE91AC5121CB2",
            "timestamp": "2021-08-26T20:13:47.667479975Z",
            "signature": "gv+D2hG7lbQgHOSxgbyQijIKdmGT28lkx4oC7+CdUPtxDpcmTc5XJuiQ4WNUsv1qMfT6bAMliEOjZxQ19bjlDQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "7E2B455523A35AE8F4064C00AF01515FEF673079",
            "timestamp": "2021-08-26T20:13:47.674429514Z",
            "signature": "gMRCyJZGM/wmIzdwXPD6rF6nR1dDlpUmy+hkYATQalLLE66sfGLtlSlxAb5UPLZTnvvd1YXs+r8tnFv/UK1FDA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "EB4717D3AD29E06B13E791B7224071E8367E5F0F",
            "timestamp": "2021-08-26T20:13:47.805376416Z",
            "signature": "SGP2bzBOPgo3Cpk9DwAtazgwG8g3y8/NaIjTzD6G3qjXaXtHx75/ShECYbiKMZud59laE4q724xRfDSQPvfYDQ=="
          }
        ]
      }
    }
  }
}`

const TX_MSG_REVOKE_BLOCK_RESULTS_RESP = `
{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "123731",
    "txs_results": [
      {
        "code": 0,
        "data": "CiEKHy9jb3Ntb3MuYXV0aHoudjFiZXRhMS5Nc2dSZXZva2U=",
        "log": "[{\"events\":[{\"type\":\"cosmos.authz.v1beta1.EventRevoke\",\"attributes\":[{\"key\":\"granter\",\"value\":\"\\\"tcro1vurfhqf0j2jgfpjahlja6g6uq2ts2r60swm2d9\\\"\"},{\"key\":\"grantee\",\"value\":\"\\\"tcro15zh5tn7xjdecu4zjclsmlnlht5ead2mx84gau2\\\"\"},{\"key\":\"msg_type_url\",\"value\":\"\\\"/cosmos.bank.v1beta1.MsgSend\\\"\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/cosmos.authz.v1beta1.MsgRevoke\"}]}]}]",
        "info": "",
        "gas_wanted": "200000",
        "gas_used": "52212",
        "events": [
          {
            "type": "coin_spent",
            "attributes": [
              {
                "key": "c3BlbmRlcg==",
                "value": "dGNybzF2dXJmaHFmMGoyamdmcGphaGxqYTZnNnVxMnRzMnI2MHN3bTJkOQ==",
                "index": true
              },
              {
                "key": "YW1vdW50",
                "value": "NTAwMDBiYXNldGNybw==",
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
                "value": "NTAwMDBiYXNldGNybw==",
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
                "value": "dGNybzF2dXJmaHFmMGoyamdmcGphaGxqYTZnNnVxMnRzMnI2MHN3bTJkOQ==",
                "index": true
              },
              {
                "key": "YW1vdW50",
                "value": "NTAwMDBiYXNldGNybw==",
                "index": true
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "c2VuZGVy",
                "value": "dGNybzF2dXJmaHFmMGoyamdmcGphaGxqYTZnNnVxMnRzMnI2MHN3bTJkOQ==",
                "index": true
              }
            ]
          },
          {
            "type": "tx",
            "attributes": [
              {
                "key": "YWNjX3NlcQ==",
                "value": "dGNybzF2dXJmaHFmMGoyamdmcGphaGxqYTZnNnVxMnRzMnI2MHN3bTJkOS8z",
                "index": true
              }
            ]
          },
          {
            "type": "tx",
            "attributes": [
              {
                "key": "c2lnbmF0dXJl",
                "value": "UTFacXRnTU5SQ3BNd2VBQmZIb0dhSnFzbVgvM0RFWElmc3ZoWkxzR3ZwRTVjYXF2NC9NRmZ6NE9oM2VwS1RCaWhFNlRzc3RWL0lrc3lTY2U0QzhHblE9PQ==",
                "index": true
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "YWN0aW9u",
                "value": "L2Nvc21vcy5hdXRoei52MWJldGExLk1zZ1Jldm9rZQ==",
                "index": true
              }
            ]
          },
          {
            "type": "cosmos.authz.v1beta1.EventRevoke",
            "attributes": [
              {
                "key": "Z3JhbnRlcg==",
                "value": "InRjcm8xdnVyZmhxZjBqMmpnZnBqYWhsamE2ZzZ1cTJ0czJyNjBzd20yZDki",
                "index": true
              },
              {
                "key": "Z3JhbnRlZQ==",
                "value": "InRjcm8xNXpoNXRuN3hqZGVjdTR6amNsc21sbmxodDVlYWQybXg4NGdhdTIi",
                "index": true
              },
              {
                "key": "bXNnX3R5cGVfdXJs",
                "value": "Ii9jb3Ntb3MuYmFuay52MWJldGExLk1zZ1NlbmQi",
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
            "value": "NDgzOTI1Mjg4M2Jhc2V0Y3Jv",
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
            "value": "NDgzOTI1Mjg4M2Jhc2V0Y3Jv",
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
            "value": "NDgzOTI1Mjg4M2Jhc2V0Y3Jv",
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
            "value": "NDgzOTI1Mjg4M2Jhc2V0Y3Jv",
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
            "value": "NDgzOTI1Mjg4M2Jhc2V0Y3Jv",
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
            "value": "MC4wMDA1OTUwOTg4Mjg3OTY0NDk=",
            "index": true
          },
          {
            "key": "aW5mbGF0aW9u",
            "value": "MC4wMTIxMTczODk2NTkyMzMxOTU=",
            "index": true
          },
          {
            "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
            "value": "MzA1NDMwNDEzNjA4ODgxNTAuNDkxNjk2MTY2MzA2MjIzMTY1",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "NDgzOTI1Mjg4Mw==",
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
            "value": "NDgzOTI1Mjg4M2Jhc2V0Y3Jv",
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
            "value": "NDgzOTI1Mjg4M2Jhc2V0Y3Jv",
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
            "value": "NDgzOTI1Mjg4M2Jhc2V0Y3Jv",
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
            "value": "MjQxOTYyNjQ0LjE1MDAwMDAwMDAwMDAwMDAwMGJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxczRnZ3EyenV6dndnNWs4dm54Mnhmd3RkbTRjejZ3dG51cWtsN2E=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MjQxOTYyNjQuNDE1MDAwMDAwMDAwMDAwMDAwYmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxczRnZ3EyenV6dndnNWs4dm54Mnhmd3RkbTRjejZ3dG51cWtsN2E=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MjQxOTYyNjQ0LjE1MDAwMDAwMDAwMDAwMDAwMGJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxczRnZ3EyenV6dndnNWs4dm54Mnhmd3RkbTRjejZ3dG51cWtsN2E=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTUzMjQzMjUzLjU1ODUzNzY2NjYwODU4MDc1MGJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNjN0djU5eXpnZXFjYXA4bHJzYTJyNHprNTgwaDhkZHI1YTBzZGQ=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTUzMjQzMjUzNS41ODUzNzY2NjYwODU4MDc1MDBiYXNldGNybw==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNjN0djU5eXpnZXFjYXA4bHJzYTJyNHprNTgwaDhkZHI1YTBzZGQ=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTUzMjQyOTQ2LjQ2MDI4NzE1NjMxOTU4ODM0NGJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNGZ6a3N6NWg3MmV0NHNzanRxcHdzbWh6NnlzazZyNG5hNXRyNjM=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTUzMjQyOTQ2NC42MDI4NzE1NjMxOTU4ODM0NDFiYXNldGNybw==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNGZ6a3N6NWg3MmV0NHNzanRxcHdzbWh6NnlzazZyNG5hNXRyNjM=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTUzMjQyNjM5Ljk3NTAwNzIwNjU5MTE0MzE2NWJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxczRnZ3EyenV6dndnNWs4dm54Mnhmd3RkbTRjejZ3dG51cWtsN2E=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTUzMjQyNjM5OS43NTAwNzIwNjU5MTE0MzE2NDliYXNldGNybw==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxczRnZ3EyenV6dndnNWs4dm54Mnhmd3RkbTRjejZ3dG51cWtsN2E=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTgzOC45MTE2Nzk2OTU2MTIyOTY5MzFiYXNldGNybw==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNTc5cGVzeDI3NDNrdnR1eDc3Z3EzaHhyMnZ3bmZ4dzl0djAzOWE=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTgzOC45MTE2Nzk2OTU2MTIyOTY5MzFiYXNldGNybw==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNTc5cGVzeDI3NDNrdnR1eDc3Z3EzaHhyMnZ3bmZ4dzl0djAzOWE=",
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

const TX_MSG_REVOKE_TXS_RESP = `{
  "tx": {
    "body": {
      "messages": [
        {
          "@type": "/cosmos.authz.v1beta1.MsgRevoke",
          "granter": "tcro1vurfhqf0j2jgfpjahlja6g6uq2ts2r60swm2d9",
          "grantee": "tcro15zh5tn7xjdecu4zjclsmlnlht5ead2mx84gau2",
          "msg_type_url": "/cosmos.bank.v1beta1.MsgSend"
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
            "key": "AxcZ9W91aKz8X9eZOQTjv4JiFcLP3JjDIc/+go8FR13m"
          },
          "mode_info": {
            "single": {
              "mode": "SIGN_MODE_DIRECT"
            }
          },
          "sequence": "3"
        }
      ],
      "fee": {
        "amount": [
          {
            "denom": "basetcro",
            "amount": "50000"
          }
        ],
        "gas_limit": "200000",
        "payer": "",
        "granter": ""
      }
    },
    "signatures": [
      "Q1ZqtgMNRCpMweABfHoGaJqsmX/3DEXIfsvhZLsGvpE5caqv4/MFfz4Oh3epKTBihE6TsstV/IksySce4C8GnQ=="
    ]
  },
  "tx_response": {
    "height": "123731",
    "txhash": "9921D7DDC530DB81B0A5FD1163678757A7B7B7D8ED78C2B4BE433BFFD30C1228",
    "codespace": "",
    "code": 0,
    "data": "0A210A1F2F636F736D6F732E617574687A2E763162657461312E4D73675265766F6B65",
    "raw_log": "[{\"events\":[{\"type\":\"cosmos.authz.v1beta1.EventRevoke\",\"attributes\":[{\"key\":\"granter\",\"value\":\"\\\"tcro1vurfhqf0j2jgfpjahlja6g6uq2ts2r60swm2d9\\\"\"},{\"key\":\"grantee\",\"value\":\"\\\"tcro15zh5tn7xjdecu4zjclsmlnlht5ead2mx84gau2\\\"\"},{\"key\":\"msg_type_url\",\"value\":\"\\\"/cosmos.bank.v1beta1.MsgSend\\\"\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/cosmos.authz.v1beta1.MsgRevoke\"}]}]}]",
    "logs": [
      {
        "msg_index": 0,
        "log": "",
        "events": [
          {
            "type": "cosmos.authz.v1beta1.EventRevoke",
            "attributes": [
              {
                "key": "granter",
                "value": "\"tcro1vurfhqf0j2jgfpjahlja6g6uq2ts2r60swm2d9\""
              },
              {
                "key": "grantee",
                "value": "\"tcro15zh5tn7xjdecu4zjclsmlnlht5ead2mx84gau2\""
              },
              {
                "key": "msg_type_url",
                "value": "\"/cosmos.bank.v1beta1.MsgSend\""
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "action",
                "value": "/cosmos.authz.v1beta1.MsgRevoke"
              }
            ]
          }
        ]
      }
    ],
    "info": "",
    "gas_wanted": "200000",
    "gas_used": "52212",
    "tx": {
      "@type": "/cosmos.tx.v1beta1.Tx",
      "body": {
        "messages": [
          {
            "@type": "/cosmos.authz.v1beta1.MsgRevoke",
            "granter": "tcro1vurfhqf0j2jgfpjahlja6g6uq2ts2r60swm2d9",
            "grantee": "tcro15zh5tn7xjdecu4zjclsmlnlht5ead2mx84gau2",
            "msg_type_url": "/cosmos.bank.v1beta1.MsgSend"
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
              "key": "AxcZ9W91aKz8X9eZOQTjv4JiFcLP3JjDIc/+go8FR13m"
            },
            "mode_info": {
              "single": {
                "mode": "SIGN_MODE_DIRECT"
              }
            },
            "sequence": "3"
          }
        ],
        "fee": {
          "amount": [
            {
              "denom": "basetcro",
              "amount": "50000"
            }
          ],
          "gas_limit": "200000",
          "payer": "",
          "granter": ""
        }
      },
      "signatures": [
        "Q1ZqtgMNRCpMweABfHoGaJqsmX/3DEXIfsvhZLsGvpE5caqv4/MFfz4Oh3epKTBihE6TsstV/IksySce4C8GnQ=="
      ]
    },
    "timestamp": "2021-08-26T20:13:47Z"
  }
}`
