package octokat

type Options struct {
	Headers Headers
	Params  interface{}
}

type Headers map[string]string
