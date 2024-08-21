package middleware

import "net/http"

func Cors()http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
		//set cors
	}
}