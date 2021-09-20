package repository

import "github.com/valyala/fasthttp"

type ICRUDRepository interface {
	GET(ctx *fasthttp.RequestCtx)
	POST(ctx *fasthttp.RequestCtx)
}
