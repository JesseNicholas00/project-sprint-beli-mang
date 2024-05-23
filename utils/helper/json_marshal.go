package helper

import "encoding/json"

// panics on marshal failure
func MustMarshalJson(obj interface{}) []byte {
	res, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	return res
}

func MustUnmarshalJson(bytes []byte, res interface{}) {
	err := json.Unmarshal(bytes, res)
	if err != nil {
		panic(err)
	}
}
