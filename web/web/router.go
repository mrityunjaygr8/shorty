package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	stdlog "log"

	"github.com/gorilla/mux"
	"github.com/mrityunjaygr8/shorty/app"

	log "github.com/go-kit/log"
)

type message struct {
	Message string `json:"message"`
}

func yo(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(message{Message: "yo"})
}

func SetupRouter(config app.Config) {
	webApp := WebApp{}
	webApp.host = config.WEB_HOST
	webApp.port = config.WEB_PORT

	webApp.app = app.Setup(config)

	webApp.router = mux.NewRouter().StrictSlash(true)
	webApp.router.HandleFunc("/yo", yo)
	webApp.router.HandleFunc("/create", webApp.createHandler)

	var logger log.Logger

	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	stdlog.SetOutput(log.NewStdlibAdapter(logger))

	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "loc", log.DefaultCaller)

	loggingMiddleware := LogginMiddleware(logger)
	loggedRouter := loggingMiddleware(webApp.router)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", webApp.port), loggedRouter); err != nil {
		logger.Log("status", "fatal", "err", err)
		os.Exit(1)
	}

}
