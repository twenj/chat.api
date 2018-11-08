package handler

import (
	"fmt"
	"net/http"
)

func Index(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "joe")
}