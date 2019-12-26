package util

import "encoding/json"

func MustEncode(enc *json.Encoder, v interface{}) {
	err := enc.Encode(v)
	if err != nil {
		panic(err)
	}
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

type GenericResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
