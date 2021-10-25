package types

type Status = string

const (
	STATUS_NOT_ESTABLISHED Status = "NotEstablished"
	STATUS_OPENED          Status = "Opened"
	STATUS_CLOSED          Status = "Closed"
)
