package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	// "github.com/mrityunjaygr8/shorty/app"
)

func (wa *WebApp) lookupHandler(w http.ResponseWriter, r *http.Request) {
	query := mux.Vars(r)

	if query["token"] == "" {
		json.NewEncoder(w).Encode(message{Message: "The token cannot be blank"})
		return
	}

	long, found, err := wa.app.Lookup(query["token"])
	if err != nil {
		json.NewEncoder(w).Encode(message{Message: fmt.Sprintf("An error has occurred. %s", err.Error())})
		return
	}

	if !found {
		json.NewEncoder(w).Encode(message{Message: fmt.Sprintf("token: %s does not exist or has expired", query["token"])})
		return
	}

	var redirect_to string

	if strings.HasPrefix(long, "http://") || strings.HasPrefix(long, "https://") {
		redirect_to = long
	} else {
		redirect_to = fmt.Sprintf("https://%s", long)
	}
	http.Redirect(w, r, redirect_to, http.StatusSeeOther)
}
