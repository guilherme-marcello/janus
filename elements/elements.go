package elements

type Recording struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Date       string `json:"date"`
	Audio      bool   `json:"audio"`
	AudioCodec string `json:"audio_codec,omitempty"`
	Video      bool   `json:"video"`
	VideoCodec string `json:"video_codec,omitempty"`
	Data       bool   `json:"data"`
}

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
