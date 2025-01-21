package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()

	type GeoLocation struct {
		Longitude float64 `json:"longitude"`
		Latitude  float64 `json:"latitude"`
	}

	type Address struct {
		Street  string `json:"street"`
		Area    string `json:"area"`
		Pincode string `json:"pincode"`
		City    string `json:"city"`
		State   string `json:"state"`
		Country string `json:"country"`
	}

	type Parking struct {
		Available bool `json:"available"`
	}

	type Enquiry struct {
		UserID    string    `json:"userId"`
		UserName  string    `json:"userName"`
		ListingID string    `json:"listingId"`
		Number    string    `json:"number"`
		CreatedAt time.Time `json:"createdAt"`
	}

	type Listing struct {
		ID               string      `json:"id"`
		Title            string      `json:"title"`
		Description      string      `json:"description"`
		Type             string      `json:"type"`
		Location         GeoLocation `json:"location"`
		ReadyToMove      bool        `json:"readyToMove"`
		GatedCommunity   bool        `json:"gatedCommunity"`
		LocationFeatures []string    `json:"locationFeatures"`
		Amenities        []string    `json:"amenities"`
		Price            int         `json:"price"`
		Phone            string      `json:"phone"`
		Address          Address     `json:"address"`
		LandArea         int         `json:"landArea"`
		Floors           int         `json:"floors"`
		DTCPApproved     bool        `json:"dtcpApproved"`
		Facing           string      `json:"facing"`
		Parking          Parking     `json:"parking"`
		Images           []string    `json:"images"`
		Tags             []string    `json:"tags"`
		CreatedAt        time.Time   `json:"createdAt"`
		UpdatedAt        time.Time   `json:"updatedAt"`
		Views            int         `json:"views"`
		Contacted        int         `json:"contacted"`
		Status           string      `json:"status"`
		Enquiries        []Enquiry   `json:"enquiries"`
	}

	jsonFile, err := os.Open("data.json")

	byteValue, _ := io.ReadAll(jsonFile)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	var listings []Listing

	json.Unmarshal(byteValue, &listings)

	// Define a simple route
	app.Get("/api/go", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, World! Go app is running!",
			"data":    listings,
		})
	})

	// CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	// Start the server
	port := ":4000"
	if err := app.Listen(port); err != nil {
		panic(err)

	}

}
