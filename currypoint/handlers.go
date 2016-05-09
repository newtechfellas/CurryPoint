package currypoint

import (
	"log"
	"net/http"
	"runtime/debug"
	"github.com/newtechfellas/CurryPoint/util"
)

// recoveryHandler is copied from Gorilla package and enhanced to log the stack trace to the log file instead of Stderr.
// This handler suppresses all panics and prevents server from going down for un-recovered panics
type RecoveryHandler struct {
	Handler http.Handler
}

func (h RecoveryHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			log.Println(util.BytesToString(debug.Stack()))
		}
	}()
	h.Handler.ServeHTTP(w, req)
}