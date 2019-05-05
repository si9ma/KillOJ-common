package judge

// the result expose to user
type OuterResult struct {
	// for compile result and run result
	ID      string `json:"id"`
	Status  string `json:"status"`
	Message string `json:"msg,omitempty"`
	StdErr  string `json:"stderr,omitempty"` // stderr from inner process

	// only for run result
	Runtime  int64  `json:"runtime,omitempty"` // time usage in ms
	Memory   int64  `json:"memory,omitempty"`  // memory usage in KB
	Input    string `json:"input,omitempty"`
	Output   string `json:"output,omitempty"`
	Expected string `json:"expected,omitempty"`
}
