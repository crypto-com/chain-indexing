package usecase_parser_test

const TX_MSG_TIER_CLEAR_POSITION_TXS_RESP = `{
  "tx": {
    "body": {
      "messages": [
        {
          "@type": "/chainmain.tieredrewards.v1.MsgClearPosition",
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
    "txhash": "B8C9D0E1F2A3B4C5D6E7F8A9B0C1D2E3F4A5B6C7D8E9F0A1B2C3D4E5F6A7B8C9",
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
            "@type": "/chainmain.tieredrewards.v1.MsgClearPosition",
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
    "events": []
  }
}`
