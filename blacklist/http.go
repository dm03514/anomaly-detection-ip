package blacklist

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Payload struct {
	IP string
}

func (p Payload) Bytes() ([]byte, error) {
	return json.Marshal(p)
}

type Blacklist interface {
	Test(ip string) (Result, error)
}

type Handler struct {
	Blacklist Blacklist
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	payload := Payload{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		msg := fmt.Sprintf("received: %q.  Expected message of format %+v",
			err, Payload{})
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	result, err := h.Blacklist.Test(payload.IP)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&result)
}
