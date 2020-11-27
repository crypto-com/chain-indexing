package genesis

type Genesis struct {
	GenesisTime     string          `json:"genesis_time"`
	ChainID         string          `json:"chain_id"`
	InitialHeight   string          `json:"initial_height"`
	ConsensusParams ConsensusParams `json:"consensus_params"`
	AppHash         string          `json:"app_hash"`
	AppState        AppState        `json:"app_state"`
}

type AppState struct {
	Auth         Auth             `json:"auth"`
	Bank         Bank             `json:"bank"`
	Chainmain    Chainmain        `json:"chainmain"`
	Distribution Distribution     `json:"distribution"`
	Evidence     AppStateEvidence `json:"evidence"`
	Genutil      Genutil          `json:"genutil"`
	Gov          Gov              `json:"gov"`
	Mint         Mint             `json:"mint"`
	Params       interface{}      `json:"params"`
	Slashing     Slashing         `json:"slashing"`
	Staking      Staking          `json:"staking"`
	Upgrade      Chainmain        `json:"upgrade"`
}

type Auth struct {
	Params   AuthParams `json:"params"`
	Accounts []Account  `json:"accounts"`
}

type Account struct {
	Type               string              `json:"@type"`
	Address            *string             `json:"address,omitempty"`
	PubKey             interface{}         `json:"pub_key"`
	AccountNumber      *string             `json:"account_number,omitempty"`
	Sequence           *string             `json:"sequence,omitempty"`
	BaseVestingAccount *BaseVestingAccount `json:"base_vesting_account,omitempty"`
}

type BaseVestingAccount struct {
	BaseAccount      BaseAccount   `json:"base_account"`
	OriginalVesting  []MinDeposit  `json:"original_vesting"`
	DelegatedFree    []interface{} `json:"delegated_free"`
	DelegatedVesting []interface{} `json:"delegated_vesting"`
	EndTime          string        `json:"end_time"`
}

type BaseAccount struct {
	Address       string      `json:"address"`
	PubKey        interface{} `json:"pub_key"`
	AccountNumber string      `json:"account_number"`
	Sequence      string      `json:"sequence"`
}

type MinDeposit struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

type AuthParams struct {
	MaxMemoCharacters      string `json:"max_memo_characters"`
	TxSigLimit             string `json:"tx_sig_limit"`
	TxSizeCostPerByte      string `json:"tx_size_cost_per_byte"`
	SigVerifyCostEd25519   string `json:"sig_verify_cost_ed25519"`
	SigVerifyCostSecp256K1 string `json:"sig_verify_cost_secp256k1"`
}

type Bank struct {
	Params        BankParams       `json:"params"`
	Balances      []Balance        `json:"balances"`
	Supply        []interface{}    `json:"supply"`
	DenomMetadata []DenomMetadatum `json:"denom_metadata"`
}

type Balance struct {
	Address string       `json:"address"`
	Coins   []MinDeposit `json:"coins"`
}

type DenomMetadatum struct {
	Description string      `json:"description"`
	DenomUnits  []DenomUnit `json:"denom_units"`
	Base        string      `json:"base"`
	Display     string      `json:"display"`
}

type DenomUnit struct {
	Denom    string   `json:"denom"`
	Exponent int64    `json:"exponent"`
	Aliases  []string `json:"aliases"`
}

type BankParams struct {
	SendEnabled        []interface{} `json:"send_enabled"`
	DefaultSendEnabled bool          `json:"default_send_enabled"`
}

type Chainmain struct {
}

type Distribution struct {
	DelegatorStartingInfos          []interface{}      `json:"delegator_starting_infos"`
	DelegatorWithdrawInfos          []interface{}      `json:"delegator_withdraw_infos"`
	FeePool                         FeePool            `json:"fee_pool"`
	OutstandingRewards              []interface{}      `json:"outstanding_rewards"`
	Params                          DistributionParams `json:"params"`
	PreviousProposer                string             `json:"previous_proposer"`
	ValidatorAccumulatedCommissions []interface{}      `json:"validator_accumulated_commissions"`
	ValidatorCurrentRewards         []interface{}      `json:"validator_current_rewards"`
	ValidatorHistoricalRewards      []interface{}      `json:"validator_historical_rewards"`
	ValidatorSlashEvents            []interface{}      `json:"validator_slash_events"`
}

type FeePool struct {
	CommunityPool []interface{} `json:"community_pool"`
}

type DistributionParams struct {
	BaseProposerReward  string `json:"base_proposer_reward"`
	BonusProposerReward string `json:"bonus_proposer_reward"`
	CommunityTax        string `json:"community_tax"`
	WithdrawAddrEnabled bool   `json:"withdraw_addr_enabled"`
}

type AppStateEvidence struct {
	Evidence []interface{} `json:"evidence"`
}

