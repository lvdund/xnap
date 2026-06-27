package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	SDTAssistantInfoSinglePacket    int64 = 0
	SDTAssistantInfoMultiplePackets int64 = 1
)

var sDTAssistantInfoConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type SDTAssistantInfo struct {
	Value int64
}

func (ie *SDTAssistantInfo) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, sDTAssistantInfoConstraints)
}

func (ie *SDTAssistantInfo) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(sDTAssistantInfoConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
