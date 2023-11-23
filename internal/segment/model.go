package segment

// Model represents detected segment model
type Model struct {
	Input    string  `json:"input"`
	Output   string  `json:"output"`
	Start    float64 `json:"start"`
	End      float64 `json:"end"`
	Duration float64 `json:"duration"`
	StartAbs float64 `json:"start_abs"`
	EndAbs   float64 `json:"end_abs"`
}
