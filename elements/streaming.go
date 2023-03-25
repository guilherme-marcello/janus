package elements

type Mountpoint struct {
	ID          int    `json:"id"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Metadata    string `json:"metadata"`
	Enabled     bool   `json:"enabled"`
	Media       []struct {
		Mid   string `json:"mid"`
		Type  string `json:"type"`
		Label string `json:"label"`
		AgeMs int    `json:"age_ms"`
	} `json:"media,omitempty"`
}
