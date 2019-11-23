package api

type Context interface {
	QueryIntArg(name string, defaultValue int) (int, error)
	PathInt64Arg(name string) (int64, error)
	QueryOffsetArg() (int, error)
	QueryLimitArg() (int, error)
	Body() []byte
}
