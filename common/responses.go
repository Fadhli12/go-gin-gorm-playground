package common

import (
	"encoding/json"
	"net/http"
)

type ResponseRequest struct {
	Data  interface{}
	Error error
}

type Responses struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

// ResponseJson :
func ResponseJson(w http.ResponseWriter, r *ResponseRequest) {
	w.Header().Set("Content-Type", "application/json")
	if r.Error != nil {
		errInfo, ok := r.Error.(*RequestError)
		if ok {
			w.WriteHeader(errInfo.StatusCode)
			json.NewEncoder(w).Encode(Responses{
				Code:   errInfo.StatusCode,
				Status: errInfo.Error(),
			})
		}
		return
	}
	w.WriteHeader(http.StatusOK)

	resp := Responses{}
	resp.Status = ResponseSuccess
	resp.Code = http.StatusOK
	if r.Data != nil {
		resp.Data = r.Data
	}
	json.NewEncoder(w).Encode(resp)
}
