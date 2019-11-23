package api

import (
	"strconv"

	"github.com/valyala/fasthttp"

	"fasthttp-project/api/errors"
)

type context struct {
	ctx *fasthttp.RequestCtx
}

func (c context) QueryIntArg(name string, defaultValue int) (int, error) {
	value, err := c.ctx.QueryArgs().GetUint(name)
	if err != nil {
		if err != fasthttp.ErrNoArgValue {
			return -1, err
		}
		value = defaultValue
	}
	return value, nil
}

func (c context) PathInt64Arg(name string) (int64, error) {
	value, err := strconv.ParseInt(c.ctx.UserValue(name).(string), 10, 64)
	if err != nil {
		return -1, err
	}
	return value, nil
}

func (c context) Body() []byte {
	return c.ctx.PostBody()
}

func (c context) QueryOffsetArg() (int, error) {
	offset, err := c.QueryIntArg("offset", 0)
	if err != nil {
		return -1, errors.ErrOffset
	}
	return offset, nil
}

func (c context) QueryLimitArg() (int, error) {
	limit, err := c.QueryIntArg("limit", 50)
	if err != nil || limit == 0 {
		return -1, errors.ErrLimit
	}
	return limit, nil
}
