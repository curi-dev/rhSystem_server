package routes

import (
	"net/http"
	appointmentsRouter "rhSystem_server/app/infrastructure/routes/appointments"
	candidatesRouter "rhSystem_server/app/infrastructure/routes/candidates"
	slotsRouter "rhSystem_server/app/infrastructure/routes/slots"
	// swaggerFiles "github.com/swaggo/files"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/api/v1/slots/", slotsRouter.Router())
	mux.Handle("/api/v1/appointments/", appointmentsRouter.Router())
	mux.Handle("/api/v1/candidates/", candidatesRouter.Router())

	return mux
}
