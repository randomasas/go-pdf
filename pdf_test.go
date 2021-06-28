package gopdf

import (
	"testing"
)

func TestAddKeywords(t *testing.T) {
	ls := []*PDFSearchLocation{
		&PDFSearchLocation{
			AddText: "31000000000",
			AbsX:    160,
			AbsY:    270,
		},
		&PDFSearchLocation{
			AddText: "陆则国",
			AbsX:    170,
			AbsY:    300.35,
		},
		// &PDFSearchLocation{
		// 	AddText: "31000000000",
		// 	AbsX:    370,
		// 	AbsY:    547.63,
		// },
	}
	t.Log(">>>>>>>>>>>")
	// f, err := os.Open("/Users/jerry.shi/Desktop/tpc-dailing.pdf")
	// defer f.Close()
	// if err != nil {
	// 	t.Log("<<<<<<<<<<<<<<<")
	// }
	err := AddKeywords(ls, "/Users/jerry.shi/Desktop/tpc_dailing.pdf", "/Users/jerry.shi/Desktop/test.pdf")

	t.Error(">>>>>>>>>>>", err)
}
