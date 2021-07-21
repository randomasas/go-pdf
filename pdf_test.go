package gopdf

import (
	"io/fs"
	"os"
	"testing"
)

// func TestAddKeywords2(t *testing.T) {
// 	ls := []*PDFSearchLocation{
// 		&PDFSearchLocation{
// 			AddText: "31000000000",
// 			AbsX:    160,
// 			AbsY:    270,
// 		},
// 		&PDFSearchLocation{
// 			AddText: "陆则国",
// 			AbsX:    170,
// 			AbsY:    300.35,
// 		},
// 		// &PDFSearchLocation{
// 		// 	AddText: "31000000000",
// 		// 	AbsX:    370,
// 		// 	AbsY:    547.63,
// 		// },
// 	}
// 	t.Log(">>>>>>>>>>>")
// 	t.Log("start 2")
// 	AddKeywords(ls, "/Users/jerry.shi/Desktop/tpc_dailing.pdf", "/Users/jerry.shi/Desktop/testfile.pdf")
// 	t.Log("end 2")
// }

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
	b, err := os.ReadFile("/Users/jerry.shi/Desktop/tpc_dailing.pdf")
	if err != nil {
		t.Log("<<<<<<<<<<<<<<<", err)
	}

	b, err = AddKeywordsBytes(ls, b)
	if err != nil {
		t.Error(">>>>>>>>>>>", err)
	}
	err = os.WriteFile("/Users/jerry.shi/Desktop/testpdf.pdf", b, fs.ModePerm)
	if err != nil {
		t.Error(">>>>>>>>>>>", err)
	}
	t.Log("end")
}
