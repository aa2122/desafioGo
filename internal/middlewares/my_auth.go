package middlewares

import (
	"crypto/sha256"
	"crypto/subtle"
	"net/http"
	// "exemplo.com/desafioGo/internal/application"
)

func MyAuth(next http.Handler) http.Handler {
	// func MyAuth(app *application.App) func(next http.Handler) http.Handler {
	// return func(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		user, pass, ok := r.BasicAuth()
		if ok {

			userHash := sha256.Sum256([]byte(user))
			passHash := sha256.Sum256([]byte(pass))
			expectedUserHash := sha256.Sum256([]byte("a.Config.MyAuthUsername"))
			expectedPassHash := sha256.Sum256([]byte("a.Config.MyAuthPassword"))
			// expectedUserHash := sha256.Sum256([]byte("olaola"))
			// expectedPassHash := sha256.Sum256([]byte("12345"))

			correctUsername := (subtle.ConstantTimeCompare(userHash[:], expectedUserHash[:]) == 1)
			correctPassword := (subtle.ConstantTimeCompare(passHash[:], expectedPassHash[:]) == 1)

			if correctUsername && correctPassword {
				next.ServeHTTP(w, r)
				// log.Printf("The authentication provided is valid")
				// h.ServeHTTP(w, r)

				return
			}

		}

		w.Header().Set("WWW-Authenticate", `Basic realm="realm", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
	// }
}
