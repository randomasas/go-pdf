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
}

type BaseColor struct {
	Red   uint8   `json:"red"`
	Green uint8   `json:"green"`
	Blue  uint8   `json:"blue"`
	Alpha float64 `json:"alpha"`
}
