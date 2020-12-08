package fashion

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWellFormed(t *testing.T) {
	var tests = []struct {
		xml    string
		output bool
	}{
		{
			xml:    "<a>test</a>",
			output: true,
		},
		{
			xml:    "<a>test<b>again</b>",
			output: false,
		},
		{
			xml:    "<a/>test",
			output: true,
		},
		{
			xml:    "<a/>test<b>again</b>",
			output: true,
		},
		{
			xml:    "<?xml version='1.0' encoding='utf-8'?><Error><Code>BucketNameUnavailable</Code><Message>The requested bucket name is not available.</Message></Error>",
			output: true,
		},
	}
	for _, test := range tests {
		r := strings.NewReader(test.xml)
		assert.Equal(t, test.output, WellFormed(r), "%s should be %v", test.xml, test.output)
	}
}
