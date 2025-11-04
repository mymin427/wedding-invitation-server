package main

import (
	"database/sql"
	"net/http"

	"github.com/mymin427/wedding-invitation-server/env"
	"github.com/mymin427/wedding-invitation-server/httphandler"
	"github.com/mymin427/wedding-invitation-server/sqldb"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
)

func main() {
	db, err := sql.Open("sqlite3", "./sql.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqldb.SetDb(db)

	mux := http.NewServeMux()
	mux.Handle("/api/guestbook", new(httphandler.GuestbookHandler))
	mux.Handle("/api/attendance", new(httphandler.AttendanceHandler))

	corHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{env.AllowOrigin},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodOptions},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
		Debug:            true,
	})

	handler := corHandler.Handler(mux)

	http.ListenAndServe(":"+env.Port, handler)
}
