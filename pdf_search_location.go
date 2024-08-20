package gopdf

type PDFSearchLocation struct {
	AbsX      float64   `json:"absX"`
	AbsY      float64   `json:"absY"`
	Page      int       `json:"page"`
	AddText   string    `json:"addText"`
	RelativeX float64   `json:"relativeX"`
	RelativeY float64   `json:"relativeY"`
	FontSize  float64   `json:"fontSize"`
	BaseColor BaseColor `json:"baseColor"`
	Image     *Image    `json:"image"`
}

type Image struct {
	FilePath string  `json:"filePath,omitempty"`
	Url      string  `json:"url,omitempty"`
	Base64   string  `json:"base64,omitempty"`
	Width    float64 `json:"width,omitempty"`
	Height   float64 `json:"height,omitempty"`
}

type BaseColor struct {
	Red   uint8   `json:"red"`
	Green uint8   `json:"green"`
	Blue  uint8   `json:"blue"`
	Alpha float64 `json:"alpha"`
}
