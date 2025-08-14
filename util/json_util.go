package util

import (
	"encoding/json"
	"go-rest-api/helper"
	"net/http"
)

func ReadFromRequest(r *http.Request, v interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(v)
	helper.PanicIfError(err)
}
