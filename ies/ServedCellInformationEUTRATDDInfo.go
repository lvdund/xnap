package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	ServedCellInformationEUTRATDDInfoSubframeAssignmnetSa0 int64 = 0
	ServedCellInformationEUTRATDDInfoSubframeAssignmnetSa1 int64 = 1
	ServedCellInformationEUTRATDDInfoSubframeAssignmnetSa2 int64 = 2
	ServedCellInformationEUTRATDDInfoSubframeAssignmnetSa3 int64 = 3
	ServedCellInformationEUTRATDDInfoSubframeAssignmnetSa4 int64 = 4
	ServedCellInformationEUTRATDDInfoSubframeAssignmnetSa5 int64 = 5
	ServedCellInformationEUTRATDDInfoSubframeAssignmnetSa6 int64 = 6
)

var servedCellInformationEUTRATDDInfoSubframeAssignmnetConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1, 2, 3, 4, 5, 6},
	ExtValues:  nil,
}

type ServedCellInformationEUTRATDDInfoSubframeAssignmnet struct {
	Value int64
}

func (ie *ServedCellInformationEUTRATDDInfoSubframeAssignmnet) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, servedCellInformationEUTRATDDInfoSubframeAssignmnetConstraints)
}

func (ie *ServedCellInformationEUTRATDDInfoSubframeAssignmnet) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(servedCellInformationEUTRATDDInfoSubframeAssignmnetConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var servedCellInformationEUTRATDDInfoConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "earfcn"},
		{Name: "e-utraTxBW"},
		{Name: "subframeAssignmnet"},
		{Name: "specialSubframeInfo"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type ServedCellInformationEUTRATDDInfo struct {
	Earfcn              EUTRAARFCN
	EUtraTxBW           EUTRATransmissionBandwidth
	SubframeAssignmnet  ServedCellInformationEUTRATDDInfoSubframeAssignmnet
	SpecialSubframeInfo SpecialSubframeInfoEUTRA
	IEExtensions        []byte
}

func (ie *ServedCellInformationEUTRATDDInfo) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(servedCellInformationEUTRATDDInfoConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.Earfcn.Encode(e); err != nil {
		return err
	}
	if err := ie.EUtraTxBW.Encode(e); err != nil {
		return err
	}
	if err := ie.SubframeAssignmnet.Encode(e); err != nil {
		return err
	}
	if err := ie.SpecialSubframeInfo.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *ServedCellInformationEUTRATDDInfo) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(servedCellInformationEUTRATDDInfoConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.Earfcn.Decode(d); err != nil {
		return err
	}
	if err := ie.EUtraTxBW.Decode(d); err != nil {
		return err
	}
	if err := ie.SubframeAssignmnet.Decode(d); err != nil {
		return err
	}
	if err := ie.SpecialSubframeInfo.Decode(d); err != nil {
		return err
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
