package main

import (
	"github.com/olegpolukhin/rvision-irp/cmd"
	"github.com/olegpolukhin/rvision-irp/config"
)

func main() {
	config.Init()
	cmd.SchedulerRun()
}
