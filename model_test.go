package tracking

import (
	"testing"
)

func TestClient_SendLog(t *testing.T) {
	client := NewClient("a", "b", "http://localhost:7777")

	client.SendLog(Tracking{
		Object: "test",
		Uid1:   "test",
	})
}
