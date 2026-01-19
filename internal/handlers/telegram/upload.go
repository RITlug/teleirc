package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
)

// MediaShareResponse represents the response from MediaShare upload endpoint.
type MediaShareResponse struct {
	Success  bool   `json:"success"`
	URL      string `json:"url"`
	RawURL   string `json:"raw_url"`
	ID       string `json:"id"`
	Filename string `json:"filename"`
	Error    string `json:"error,omitempty"`
}

// uploadToMediaShare uploads a file from a URL to MediaShare service.
func uploadToMediaShare(tg *Client, fileURL string, filename string, username string) string {
	if tg.MediaShareSettings == nil || !tg.MediaShareSettings.Enabled {
		return ""
	}

	if tg.MediaShareSettings.Endpoint == "" {
		tg.logger.LogError("MediaShare endpoint not configured")
		return ""
	}

	// Download file from Telegram
	resp, err := http.Get(fileURL)
	if err != nil {
		tg.logger.LogError("Failed to download file from Telegram:", err)
		return ""
	}
	defer resp.Body.Close()

	// Read the file content
	fileContent, err := io.ReadAll(resp.Body)
	if err != nil {
		tg.logger.LogError("Failed to read file content:", err)
		return ""
	}

	// Create multipart form
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		tg.logger.LogError("Failed to create form file:", err)
		return ""
	}

	if _, err := part.Write(fileContent); err != nil {
		tg.logger.LogError("Failed to write file to form:", err)
		return ""
	}

	// Add username field
	if username != "" {
		if err := writer.WriteField("username", username); err != nil {
			tg.logger.LogError("Failed to write username field:", err)
			return ""
		}
	}

	if err := writer.Close(); err != nil {
		tg.logger.LogError("Failed to close multipart writer:", err)
		return ""
	}

	// Send to MediaShare
	endpoint := strings.TrimSuffix(tg.MediaShareSettings.Endpoint, "/") + "/upload"
	req, err := http.NewRequest("POST", endpoint, body)
	if err != nil {
		tg.logger.LogError("Failed to create request:", err)
		return ""
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	if tg.MediaShareSettings.APIKey != "" {
		req.Header.Set("X-API-Key", tg.MediaShareSettings.APIKey)
	}

	client := &http.Client{}
	uploadResp, err := client.Do(req)
	if err != nil {
		tg.logger.LogError("Failed to upload to MediaShare:", err)
		return ""
	}
	defer uploadResp.Body.Close()

	// Parse response
	var result MediaShareResponse
	if err := json.NewDecoder(uploadResp.Body).Decode(&result); err != nil {
		tg.logger.LogError("Failed to parse MediaShare response:", err)
		return ""
	}

	if !result.Success {
		tg.logger.LogError("MediaShare upload failed:", result.Error)
		return ""
	}

	tg.logger.LogDebug("Uploaded to MediaShare:", result.URL)
	return result.URL
}

// uploadFile uploads a file to MediaShare (or falls back to Imgur for images).
// Returns the URL of the uploaded file.
func uploadFile(tg *Client, fileID string, filename string, username string) string {
	// Check if API client is available
	if tg.api == nil {
		return ""
	}

	// Get Telegram file URL
	fileURL, err := tg.api.GetFileDirectURL(fileID)
	if err != nil {
		tg.logger.LogError("Failed to get Telegram file URL:", err)
		return ""
	}

	// Try MediaShare first if enabled
	if tg.MediaShareSettings != nil && tg.MediaShareSettings.Enabled {
		if url := uploadToMediaShare(tg, fileURL, filename, username); url != "" {
			return url
		}
	}

	// Fall back to Imgur for images only
	ext := strings.ToLower(filepath.Ext(filename))
	if ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif" || ext == ".webp" {
		return getImgurLink(tg, fileURL)
	}

	return ""
}

// uploadVideo uploads a video file and returns the URL.
func uploadVideo(tg *Client, fileID string, filename string, username string) string {
	if filename == "" {
		filename = "video.mp4"
	}
	return uploadFile(tg, fileID, filename, username)
}

// uploadVoice uploads a voice message and returns the URL.
func uploadVoice(tg *Client, fileID string, username string) string {
	return uploadFile(tg, fileID, "voice.ogg", username)
}

// uploadPhoto uploads a photo and returns the URL.
// Uses the existing Imgur integration as fallback.
func uploadPhoto(tg *Client, fileID string, username string) string {
	// Check if API client is available
	if tg.api == nil {
		return ""
	}

	// Get Telegram file URL
	fileURL, err := tg.api.GetFileDirectURL(fileID)
	if err != nil {
		tg.logger.LogError("Failed to get Telegram photo URL:", err)
		return ""
	}

	// Try MediaShare first if enabled
	if tg.MediaShareSettings != nil && tg.MediaShareSettings.Enabled {
		if url := uploadToMediaShare(tg, fileURL, "photo.jpg", username); url != "" {
			return url
		}
	}

	// Fall back to Imgur
	return getImgurLink(tg, fileURL)
}

// formatMediaMessage formats a message for media sharing.
func formatMediaMessage(username string, mediaType string, caption string, url string) string {
	if url == "" {
		formatted := username + " shared a " + mediaType + " on Telegram"
		if caption != "" {
			formatted += " with caption: '" + caption + "'."
		} else {
			formatted += "."
		}
		return formatted
	}

	if caption != "" {
		return fmt.Sprintf("'%s' %s by %s: %s", caption, mediaType, username, url)
	}
	return fmt.Sprintf("%s shared by %s: %s", mediaType, username, url)
}
