package ies

import (
	"github.com/lvdund/asn1go/per"
)

var raReportIndicationListItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "m-NG-RAN-node-UE-XnAP-ID"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type RaReportIndicationListItem struct {
	MNGRANNodeUEXnAPID NGRANnodeUEXnAPID
	IEExtensions       []byte
}

func (ie *RaReportIndicationListItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(raReportIndicationListItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.MNGRANNodeUEXnAPID.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *RaReportIndicationListItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(raReportIndicationListItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.MNGRANNodeUEXnAPID.Decode(d); err != nil {
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
