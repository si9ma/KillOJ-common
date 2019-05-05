package judge

const (
	CompileResType = "COMPILE"
	RunResType     = "RUN"
)

const (
	SUCCESS = iota
	FAIL
)

// result struct in inner system
type InnerResult struct {
	// for compile result and run result
	ID         string `json:"id"`
	ResultType string `json:"resType"`
	Status     int64  `json:"status"`
	Errno      int64  `json:"errno,omitempty"`
	Message    string `json:"msg,omitempty"`
	StdErr     string `json:"stderr,omitempty"` // stderr from inner process

	// only for run result
	Runtime  int64  `json:"runtime,omitempty"` // time usage in ms
	Memory   int64  `json:"memory,omitempty"`  // memory usage in KB
	Input    string `json:"input,omitempty"`
	Output   string `json:"output,omitempty"`
	Expected string `json:"expected,omitempty"`
}
