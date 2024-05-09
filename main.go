package main

import (
    "context"
    "database/sql"
    "fmt"
    "log"
    "time"

    _ "github.com/lib/pq"
)

type Row struct {
    ID        int
    Generated int
}

func main() {
    db, err := sql.Open("postgres", "user=postgres password=vakhaboff host=localhost dbname=shaxboz sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    cont, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    rows, err := db.QueryContext(cont, "SELECT id, generated FROM large_dataset")
    if err != nil {
        log.Println("Error executing query:", err)
        return
    }
    defer rows.Close()

    var natija []Row
    for rows.Next() {
        var rw Row
        err = rows.Scan(&rw.ID, &rw.Generated)
        if err != nil {
            log.Println("Error scanning rw:", err)
            continue
        }
        natija = append(natija, rw)
    }

    if err = rows.Err(); err != nil {
        log.Println("Error iterating rows:", err)
    }

    fmt.Println("Number of rows:", len(natija))
}
