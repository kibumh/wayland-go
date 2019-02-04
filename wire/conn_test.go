package wire_test

import (
	"testing"

	"github.com/kibumh/wayland-go/protocol"
)

func TestConnect(t *testing.T) {
	addr := "/run/user/1000/wayland-100"
	_, err := protocol.Dial(addr)
	if err == nil {
		t.Errorf("The address '%s' does not exist... but Dial succeeded", addr)
	}

	addr = "/run/user/1000/wayland-0"
	conn, err := protocol.Dial(addr)
	if err != nil {
		t.Errorf("Dial failed: %v.", err)
	}
	err = conn.Close()
	if err != nil {
		t.Errorf("Close failed: %v", err)
	}
}
