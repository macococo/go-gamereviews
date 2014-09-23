package utils

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

func HandleError(err error) error {
	if err != nil {
		log.Println(err)
	}
	return err
}

func GetParam(r *http.Request, name string, def string) string {
	param := r.FormValue(name)
	if param == "" {
		return def
	}
	return param
}

func GetParamInt(r *http.Request, name string, def int) int {
	str := r.FormValue(name)
	if str == "" {
		return def
	}

	param, err := strconv.Atoi(str)
	if HandleError(err) != nil {
		return def
	}
	return param
}

func ToJsonBytes(v interface{}) []byte {
	content, err := json.Marshal(v)
	HandleError(err)

	return content
}

func FromJsonBytes(data []byte, v interface{}) {
	err := json.Unmarshal(data, v)
	HandleError(err)
}

func WriteJson(w http.ResponseWriter, response interface{}) {
	WriteJsonBytes(w, ToJsonBytes(response))
}

func WriteJsonBytes(w http.ResponseWriter, response []byte) {
	w.Header()["Content-Type"] = []string{"application/json;charset=UTF-8"}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func GetGUID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
