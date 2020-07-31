package datasource

import (
	"github.com/olegpolukhin/rvision-irp/domain/model"
)


func (pg *Postgres) UsersList(s *[]model.User) error {
	err := pg.tx.SelectFrom("users").All(s)
	return PanicOnError(err)
}

func (pg *Postgres) OrganizationsList(s *[]model.Organization) error {
	err := pg.tx.SelectFrom("organizations").All(s)
	return PanicOnError(err)
}

func (pg *Postgres) NetworkSegmentsList(s *[]model.NetworkSegment) error {
	err := pg.tx.SelectFrom("network_segments").All(s)
	return PanicOnError(err)
}

func (pg *Postgres) SziList(s *[]model.Szi) error {
	err := pg.tx.SelectFrom("szis").All(s)
	return PanicOnError(err)
}

func (pg *Postgres) SziManufacturerList(s *[]model.SziManufacture) error {
	err := pg.tx.SelectFrom("szi_manufacturers").All(s)
	return PanicOnError(err)
}

func (pg *Postgres) SziTypeList(s *[]model.SziType) error {
	err := pg.tx.SelectFrom("szi_types").All(s)
	return PanicOnError(err)
}

func (pg *Postgres) KtmList(s *[]model.KtmType) error {
	err := pg.tx.SelectFrom("ktm_types").All(s)
	return PanicOnError(err)
}

func (pg *Postgres) IncidentPriorityList(s *[]model.IncidentPriority) error {
	err := pg.tx.SelectFrom("incident_priorities").All(s)
	return PanicOnError(err)
}

func (pg *Postgres) IncidentCategoryList(s *[]model.IncidentCategory) error {
	err := pg.tx.SelectFrom("incident_categories").All(s)
	return PanicOnError(err)
}

func (pg *Postgres) IncidentStatusList(s *[]model.IncidentStatus) error {
	err := pg.tx.SelectFrom("incident_statuses").All(s)
	return PanicOnError(err)
}

func (pg *Postgres) IncidentTypeList(s *[]model.IncidentType) error {
	err := pg.tx.SelectFrom("incident_types").All(s)
	return PanicOnError(err)
}
