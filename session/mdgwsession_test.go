package mdgwsession

import (
	vssconf "ssevss/configs"
	"testing"
)

func TestConnMDGW(t *testing.T) {
	value := ConnMDGW(vssconf.VssConf)
	want := 0
	if value != want {
		t.Errorf("get %q want %q", value, want)
	}
}