type Genutil struct {
	GenTxs []GenTx `json:"gen_txs"`
}

type GenTx struct {
	Body       Body     `json:"body"`
	AuthInfo   AuthInfo `json:"auth_info"`
	Signatures []string `json:"signatures"`
}

type AuthInfo struct {
	SignerInfos []SignerInfo `json:"signer_infos"`
	Fee         Fee          `json:"fee"`
}

type Fee struct {
	Amount   []interface{} `json:"amount"`
	GasLimit string        `json:"gas_limit"`
	Payer    string        `json:"payer"`
	Granter  string        `json:"granter"`
}

type SignerInfo struct {
	PublicKey PublicKey `json:"public_key"`
	ModeInfo  ModeInfo  `json:"mode_info"`
	Sequence  string    `json:"sequence"`
}

type ModeInfo struct {
	Single Single `json:"single"`
}

type Single struct {
	Mode string `json:"mode"`
}

type PublicKey struct {
	Type string `json:"@type"`
	Key  string `json:"key"`
}

type Body struct {
	Messages                    []map[string]interface{} `json:"messages"`
	Memo                        string                   `json:"memo"`
	TimeoutHeight               string                   `json:"timeout_height"`
	ExtensionOptions            []interface{}            `json:"extension_options"`
	NonCriticalExtensionOptions []interface{}            `json:"non_critical_extension_options"`
}

type Gov struct {
	DepositParams      DepositParams `json:"deposit_params"`
	Deposits           []interface{} `json:"deposits"`
	Proposals          []interface{} `json:"proposals"`
	StartingProposalID string        `json:"starting_proposal_id"`
	TallyParams        TallyParams   `json:"tally_params"`
	Votes              []interface{} `json:"votes"`
	VotingParams       VotingParams  `json:"voting_params"`
}

type DepositParams struct {
	MaxDepositPeriod string       `json:"max_deposit_period"`
	MinDeposit       []MinDeposit `json:"min_deposit"`
}

type TallyParams struct {
	Quorum        string `json:"quorum"`
	Threshold     string `json:"threshold"`
	VetoThreshold string `json:"veto_threshold"`
}

type VotingParams struct {
	VotingPeriod string `json:"voting_period"`
}

type Mint struct {
	Minter Minter     `json:"minter"`
	Params MintParams `json:"params"`
}

type Minter struct {
	AnnualProvisions string `json:"annual_provisions"`
	Inflation        string `json:"inflation"`
}

type MintParams struct {
	BlocksPerYear       string `json:"blocks_per_year"`
	GoalBonded          string `json:"goal_bonded"`
	InflationMax        string `json:"inflation_max"`
	InflationMin        string `json:"inflation_min"`
	InflationRateChange string `json:"inflation_rate_change"`
	MintDenom           string `json:"mint_denom"`
}

type Slashing struct {
	MissedBlocks []interface{}  `json:"missed_blocks"`
	Params       SlashingParams `json:"params"`
	SigningInfos []interface{}  `json:"signing_infos"`
}

type SlashingParams struct {
	DowntimeJailDuration    string `json:"downtime_jail_duration"`
	MinSignedPerWindow      string `json:"min_signed_per_window"`
	SignedBlocksWindow      string `json:"signed_blocks_window"`
	SlashFractionDoubleSign string `json:"slash_fraction_double_sign"`
	SlashFractionDowntime   string `json:"slash_fraction_downtime"`
}

type Staking struct {
	Delegations          []interface{} `json:"delegations"`
	Exported             bool          `json:"exported"`
	LastTotalPower       string        `json:"last_total_power"`
	LastValidatorPowers  []interface{} `json:"last_validator_powers"`
	Params               StakingParams `json:"params"`
	Redelegations        []interface{} `json:"redelegations"`
	UnbondingDelegations []interface{} `json:"unbonding_delegations"`
	Validators           []interface{} `json:"validators"`
}

type StakingParams struct {
	BondDenom         string `json:"bond_denom"`
	HistoricalEntries int64  `json:"historical_entries"`
	MaxEntries        int64  `json:"max_entries"`
	MaxValidators     int64  `json:"max_validators"`
	UnbondingTime     string `json:"unbonding_time"`
}

type ConsensusParams struct {
	Block     Block                   `json:"block"`
	Evidence  ConsensusParamsEvidence `json:"evidence"`
	Validator Validator               `json:"validator"`
	Version   Chainmain               `json:"version"`
}

type Block struct {
	MaxBytes   string `json:"max_bytes"`
	MaxGas     string `json:"max_gas"`
	TimeIotaMS string `json:"time_iota_ms"`
}

type ConsensusParamsEvidence struct {
	MaxAgeNumBlocks string `json:"max_age_num_blocks"`
	MaxAgeDuration  string `json:"max_age_duration"`
}

type Validator struct {
	PubKeyTypes []string `json:"pub_key_types"`
}
