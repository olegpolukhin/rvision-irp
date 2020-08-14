package helper

import (
	"context"
	"github.com/olegpolukhin/rvision-irp/pkg/ping"
	"github.com/spf13/viper"
	log "github.com/sirupsen/logrus"
)

// EventPingRevisionAPI чекаем доступность API
func EventPingRevisionAPI(ctx context.Context) {
	err := ping.NewPinger(viper.GetString("api_url"))
	if err != nil {
		log.Error("NewPinger %v", err)
	}

	return
}
