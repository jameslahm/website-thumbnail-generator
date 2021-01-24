package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

type thumbnailRequest struct {
	Url string `json:"url"`
}

type screenshotApiRequest struct {
	Token          string `json:"token"`
	Url            string `json:"url"`
	Output         string `json:"output"`
	Width          int    `json:"width"`
	Height         int    `json:"height"`
	ThumbnailWidth int    `json:"thumbnail_width"`
}

func thumbnailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	var decoded thumbnailRequest

	err := json.NewDecoder(r.Body).Decode(&decoded)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	apiRequest := screenshotApiRequest{
		Token:          "VZB3CGJHEAQC93GVWAHRNYICHB843HGE",
		Url:            decoded.Url,
		Output:         "json",
		Width:          1920,
		Height:         1080,
		ThumbnailWidth: 300,
	}

	jsonString, err := json.Marshal(apiRequest)
	checkError(err)
	req, err := http.NewRequest("POST", "https://screenshotapi.net/api/v1/screenshot", bytes.NewBuffer(jsonString))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	checkError(err)

	defer response.Body.Close()

	type screenshotAPIResponse struct {
		Screenshot string `json:"screenshot"`
	}

	var apiResponse screenshotAPIResponse
	err = json.NewDecoder(response.Body).Decode(&apiResponse)
	checkError(err)

	_, err = fmt.Fprintf(w, `{"screenshot":"%s"}`, apiResponse.Screenshot)
	checkError(err)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/thumbnail", thumbnailHandler)
	handler := cors.Default().Handler(mux)
	fmt.Println("Server listen on port 3000")
	log.Panic(http.ListenAndServe(":3000", handler))
}
