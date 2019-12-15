package util

import "encoding/json"

func MustEncode(enc *json.Encoder, v interface{}) {
	err := enc.Encode(v)
	if err != nil {
		panic(err)
	}
}
