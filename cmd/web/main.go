package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// define app struct
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// add command flag to address
	addr := flag.String("addr", ":4000", "HTTP network address")
	// flag for mysql
	dsn := flag.String("dsn", "root:1234@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	// create new logs
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Open DB
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

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
	err = server.ListenAndServe()
	errorLog.Fatal(err.Error())
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
