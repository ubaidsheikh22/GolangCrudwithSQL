package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var server = "localhost"
var port = 1433
var user = "sa"
var pass = "ubaidbm"
var database = "Golangdb"
var db *sql.DB

func main() {
	//fmt.Println("CRUD in Golang with MSSQL")
	var err error

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, pass, port, database)
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error Connecting SQL server " + err.Error())
	}
	log.Printf("Connected to Sql server")

	// //create User in database
	// userid, err := CreateUser("Alex", "Software Engineer")
	// fmt.Println("New User Created with ID %d", userid)

	//Updating User in Database//
	// userid, err := UpdateUser("ubaid", "Software Developer")
	// fmt.Println("User Details has been Updated ", userid)

	//Delete User Details from Databse //
	// userid, err := Deleteuser("Ubaid")
	// fmt.Println("User Details has been deleted", userid)
	// // get Users Records from Database//
	// count, err := Getusers()
	// fmt.Println("Get Users Record Count ", count)
}

///Create User Function to insert User data in db table//
func CreateUser(name string, position string) (int64, error) {
	ctx := context.Background()
	var err error
	if db == nil {
		log.Fatal("Db not alive")
	}
	//check db is alive through ping//
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal("Error Pinging Database", err.Error())
	}
	tsql := fmt.Sprintf("insert into tbluser(Name,Position) values (@Name,@Position);")

	result, err := db.ExecContext(
		ctx,
		tsql,
		sql.Named("Name", name),
		sql.Named("Position", position))
	if err != nil {
		log.Fatal("Error inserting Data", err.Error())
		return -1, err
	}
	return result.LastInsertId()
}

//function to get Users information from database//
func Getusers() (int64, error) {
	ctx := context.Background()
	var err error
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal("Error Connecting Database ", err.Error())
	}
	tsql := fmt.Sprintf("select * from tbluser;")
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		log.Fatal("Error Reading Data", err.Error())
		return -1, err
	}
	defer rows.Close()
	var count int = 0
	for rows.Next() {
		var name, position string
		var id int

		err := rows.Scan(&id, &name, &position)
		{
			if err != nil {
				log.Fatal("Error Occured", err.Error())
				return -1, err
			}
		}
		fmt.Println(id, name, position)
		count++
	}
	return int64(count), nil

}

//Delete User Record in Database//
func Deleteuser(name string) (int64, error) {
	ctx := context.Background()
	var err error
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal("Error Connecting Databse", err.Error())
	}
	tsql := fmt.Sprintf("delete from tbluser where name=@Name;")
	result, err := db.ExecContext(
		ctx,
		tsql,
		sql.Named("Name", name))
	if err != nil {
		log.Fatal("Error Deleting User Delete", err.Error())
		return -1, err
	}
	return result.LastInsertId()

}

//For Updating User in Database//

func UpdateUser(name string, position string) (int64, error) {
	ctx := context.Background()
	var err error
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal("Error Pinging Database", err.Error())
	}
	tsql := fmt.Sprintf("update tbluser set position=@Position where name=@Name;")
	result, err := db.ExecContext(
		ctx,
		tsql,
		sql.Named("Name", name),
		sql.Named("Position", position))
	if err != nil {
		log.Fatal("Error Updating Values in Database", err.Error())
		return -1, err
	}
	return result.LastInsertId()

}
