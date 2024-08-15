package main
import (
    "net/http"
    "encoding/json"
    "log"
)

func respError(w http.ResponseWriter, errInfo error, status int) {
    resp := struct {
	Error string `json:"error"`
    }{errInfo.Error()}
    jsonResp, err := json.Marshal(resp)
    if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    w.WriteHeader(status)
    w.Header().Set("Content-Type", "application/json")
    _, err = w.Write(jsonResp)
    if err != nil {
	log.Println(err)
    }
}

func respResult(w http.ResponseWriter, resInfo string, evs []Event, status int) {
    resp := struct {
	Info string `json:"info"`
	Events []Event `json:"events"`
    }{resInfo, evs}
    jsonResp, err := json.Marshal(resp)
    if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    w.WriteHeader(status)
    w.Header().Set("Content-Type", "application/json")
    _, err = w.Write(jsonResp)
    if err != nil {
	log.Println(err)
    }
}
