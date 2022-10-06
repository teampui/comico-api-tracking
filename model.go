package tracking

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
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
	key = "YOUR_KEY"
	secret = "YOUR_SECRET"
	host = "YOUR_HOST"
	//key = os.Getenv("API_KEY")
	//if key == "" {
	//	panic("API_KEY is not set")
	//}
	//
	//secret = os.Getenv("API_SECRET")
	//if secret == "" {
	//	panic("API_SECRET is not set")
	//}
	//
	//host = os.Getenv("API_HOST")
	//if host == "" {
	//	panic("API_HOST is not set")
	//}
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

	today := time.Now().UTC().Add(8 * time.Hour).Format("20060102")

	signature := hmacSha256(secret, today)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Comico-Key", key)
	req.Header.Set("X-Comico-Date", today)
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

func hmacSha256(secret, data string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(data))
	return hex.EncodeToString(mac.Sum(nil))
}
