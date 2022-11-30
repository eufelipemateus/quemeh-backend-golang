package controllers

import (
	//"fmt"
	"log"
	//"github.com/eufelipemateus/quemeh-backend-golang/configs"
	"github.com/gofiber/fiber/v2"
	"github.com/undiabler/golang-whois"
)


type whoisType struct {
	Domain string `json:"domain" xml:"domain" form:"domain"`
}

func Whois (c *fiber.Ctx) error{

	var body whoisType

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	result, err := whois.GetWhois(body.Domain)
	if err != nil {
		log.Printf("Error in whois lookup : %v \n", err)
		c.SendStatus(fiber.StatusBadRequest)

		return 	c.JSON(&fiber.Map{
			"status": "error",
			"message":  "Error in whois lookup!",
		})
	} 

	return 	c.JSON(&fiber.Map{
        "status": "success",
        "data":  result,
    })

}
