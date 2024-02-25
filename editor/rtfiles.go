package editor

type Syntax struct {
	Name   string `json:"name"`
	Detect struct {
		Filename string `json:"filename"`
		Header   string `json:"header,omitempty"`
	} `json:"detect"`
}
