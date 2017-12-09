package handlers

import (
	"fmt"
	"net/http"
)

func SubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("hub.verify_token") == "moi" {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, r.URL.Query().Get("hub.challenge"))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
