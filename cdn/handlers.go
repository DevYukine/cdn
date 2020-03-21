package cdn

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func optionalApplyProxyHandler(handler http.Handler, apply bool) http.Handler {
	if apply {
		return handlers.ProxyHeaders(handler)
	}
	return handler
}

func applyHandlers(router *mux.Router, behindProxy bool) http.Handler {
	return handlers.CORS()(handlers.CompressHandler(optionalApplyProxyHandler(handlers.CombinedLoggingHandler(os.Stdout, router), behindProxy)))
}
