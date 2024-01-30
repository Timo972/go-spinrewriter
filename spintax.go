package spinrewriter

type SpintaxFormat string

func (f SpintaxFormat) String() string {
	return string(f)
}

const (
	FormatPipeBraces   SpintaxFormat = "{|}"
	FormatTildBraces   SpintaxFormat = "{~}"
	FormatPipeBrackets SpintaxFormat = "[|]"
	FormatSpinBrackets SpintaxFormat = "[spin]"
	FormatHashSpin     SpintaxFormat = "#SPIN"
)

type Spintax struct {
	Quota
	format  SpintaxFormat
	spintax string
}

func NewSpintax(spintax string, format SpintaxFormat, reqsMade int, reqsAvailable int) *Spintax {
	return &Spintax{
		format:  format,
		spintax: spintax,
		Quota: Quota{
			ReqsMade:      reqsMade,
			ReqsAvailable: reqsAvailable,
		},
	}
}

func (s *Spintax) String() string {
	return s.spintax
}
