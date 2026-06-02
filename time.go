package main

import (
	"encoding/xml"
	"time"
)

type DesTime struct{ time.Time }

func (t *DesTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	fmts := [...]string{
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		"Mon, 2 Jan 2006 15:04:05 MST",   // RFC1123 with single digit day
		"Mon, 2 Jan 2006 15:04:05 -0700", // RFC1123Z "
	}
	for _, f := range fmts {
		parsed, err := time.Parse(f, s)
		if err == nil {
			t.Time = parsed
			return nil
		}
	}
	return errorf("unable to parse time: %s", s)
}
