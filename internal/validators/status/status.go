package status

import "fmt"

type status struct {
	totalJobs    []string
	completeJobs []string
	errJobs      []string
	succeeded    bool
}

func (s *status) Detail() string {
	return fmt.Sprintf(
		`%d out of %d

  total job count: %d
    jobs: %+q
  completed job count: %d
    jobs: %+q
  failed job count: %d
    jobs: %+q`,
		len(s.completeJobs), len(s.totalJobs),
		len(s.totalJobs), s.totalJobs,
		len(s.completeJobs), s.completeJobs,
		len(s.errJobs), s.errJobs,
	)
}

func (s *status) IsSuccess() bool {
	return s.succeeded
}
