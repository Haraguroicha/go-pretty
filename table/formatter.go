package table

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-openapi/strfmt"
	"github.com/jedib0t/go-pretty/text"
	"reflect"
	"strings"
	"time"
)

// Formatter related constants
const (
	unixTimeMinMilliseconds = int64(10000000000)
	unixTimeMinMicroseconds = unixTimeMinMilliseconds * 1000
	unixTimeMinNanoSeconds  = unixTimeMinMicroseconds * 1000
)

// Formatter related variables
var (
	colorsNumberPositive = text.Colors{text.FgHiGreen}
	colorsNumberNegative = text.Colors{text.FgHiRed}
	colorsNumberZero     = text.Colors{}
	colorsURL            = text.Colors{text.Underline, text.FgBlue}
)

// Formatter helps format the contents of a Column to the user's liking.
type Formatter func(val interface{}) string

// FormatNumber returns a number Formatter that:
//   * formats the number as directed by 'format' (ex.: %.2f)
//   * colors negative values Red
//   * colors positive values Green
func FormatNumber(format string) Formatter {
	return func(val interface{}) string {
		switch reflect.TypeOf(val).Kind() {
		case reflect.Int:
			if number, ok := val.(int); ok {
				if number < 0 {
					return colorsNumberNegative.Sprintf("-"+format, -number)
				}
				if number > 0 {
					return colorsNumberPositive.Sprintf(format, number)
				}
				return colorsNumberZero.Sprintf(format, number)
			}
		case reflect.Int8:
			if number, ok := val.(int8); ok {
				if number < 0 {
					return colorsNumberNegative.Sprintf("-"+format, -number)
				}
				if number > 0 {
					return colorsNumberPositive.Sprintf(format, number)
				}
				return colorsNumberZero.Sprintf(format, number)
			}
		case reflect.Int16:
			if number, ok := val.(int16); ok {
				if number < 0 {
					return colorsNumberNegative.Sprintf("-"+format, -number)
				}
				if number > 0 {
					return colorsNumberPositive.Sprintf(format, number)
				}
				return colorsNumberZero.Sprintf(format, number)
			}
		case reflect.Int32:
			if number, ok := val.(int32); ok {
				if number < 0 {
					return colorsNumberNegative.Sprintf("-"+format, -number)
				}
				if number > 0 {
					return colorsNumberPositive.Sprintf(format, number)
				}
				return colorsNumberZero.Sprintf(format, number)
			}
		case reflect.Int64:
			if number, ok := val.(int64); ok {
				if number < 0 {
					return colorsNumberNegative.Sprintf("-"+format, -number)
				}
				if number > 0 {
					return colorsNumberPositive.Sprintf(format, number)
				}
				return colorsNumberZero.Sprintf(format, number)
			}
		case reflect.Uint:
			if number, ok := val.(uint); ok {
				if number > 0 {
					return colorsNumberPositive.Sprintf(format, number)
				}
				return colorsNumberZero.Sprintf(format, number)
			}
		case reflect.Uint8:
			if number, ok := val.(uint8); ok {
				if number > 0 {
					return colorsNumberPositive.Sprintf(format, number)
				}
				return colorsNumberZero.Sprintf(format, number)
			}
		case reflect.Uint16:
			if number, ok := val.(uint16); ok {
				if number > 0 {
					return colorsNumberPositive.Sprintf(format, number)
				}
				return colorsNumberZero.Sprintf(format, number)
			}
		case reflect.Uint32:
			if number, ok := val.(uint32); ok {
				if number > 0 {
					return colorsNumberPositive.Sprintf(format, number)
				}
				return colorsNumberZero.Sprintf(format, number)
			}
		case reflect.Uint64:
			if number, ok := val.(uint64); ok {
				if number > 0 {
					return colorsNumberPositive.Sprintf(format, number)
				}
				return colorsNumberZero.Sprintf(format, number)
			}
		case reflect.Float32:
			if number, ok := val.(float32); ok {
				if number < 0 {
					return colorsNumberNegative.Sprintf("-"+format, -number)
				}
				if number > 0 {
					return colorsNumberPositive.Sprintf(format, number)
				}
				return colorsNumberZero.Sprintf(format, number)
			}
		case reflect.Float64:
			if number, ok := val.(float64); ok {
				if number < 0 {
					return colorsNumberNegative.Sprintf("-"+format, -number)
				}
				if number > 0 {
					return colorsNumberPositive.Sprintf(format, number)
				}
				return colorsNumberZero.Sprintf(format, number)
			}
		}
		// ultimate fallback
		return fmt.Sprint(val)
	}
}

// FormatJSON returns a Formatter that can format the following into
// pretty-indented JSON-strings:
//    * strings with JSON content
//    * structs
func FormatJSON(prefix string, indent string) Formatter {
	return func(val interface{}) string {
		if valStr, ok := val.(string); ok {
			var b bytes.Buffer
			if err := json.Indent(&b, []byte(strings.TrimSpace(valStr)), prefix, indent); err == nil {
				return string(b.Bytes())
			}
		} else if b, err := json.MarshalIndent(val, prefix, indent); err == nil {
			return string(b)
		}
		return fmt.Sprintf("%#v", val)
	}
}

// FormatTime returns a Formatter that can format a timestamp (time.Time or
// strfmt.DateTime) into a well-defined time format defined using layout. If a
// non-nil location value is provided, the time will be converted to that
// location's time (use time.Local to get localized timestamps).
func FormatTime(layout string, location *time.Location) Formatter {
	return func(val interface{}) string {
		formatTime := func(t time.Time) string {
			rsp := ""
			if t.Unix() > 0 {
				if location != nil {
					t = t.In(location)
				}
				rsp = t.Format(layout)
			}
			return rsp
		}

		rsp := fmt.Sprint(val)
		if valDate, ok := val.(strfmt.DateTime); ok {
			rsp = formatTime(time.Time(valDate))
		} else if valTime, ok := val.(time.Time); ok {
			rsp = formatTime(valTime)
		} else if valStr, ok := val.(string); ok {
			if valTime, err := time.Parse(time.RFC3339, valStr); err == nil {
				rsp = formatTime(valTime)
			}
		}
		return rsp
	}
}

// FormatUnixTime returns a Formatter that can format a unix-timestamp into a
// well-defined time format defined using layout. If a non-nil location value is
// provided, the time will be converted to that location's time (use time.Local
// to get localized timestamps).
func FormatUnixTime(layout string, location *time.Location) Formatter {
	timeFormatter := FormatTime(layout, location)

	return func(val interface{}) string {
		if unixTime, ok := val.(int64); ok {
			if unixTime >= unixTimeMinNanoSeconds {
				unixTime = unixTime / time.Second.Nanoseconds()
			} else if unixTime >= unixTimeMinMicroseconds {
				unixTime = unixTime / (time.Second.Nanoseconds() / 1000)
			} else if unixTime >= unixTimeMinMilliseconds {
				unixTime = unixTime / (time.Second.Nanoseconds() / 1000000)
			}
			return timeFormatter(time.Unix(unixTime, 0))
		}
		return fmt.Sprint(val)
	}
}

// FormatURL prints the value which is assumed to be an URL string with the
// string underlined and colored Blue.
func FormatURL(val interface{}) string {
	return colorsURL.Sprint(val)
}
