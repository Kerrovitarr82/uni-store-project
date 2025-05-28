package main

import "uniStore/internal/database"

func main() {
	database.ConnectToDB()
	database.MigrateDB()
	database.CheckRoles()
}
