package middlewares

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	sentry "github.com/getsentry/sentry-go"
	log "github.com/sirupsen/logrus"
)

// Recovery : global recovery middleware
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body interface{}
		var reqBody []byte
		reqBody, _ = ioutil.ReadAll(r.Body)
		_ = json.Unmarshal(reqBody, &body)
		r.Body.Close()
		r.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
		defer func() {
			err := recover()
			if err != nil {
				// send to sentry
				errString := fmt.Sprintf("%v", err)
				go func() {
					sentry.WithScope(func(scope *sentry.Scope) {
						scope.SetTag("project", "Helathmart")
						sentry.CaptureException(errors.New(string(errString)))
						sentry.Flush(time.Second * 5)
						fmt.Println(errString)
					})
				}()
				errResp := map[string]string{
					"error": "Internal Server Error",
				}
				jsonBody, _ := json.Marshal(errResp)
				log.WithFields(log.Fields{
					"path":          r.URL.Path,
					"method":        r.Method,
					"body":          body,
					"useragent":     r.Header.Get("User-Agent"),
					"status":        http.StatusInternalServerError,
					"response_body": errResp,
					"error":         errString,
				}).Info("error data")
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(jsonBody)
			}

		}()

		next.ServeHTTP(w, r)

	})
}
