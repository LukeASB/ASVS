package db

import (
	"database/sql"
	"fmt"
	"secureCodingCourse/data"

	_ "github.com/denisenkom/go-mssqldb" // MS SQL driver
)

type DB struct {
	database *sql.DB
}

type IDB interface {
	UnsafeRetrievePatients(input any) (any, error)
	SafeRetrievePatients(input any) (any, error)
	RetrieveUsers(input any) (any, error)
	Close()
}

func NewDB(table string) (*DB, error) {
	database, err := Connect(table)

	if err != nil {
		return nil, err
	}

	return &DB{database: database}, nil
}

func Connect(table string) (*sql.DB, error) {
	// Connection string for SQL Server
	connString := fmt.Sprintf("sqlserver://localhost:1433?database=%s&trusted_connection=true", table)
	var err error
	// Open the connection
	database, err := sql.Open("sqlserver", connString)
	if err != nil {
		fmt.Println("Error creating connection pool: ", err)
		return nil, fmt.Errorf("error creating connection pool: %s", err)
	}

	// Ping to test connection
	err = database.Ping()
	if err != nil {
		fmt.Println("Cannot connect to database: ", err)
		return nil, fmt.Errorf("cannot connect to database: %s", err)
	}
	fmt.Println("Connected to SQL Server!")

	return database, nil
}

func (db *DB) UnsafeRetrievePatients(input any) (any, error) {
	rows, err := db.database.Query(fmt.Sprintf("SELECT [id], [name], [surname], [age], [gender] FROM [dbo].[Patients] WHERE age = %s", input))

	if err != nil {
		fmt.Println("Error creating connection pool: ", err)
		return nil, fmt.Errorf("error creating connection pool: %s", err)
	}

	defer rows.Close()

	var results []data.Patient

	for rows.Next() {
		var p data.Patient

		// Scan the row values into variables
		err := rows.Scan(&p.Id, &p.Name, &p.Surname, &p.Age, &p.Gender)
		if err != nil {
			fmt.Println("Error scanning row: ", err)
			return nil, fmt.Errorf("error scanning row: %s", err)
		}

		// Print each row
		fmt.Printf("ID: %d, Name: %s %s, Age: %d, Gender: %s\n", p.Id, p.Name, p.Surname, p.Age, p.Gender)
		results = append(results, p)
	}

	// Check for any errors during iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (db *DB) SafeRetrievePatients(input any) (any, error) {
	stmt, err := db.database.Prepare("SELECT [id], [name], [surname], [age], [gender] FROM [dbo].[Patients] WHERE age = @p1")

	if err != nil {
		fmt.Println("Error scanning row: ", err)
		return nil, fmt.Errorf("error scanning row: %s", err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(input)

	if err != nil {
		fmt.Println("Fail: ", err)
		return nil, fmt.Errorf("%s", err)
	}

	defer rows.Close()

	var results []data.Patient

	for rows.Next() {
		var p data.Patient

		// Scan the row values into variables
		err := rows.Scan(&p.Id, &p.Name, &p.Surname, &p.Age, &p.Gender)
		if err != nil {
			fmt.Println("Error scanning row: ", err)
			return nil, fmt.Errorf("error scanning row: %s", err)
		}

		// Print each row
		fmt.Printf("ID: %d, Name: %s %s, Age: %d, Gender: %s\n", p.Id, p.Name, p.Surname, p.Age, p.Gender)
		results = append(results, p)
	}

	// Check for any errors during iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (db *DB) RetrieveUsers(input any) (any, error) {
	stmt, err := db.database.Prepare("SELECT * FROM [dbo].[Users] WHERE username = @p1")

	if err != nil {
		fmt.Println("Error scanning row: ", err)
		return nil, fmt.Errorf("error scanning row: %s", err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(input)

	if err != nil {
		fmt.Println("Fail: ", err)
		return nil, fmt.Errorf("%s", err)
	}

	defer rows.Close()

	var results []data.User

	for rows.Next() {
		var u data.User

		// Scan the row values into variables
		err := rows.Scan(&u.Id, &u.UserName, &u.Password)
		if err != nil {
			fmt.Println("Error scanning row: ", err)
			return nil, fmt.Errorf("error scanning row: %s", err)
		}

		// Print each row
		fmt.Printf("ID: %d, Username: %s, Password: %s\n", u.Id, u.UserName, u.Password)
		results = append(results, u)
	}

	// Check for any errors during iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (db *DB) Close() {
	defer db.database.Close()
}
