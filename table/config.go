package table

import (
	"github.com/jedib0t/go-pretty/text"
)

// ColumnConfig contains configurations that determine and modify the way the
// contents of the column get rendered.
type ColumnConfig struct {
	// Name is the name of the Column as it appears in the first Header row.
	// If a Header is not provided, or the name is not found in the header, this
	// will not work.
	Name string
	// Number is the Column # from left. When specified, it overrides the Name
	// property. If you know the exact Column number, use this instead of Name.
	Number int

	// Align defines the horizontal alignment
	Align text.Align
	// AlignFooter defines the horizontal alignment of Footer rows
	AlignFooter text.Align
	// AlignFooter defines the horizontal alignment of Header rows
	AlignHeader text.Align

	// Colors defines the colors to be used on the column
	Colors text.Colors
	// ColorsFooter defines the colors to be used on the column in Footer rows
	ColorsFooter text.Colors
	// ColorsFooter defines the colors to be used on the column in Header rows
	ColorsHeader text.Colors

	// Formatter is a custom-function that changes the way the value gets
	// rendered to the console. Refer to text/formatter.go for ready-to-use
	// Formatter functions.
	Formatter text.Formatter
	// FormatterFooter is like Formatter for Footer rows
	FormatterFooter text.Formatter
	// FormatterFooter is like Formatter for Header rows
	FormatterHeader text.Formatter

	// VAlign defines the vertical alignment
	VAlign text.VAlign
	// VAlignFooter defines the vertical alignment in Footer rows
	VAlignFooter text.VAlign
	// VAlignHeader defines the vertical alignment in Header rows
	VAlignHeader text.VAlign

	// WidthMin defines the minimum character length of the column
	WidthMin int
	// WidthMax defines the maximum character length of the column
	WidthMax int
}
