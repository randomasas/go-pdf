package gopdf

import (
	"bytes"
	"errors"
	"io"
	"log/slog"
	"sync"

	"github.com/skirrund/go-pdf/font"

	"os"

	"github.com/phpdave11/gofpdi"
	"github.com/signintech/gopdf"
)

var fontBytes []byte

var bufferPool = &sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

func getByteBuffer() *bytes.Buffer {
	return bufferPool.Get().(*bytes.Buffer)
}

func releaseByteBuffer(buffer *bytes.Buffer) {
	buffer.Reset()
	bufferPool.Put(buffer)
}

func init() {
	var err error
	fontBytes, err = font.Asset("font/font.ttf")
	d, _ := os.Getwd()

	//fontBytes, err = os.ReadFile(d + "/font/font.ttf")
	if err != nil {
		slog.Error("[PDF] can not find font:" + d + "/font/font.ttf")
	}
}

func readTempFile(templateFile string) (*io.ReadSeeker, error) {
	ex, err := Exist(templateFile)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("templateFile not exist" + templateFile)
	}
	bs, err := os.ReadFile(templateFile)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(bs)
	rd := io.ReadSeeker(reader)
	return &rd, nil
}

func readTempFileBytes(templateFile []byte) (*io.ReadSeeker, error) {
	if len(templateFile) == 0 {
		return nil, errors.New("templateFile error length is 0")
	}
	reader := bytes.NewReader(templateFile)
	rd := io.ReadSeeker(reader)
	return &rd, nil
}

func AddKeywords(locations []*PDFSearchLocation, templateFile string, saveasFilepath string) (err error) {
	defer func() {
		if err := recover(); err != nil {
			slog.Error("[PDF] recover :", err)
		}
	}()
	pdf := &gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	if len(fontBytes) == 0 {
		return errors.New("[PDF] can not find font file")
	}
	err = pdf.AddTTFFontByReader("song", bytes.NewReader(fontBytes))
	if err != nil {
		//logger.Logger.Error("[PDF] AddTTFFontByReader error", err)
		return
	}
	rd, err := readTempFile(templateFile)
	if err != nil {
		//logger.Logger.Error("[PDF] readTempFile", err)
		return err
	}

	importer := gofpdi.NewImporter()
	importer.SetSourceStream(rd)
	num := importer.GetNumPages()
	for i := 1; i <= num; i++ {
		pdf.AddPage()
		tpl := pdf.ImportPageStream(rd, i, "/MediaBox")
		pdf.UseImportedTemplate(tpl, 0, 0, float64(gopdf.PageSizeA4.W), float64(gopdf.PageSizeA4.H))
		for _, sl := range locations {
			if sl.Page <= 0 {
				sl.Page = 1
			}
			if sl.Page == i {
				if sl.FontSize <= 0 {
					sl.FontSize = 10.0
				}
				err = pdf.SetFont("song", "", int(sl.FontSize))
				if err != nil {
					//logger.Logger.Error("[PDF] SetFont error", err)
					return err
				}

				pdf.SetTextColor(sl.BaseColor.Red, sl.BaseColor.Green, sl.BaseColor.Blue)
				if sl.BaseColor.Alpha == 0 {
					sl.BaseColor.Alpha = 1
				}
				pdf.SetTransparency(gopdf.Transparency{
					Alpha: sl.BaseColor.Alpha,
				})

				pdf.SetX(sl.AbsX)
				pdf.SetY(gopdf.PageSizeA4.H - sl.AbsY)
				pdf.Text(sl.AddText)
			}
		}
	}

	outFile, err := os.Create(saveasFilepath)
	if err != nil {
		//logger.Logger.Error("[PDF] Create saveasFilepath error", err)
		return err
	}
	defer outFile.Close()
	err = pdf.Write(outFile)
	return err
}

func AddKeywordsBytes(locations []*PDFSearchLocation, templateFile []byte) (bs []byte, err error) {
	defer func() {
		if err := recover(); err != nil {
			slog.Error("[PDF] recover :", err)
		}
	}()
	pdf := &gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	if len(fontBytes) == 0 {
		return templateFile, errors.New("[PDF] can not find font file")
	}
	err = pdf.AddTTFFontByReader("song", bytes.NewReader(fontBytes))
	if err != nil {
		//logger.Logger.Error("[PDF] AddTTFFontByReader error", err)
		return
	}
	rd, err := readTempFileBytes(templateFile)
	if err != nil {
		//logger.Logger.Error("[PDF] readTempFile", err)
		return templateFile, err
	}

	importer := gofpdi.NewImporter()
	importer.SetSourceStream(rd)
	num := importer.GetNumPages()
	for i := 1; i <= num; i++ {
		pdf.AddPage()
		tpl := pdf.ImportPageStream(rd, i, "/MediaBox")
		pdf.UseImportedTemplate(tpl, 0, 0, float64(gopdf.PageSizeA4.W), float64(gopdf.PageSizeA4.H))
		for _, sl := range locations {
			if sl.Page <= 0 {
				sl.Page = 1
			}
			if sl.Page == i {
				if sl.FontSize <= 0 {
					sl.FontSize = 10.0
				}
				err = pdf.SetFont("song", "", int(sl.FontSize))
				if err != nil {
					//logger.Logger.Error("[PDF] SetFont error", err)
					return templateFile, err
				}

				pdf.SetTextColor(sl.BaseColor.Red, sl.BaseColor.Green, sl.BaseColor.Blue)
				if sl.BaseColor.Alpha == 0 {
					sl.BaseColor.Alpha = 1
				}
				pdf.SetTransparency(gopdf.Transparency{
					Alpha: sl.BaseColor.Alpha,
				})

				pdf.SetX(sl.AbsX)
				pdf.SetY(gopdf.PageSizeA4.H - sl.AbsY)
				pdf.Text(sl.AddText)
			}
		}
	}
	buffer := getByteBuffer()
	defer releaseByteBuffer(buffer)
	err = pdf.Write(buffer)
	if err != nil {
		return templateFile, nil
	}
	bs = make([]byte, buffer.Len())
	buffer.Read(bs)
	return bs, nil
}
