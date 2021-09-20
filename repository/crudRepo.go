package repository

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	c "github.com/transybao1393/go-bigcache/controller"
	"github.com/valyala/fasthttp"
)

var (
	controller = c.BigCacheControllerInit()
)

//- implement new class
type CRUDRepo struct{}

type BCBodyStruct struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func InitCRUDRepo() ICRUDRepository {
	return &CRUDRepo{}
}

func (cr *CRUDRepo) GET(ctx *fasthttp.RequestCtx) {
	byteParam := ctx.QueryArgs().Peek("key")
	fmt.Println("byte params", string(byteParam))
	data := controller.GetCacheData(string(byteParam))
	resp, err := jsonResponseHandler(ctx, 200, data)
	if err != nil {
		log.Fatal(err)
		// panic(err)
	}
	ctx.Write(resp)
}
func (cr *CRUDRepo) POST(ctx *fasthttp.RequestCtx) {
	//- add value to cache

	//- get body data
	bodyBytes := ctx.Request.Body()
	fmt.Println(string(bodyBytes))

	//- json transform
	bcData := BCBodyStruct{}
	json.Unmarshal(bodyBytes, &bcData)
	fmt.Printf("%s", bcData)
	controller.SetCacheData(bcData.Name, bcData.Value)

	//- success response
	resp, err := jsonResponseHandler(ctx, 200, "POST success")
	if err != nil {
		log.Fatal(err)
		// panic(err)
	}
	ctx.Write(resp)
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
