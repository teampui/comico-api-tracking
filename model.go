package tracking

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Tracking struct {
	ID          int64
	Referrer    string
	Platform    string
	Event       string
	EventSource string
	Object      string
	Uid1        string
	Uid2        string
	IP          string
	Version     string
	CreatedAt   time.Time
}

var (
	authorization string
	signature     string
	host          string

	agent *fiber.Agent
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

	agent = fiber.AcquireAgent()

	req := agent.Request()
	req.Header.Set("Authorization", authorization)
	req.Header.Set("X-Comico-Signature", signature)
	req.Header.SetMethod(fiber.MethodPost)
	req.SetRequestURI(host + "/api/v1/logs")
}

func SendLog(referrer string, ip string, trackingIdentifier string, version string) {
	// 發起 request

	agent.JSON(Tracking{
		Referrer: referrer,
		Platform: "android",
		Event:    "download",
		Version:  version,
		Uid1:     trackingIdentifier,
		IP:       ip,
	})

	if err := agent.Parse(); err != nil {
		fmt.Printf("[error] %v", err)
		return
	}

	_, _, _ = agent.String()
}
