package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"notificator/connections"
	"strconv"
)

type resCheckCall struct {
	Id     int `json:"id"`
	Status int `json:"status"`
}

func HandlerCheckStatusCall(c *fiber.Ctx) error {
	//fmt.Printf("\n\nHandlerCheckCall")
	var request []int
	err := json.Unmarshal(c.Body(), &request)
	if err != nil {
		fmt.Printf("\nError parse body %s", err)
		return c.SendString("Error Data")
	}

	if len(request) == 0 {
		return nil
	}

	constr := ""
	for _, i := range request {
		constr = constr + strconv.Itoa(i) + ","
	}
	constr = "{" + constr[:len(constr)-1] + "}"
	//fmt.Printf("\n	items: %s", constr)

	pgRes, err := connections.PostgresConn.Query(context.Background(),
		"select id,status from calls where id = any($1)", constr)
	if err != nil {
		fmt.Printf("\nError select %s", err)
		return c.SendString("Error select")
	}

	var result []resCheckCall
	for pgRes.Next() {
		var item resCheckCall
		err = pgRes.Scan(&item.Id, &item.Status)
		if err != nil {
			fmt.Printf("\nError Scan %s", err)
			return c.SendString("Error Scan")
		}

		result = append(result, item)
		//fmt.Printf("\n	check item id: %d, status: %d", item.Id, item.Status)
	}

	return c.JSON(result)
}
