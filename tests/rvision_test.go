package tests

import (
	gorvison "github.com/olegpolukhin/rvision-irp/cmd"
	"log"
	"testing"
)

func NewRevisionWithTestData() *gorvison.Revision {
	//auth := gorvison.Auth{
	//	Username: "111",
	//	APIToken: "222",
	//}
	var auth gorvison.Auth
	return gorvison.NewRevision(&auth, "https://example")
}

func TestJobs(t *testing.T) {
	Revision := NewRevisionWithTestData()
	jobs, err := Revision.GetJobs()

	if err != nil {
		t.Errorf("error %v\n", err)
	}

	if len(jobs) == 0 {
		t.Errorf("return no jobs\n")
	}

	log.Println(jobs)
}

func TestJob(t *testing.T) {
	Revision := NewRevisionWithTestData()
	job, err := Revision.GetJob("one")

	if err != nil {
		t.Errorf("error %v\n", err)
	}

	if len(job.Name) == 0 {
		t.Errorf("return no job\n")
	}
}