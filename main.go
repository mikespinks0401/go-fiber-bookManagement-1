package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/mikespinks0401/go-fiber-crm/lead"
	"github.com/mikespinks0401/go-fiber-crm/database"
	"github.com/jinzhu/gorm"
	"log"
)

func setupRoutes(app *fiber.App){
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id",lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase(){
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil{
		panic("Failed to connect database")
	}
	fmt.Println("Connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main(){
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	err := app.Listen(3000)
	if err != nil {
		log.Fatal("Failed to start app on port 3000")
	}
	defer database.DBConn.Close()
}