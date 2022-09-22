package tracking

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
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
	key    string
	secret string
	host   string
)

func init() {
	key = os.Getenv("TRACKING_KEY")
	if key == "" {
		panic("TRACKING_KEY is not set")
	}

	secret = os.Getenv("TRACKING_SECRET")
	if secret == "" {
		panic("TRACKING_SECRET is not set")
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
	req.Header.Set("KEY", key)
	req.Header.Set("X-Comico-Signature", Signature())

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

func Signature() string {
	hmacNew := hmac.New(sha256.New, []byte(key))
	hmacNew.Write([]byte(secret))
	return fmt.Sprintf("%x", hmacNew.Sum(nil))
}

func CheckSignature(signature string) bool {
	return hmac.Equal([]byte(signature), []byte(Signature()))
}
