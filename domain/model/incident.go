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
	ID               string       `db:"id,omitempty" json:"id"`
	UserID           string       `db:"user_id" json:"user_id"`
	ResponsibleID    string       `db:"responsible_id" json:"responsible_id"`
	OrganizationID   string       `db:"organization_id" json:"organization_id"`
	NetworkSegmentID string       `db:"network_segment_id" json:"network_segment_id"`
	SziID            string       `db:"szi_id" json:"szi_id"`
	CategoryID       string       `db:"category_id" json:"category_id"`
	PriorityID       string       `db:"priority_id" json:"priority_id"`
	StatusID         string       `db:"status_id" json:"status_id"`
	TypeID           string       `db:"type_id" json:"type_id"`
	Identifier       string       `db:"identifier" json:"identifier"`
	DetectName       string       `db:"detect_name" json:"detect_name"`
	Description      string       `db:"description" json:"description"`
	SavzResult       string       `db:"savz_result" json:"savz_result"`
	Path             string       `db:"path" json:"path"`
	CreatedDate      time.Time    `db:"created_date" json:"created_date"`
	CustomFields     *interface{} `db:"custom_fields,omitempty" json:"custom_fields"`
	PermissionLevel  uint16       `db:"permission_level"  json:"-"`
	Archived         bool         `db:"archived" json:"-"`
}
