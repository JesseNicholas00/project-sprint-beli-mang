package helper

func StructToMap(obj interface{}) (res map[string]interface{}) {
	a := MustMarshalJson(obj)
	MustUnmarshalJson(a, &res)
	return
}
