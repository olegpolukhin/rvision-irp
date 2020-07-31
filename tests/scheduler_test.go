package tests

import (
	"context"
	"github.com/olegpolukhin/rvision-irp/cmd/helper"
	"github.com/olegpolukhin/rvision-irp/config"
	"testing"
)

func TestEventIncidentListToDB(t *testing.T) {
	config.Init()

	helper.EventIncidentListToDB(context.Background())
}
