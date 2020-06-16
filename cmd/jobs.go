package gorvison

import (
	"fmt"

	"github.com/olegpolukhin/rvision-irp/models"
)

// GetJobs returns all jobs you can read.
func (r *Revision) GetJobs() ([]models.Job, error) {
	var payload = struct {
		Jobs []models.Job `json:"jobs"`
	}{}

	if err := r.get("", nil, &payload); err != nil {
		return nil, err
	}

	return payload.Jobs, nil
}

// GetJob returns a job which has specified name.
func (r *Revision) GetJob(name string) (job models.Job, err error) {
	err = r.get(fmt.Sprintf("/job/%s", name), nil, &job)
	return
}
