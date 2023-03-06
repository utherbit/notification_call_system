package utilities

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadFile(url string, outPath string) error {
	fmt.Printf("\nDownloadFile %s", outPath)
	resp, err := http.Get(url)

	PanicIfErr(err)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		PanicIfErr(err)
	}(resp.Body)

	out, err := os.Create(outPath)

	PanicIfErr(err)

	defer func(out *os.File) {
		err := out.Close()
		PanicIfErr(err)
	}(out)

	_, err = io.Copy(out, resp.Body)

	PanicIfErr(err)

	return nil
}
