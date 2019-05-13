package table

import (
	"io"

	"github.com/jedib0t/go-pretty/text"
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
	SetAllowedRowLength(length int)
	SetAutoIndex(autoIndex bool)
	SetCaption(format string, a ...interface{})
	SetColumnConfigs(configs []ColumnConfig)
	SetHTMLCSSClass(cssClass string)
	SetIndexColumn(colNum int)
	SetOutputMirror(mirror io.Writer)
	SetPageSize(numLines int)
	SetRowPainter(painter RowPainter)
	SetStyle(style Style)
	SortBy(sortBy []SortBy)
	Style() *Style

	// deprecated; use SetColumnOptions instead
	SetAlign(align []text.Align)
	// deprecated; use SetColumnOptions instead
	SetAlignFooter(align []text.Align)
	// deprecated; use SetColumnOptions instead
	SetAlignHeader(align []text.Align)
	// deprecated; use SetColumnOptions instead
	SetAllowedColumnLengths(lengths []int)
	// deprecated; use SetColumnOptions instead
	SetColors(colors []text.Colors)
	// deprecated; use SetColumnOptions instead
	SetColorsFooter(colors []text.Colors)
	// deprecated; use SetColumnOptions instead
	SetColorsHeader(colors []text.Colors)
	// deprecated; use SetColumnOptions instead
	SetVAlign(vAlign []text.VAlign)
	// deprecated; use SetColumnOptions instead
	SetVAlignFooter(vAlign []text.VAlign)
	// deprecated; use SetColumnOptions instead
	SetVAlignHeader(vAlign []text.VAlign)
}

// NewWriter initializes and returns a Writer.
func NewWriter() Writer {
	return &Table{}
}
