package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	LoggingIntervalMs320    int64 = 0
	LoggingIntervalMs640    int64 = 1
	LoggingIntervalMs1280   int64 = 2
	LoggingIntervalMs2560   int64 = 3
	LoggingIntervalMs5120   int64 = 4
	LoggingIntervalMs10240  int64 = 5
	LoggingIntervalMs20480  int64 = 6
	LoggingIntervalMs30720  int64 = 7
	LoggingIntervalMs40960  int64 = 8
	LoggingIntervalMs61440  int64 = 9
	LoggingIntervalInfinity int64 = 10
)

var loggingIntervalConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	ExtValues:  nil,
}

type LoggingInterval struct {
	Value int64
}

func (ie *LoggingInterval) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, loggingIntervalConstraints)
}

func (ie *LoggingInterval) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(loggingIntervalConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
