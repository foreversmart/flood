package types

// Report is for a single task result record report
type Report struct {
	TaskID          string          `json:"task_id"`       // task id
	ResponseTime    map[int64]int64 `json:"response_time"` // response time per second
	ResponseNum     map[int64]int64 `json:"response_num"`  // response num per second
	ResponseSuccess map[int64]int64 `json:"response_ok"`   // response success num per second
}
