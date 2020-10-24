package gocsv

import "encoding/csv"

type DQuotCSVWriter struct {
	w *SafeCSVWriter
}

func NewDoubleQuotSafeCSVWriter(original *csv.Writer) CSVWriter {
	return &DQuotCSVWriter{
		w: NewSafeCSVWriter(original),
	}
}

//Override write
func (d *DQuotCSVWriter) Write(row []string) error {
	d.w.m.Lock()
	defer d.w.m.Unlock()
	//add double quot
	for i, field := range row {
		row[i] = "\"" + field + "\""
	}

	return d.w.Writer.Write(row)
}

//Override flush
func (d *DQuotCSVWriter) Flush() {
	d.w.Flush()
}

//Override error
func (d *DQuotCSVWriter) Error() error {
	return d.w.Error()
}
