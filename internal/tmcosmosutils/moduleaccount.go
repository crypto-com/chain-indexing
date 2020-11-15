package tmcosmosutils

type ModuleAccounts struct {
	FeeCollector        string
	Mint                string
	Distribution        string
	Gov                 string
	BondedTokensPool    string
	NotBondedTokensPool string
}

func NewModuleAccounts(accountAddressPrefix string) ModuleAccounts {
	if accountAddressPrefix == "trco" {
		return ModuleAccounts{
			FeeCollector:        "tcro17xpfvakm2amg962yls6f84z3kell8c5lxhzaha",
			Mint:                "tcro1m3h30wlvsf8llruxtpukdvsy0km2kum87lx9mq",
			Distribution:        "tcro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8339p4l",
			Gov:                 "tcro10d07y265gmmuvt4z0w9aw880jnsr700jvvjc2n",
			BondedTokensPool:    "tcro1fl48vsnmsdzcv85q5d2q4z5ajdha8yu3r4gj9h",
			NotBondedTokensPool: "tcro1tygms3xhhs3yv487phx3dw4a95jn7t7lh45rnr",
		}
	}
	// TODO: Dynamically hash and encode the module accounts addresses
	panic("unimplemented ModuleAccounts for non-tcro prefix")
}
