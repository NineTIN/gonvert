package converter

type Converter interface {
	Convert() (string, error)
}

type ConversionPattern struct {
	Text string
}