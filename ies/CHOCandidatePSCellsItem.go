package ies

import (
	"github.com/lvdund/asn1go/per"
)

var cHOCandidatePSCellsItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pscell-id"},
		{Name: "target2source-NG-RANNode-Container"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type CHOCandidatePSCellsItem struct {
	PscellId                        NRCGI
	Target2sourceNGRANNodeContainer []byte
	IEExtensions                    []byte
}

func (ie *CHOCandidatePSCellsItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cHOCandidatePSCellsItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.PscellId.Encode(e); err != nil {
		return err
	}
	if err := e.EncodeOctetString(ie.Target2sourceNGRANNodeContainer, per.SizeConstraints{
		Extensible: false,
		Min:        nil,
		Max:        nil,
	}); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *CHOCandidatePSCellsItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cHOCandidatePSCellsItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.PscellId.Decode(d); err != nil {
		return err
	}
	{
		val, err := d.DecodeOctetString(per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		})
		if err != nil {
			return err
		}
		ie.Target2sourceNGRANNodeContainer = val
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
