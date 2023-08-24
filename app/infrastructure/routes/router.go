package routes

import (
	"net/http"
	slotsRouter "rhSystem_server/app/infrastructure/routes/slots"
	// swaggerFiles "github.com/swaggo/files"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/api/v1/slots/", slotsRouter.Router())
	// mux.Handle("/api/v1/docs/", swaggerFiles.Handler)

	return mux
}
