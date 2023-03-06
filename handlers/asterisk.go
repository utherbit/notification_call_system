package handlers

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"notificator/connections"
	"notificator/utilities"
)

type AsteriskResultModel struct {
	CallId int `json:"call_id"`
	Status int `json:"status"`
}

func HandlerAsteriskResult(c *fiber.Ctx) error {
	fmt.Printf("\n\nHandlerAsteriskResult")
	var request AsteriskResultModel
	bpErr := c.BodyParser(&request)
	utilities.PanicIfErr(bpErr)

	// Status 5 фейл
	// Status 4 успех

	_, err := connections.PostgresConn.Exec(context.Background(),
		"update calls set status = $2 where id = $1", request.CallId, request.Status)
	if err != nil {
		fmt.Printf("\nERROR %s", err)
	}

	return c.SendString("Success")
}
