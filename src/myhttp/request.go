package myhttp

import(
	"net/http"
)


func GetVal(request *http.Request,param_name string) string {
	value := request.FormValue(param_name)

	return value
}

func GetMethod(request *http.Request) string{
	return request.Method
}