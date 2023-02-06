package icaauth

type CosmosTx struct {
	Messages []Message
}

type Message struct {
	Type string
}
