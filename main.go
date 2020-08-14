package main

import (
	"github.com/olegpolukhin/rvision-irp/cmd"
	"github.com/olegpolukhin/rvision-irp/config"
	log "github.com/sirupsen/logrus"
)

func main() {
	config.InitLogRus()
	config.Init()
	cmd.SchedulerRun()
	log.Info("Start service")
}
