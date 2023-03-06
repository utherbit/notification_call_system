package handlers

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"notificator/connections"
)

type dataHandlerCall struct {
	Id     int    `json:"id"`
	Text   string `json:"text"`
	Phone  string `json:"phone"`
	Status int    `json:"status"`
}

func HandlerNewCall(c *fiber.Ctx) error {
	fmt.Printf("\n\nHandlerNewCall")
	var request []dataHandlerCall
	err := c.BodyParser(&request)
	if err != nil {
		fmt.Printf("\nError parse body %s", err)
		return c.SendString("Error Data")
	}

	var result []dataHandlerCall
	for _, item := range request {
		if item.Id != 0 {
			result = append(result, item)
			continue
		}

		var resItem dataHandlerCall

		resItem.Text = item.Text
		resItem.Phone = item.Phone

		var audioNotExist bool
		pgErr := connections.PostgresConn.QueryRow(context.Background(),
			"select * from insert_call($1)", item.Text).Scan(&resItem.Id, &resItem.Status, &audioNotExist)
		if pgErr != nil {
			fmt.Printf("\nError select %s", pgErr)
			return c.SendString("Error select")
		}

		if audioNotExist {
			connections.SynthesizerAdd(connections.DataSynthesesAudio{
				Id:   resItem.Id,
				Text: resItem.Text})
		}

		result = append(result, resItem)
	}

	return c.JSON(result)
}
