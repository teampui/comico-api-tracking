package tracking

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Tracking struct {
	Referrer    string `json:"referrer"`
	Platform    string `json:"platform"`
	Event       string `json:"event"`
	EventSource string `json:"event_source,omitempty"`
	Object      string `json:"object,omitempty"`
	Uid1        string `json:"uid1"`
	Uid2        string `json:"uid2,omitempty"`
	IP          string `json:"ip"`
	Version     string `json:"version,omitempty"`
}

var (
	authorization string
	signature     string
	host          string
)

func init() {
	authorization = os.Getenv("TRACKING_AUTHORIZATION")
	if authorization == "" {
		panic("TRACKING_AUTHORIZATION is not set")
	}

	signature = os.Getenv("TRACKING_SIGNATURE")
	if signature == "" {
		panic("TRACKING_SIGNATURE is not set")
	}

	host = os.Getenv("TRACKING_HOST")
	if host == "" {
		panic("TRACKING_HOST is not set")
	}

}

func SendLog(track Tracking) {
	// 發起 request
	jsonValue, err := json.Marshal(track)
	if err != nil {
		fmt.Printf("[error] %v", err)
		return
	}

	req, err := http.NewRequest(http.MethodPost, host+"/api/v1/logs", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("[error] %v", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authorization)
	req.Header.Set("X-Comico-Signature", signature)

	// 逾時設定
	timeout := 30 * time.Second

	//adding the Transport object to the http Client
	client := http.Client{
		Timeout: timeout,
	}

	_, err = client.Do(req)
	if err != nil {
		fmt.Printf("[error] %v", err)
		return
	}

}
