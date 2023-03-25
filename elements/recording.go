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
