package cdn

import (
	"encoding/json"
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
	return handlers.CORS()(handlers.CompressHandler(handlers.CombinedLoggingHandler(os.Stdout, optionalApplyProxyHandler(router, behindProxy))))
}

func okWithContent(content []byte, w http.ResponseWriter) {
	ok(w)
	w.Write(content)
}

func ok(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

func noContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func badRequest(message string, w http.ResponseWriter) {
	exception(http.StatusBadRequest, message, w)
}

func internalServerError(message string, w http.ResponseWriter) {
	exception(http.StatusInternalServerError, message, w)
}

func unauthorized(message string, w http.ResponseWriter) {
	exception(http.StatusUnauthorized, message, w)
}

func forbidden(message string, w http.ResponseWriter) {
	exception(http.StatusForbidden, message, w)
}

func exception(statusCode int, message string, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	val := &errorResponse{
		Message: message,
	}
	data, _ := json.Marshal(val)
	w.Write(data)
}
