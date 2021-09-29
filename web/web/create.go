package web

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateResponse struct {
	Long_URL string `json:"long_url"`
	Token    string `json:"token"`
}

func (wa *WebApp) createHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("url")
	if query == "" {
		json.NewEncoder(w).Encode(message{Message: "You need to pass a url"})
		return
	}

	response := CreateResponse{}
	response.Long_URL = query
	token, err := wa.app.Create(query)
	if err != nil {
		json.NewEncoder(w).Encode(message{Message: fmt.Sprintf("An error has occurred. %s", err.Error())})
		return
	}

	response.Token = fmt.Sprintf("%s/%s", wa.host, token)
	json.NewEncoder(w).Encode(response)
}
