package api

import (
	"net/http"

	"github.com/dshum/school/internal/utils"
)

func (s *Server) ApiStatus(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status": "success",
		"data":   "School API is running smoothly",
	}
	utils.JSONResponse(w, http.StatusOK, response)
}
