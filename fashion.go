package fashion

import (
	"encoding/xml"
	"io"
)

// WellFormed uses xml.Decoder to ensure that Start and End elements
// are properly nested and matched.
func WellFormed(r io.Reader) bool {
	d := xml.NewDecoder(r)
	for {
		token, err := d.Token()
		if token == nil && err == io.EOF {
			break
		}
		if err != nil {
			return false
		}
	}
	return true
}

// WellFormedUnmarshal uses xml.Unmarshal to ensure xml is well formed
func WellFormedUnmarshal(b []byte) bool {
	err := xml.Unmarshal(b, new(interface{}))
	if err != nil {
		return false
	}
	return true
}
