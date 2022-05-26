package main

import (
	"net/http"

	"github.com/Tozinoo/web/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
