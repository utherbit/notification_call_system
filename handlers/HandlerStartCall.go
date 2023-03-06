package handlers

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"notificator/connections"
	"notificator/utilities"
	"os"
	"strconv"
	"time"
)

type reqHandlerStartCall struct {
	Id    int    `json:"id"`
	Phone string `json:"phone"`
}

func HandlerStartCall(c *fiber.Ctx) error {
	fmt.Printf("\n\nHandlerStartCall")

	var request []reqHandlerStartCall
	bpErr := c.BodyParser(&request)
	utilities.PanicIfErr(bpErr)

	for _, itemRequest := range request {
		var audioUrl string
		var audioId int
		var duration int
		pgErr := connections.PostgresConn.QueryRow(context.Background(),
			"select * from at_the_start_of_a_call($1)", itemRequest.Id).Scan(
			&audioUrl, &audioId, &duration)
		if pgErr != nil {
			fmt.Printf("\nERROR %s", pgErr)
			continue
		}

		file := "dial_" + strconv.Itoa(audioId) + ".wav"

		fmt.Printf("\nbefor if")
		if _, err := os.Stat(connections.PathTempOutput + file); os.IsNotExist(err) {
			err := utilities.DownloadFile(audioUrl, connections.PathTempOutput+file)
			if err != nil {
				fmt.Printf("\nERROR: %s", err)
				return err
			}
			d, _ := time.ParseDuration("1s")
			time.Sleep(d)
			err = utilities.ConvertToAlaw(connections.PathTempOutput+file, connections.PathOutputMoh)
			if err != nil {
				fmt.Printf("\nERROR %s", err)
				return err
			}
		}

		time.Sleep(1)
		connections.DialAsterisk(connections.DialModel{
			Phone:    itemRequest.Phone,
			CallId:   itemRequest.Id,
			RecordId: audioId,
			Duration: duration})
	}

	return c.SendString("Success")
}
