package middlewares

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/imrancluster/goboiler/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/imrancluster/goboiler/config"
)

// AuthMiddleware : token verification and whitelist middleware
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// read request body for logging
		start := time.Now().UTC()
		var body interface{}
		var reqBody []byte
		reqBody, _ = ioutil.ReadAll(r.Body)
		_ = json.Unmarshal(reqBody, &body)
		r.Body.Close()
		r.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))

		cfg := config.App()
		jwtKey := []byte(cfg.AppKey)
		tokenStr := r.Header.Get("Authorization")
		if tokenStr == "" {
			utils.RespondWithError(w, http.StatusUnauthorized, "authorization token is required")
			return
		}
		claims := &utils.JwtClaim{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(tkn *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				utils.RespondWithError(w, http.StatusUnauthorized, "the signature is not valid")
				return
			}
			utils.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		if !token.Valid {
			utils.RespondWithError(w, http.StatusUnauthorized, "invalid token")
			return
		}
		ctx := context.WithValue(r.Context(), string("userInfo"), claims)

		// sending req and response write to handler
		lr := &utils.ResponseWriterWithLog{ResponseWriter: w}
		next.ServeHTTP(lr, r.WithContext(ctx))

		// unmarsha the response body
		var resData interface{}
		_ = json.Unmarshal(lr.Body, &resData)

		// request logging
		log.WithFields(log.Fields{
			"path":           r.URL.Path,
			"method":         r.Method,
			"body":           body,
			"auth_user":      claims,
			"latency":        time.Now().UTC().Sub(start),
			"useragent":      r.Header.Get("User-Agent"),
			"status":         lr.Status,
			"response_body":  resData,
			"content_length": lr.Length,
		}).Info("request data")
	})
}
