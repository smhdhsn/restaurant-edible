package helper

import (
	"encoding/json"
	"net/http"
)

// RespBody is the struct containing response's structure.
type RespBody struct {
	Data    interface{} `json:"data"`
	Message *string     `json:"message"`
	Error   *string     `json:"errors"`
}

// SetData fills the data section of the response schema.
func (r *RespBody) SetData(data interface{}) *RespBody {
	r.Data = data

	return r
}

// SetMessage fills the message section of the response schema.
func (r *RespBody) SetMessage(message string) *RespBody {
	r.Message = &message

	return r
}

// SetError fills the error section of the response schema.
func (r *RespBody) SetError(err error) *RespBody {
	errMsg := err.Error()
	r.Error = &errMsg

	return r
}

// Json writes a JSON formatted response to the response writer.
func (r *RespBody) Json(rw http.ResponseWriter, code int) {
	res, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(code)
	rw.Write(res)
}
