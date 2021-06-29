package telegram

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

type imgurDataWrapper struct {
	Data    *imageData `json:"data"`
	Success bool       `json:"success"`
	Status  int        `json:"status"`
}

type imageData struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Datetime    int    `json:"datetime"`
	MimeType    string `json:"type"`
	Animated    bool   `json:"animated"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	Size        int    `json:"size"`
	Views       int    `json:"views"`
	Bandwidth   int    `json:"bandwidth"`
	Deletehash  string `json:"deletehash,omitempty"`
	Name        string `json:"name,omitempty"`
	Section     string `json:"section"`
	Link        string `json:"link"`
	Gifv        string `json:"gifv,omitempty"`
	Mp4         string `json:"mp4,omitempty"`
	Mp4Size     int    `json:"mp4_size,omitempty"`
	Looping     bool   `json:"looping,omitempty"`
	Favorite    bool   `json:"favorite"`
	Nsfw        bool   `json:"nsfw"`
	Vote        string `json:"vote"`
	InGallery   bool   `json:"in_gallery"`
}

func getImgurLink(tg *Client, tgLink string) string {
	url := "https://api.imgur.com/3/image"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("image", tgLink)
	err := writer.Close()
	if err != nil {
		tg.logger.LogError("Could not close Writer:", err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		tg.logger.LogError("Could not build HTTP request:", err)
	}
	clientID := tg.ImgurSettings.ImgurClientID
	req.Header.Add("Authorization", "Client-ID "+clientID)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	result, clientErr := client.Do(req)
	if err != nil {
		tg.logger.LogError("Could not send imgur request:", clientErr)
	}
	defer result.Body.Close()

	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		tg.logger.LogError("Could not read imgur response:", err)
	}

	var resp imgurDataWrapper
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		tg.logger.LogError("Could not retrieve imgur json data:", err)
	}

	return resp.Data.Link
}
