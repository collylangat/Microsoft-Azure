// Go connection Sample Code:
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/microsoft/go-mssqldb"
)

var db *sql.DB
var server = "database.windows.net"
var port = 1433
var user = ""
var password = ""
var database = ""

func main() {
	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!")
}

/*

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/microsoft/go-mssqldb"
)

func main() {
	// Replace with your Azure SQL details
	server := "sqlserver-todolistngenii.database.windows.net"
	port := 1433
	user := "ngenii"
	password := "Kipngeno33"
	database := "rg-todolist"
	// Build connection string
	""
	//connString := fmt.Sprintf("server=%s;userid=%s;password=%s;port=%d;database=%s", server, user, password, port, database)
	// Open connection
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool:", err.Error())
	}
	defer db.Close()
	// Test connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Ping failed:", err.Error())
	}
	fmt.Println("✅ Connected to Azure SQL Database!")
	// Example query
	rows, err := db.Query("SELECT TOP 3 name FROM sys.databases")
	if err != nil {
		log.Fatal("Query failed:", err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var dbName string
		err := rows.Scan(&dbName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Database:", dbName)
	}
}
*/
