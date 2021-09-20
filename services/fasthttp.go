package services

import (
	"fmt"

	"github.com/transybao1393/go-bigcache/repository"
	"github.com/valyala/fasthttp"
)

var (
	bcr = repository.InitCRUDRepo()
)

type MyHandler struct {
	foobar string
}

func rootHandler(ctx *fasthttp.RequestCtx) {

	switch string(ctx.Method()) {
	case "GET":
		bcr.GET(ctx)
	case "POST":
		bcr.POST(ctx)
	}
}

func defaultHandler() func(*fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/":
			rootHandler(ctx)
		case "/bc":
			// barHandler(ctx)
		default:
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}
	}
}

// request handler in net/http style, i.e. method bound to MyHandler struct.
func (h *MyHandler) HandleFastHTTP(ctx *fasthttp.RequestCtx) {
	// notice that we may access MyHandler properties here - see h.foobar.
	fmt.Fprintf(ctx, "Hello, world! Requested path is %q. Foobar is %q",
		ctx.Path(), h.foobar)
}

// request handler in fasthttp style, i.e. just plain function.
func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hi there! RequestURI is %q", ctx.RequestURI())
}

func FHExecute() {

	// pass bound struct method to fasthttp
	// myHandler := &MyHandler{
	// 	foobar: "foobar",
	// }
	// fmt.Println("FastHTTP bound struct server is running on port 8080")
	// fasthttp.ListenAndServe(":8080", myHandler.HandleFastHTTP)

	// pass plain function to fasthttp
	fmt.Println("FastHTTP normal server is running on port 8081")
	fasthttp.ListenAndServe(":8081", defaultHandler())

}
