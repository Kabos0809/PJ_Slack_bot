package Middleware

import (
	"bytes"
	"log"
	"io"
	"net/http"
	"os"
	
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

const version = "v0.0.0.1"
var revision string = "HEAD"

//http.NewServeMuxを使用するときのMiddleWare
type SecretsVerifierMiddleware struct {
	handler http.Handler
}

func NewSecretsVerifierMiddleware(h http.Handler) *SecretsVerifierMiddleware {
	return &SecretsVerifierMiddleware{h}
}

func (v *SecretsVerifierMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("[INFO] Start ServeHTTP")
	switch r.Method {
	case http.MethodGet:
		log.Printf("[INFO] Hello! (Will_Bot version: %s, revision: %s)", version, revision)
		w.WriteHeader(200)
		return
	case http.MethodPost:
		var slack_signing_secret string = os.Getenv("SLACK_SIGNING_SECRET")
		log.Printf("%v", slack_signing_secret)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("[ERROR] Failed to read Response body: %v", err)
			log.Printf("[ERROR] Failed to ServeHTTP")
			w.WriteHeader(200)
			return
		}
		r.Body.Close()
		r.Body = io.NopCloser(bytes.NewBuffer(body))

		sv, err := slack.NewSecretsVerifier(r.Header, slack_signing_secret)
		if err != nil {
			log.Println("[ERROR] Failed to build SecretsVerifier: %v", err)
			log.Printf("[ERROR] Failed to ServeHTTP")
			w.WriteHeader(200)
			return
		}

		if _, err = sv.Write(body); err != nil {
			log.Printf("[ERROR] Failed to write body to SecretsVerifier: %v", err)
			log.Printf("[ERROR] Failed to ServeHTTP")
			w.WriteHeader(200)
			return
		}

		if err := sv.Ensure(); err != nil {
			log.Printf("[ERROR] Token is Unauthorized: %v", err)
			log.Printf("[ERROR] Failed to ServeHTTP")
			w.WriteHeader(200)
			return
		}


		v.handler.ServeHTTP(w, r)
		log.Printf("[INFO] End ServeHTTP")
	}	
}


//http.HandleFuncを使うときのMiddleWare
func Verify(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			log.Printf("[INFO] Hello! (Will_Bot version: %s, revision: %s)", version, revision)
			w.WriteHeader(200)
			return
		case http.MethodPost:
			godotenv.Load("../.env")
			var slack_signing_secret string = os.Getenv("SLACK_SIGNING_SECRET")
			body, err := io.ReadAll(r.Body)
			if err != nil {
				log.Println("[ERROR] Failed to read Response body: %v", err)
				w.WriteHeader(200)
				return
			}
			r.Body.Close()
			r.Body = io.NopCloser(bytes.NewBuffer(body))

			sv, err := slack.NewSecretsVerifier(r.Header, slack_signing_secret)
			if err != nil {
				log.Println("[ERROR] Failed to build SecretsVerifier: %v", err)
				w.WriteHeader(200)
				return
			}

			if _, err = sv.Write(body); err != nil {
				log.Printf("[ERROR] Failed to write body to SecretsVerifier: %v", err)
				w.WriteHeader(200)
				return
			}

			if err := sv.Ensure(); err != nil {
				log.Printf("[ERROR] Failed to discriminate that valid request: %v", err)
				w.WriteHeader(200)
				return
			}


			next.ServeHTTP(w, r)
		}
	}
}