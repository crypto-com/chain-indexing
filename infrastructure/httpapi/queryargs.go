package httpapi

import "github.com/valyala/fasthttp"

// QueryArgs is a helper on top of fasthttp.QueryArgs
type QueryArgs struct {
	*fasthttp.Args
}

func NewQueryArgs(args *fasthttp.Args) *QueryArgs {
	return &QueryArgs{args}
}

func (args *QueryArgs) Has(key string) bool {
	return args.Args.Has(key)
}

// Return the arguments by key, empty string if not exist
func (args *QueryArgs) Get(key string) string {
	if !args.Has(key) {
		return ""
	}

	return string(args.Args.Peek(key))
}
