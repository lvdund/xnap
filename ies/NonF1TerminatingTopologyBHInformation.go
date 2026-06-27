package ies

import (
	"github.com/lvdund/asn1go/per"
)

var nonF1TerminatingTopologyBHInformationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nonF1TerminatingBHInformation-List"},
		{Name: "bAPControlPDURLCCH-List", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type NonF1TerminatingTopologyBHInformation struct {
	NonF1TerminatingBHInformationList NonF1TerminatingBHInformationList
	BAPControlPDURLCCHList            *BAPControlPDURLCCHList
	IEExtensions                      []byte
}

func (ie *NonF1TerminatingTopologyBHInformation) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nonF1TerminatingTopologyBHInformationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.BAPControlPDURLCCHList != nil, false}); err != nil {
		return err
	}
	if err := ie.NonF1TerminatingBHInformationList.Encode(e); err != nil {
		return err
	}
	if ie.BAPControlPDURLCCHList != nil {
		if err := ie.BAPControlPDURLCCHList.Encode(e); err != nil {
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

func (ie *NonF1TerminatingTopologyBHInformation) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nonF1TerminatingTopologyBHInformationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.NonF1TerminatingBHInformationList.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.BAPControlPDURLCCHList = new(BAPControlPDURLCCHList)
		if err := ie.BAPControlPDURLCCHList.Decode(d); err != nil {
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
