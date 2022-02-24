package response

import (
	"encoding/json"
	"net/http"
)

func JSONResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)

}

// func ErrorResponse(w http.ResponseWriter, error string, code int) {
// 	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
// 	w.Header().Set("X-Content-Type-Options", "nosniff")
// 	w.WriteHeader(code)
// 	fmt.Fprintln(w, error)
// }

func ErrorResponse(w http.ResponseWriter, code int, msg string) {
	JSONResponse(w, code, map[string]string{"Error Message": msg})
}
