package gopdf

type PDFSearchLocation struct {
	AbsX      float64
	AbsY      float64
	Page      int
	AddText   string
	RelativeX float64
	RelativeY float64
	FontSize  float64
	BaseColor BaseColor
}

type BaseColor struct {
	Red   uint8
	Green uint8
	Blue  uint8
	Alpha float64
}
