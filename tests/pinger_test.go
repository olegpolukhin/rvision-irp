package tests

import (
	"github.com/olegpolukhin/rvision-irp/config"
	"github.com/olegpolukhin/rvision-irp/pkg/ping"
	"github.com/spf13/viper"
	"testing"
)

func TestPinger(t *testing.T) {
	config.Init()

	err := ping.NewPinger(viper.GetString("api_url"))
	if err != nil {
		t.Logf("pinger error %v", err)
		return
	}
}
