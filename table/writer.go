package table

import (
	"github.com/jedib0t/go-pretty/text"
	"io"
)

// Writer declares the interfaces that can be used to setup and render a table.
type Writer interface {
	AppendFooter(row Row)
	AppendHeader(row Row)
	AppendRow(row Row)
	AppendRows(rows []Row)
	Length() int
	Render() string
	RenderCSV() string
	RenderHTML() string
	RenderMarkdown() string
	// deprecated
	SetAlign(align []text.Align)
	// deprecated
	SetAlignFooter(align []text.Align)
	// deprecated
	SetAlignHeader(align []text.Align)
	// deprecated
	SetAllowedColumnLengths(lengths []int)
	SetAllowedRowLength(length int)
	SetAutoIndex(autoIndex bool)
	SetCaption(format string, a ...interface{})
	// deprecated
	SetColors(colors []text.Colors)
	// deprecated
	SetColorsFooter(colors []text.Colors)
	// deprecated
	SetColorsHeader(colors []text.Colors)
	SetColumnConfigs(configs []ColumnConfig)
	SetHTMLCSSClass(cssClass string)
	SetIndexColumn(colNum int)
	SetOutputMirror(mirror io.Writer)
	SetPageSize(numLines int)
	SetStyle(style Style)
	// deprecated
	SetVAlign(vAlign []text.VAlign)
	// deprecated
	SetVAlignFooter(vAlign []text.VAlign)
	// deprecated
	SetVAlignHeader(vAlign []text.VAlign)
	SortBy(sortBy []SortBy)
	Style() *Style
}

// NewWriter initializes and returns a Writer.
func NewWriter() Writer {
	return &Table{}
}
