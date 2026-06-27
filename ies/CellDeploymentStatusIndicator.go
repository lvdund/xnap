package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	CellDeploymentStatusIndicatorPreChangeNotification int64 = 0
)

var cellDeploymentStatusIndicatorConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type CellDeploymentStatusIndicator struct {
	Value int64
}

func (ie *CellDeploymentStatusIndicator) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, cellDeploymentStatusIndicatorConstraints)
}

func (ie *CellDeploymentStatusIndicator) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(cellDeploymentStatusIndicatorConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
