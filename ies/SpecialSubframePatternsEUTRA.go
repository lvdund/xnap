package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	SpecialSubframePatternsEUTRASsp0  int64 = 0
	SpecialSubframePatternsEUTRASsp1  int64 = 1
	SpecialSubframePatternsEUTRASsp2  int64 = 2
	SpecialSubframePatternsEUTRASsp3  int64 = 3
	SpecialSubframePatternsEUTRASsp4  int64 = 4
	SpecialSubframePatternsEUTRASsp5  int64 = 5
	SpecialSubframePatternsEUTRASsp6  int64 = 6
	SpecialSubframePatternsEUTRASsp7  int64 = 7
	SpecialSubframePatternsEUTRASsp8  int64 = 8
	SpecialSubframePatternsEUTRASsp9  int64 = 9
	SpecialSubframePatternsEUTRASsp10 int64 = 10
)

var specialSubframePatternsEUTRAConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	ExtValues:  nil,
}

type SpecialSubframePatternsEUTRA struct {
	Value int64
}

func (ie *SpecialSubframePatternsEUTRA) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, specialSubframePatternsEUTRAConstraints)
}

func (ie *SpecialSubframePatternsEUTRA) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(specialSubframePatternsEUTRAConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
