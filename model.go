package tracking

import "time"

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
