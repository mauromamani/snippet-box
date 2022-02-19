package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// define app struct
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// add command flag to address
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// create new logs
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// initialize app instance
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// create new server struct
	server := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := server.ListenAndServe()
	errorLog.Fatal(err.Error())
}
