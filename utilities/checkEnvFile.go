package utilities

import (
	"github.com/joho/godotenv"
	"os"
)

func CheckEnvFile() {
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
	InitCheckPaths()
}

var (
	PathFfmpegExe string
)

func InitCheckPaths() {
	EnvPathAsteriskOutgoingCall, errb := os.LookupEnv("PATH_FFMPEG")
	if !errb {
		PathFfmpegExe = ""
	}
	PathFfmpegExe = EnvPathAsteriskOutgoingCall
}
