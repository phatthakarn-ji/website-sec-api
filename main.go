package main

import (
	"fmt"
	"log"

	"github.com/phatthakarn-ji/website-sec-api/database"
	"github.com/phatthakarn-ji/website-sec-api/models"
	"github.com/phatthakarn-ji/website-sec-api/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func initDatabase() {
	var err error
	database.DB, err = gorm.Open("sqlite3", "database.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	contact := new(models.Contact)
	database.DB.Find(&contact, 1)
	if contact.Company == "" {
		database.DB.AutoMigrate(&models.Contact{})
		// Create
		database.DB.Create(&models.Contact{Company: "บริษัทหลักทรัพย์ เคทีบีเอสที จำกัด (มหาชน)", TelText: "02 351 1800 กด 1", TelValue: "023511800", Email: "customerservice@ktbst.co.th", Address: "เลขที่ 87/2 อาคารซีอาร์ซี ทาวเวอร์ ออลซีซั่นส์ เพลส ชั้นที่ 9, 18, 39 และ 52 ถนนวิทยุ แขวงลุมพินี เขตปทุมวัน กรุงเทพฯ 10330"})
	}
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	initDatabase()
	defer database.DB.Close()

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":5000"))
}
