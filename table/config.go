package table

import (
	"github.com/jedib0t/go-pretty/text"
)

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

	// IsIndexColumn sets the column as the Index column and enables it to be
	// styled differently using Style.Color (ColorOptions).
	IsIndexColumn bool

	// Formatter is a custom-function that changes the way the value gets
	// rendered to the console. Refer to formatter.go for ready-to-use Formatter
	// functions.
	Formatter Formatter

	// VAlign defines the vertical alignment
	VAlign text.VAlign
}
