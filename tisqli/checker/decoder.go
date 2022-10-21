package checker

import (
	"net/url"
)

// URLDecode decodes a string in URL encoding.
func URLDecode(s string) string {
	if result, err := url.QueryUnescape(s); err == nil {
		return result
	}
	return s
}

// Decoder is a chain of decoders for SQL statements.
type Decoder struct {
	Decoders []func(string) string
}

// NewDecoder creates a new decoder with the given decoders.
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

// DefaultDecoders returns the default decoders, with URLDecode.
func DefaultDecoders() *Decoder {
	return NewDecoder(
		URLDecode,
	)
}
