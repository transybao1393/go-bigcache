package services

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/valyala/fasthttp"
)

type MyHandler struct {
	foobar string
}

type BCBodyStruct struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func jsonResponseHandler(ctx *fasthttp.RequestCtx, statusCode int, data string) ([]byte, error) {
	//- set response headers
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.Response.SetStatusCode(statusCode)

	return json.MarshalIndent(map[string]string{
		"statusCode": strconv.Itoa(statusCode),
		"message":    fasthttp.StatusMessage(statusCode),
		"data":       data,
	}, "", "  ")
}

func rootHandler(ctx *fasthttp.RequestCtx) {

	switch string(ctx.Method()) {
	case "GET":
		byteParam := ctx.QueryArgs().Peek("key")
		fmt.Println("byte params", string(byteParam))
		data := GetCacheData(string(byteParam))
		resp, err := jsonResponseHandler(ctx, 200, data)
		if err != nil {
			log.Fatal(err)
			// panic(err)
		}
		ctx.Write(resp)
	case "POST":
		//- add value to cache

		//- get body data
		bodyBytes := ctx.Request.Body()
		fmt.Println(string(bodyBytes))

		//- json transform
		bcData := BCBodyStruct{}
		json.Unmarshal(bodyBytes, &bcData)
		fmt.Printf("%s", bcData)

		SetCacheData(bcData.Name, bcData.Value)

		//- success response
		resp, err := jsonResponseHandler(ctx, 200, "POST success")
		if err != nil {
			log.Fatal(err)
			// panic(err)
		}
		ctx.Write(resp)
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
