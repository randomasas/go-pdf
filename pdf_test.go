package gopdf

import (
	"log/slog"
	"testing"
)

func TestT(t *testing.T) {
	obj := PDFSearchLocation{
		Page:     2,
		AddText:  "测试",
		AbsX:     100,
		AbsY:     585,
		FontSize: 10,
	}
	slog.Error("obj:", slog.Attr{Key: "obj", Value: slog.AnyValue(obj)})
}

func TestAddKeywords(t *testing.T) {
	ls := []*PDFSearchLocation{
		&PDFSearchLocation{
			Page:     2,
			AddText:  "姓名",
			AbsX:     180,
			AbsY:     580,
			FontSize: 12,
		},
		&PDFSearchLocation{
			Page:     2,
			AddText:  "男",
			AbsX:     430,
			AbsY:     580,
			FontSize: 12,
		},
		&PDFSearchLocation{
			Page:     2,
			AddText:  "31000000000",
			AbsX:     180,
			AbsY:     535,
			FontSize: 12,
		},
		&PDFSearchLocation{
			Page:     2,
			AddText:  "2023",
			AbsX:     166,
			AbsY:     400,
			FontSize: 12,
		},
		&PDFSearchLocation{
			Page:     2,
			AddText:  "02",
			AbsX:     220,
			AbsY:     400,
			FontSize: 12,
		},
		&PDFSearchLocation{
			Page:     2,
			AddText:  "31",
			AbsX:     254,
			AbsY:     400,
			FontSize: 12,
		},
		&PDFSearchLocation{
			Page:     2,
			AddText:  "13321954022",
			AbsX:     180,
			AbsY:     490,
			FontSize: 12,
		},
		&PDFSearchLocation{
			Page:     2,
			AddText:  "方案名称",
			AbsX:     180,
			AbsY:     444,
			FontSize: 12,
		},
	}
	t.Log(">>>>>>>>>>>")
	// b, err := os.ReadFile("/Users/jerry.shi/Desktop/240_s.pdf")
	// if err != nil {
	// 	t.Log("<<<<<<<<<<<<<<<", err)
	// }

	//b, err = AddKeywordsBytes(ls, b, true)
	err := AddKeywords(ls, "/Users/jerry.shi/Desktop/240_s.pdf", "/Users/jerry.shi/Desktop/test_s.pdf", true)
	if err != nil {
		t.Error(">>>>>>>>>>>", err)
	}
	// err = os.WriteFile("/Users/jerry.shi/Desktop/test_s.pdf", b, fs.ModePerm)
	// if err != nil {
	// 	t.Error(">>>>>>>>>>>", err)
	// }
	t.Log("end")
}
