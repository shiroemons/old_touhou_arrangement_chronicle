package ctxkey

type key int

const (
	GinCtxKey key = iota
	TxCtxKey
)
