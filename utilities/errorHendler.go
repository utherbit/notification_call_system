package utilities

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func PrintErr(err error) {
	log.Panicf("CRITICAL [%s] %v\n", returnPrettyDate(), err)
}

func LogStrErr(prefix string, text string) {
	log.Panicf("%s [%s] %s\n", prefix, returnPrettyDate(), text)
}

func LogStrInfo(text string) {
	fmt.Printf("INFO [%s] %s\n", returnPrettyDate(), text)
}
func LogStr(prefix, text string) {
	fmt.Printf("%s [%s] %s\n", prefix, returnPrettyDate(), text)
}

func PanicIfErr(err error) {
	if err != nil {
		PrintErr(err)
	}
}

func returnPrettyDate() string {
	return fmt.Sprintf("%d/%d/%d %d:%d:%d", time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second())
}

func PrintJSON(payload interface{}) {
	response, _ := json.Marshal(payload)

	fmt.Printf("%s\n", response)
}
