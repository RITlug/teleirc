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

type accessToken struct {
	AccountID       int    `json:"account_id"`
	AccountUsername string `json:"account_username"`
	RefreshToken    string `json:"refresh_token"`
	AccessToken     string `json:"access_token"`
	TokenType       string `json:"token_type"`
	ExpiresIn       int    `json:"expires_in"`
}

func getImgurLink(tg *Client, tgLink string) string {
	url := "https://api.imgur.com/3/image"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("image", tgLink)
	if tg.ImgurSettings.ImgurAlbumHash != "" {
		_ = writer.WriteField("album", tg.ImgurSettings.ImgurAlbumHash)
	}
	err := writer.Close()
	if err != nil {
		tg.logger.LogError("Could not close Writer:", err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		tg.logger.LogError("Could not build HTTP request:", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	if tg.ImgurSettings.ImgurRefreshToken != "" {
		if tg.ImgurSettings.ImgurAccessToken == "" {
			getImgurAccessToken(tg)
		}
		req.Header.Add("Authorization", "Bearer "+tg.ImgurSettings.ImgurAccessToken)
	} else {
		clientID := tg.ImgurSettings.ImgurClientID
		req.Header.Add("Authorization", "Client-ID "+clientID)
	}

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

    if resp.Data == nil {
        tg.logger.LogError("Imgur response data was null. Possible API rate limiting.")
        return ""
    }

	if resp.Data.Deletehash != "" {
		deleteLink := "https://imgur.com/delete/" + resp.Data.Deletehash
		tg.logger.LogInfo("Deletion link for", resp.Data.Link, "is", deleteLink)
	}

	return resp.Data.Link
}

func getImgurAccessToken(tg *Client) {
	if tg.ImgurSettings.ImgurClientID == "" || tg.ImgurSettings.ImgurRefreshToken == "" {
		tg.logger.LogError("Imgur client secret and refresh token must be set")
		return
	}

	url := "https://api.imgur.com/oauth2/token"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("refresh_token", tg.ImgurSettings.ImgurRefreshToken)
	_ = writer.WriteField("client_id", tg.ImgurSettings.ImgurClientID)
	_ = writer.WriteField("client_secret", tg.ImgurSettings.ImgurClientSecret)
	_ = writer.WriteField("grant_type", "refresh_token")
	err := writer.Close()
	if err != nil {
		tg.logger.LogError("Could not close Writer:", err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		tg.logger.LogError("Could not build HTTP request:", err)
		return
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		tg.logger.LogError("Could not send request:", err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		tg.logger.LogError("Could not read response:", err)
		return
	}

	var data accessToken
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		tg.logger.LogError("Couldn't unmarshal json:", err)
		return
	}

	tg.ImgurSettings.ImgurAccessToken = data.AccessToken
}
