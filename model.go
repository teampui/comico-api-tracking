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

type Client struct {
	key    string
	secret string
	host   string
}

func NewClient(key, secret, host string) *Client {
	return &Client{
		key:    key,
		secret: secret,
		host:   host,
	}
}

func (c *Client) SendLog(track Tracking) {
	// 發起 request
	jsonValue, err := json.Marshal(track)
	if err != nil {
		fmt.Printf("[error] %v", err)
		return
	}

	req, err := http.NewRequest(http.MethodPost, c.host+"/api/keyv1/logs", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("[error] %v", err)
		return
	}

	today := time.Now().UTC().Add(8 * time.Hour).Format("20060102")

	signature := hmacSha256(c.secret, today)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Comico-Key", c.key)
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
