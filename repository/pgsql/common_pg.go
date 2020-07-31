package pgsql

import (
	"fmt"
	"github.com/olegpolukhin/rvision-irp/cmd/server"
	"github.com/olegpolukhin/rvision-irp/domain/model"
	"github.com/olegpolukhin/rvision-irp/pkg/datasource"
)

type CommonRepository struct {
	db *datasource.Postgres
}

func NewCommonRepository(app *server.App) (helper *CommonRepository) {
	return &CommonRepository{app.Postgres}
}

func (r *CommonRepository) UserList() (list []model.User, err error) {
	if err := r.db.Connect(); err != nil {
		return list, fmt.Errorf("connect to DB error %s", err.Error())
	}

	defer r.db.Close(&err)

	if err := r.db.UsersList(&list); err != nil {
		return list, fmt.Errorf("UsersList error %s", err.Error())
	}

	if len(list) == 0 {
		return list, fmt.Errorf("UsersList error. No Items")
	}

	return
}

func (r *CommonRepository) KtmTypeList() (list []model.KtmType, err error) {
	if err := r.db.Connect(); err != nil {
		return list, fmt.Errorf("connect to DB error %s", err.Error())
	}

	defer r.db.Close(&err)

	if err := r.db.KtmList(&list); err != nil {
		return list, fmt.Errorf("KtmList error %s", err.Error())
	}

	if len(list) == 0 {
		return list, fmt.Errorf("KtmList error. No Items")
	}

	return
}

func (r *CommonRepository) SziTypeList() (list []model.SziType, err error) {
	if err := r.db.Connect(); err != nil {
		return list, fmt.Errorf("connect to DB error %s", err.Error())
	}

	defer r.db.Close(&err)

	if err := r.db.SziTypeList(&list); err != nil {
		return list, fmt.Errorf("SziTypeList error %s", err.Error())
	}

	if len(list) == 0 {
		return list, fmt.Errorf("SziTypeList error. No Items")
	}

	return
}

func (r *CommonRepository) SziList() (list []model.Szi, err error) {
	if err := r.db.Connect(); err != nil {
		return list, fmt.Errorf("connect to DB error %s", err.Error())
	}

	defer r.db.Close(&err)

	if err := r.db.SziList(&list); err != nil {
		return list, fmt.Errorf("SziList error %s", err.Error())
	}

	if len(list) == 0 {
		return list, fmt.Errorf("SziList error. No Items")
	}

	return
}

func (r *CommonRepository) SziManufactureList() (list []model.SziManufacture, err error) {
	if err := r.db.Connect(); err != nil {
		return list, fmt.Errorf("connect to DB error %s", err.Error())
	}

	defer r.db.Close(&err)

	if err := r.db.SziManufacturerList(&list); err != nil {
		return list, fmt.Errorf("SziManufacturerList error %s", err.Error())
	}

	if len(list) == 0 {
		return list, fmt.Errorf("SziManufacturerList error. No Items")
	}

	return
}

func (r *CommonRepository) NetworkSegmentList() (list []model.NetworkSegment, err error) {
	if err := r.db.Connect(); err != nil {
		return list, fmt.Errorf("connect to DB error %s", err.Error())
	}

	defer r.db.Close(&err)

	if err := r.db.NetworkSegmentsList(&list); err != nil {
		return list, fmt.Errorf("SziManufacturerList error %s", err.Error())
	}

	if len(list) == 0 {
		return list, fmt.Errorf("SziManufacturerList error. No Items")
	}

	return
}

func (r *CommonRepository) OrganizationList() (list []model.Organization, err error) {
	if err := r.db.Connect(); err != nil {
		return list, fmt.Errorf("connect to DB error %s", err.Error())
	}

	defer r.db.Close(&err)

	if err := r.db.OrganizationsList(&list); err != nil {
		return list, fmt.Errorf("OrganizationsList error %s", err.Error())
	}

	if len(list) == 0 {
		return list, fmt.Errorf("OrganizationsList error. No Items")
	}

	return
}

func (r *CommonRepository) IncidentCategoryList() (list []model.IncidentCategory, err error) {
	if err := r.db.Connect(); err != nil {
		return list, fmt.Errorf("connect to DB error %s", err.Error())
	}

	defer r.db.Close(&err)

	if err := r.db.IncidentCategoryList(&list); err != nil {
		return list, fmt.Errorf("IncidentCategoryList error %s", err.Error())
	}

	if len(list) == 0 {
		return list, fmt.Errorf("IncidentCategoryList error. No Items")
	}

	return
}

func (r *CommonRepository) IncidentPriorityList() (list []model.IncidentPriority, err error) {
	if err := r.db.Connect(); err != nil {
		return list, fmt.Errorf("connect to DB error %s", err.Error())
	}

	defer r.db.Close(&err)

	if err := r.db.IncidentPriorityList(&list); err != nil {
		return list, fmt.Errorf("IncidentPriorityList error %s", err.Error())
	}

	if len(list) == 0 {
		return list, fmt.Errorf("IncidentPriorityList error. No Items")
	}

	return
}

func (r *CommonRepository) IncidentStatusList() (list []model.IncidentStatus, err error) {
	if err := r.db.Connect(); err != nil {
		return list, fmt.Errorf("connect to DB error %s", err.Error())
	}

	defer r.db.Close(&err)

	if err := r.db.IncidentStatusList(&list); err != nil {
		return list, fmt.Errorf("IncidentStatusList error %s", err.Error())
	}

	if len(list) == 0 {
		return list, fmt.Errorf("IncidentStatusList error. No Items")
	}

	return
}

func (r *CommonRepository) IncidentTypeList() (list []model.IncidentType, err error) {
	if err := r.db.Connect(); err != nil {
		return list, fmt.Errorf("connect to DB error %s", err.Error())
	}

	defer r.db.Close(&err)

	if err := r.db.IncidentTypeList(&list); err != nil {
		return list, fmt.Errorf("IncidentTypeList error %s", err.Error())
	}

	if len(list) == 0 {
		return list, fmt.Errorf("IncidentTypeList error. No Items")
	}

	return
}
