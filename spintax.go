package spinrewriter

import "strings"

type SpintaxFormat string

func (f SpintaxFormat) String() string {
	return string(f)
}

const (
	SpintaxFormatPipeBraces   SpintaxFormat = "{|}"
	SpintaxFormatTildBraces   SpintaxFormat = "{~}"
	SpintaxFormatPipeBrackets SpintaxFormat = "[|]"
	SpintaxFormatSpinBrackets SpintaxFormat = "[spin]"
	SpintaxFormatHashSpin     SpintaxFormat = "#SPIN"
)

type Spintax struct {
	options []string
	Format  SpintaxFormat
}

func (s *Spintax) String() string {
	return ""
}

func parseOptions(str string, format SpintaxFormat) []string {
	switch format {
	case SpintaxFormatPipeBraces:
		fallthrough
	case SpintaxFormatTildBraces:
		fallthrough
	case SpintaxFormatPipeBrackets:
		// remove first char (opening bracket / brace)
		// remove last char (closing bracket / brace)
		// spit options by format seperator
		return strings.Split(str[1:len(str)-1], string(format.String()[1]))
	case SpintaxFormatSpinBrackets:
		return strings.Split(str[6:len(str)-7], "|")
	case SpintaxFormatHashSpin:
		return strings.Split(strings.Trim(str[7:len(str)-2], " "), " || ")
	default:
		return nil
	}
}

func NewSpintax(str string, format SpintaxFormat) *Spintax {
	return &Spintax{
		Format:  format,
		options: parseOptions(str, format),
	}
}

func (s *Spintax) Options() []string {
	return s.options
}

func (s *Spintax) Option(i int) (string, bool) {
	if i < 0 || i >= len(s.options) {
		return "", false
	}

	return s.options[i], true
}

func (s *Spintax) NumOptions() int {
	return len(s.options)
}
