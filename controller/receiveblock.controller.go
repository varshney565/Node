package controller

import (
	"fmt"
	"node/model"

	"github.com/gofiber/fiber/v2"
)

func Receive(c *fiber.Ctx) error {
	gid := c.Get("group-id")
	pid := c.Get("param-id")
	aip := c.Get("admin-ip")
	blockhash := c.Get("block-hash")

	block := model.Block{}
	if err := c.BodyParser(&block); err != nil {
		fmt.Println("Error fething data of block :", err)
		return c.Status(400).JSON(fiber.Map{
			"Error": err,
		})
	}

	fmt.Printf("Group-id : %v\nParam-id : %v\nAdmin-IP : %v\nblockhash : %v\n", gid, pid, aip, blockhash)
	fmt.Printf("Block : %+v\n", block)
	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
	})
}
