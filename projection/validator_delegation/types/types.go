package types

type ValidatorStatus string

const (
	// UNBONDED defines a validator that is not bonded.
	UNBONDED ValidatorStatus = "Unbonded"
	// UNBONDING defines a validator that is unbonding.
	UNBONDING ValidatorStatus = "Unbonding"
	// BONDED defines a validator that is bonded.
	BONDED ValidatorStatus = "Bonded"
)

type SlashReason string

const (
	DOUBLE_SIGN       SlashReason = "double_sign"
	MISSING_SIGNATURE SlashReason = "missing_signature"
)
