package controllers

import (
	"log"
	"strconv"
	"time"

	"github.com/RyotaKITA-12/locatask.git/database"
	"github.com/RyotaKITA-12/locatask.git/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	geo "github.com/martinlindhe/google-geolocate"
)

func RegisterPost(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		log.Fatalln(err)
	}

	user_id, _ := strconv.Atoi(data["user_id"])
	client := geo.NewGoogleGeo("AIzaSyCb7FqJb6W_n39blFQMfna9u1wUcDD-ABk")
	res, _ := client.Geocode(data["address"])
	// longitude, _ := strconv.ParseFloat(data["longitude"], 64)
	// latitude, _ := strconv.ParseFloat(data["latitude"], 64)
	period, _ := time.Parse("2006-01-02", data["period"])

	post := models.Post{
		UserID:    user_id,
		Title:     data["title"],
		Period:    period,
		Address:   data["address"],
		Longitude: res.Lng,
		Latitude:  res.Lat,
	}

	database.DB.Create(&post)

	return c.JSON(post)
}

func Post(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	log.Println(cookie)
	token, err := jwt.ParseWithClaims(cookie, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil || !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Un autenticated."})
	}
	claims := token.Claims.(*Claims)
	id := claims.Issuer

	var posts []models.Post
	database.DB.Where("user_id = ?", id).Find(&posts)

	return c.JSON(posts)
}
