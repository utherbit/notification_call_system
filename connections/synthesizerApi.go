package connections

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"notificator/utilities"
	"os"
	"strconv"
	"time"
)

var (
	SynthesizerApi    []DataSynthesesAudio
	UrlSynthesizerApi string
)

func InitSynthesizerApi() {
	url, err := os.LookupEnv("URL_SYNTHESIZER_API")
	if !err {
		UrlSynthesizerApi = "http://192.168.25.185:8899"
	}
	UrlSynthesizerApi = url

	go SynthesesApi()
}

func SynthesizerAdd(data DataSynthesesAudio) {
	SynthesizerApi = append(SynthesizerApi, data)
}

type synthsPostData struct {
	Text   string `json:"text"`
	Voice  string `json:"voice"`
	Rate   string `json:"rate"`
	Pitch  string `json:"pitch"`
	Volume string `json:"volume"`
}

func newSynthsPostData(text string) synthsPostData {
	data := synthsPostData{}
	data.Text = text
	data.Voice = "Natasha"
	data.Rate = "0.8"
	data.Pitch = "1"
	data.Volume = "0"
	return data
}

type synthsResData struct {
	Response []struct {
		DurationS        float64 `json:"duration_s"`
		ResponseAudio    string  `json:"response_audio"`
		ResponseAudioUrl string  `json:"response_audio_url"`
		SampleRate       int     `json:"sample_rate"`
		SynthesisTime    float64 `json:"synthesis_time"`
		Voice            string  `json:"voice"`
	} `json:"response"`
	ResponseCode int `json:"response_code"`
}

type DataSynthesesAudio struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
}

func SynthesesApi() {
	for {
		time.Sleep(1000)
		lenS := len(SynthesizerApi)
		if lenS != 0 {
			item := SynthesizerApi[0]
			if lenS > 0 {
				SynthesizerApi = SynthesizerApi[1:lenS]
			} else {
				SynthesizerApi = []DataSynthesesAudio{}
			}
			fmt.Printf("\nSyntheses id: %d, text: '%s'", item.Id, item.Text)
			PostSyntheses(item)
		}
	}
}

func PostSyntheses(dataItem DataSynthesesAudio) {
	body := newSynthsPostData(dataItem.Text)
	jsonData, err := json.Marshal(body)
	if err != nil {
		fmt.Printf("\nError Marshal post: %s", err)
	}

	resp, err := http.Post(UrlSynthesizerApi+"/synthesize", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("\nError http post: %s\n\n\n", err)
	}

	var synthsRes synthsResData
	err = json.NewDecoder(resp.Body).Decode(&synthsRes)
	if err != nil {
		fmt.Printf("\nError Decode: %s", resp.Body)
	}

	audioUrl := UrlSynthesizerApi + synthsRes.Response[0].ResponseAudioUrl
	DurationS := int(synthsRes.Response[0].DurationS + 1)

	var audioId int
	err = PostgresConn2.QueryRow(context.Background(), "select insert_audio($1, $2, $3)", dataItem.Id, audioUrl, DurationS+2).Scan(&audioId)
	if err != nil {
		fmt.Printf("\nPg Error: %s", err)
	}

	fileName := "dial_" + strconv.Itoa(audioId) + ".wav"

	f, err := os.OpenFile(PathTempOutput+fileName, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		panic(err)
	}

	d, err := base64.StdEncoding.DecodeString(synthsRes.Response[0].ResponseAudio)

	r := bytes.NewReader(d)

	_, err = io.Copy(f, r)

	if err != nil {
		panic(err)
	}

	err = utilities.ConvertToAlaw(PathTempOutput+fileName, PathOutputMoh)
	if err != nil {
		fmt.Printf("\nERROR ConvertToAlawUlaw%s", err)
	}
}

//func SynthesesApi(request []DataSynthesesAudio) {
//	fmt.Printf("\n\nfunc PostSynthesesApi")
//	for _, item := range request {
//		fmt.Printf("\n	synthes audio id: %d text: %s process: ", item.Id, item.Text)
//		PostSyntheses(item)
//	}
//}
