package utilities

import (
	"fmt"
	"path/filepath"
	"strings"

	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

func ConvertToAlaw(inputFilePath string, outputDirectory string) error {
	fmt.Printf("\nConvertToAlawUlaw %s", outputDirectory)

	err := ffmpeg_go.Input(inputFilePath).
		Output(
			outputDirectory+strings.Split(filepath.Base(inputFilePath), ".")[0]+".alaw",
			ffmpeg_go.KwArgs{"af": "adelay=1000|1000", "f": "alaw", "ar": "8000"}).
		OverWriteOutput().
		Run()

	if err != nil {
		panic(err)
	}

	return nil
}
