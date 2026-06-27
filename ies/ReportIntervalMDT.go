package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	ReportIntervalMDTMs120   int64 = 0
	ReportIntervalMDTMs240   int64 = 1
	ReportIntervalMDTMs480   int64 = 2
	ReportIntervalMDTMs640   int64 = 3
	ReportIntervalMDTMs1024  int64 = 4
	ReportIntervalMDTMs2048  int64 = 5
	ReportIntervalMDTMs5120  int64 = 6
	ReportIntervalMDTMs10240 int64 = 7
	ReportIntervalMDTMin1    int64 = 8
	ReportIntervalMDTMin6    int64 = 9
	ReportIntervalMDTMin12   int64 = 10
	ReportIntervalMDTMin30   int64 = 11
	ReportIntervalMDTMin60   int64 = 12
)

var reportIntervalMDTConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
	ExtValues:  nil,
}

type ReportIntervalMDT struct {
	Value int64
}

func (ie *ReportIntervalMDT) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, reportIntervalMDTConstraints)
}

func (ie *ReportIntervalMDT) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(reportIntervalMDTConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
