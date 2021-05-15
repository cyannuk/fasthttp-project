package api

import (
	"strconv"

	"fasthttp-project/api/errors"
	"github.com/goccy/go-json"
	"github.com/valyala/fasthttp"
	"gopkg.in/reform.v1"
)

func getPathParameter(ctx *fasthttp.RequestCtx, name string) (int64, error) {
	parameter := ctx.UserValue(name)
	if parameter == nil {
		return 0, fasthttp.ErrNoArgValue
	}
	return strconv.ParseInt(parameter.(string), 10, 64)
}

func getQueryOffset(ctx *fasthttp.RequestCtx) (int64, error) {
	value, err := ctx.QueryArgs().GetUint("offset")
	if err == fasthttp.ErrNoArgValue {
		return 0, nil
	}
	if value < 0 {
		return 0, errors.ErrOffset
	}
	return int64(value), nil
}

func getQueryLimit(ctx *fasthttp.RequestCtx) (int64, error) {
	value, err := ctx.QueryArgs().GetUint("limit")
	if err == fasthttp.ErrNoArgValue {
		return 50, nil
	}
	if value <= 0 || value > 50 {
		return 0, errors.ErrLimit
	}
	return int64(value), nil
}

func sendError(err error, ctx *fasthttp.RequestCtx) {
	var statusCode int
	switch err {
	case errors.ErrLimit, errors.ErrOffset, errors.ErrUserId, errors.ErrOrderId, fasthttp.ErrNoArgValue:
		statusCode = fasthttp.StatusBadRequest
	case errors.ErrForbidden:
		statusCode = fasthttp.StatusForbidden
	case reform.ErrNoRows:
		statusCode = fasthttp.StatusNotFound
	default:
		statusCode = fasthttp.StatusInternalServerError
	}
	ctx.Error(err.Error(), statusCode)
}

func sendData(data interface{}, ctx *fasthttp.RequestCtx) {
	body, err := json.Marshal(data)
	if err != nil {
		sendError(err, ctx)
		return
	}
	ctx.SetContentType("application/json")
	ctx.Response.AppendBody(body)
}
