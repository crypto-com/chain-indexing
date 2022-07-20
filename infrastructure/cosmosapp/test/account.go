package infrastructure_cosmosapp_test

const AUTH_ACCOUNT_1_JSON = `{
	"account": {
	  "@type": "/cosmos.auth.v1beta1.BaseAccount",
	  "address": "cro1rrnm2rkhrkgelj7j9pxeu7sggv8fde3tkymym7",
	  "pub_key": {
		"@type": "/cosmos.crypto.secp256k1.PubKey",
		"key": "AomWNLM+dBB76InhfghTzlUDOPevNG2AClk286KuSODS"
	  },
	  "account_number": "3",
	  "sequence": "1"
	}
  }`
const AUTH_ACCOUNT_2_JSON = `{
	"account": {
	  "@type": "/cosmos.vesting.v1beta1.DelayedVestingAccount",
	  "base_vesting_account": {
		"base_account": {
		  "address": "cro1442fdq2t62vqchraj6ujxnhq3gkzq3ra9nt4lc",
		  "pub_key": {
			"@type": "/cosmos.crypto.secp256k1.PubKey",
			"key": "ApTiQlAs/mTr9a1RQwmm5G+bXe2MvGRypncY9pAHcWKO"
		  },
		  "account_number": "4",
		  "sequence": "1"
		},
		"original_vesting": [
		  {
			"denom": "basecro",
			"amount": "20000000000"
		  }
		],
		"delegated_free": [],
		"delegated_vesting": [],
		"end_time": "1658221100"
	  }
	}
  }`
