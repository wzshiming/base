package base

import (
	"testing"
)

func Test_ip(t *testing.T) {
	if ip := getLocalIP(); ip == "127.0.0.1" {
		t.Fail()
	} else {
		t.Log(ip)
	}
}
