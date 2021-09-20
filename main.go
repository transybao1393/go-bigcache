// package file_validation_package

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"

	s "github.com/transybao1393/go-bigcache/services"
)

type FileStruct struct {
	ContentDisposition map[string]string
	ContentType        string
}

type BCRequestStruct struct {
	name  string
	value string
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "form.html")
	case "POST":

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		name := r.FormValue("name")
		occupation := r.FormValue("occupation")

		//- file handler
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("filename")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "file header %v", handler.Header)

		fmt.Fprintf(w, "%s is a %s\n", name, occupation)

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func bigCacheHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/bc" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {

	case "GET":
		fmt.Println("124")
		//- get bigcache by name

	case "POST":
		var bcr BCRequestStruct
		fmt.Println("body", r.Body)

		// dec := json.NewDecoder(r.Body)
		// dec.DisallowUnknownFields()

		err := json.NewDecoder(r.Body).Decode(&bcr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "bigcache request body: %+v", bcr)
	}

}

func main() {

	//- load env file
	config, error := s.LoadConfig(".")
	if error != nil {
		log.Fatal("cannot load config:", error)
	}
	fmt.Println("config", reflect.TypeOf(config))
	fmt.Println("type of verbose", reflect.TypeOf(config.Verbose))

	//- fasthttp
	s.FHExecute()
}
