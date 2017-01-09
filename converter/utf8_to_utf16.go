package converter

import (
	"bytes"
	"golang.org/x/text/encoding/japanese"
)

type UTF8ToUTF16Converter ConversionPattern

func (c *UTF8ToUTF16Converter) Convert() (string, error) {
	return transformEncoding(bytes.NewReader(c.TextByte), japanese.ShiftJIS.NewEncoder())
}