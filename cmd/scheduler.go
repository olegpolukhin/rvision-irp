package cmd

import (
	"context"
	"github.com/olegpolukhin/rvision-irp/cmd/helper"
	"github.com/olegpolukhin/rvision-irp/pkg/scheduler"
	"os"
	"os/signal"
	"time"
)

const timeDelayCommon = time.Second * 20
const timeDelayIncident = time.Second * 20

func SchedulerRun() {
	ctx := context.Background()

	worker := scheduler.NewScheduler()
	worker.Add(ctx, helper.EventCommon, timeDelayCommon)
	worker.Add(ctx, helper.EventIncidentListToDB, timeDelayIncident)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit
	worker.Stop()
}
