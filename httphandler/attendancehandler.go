package httphandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/mymin427/wedding-invitation-server/sqldb"
	"github.com/mymin427/wedding-invitation-server/types"
)

type AttendanceHandler struct {
	http.Handler
}

func (h *AttendanceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Parse query
		offsetQ := r.URL.Query().Get("offset")
		limitQ := r.URL.Query().Get("limit")
		if offsetQ == "" {
			offsetQ = "0"
		}
		if limitQ == "" {
			limitQ = "20"
		}
		offset, err := strconv.Atoi(offsetQ)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("BadRequest"))
			return
		}
		limit, err := strconv.Atoi(limitQ)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("BadRequest"))
			return
		}

		items, err := sqldb.GetAttendance(offset, limit)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("InternalServerError"))
			return
		}
		total, err := sqldb.CountAttendance()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("InternalServerError"))
			return
		}
		resp := types.AttendanceListResponse{Items: items, Total: total}
		pbytes, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("InternalServerError"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(pbytes)
	} else if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		var attendance types.AttendanceCreate
		err := decoder.Decode(&attendance)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("BadRequest"))
			return
		}

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("InternalServerError"))
			return
		}

		err = sqldb.CreateAttendance(attendance.Side, attendance.Name, attendance.Meal, attendance.Count)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("InternalServerError"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
	}
}
