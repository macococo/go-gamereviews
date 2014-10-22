// Controllers 1.
package controllers

import (
	"net/http"
)

func RoutingController(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(r.URL.Query().Get(":id") + ", " + r.URL.Query().Get(":name")))
}
