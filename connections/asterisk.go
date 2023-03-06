package connections

import (
	"fmt"
	"notificator/utilities"
	"os"
	"strconv"
)

var (
	PathAsteriskOutgoingCall string
	PathTempOutput           string
	PathOutputMoh            string
	CallExt                  string
)

func InitAsterisk() {
	envPathAsteriskOutgoingCall, errb := os.LookupEnv("PATH_ASTERISK_OUTGOING_CALL")
	if !errb {
		PathAsteriskOutgoingCall = "C:\\asterisk\\"
	}
	PathAsteriskOutgoingCall = envPathAsteriskOutgoingCall

	envTempOutputPath, errb := os.LookupEnv("PATH_TEMP_OUTPUT_WAV")
	if !errb {
		envTempOutputPath = "C:\\asterisk\\temp\\"
	}
	PathTempOutput = envTempOutputPath

	envOutputPathMoh, errb := os.LookupEnv("PATH_OUTPUT_MOH")
	if !errb {
		envOutputPathMoh = "C:\\asterisk\\moh\\"
	}
	PathOutputMoh = envOutputPathMoh

	envCallExt, errb := os.LookupEnv("CALL_EXTENSION")
	if !errb {
		envCallExt = ""
	}
	CallExt = envCallExt

}

type DialModel struct {
	Phone    string `json:"phone"`
	CallId   int    `json:"callId"`
	RecordId int    `json:"recordId"`
	Duration int    `json:"duration"`
}

func DialAsterisk(info DialModel) {
	fmt.Printf("\n\nDialAsterisk ")
	callFile := "dial-" + info.Phone + "-" + strconv.Itoa(info.CallId) + ".call"

	filename := PathAsteriskOutgoingCall + callFile
	fmt.Printf("\nCall file: %s", filename)
	f, err := os.Create(filename)
	utilities.PanicIfErr(err)

	defer func(f *os.File) {
		err := f.Close()
		utilities.PanicIfErr(err)
	}(f)

	text := fmt.Sprintf(
		`Channel: Local/%s*%d*%d*%d@%s-out-robot
MaxRetries: 0
RetryTime: 0
WaitTime: 30
Context: test
Extension: checkphone
`, info.Phone, info.CallId, info.RecordId, info.Duration, CallExt)

	_, err = f.WriteString(text)
	utilities.PanicIfErr(err)

	println(text)
}
