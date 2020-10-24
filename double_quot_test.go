package gocsv

import (
	"encoding/csv"
	"io"
	"testing"
)

func Test_double_quot(t *testing.T) {
	blah := 2
	sptr := "*string"
	s := []Sample{
		{Foo: "f", Bar: 1, Baz: "baz", Frop: 0.1, Blah: &blah, SPtr: &sptr},
		{Foo: "e", Bar: 3, Baz: "b", Frop: 6.0 / 13, Blah: nil, SPtr: nil},
	}
	SetNewCSVWriter(func(out io.Writer) CSVWriter {
		return NewDoubleQuotSafeCSVWriter(csv.NewWriter(out))
	})
	defer SetNewCSVWriter(DefaultCSVWriter)

	csvContent, err := MarshalString(&s)
	if err != nil {
		t.Fatal(err)
	}
	if csvContent != "\"\"\"foo\"\"\",\"\"\"BAR\"\"\",\"\"\"Baz\"\"\",\"\"\"Quux\"\"\",\"\"\"Blah\"\"\",\"\"\"SPtr\"\"\",\"\"\"Omit\"\"\"\n\"\"\"f\"\"\",\"\"\"1\"\"\",\"\"\"baz\"\"\",\"\"\"0.1\"\"\",\"\"\"2\"\"\",\"\"\"*string\"\"\",\"\"\"\"\"\"\n\"\"\"e\"\"\",\"\"\"3\"\"\",\"\"\"b\"\"\",\"\"\"0.46153846153846156\"\"\",\"\"\"\"\"\",\"\"\"\"\"\",\"\"\"\"\"\"\n" {
		t.Fatalf("Error marshaling struct with double quotation. Expected \n\"\"\"foo\"\"\",\"\"\"BAR\"\"\",\"\"\"Baz\"\"\",\"\"\"Quux\"\"\",\"\"\"Blah\"\"\",\"\"\"SPtr\"\"\",\"\"\"Omit\"\"\"\n\"\"\"f\"\"\",\"\"\"1\"\"\",\"\"\"baz\"\"\",\"\"\"0.1\"\"\",\"\"\"2\"\"\",\"\"\"*string\"\"\",\"\"\"\"\"\"\n\"\"\"e\"\"\",\"\"\"3\"\"\",\"\"\"b\"\"\",\"\"\"0.46153846153846156\"\"\",\"\"\"\"\"\",\"\"\"\"\"\",\"\"\"\"\"\"\ngot:\n%v", csvContent)
	}
}
