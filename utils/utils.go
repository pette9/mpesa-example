package utils

import (
	"encoding/json"
	"net/http"
	"strconv"
)

//WriteJSONResponse ...
func WriteJSONResponse(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(status)
	w.Write(data)
}

//WriteErrorResponse ...
func WriteErrorResponse(w http.ResponseWriter, status int, errorMessage string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	httpError := map[string]interface{}{"error": errorMessage}
	data, _ := json.Marshal(httpError)
	w.Write(data)
}
