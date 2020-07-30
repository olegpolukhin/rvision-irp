package model

import "time"

type Incident struct {
	ID           int       `json:"id"`
	Identifier   string    `json:"identifier"`
	Creation     time.Time `json:"creation"`
	Updated      time.Time `json:"updated"`
	IncidentType string    `json:"incident_type"`
	Category     string    `json:"category"`
	Status       string    `json:"status"`
	Level        string    `json:"level"`
	Owner        string    `json:"owner"`
	Responsible  string    `json:"responsible"`
	Description  string    `json:"description"`
}

type IncidentList struct {
	Total     int        `json:"total"`
	Incidents []Incident `json:"incidents"`
}

type IncidentDB struct {
	ID               string       `db:"id,omitempty" json:"id" front_type:"string" label:"-" required:"false"`
	UserID           string       `db:"user_id" json:"user_id" front_type:"select" label:"Пользователь" required:"true" list:"user"`
	ResponsibleID    string       `db:"responsible_id" json:"responsible_id" front_type:"select" label:"Ответсвенный" required:"true" list:"user"`
	OrganizationID   string       `db:"organization_id" json:"organization_id" front_type:"select" label:"Организация в/ч" required:"true" list:"organization"`
	NetworkSegmentID string       `db:"network_segment_id" json:"network_segment_id" front_type:"select" label:"Сегмент сети ЗС" required:"true" list:"network_segment"`
	SziID            string       `db:"szi_id" json:"szi_id" front_type:"select" label:"СЗИ" required:"true" list:"szi"`
	CategoryID       string       `db:"category_id" json:"category_id" front_type:"select" label:"Категория" required:"true" list:"incident_category"`
	PriorityID       string       `db:"priority_id" json:"priority_id" front_type:"select" label:"Уровень" required:"true" list:"incident_priority"`
	StatusID         string       `db:"status_id" json:"status_id" front_type:"select" label:"Статус" required:"true" list:"incident_status"`
	TypeID           string       `db:"type_id" json:"type_id" front_type:"select" label:"Тип" required:"true" list:"incident_type"`
	Identifier       string       `db:"identifier" json:"identifier" front_type:"string" label:"Идентификатор" required:"false"`
	DetectName       string       `db:"detect_name" json:"detect_name" front_type:"string" label:"Наименование детекта" required:"true"`
	Description      string       `db:"description" json:"description" front_type:"text" label:"Краткое описание инцидента" required:"false"`
	SavzResult       string       `db:"savz_result" json:"savz_result" front_type:"string" label:"Результат работы САВЗ" required:"false"`
	Path             string       `db:"path" json:"path" front_type:"string" label:"Путь до вредоносного объекта" required:"false"`
	CreatedDate      time.Time    `db:"created_date" json:"created_date" front_type:"date" label:"Дата создания" required:"true"`
	CustomFields     *interface{} `db:"custom_fields,omitempty" json:"custom_fields" front_type:"-"  label:"-" required:"false"`
	PermissionLevel  uint16       `db:"permission_level"  json:"-" front_type:"-"  label:"-" required:"false"`
	Archived         bool         `db:"archived" json:"-" front_type:"-"  label:"-" required:"false"`
}
