package fashion

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sampleXml = "<?xml version=\"1.0\"?><catalog><book id=\"bk101\"><author>Gambardella, Matthew</author><title>XML Developer's Guide</title><genre>Computer</genre><price>44.95</price><publish_date>2000-10-01</publish_date><description>An in-depth look at creating applications with XML.</description></book><book id=\"bk102\"><author>Ralls, Kim</author><title>Midnight Rain</title><genre>Fantasy</genre><price>5.95</price><publish_date>2000-12-16</publish_date><description>A former architect battles corporate zombies, an evil sorceress, and her own childhood to become queen of the world.</description></book><book id=\"bk103\"><author>Corets, Eva</author><title>Maeve Ascendant</title><genre>Fantasy</genre><price>5.95</price><publish_date>2000-11-17</publish_date><description>After the collapse of a nanotechnology society in England, the young survivors lay the foundation for a new society.</description></book><book id=\"bk104\"><author>Corets, Eva</author><title>Oberon's Legacy</title><genre>Fantasy</genre><price>5.95</price><publish_date>2001-03-10</publish_date><description>In post-apocalypse England, the mysterious agent known only as Oberon helps to create a new life for the inhabitants of London. Sequel to Maeve Ascendant.</description></book><book id=\"bk105\"><author>Corets, Eva</author><title>The Sundered Grail</title><genre>Fantasy</genre><price>5.95</price><publish_date>2001-09-10</publish_date><description>The two daughters of Maeve, half-sisters, battle one another for control of England. Sequel to Oberon's Legacy.</description></book><book id=\"bk106\"><author>Randall, Cynthia</author><title>Lover Birds</title><genre>Romance</genre><price>4.95</price><publish_date>2000-09-02</publish_date><description>When Carla meets Paul at an ornithology conference, tempers fly as feathers get ruffled.</description></book><book id=\"bk107\"><author>Thurman, Paula</author><title>Splish Splash</title><genre>Romance</genre><price>4.95</price><publish_date>2000-11-02</publish_date><description>A deep sea diver finds true love twenty thousand leagues beneath the sea.</description></book><book id=\"bk108\"><author>Knorr, Stefan</author><title>Creepy Crawlies</title><genre>Horror</genre><price>4.95</price><publish_date>2000-12-06</publish_date><description>An anthology of horror stories about roaches, centipedes, scorpions  and other insects.</description></book><book id=\"bk109\"><author>Kress, Peter</author><title>Paradox Lost</title><genre>Science Fiction</genre><price>6.95</price><publish_date>2000-11-02</publish_date><description>After an inadvertant trip through a Heisenberg Uncertainty Device, James Salway discovers the problems of being quantum.</description></book><book id=\"bk110\"><author>O'Brien, Tim</author><title>Microsoft .NET: The Programming Bible</title><genre>Computer</genre><price>36.95</price><publish_date>2000-12-09</publish_date><description>Microsoft's .NET initiative is explored in detail in this deep programmer's reference.</description></book><book id=\"bk111\"><author>O'Brien, Tim</author><title>MSXML3: A Comprehensive Guide</title><genre>Computer</genre><price>36.95</price><publish_date>2000-12-01</publish_date><description>The Microsoft MSXML3 parser is covered in detail, with attention to XML DOM interfaces, XSLT processing, SAX and more.</description></book><book id=\"bk112\"><author>Galos, Mike</author><title>Visual Studio 7: A Comprehensive Guide</title><genre>Computer</genre><price>49.95</price><publish_date>2001-04-16</publish_date><description>Microsoft Visual Studio 7 is explored in depth, looking at how Visual Basic, Visual C++, C#, and ASP+ are integrated into a comprehensive development environment.</description></book></catalog>"

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
		{
			xml:    sampleXml,
			output: true,
		},
	}
	for _, test := range tests {
		r := strings.NewReader(test.xml)
		assert.Equal(t, test.output, WellFormed(r), "%s should be %v", test.xml, test.output)
	}
}

func TestWellFormedUnmarshal(t *testing.T) {
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
		{
			xml:    sampleXml,
			output: true,
		},
	}
	for _, test := range tests {
		r := []byte(test.xml)
		assert.Equal(t, test.output, WellFormedUnmarshal(r), "%s should be %v", test.xml, test.output)
	}
}
func BenchmarkWellFormed(b *testing.B) {
	r := strings.NewReader(sampleXml)
	for i := 0; i < b.N; i++ {
		r.Seek(0, 0)
		WellFormed(r)
	}
}

func BenchmarkWellFormedUnmarshal(b *testing.B) {
	r := []byte(sampleXml)
	for i := 0; i < b.N; i++ {
		WellFormedUnmarshal(r)
	}
}
