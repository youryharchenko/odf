package odf

type DocumentType int

const (
	Text DocumentType = iota
	Spreadsheet
	Presentation
	Drawing
	Chart
	Image
	Formula
)

func (docType DocumentType) Name() string {
	switch docType {
	case Text:
		return "Text"
	case Spreadsheet:
		return "Spreadsheet"
	case Presentation:
		return "Presentation"
	case Drawing:
		return "Drawing"
	case Chart:
		return "Chart"
	case Image:
		return "Image"
	case Formula:
		return "Formula"
	}
	return "Undefined"
}

func (docType DocumentType) Template() string {
	switch docType {
	case Text:
		return "text.ott"
	case Spreadsheet:
		return "spreadsheet.ots"
	case Presentation:
		return "presentation.otp"
	case Drawing:
		return "drawing.otg"
	case Chart:
		return "chart.otc"
	case Image:
		return "image.oti"
	case Formula:
		return "formula.otf"
	}
	return "undefined.otu"
}
