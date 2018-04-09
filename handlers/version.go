package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Abdujabbor/minikube-go-microservice/version"
	"github.com/julienschmidt/httprouter"
)

func versionHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	jsonValue, err := json.Marshal(struct {
		Commit    string
		Release   string
		BuildDate string
	}{
		version.Commit,
		version.Release,
		version.BuildDate,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Write(jsonValue)
}
