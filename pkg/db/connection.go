package db

func DatabaseOpen() Database {
	var database Database
	database = &PostgreSQL{}
	dataSourceName := "host=localhost port=5432 user=postgres password=94508443r dbname=ToDoApp sslmode=disable"

	database.InitDB(dataSourceName)
	database.CreateTables()
	return database

}
