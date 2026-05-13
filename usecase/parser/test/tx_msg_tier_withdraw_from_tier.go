package usecase_parser_test

const TX_MSG_TIER_WITHDRAW_FROM_TIER_TXS_RESP = `{
  "tx": {
    "body": {
      "messages": [
        {
          "@type": "/chainmain.tieredrewards.v1.MsgWithdrawFromTier",
          "owner": "cro1dulwqgcdpemn8c34sjd92fxepz5p0sqpeevw84",
          "position_id": "789"
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
    "txhash": "C9D0E1F2A3B4C5D6E7F8A9B0C1D2E3F4A5B6C7D8E9F0A1B2C3D4E5F6A7B8C9D0",
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
            "@type": "/chainmain.tieredrewards.v1.MsgWithdrawFromTier",
            "owner": "cro1dulwqgcdpemn8c34sjd92fxepz5p0sqpeevw84",
            "position_id": "789"
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
    "events": [
      {
        "type": "chainmain.tieredrewards.v1.EventPositionWithdrawn",
        "attributes": [
            {
                "key": "amount",
                "value": "{\"denom\":\"basetcro\",\"amount\":\"1000000000\"}",
                "index": true
            },
            {
                "key": "position",
                "value": "{\"id\":\"272\",\"owner\":\"tcro1vzgslnayeum4t6qt7j2cm79nrppaw4w4se038w\",\"tier_id\":3,\"amount\":\"1000000000\",\"validator\":\"\",\"delegated_shares\":\"0.000000000000000000\",\"base_rewards_per_share\":[],\"last_bonus_accrual\":\"0001-01-01T00:00:00Z\",\"exit_triggered_at\":\"2026-04-28T10:07:43.510505520Z\",\"exit_unlock_at\":\"2026-04-28T10:11:43.510505520Z\",\"created_at_height\":\"27374591\",\"created_at_time\":\"2026-04-28T10:03:12.278738851Z\",\"last_event_seq\":\"0\",\"last_known_bonded\":false}",
                "index": true
            },
            {
                "key": "msg_index",
                "value": "0",
                "index": true
            }
        ]
      }
    ]
  }
}`
