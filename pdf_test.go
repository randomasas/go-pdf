package gopdf

import (
	"io/fs"
	"log/slog"
	"os"
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
			AddText:  "测试",
			AbsX:     100,
			AbsY:     585,
			FontSize: 10,
		},
		&PDFSearchLocation{
			Page:     2,
			AddText:  "男",
			AbsX:     220,
			AbsY:     585,
			FontSize: 10,
		},
		&PDFSearchLocation{
			Page:     2,
			AddText:  "31000000000",
			AbsX:     100,
			AbsY:     542,
			FontSize: 10,
		},
		&PDFSearchLocation{
			Page:     2,
			AddText:  "2023",
			AbsX:     79,
			AbsY:     409,
			FontSize: 10,
		},
		&PDFSearchLocation{
			Page:     2,
			AddText:  "02",
			AbsX:     105,
			AbsY:     409,
			FontSize: 10,
		},
		&PDFSearchLocation{
			Page:     2,
			AddText:  "31",
			AbsX:     124,
			AbsY:     409,
			FontSize: 10,
		},
		&PDFSearchLocation{
			Page:     2,
			AddText:  "13321954022",
			AbsX:     100,
			AbsY:     499,
			FontSize: 10,
		},
		&PDFSearchLocation{
			Page:     2,
			AddText:  "方案名称",
			AbsX:     100,
			AbsY:     454,
			FontSize: 10,
		},
	}
	t.Log(">>>>>>>>>>>")
	b, err := os.ReadFile("/Users/jerry.shi/Desktop/简版-特药服务（海内外）权益服务手册.pdf")
	if err != nil {
		t.Log("<<<<<<<<<<<<<<<", err)
	}

	b, err = AddKeywordsBytes(ls, b)
	if err != nil {
		t.Error(">>>>>>>>>>>", err)
	}
	err = os.WriteFile("/Users/jerry.shi/Desktop/test_s.pdf", b, fs.ModePerm)
	if err != nil {
		t.Error(">>>>>>>>>>>", err)
	}
	t.Log("end")
}
