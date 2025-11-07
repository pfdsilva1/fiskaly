package api

import (
	"encoding/json"
	"net/http"

	"github.com/pfdsilva1/fiskaly/signing-service-challenge-go/persistence"
	"github.com/pfdsilva1/fiskaly/signing-service-challenge-go/service"
)

// Response is the generic API response container.
type Response struct {
	Data interface{} `json:"data"`
}

// ErrorResponse is the generic error API response container.
type ErrorResponse struct {
	Errors []string `json:"errors"`
}

// Server manages HTTP requests and dispatches them to the appropriate services.
type Server struct {
	listenAddress string
	service       service.SignatureDeviceService
}

// NewServer is a factory to instantiate a new Server.
func NewServer(listenAddress string) *Server {
	repo := persistence.NewInMemorySignatureDeviceRepository()
	svc := service.NewSignatureService(repo)
	return &Server{
		listenAddress: listenAddress,
		service:       svc,
	}
}

// Run registers all HandlerFuncs for the existing HTTP routes and starts the Server.
func (s *Server) Run() error {
	mux := http.NewServeMux()

	mux.Handle("/api/v0/health", http.HandlerFunc(s.Health))
	mux.Handle("/api/v0/devices", http.HandlerFunc(s.SignatureDevices))
	mux.Handle("/api/v0/devices/{device_id}/sign", http.HandlerFunc(s.SignTransaction))
	mux.Handle("/api/v0/devices/{device_id}/signatures", http.HandlerFunc(s.DeviceSignatures))

	return http.ListenAndServe(s.listenAddress, mux)
}

// WriteInternalError writes a default internal error message as an HTTP response.
func WriteInternalError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
}

// WriteErrorResponse takes an HTTP status code and a slice of errors
// and writes those as an HTTP error response in a structured format.
func WriteErrorResponse(w http.ResponseWriter, code int, errors []string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	errorResponse := ErrorResponse{
		Errors: errors,
	}

	bytes, err := json.Marshal(errorResponse)
	if err != nil {
		WriteInternalError(w)
	}

	w.Write(bytes)
}

// WriteAPIResponse takes an HTTP status code and a generic data struct
// and writes those as an HTTP response in a structured format.
func WriteAPIResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	response := Response{
		Data: data,
	}

	bytes, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		WriteInternalError(w)
	}

	w.Write(bytes)
}
