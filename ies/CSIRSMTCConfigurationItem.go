package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

const (
	CSIRSMTCConfigurationItemCsiRSStatusActivated   int64 = 0
	CSIRSMTCConfigurationItemCsiRSStatusDeactivated int64 = 1
)

var cSIRSMTCConfigurationItemCsiRSStatusConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type CSIRSMTCConfigurationItemCsiRSStatus struct {
	Value int64
}

func (ie *CSIRSMTCConfigurationItemCsiRSStatus) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, cSIRSMTCConfigurationItemCsiRSStatusConstraints)
}

func (ie *CSIRSMTCConfigurationItemCsiRSStatus) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(cSIRSMTCConfigurationItemCsiRSStatusConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var cSIRSMTCConfigurationItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "csi-RS-Index"},
		{Name: "csi-RS-Status"},
		{Name: "csi-RS-Neighbour-List", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type CSIRSMTCConfigurationItem struct {
	CsiRSIndex         int64
	CsiRSStatus        CSIRSMTCConfigurationItemCsiRSStatus
	CsiRSNeighbourList *CSIRSNeighbourList
	IEExtensions       []byte
}

func (ie *CSIRSMTCConfigurationItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cSIRSMTCConfigurationItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.CsiRSNeighbourList != nil, false}); err != nil {
		return err
	}
	if err := e.EncodeInteger(ie.CsiRSIndex, per.IntegerConstraints{
		Extensible: false,
		LowerBound: common.Ptr(int64(0)),
		UpperBound: common.Ptr(int64(95)),
	}); err != nil {
		return err
	}
	if err := ie.CsiRSStatus.Encode(e); err != nil {
		return err
	}
	if ie.CsiRSNeighbourList != nil {
		if err := ie.CsiRSNeighbourList.Encode(e); err != nil {
			return err
		}
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *CSIRSMTCConfigurationItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cSIRSMTCConfigurationItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	{
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(0)),
			UpperBound: common.Ptr(int64(95)),
		})
		if err != nil {
			return err
		}
		ie.CsiRSIndex = val
	}
	if err := ie.CsiRSStatus.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(2) {
		ie.CsiRSNeighbourList = new(CSIRSNeighbourList)
		if err := ie.CsiRSNeighbourList.Decode(d); err != nil {
			return err
		}
	}
	extBytes, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	if len(extBytes) > 0 {
		ie.IEExtensions = extBytes[0]
	}
	return nil
}
