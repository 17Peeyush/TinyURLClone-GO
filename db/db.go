package db

import (
	"database/sql"
	"fmt"

	//"github.com/17Peeyush/TinyURLClone-GO/models"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	//connStr := "postgres://username:password@localhost/db?sslmode=verify-full"
	connStr := "postgres://root:secret@localhost:5432/url_db?sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Connection to db established successfully...")
	//DB.SetMaxOpenConns(10)
	//DB.SetMaxIdleConns(5)
	createTables()
	
}

func createTables(){
	createUrlTable := `
	CREATE TABLE IF NOT EXISTS Url(
		id SERIAL PRIMARY KEY,
		hash_url VARCHAR(7) UNIQUE,
		link VARCHAR(500),
		created_at TIMESTAMP DEFAULT current_timestamp,
		expire_at TIMESTAMP DEFAULT current_timestamp + INTERVAL '7 days'
	)
	`
	_, err := DB.Exec(createUrlTable)
	if err != nil{
		fmt.Println(err.Error())
		panic("Could not create url table.")
	}else{
		fmt.Println("URL table created successfully.")
	}
}

// func InsertData(m models.Meet){
// 	fmt.Print("INSIDE INSERT")
// 	// query1 :=`
// 	// INSERT INTO Url(hash_url,link)
// 	// VALUES($1,$2) RETURNING id`
// 	query2 :=`
// 	INSERT INTO Url(hash_url,link)
// 	VALUES($1,$2) RETURNING id`

// 	/* Option 1
// 	stmt, err := DB.Prepare(query)
// 	if err != nil{
// 		//fmt.Println("Not able to insert data into table")
// 		fmt.Println(err.Error())
// 	}
// 	defer stmt.Close()
// 	_, err = stmt.Exec(m.UniqueID, m.Url)
// 	if err != nil{
// 		fmt.Println( "Not able to insert data into table")
// 		fmt.Println(err.Error())
// 	}*/
// 	lastInsertId := 0
// 	err := DB.QueryRow(query2, m.UniqueID, m.Url).Scan(&lastInsertId)
// 	if err != nil{
// 		fmt.Println( "Not able to insert data into table")
// 		fmt.Println(err.Error())
// 	}
// 	fmt.Println("LastinsertedID: ", lastInsertId)
// }