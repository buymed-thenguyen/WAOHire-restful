package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type WSClient struct {
	BaseURL string
	Client  *http.Client
}

func NewWSClient(host, port string) *WSClient {
	if port != "" {
		port = ":" + port
	}
	return &WSClient{
		BaseURL: fmt.Sprintf("%s%s", host, port),
		Client:  &http.Client{},
	}
}

func (c *WSClient) Post(path string, payload interface{}) error {
	body, err := json.Marshal(payload)
	if err != nil {
		log.Printf("marshal error: %w", err)
		return err
	}

	url := fmt.Sprintf("%s%s", c.BaseURL, path)
	log.Printf("[POST] %s", url)
	resp, err := c.Client.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("post failed: %w", err)
		return err
	}
	defer resp.Body.Close()
	log.Printf("✅ HTTP POST %s → status: %s", url, resp.Status)

	return nil
}
