package ies

import (
	"github.com/lvdund/asn1go/per"
)

var dataCollectionIDConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nGRAN-Node1-Measurement-ID"},
		{Name: "nGRAN-Node2-Measurement-ID"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type DataCollectionID struct {
	NGRANNode1MeasurementID MeasurementID
	NGRANNode2MeasurementID MeasurementID
	IEExtensions            []byte
}

func (ie *DataCollectionID) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dataCollectionIDConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.NGRANNode1MeasurementID.Encode(e); err != nil {
		return err
	}
	if err := ie.NGRANNode2MeasurementID.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *DataCollectionID) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dataCollectionIDConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.NGRANNode1MeasurementID.Decode(d); err != nil {
		return err
	}
	if err := ie.NGRANNode2MeasurementID.Decode(d); err != nil {
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
