package checker

import (
	"net/url"
)

const (
	DecodeNothing   = 0x00
	DecodeURLEncode = 0x01
)

func URLDecode(s string) string {
	if result, err := url.QueryUnescape(s); err == nil {
		return result
	}
	return s
}

type Decoder struct {
	Decoders []func(string) string
}

func NewDecoder(decoder ...func(string) string) *Decoder {
	d := &Decoder{
		Decoders: decoder,
	}
	return d
}

func (d *Decoder) Decode(s string) string {
	if d == nil {
		return s
	}

	for _, decoder := range d.Decoders {
		s = decoder(s)
	}
	return s
}

func DefaultDecoders() *Decoder {
	return NewDecoder(
		URLDecode,
	)
}
