package model



// AudioSegment contains audio segment detail
type AudioSegment struct {
	Filename string  `json:"filename"`
	Start    float64 `json:"start"`
	End      float64 `json:"end"`
	Duration float64 `json:"duration"`
	StartAbs float64 `json:"start_abs"`
	EndAbs   float64 `json:"end_abs"`
}

// VideoSegment contains video segment detail
type VideoSegment struct {
	Filename string  `json:"filename"`
	Start    float64 `json:"start"`
	End      float64 `json:"end"`
	Duration float64 `json:"duration"`
	StartAbs float64 `json:"start_abs"`
	EndAbs   float64 `json:"end_abs"`
}
