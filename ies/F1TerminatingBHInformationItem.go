package ies

import (
	"github.com/lvdund/asn1go/per"
)

var f1TerminatingBHInformationItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "bHInfoIndex"},
		{Name: "dLTNLAddress"},
		{Name: "dlF1TerminatingBHInfo", Optional: true},
		{Name: "ulF1TerminatingBHInfo", Optional: true},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type F1TerminatingBHInformationItem struct {
	BHInfoIndex           BHInfoIndex
	DLTNLAddress          IABTNLAddress
	DlF1TerminatingBHInfo *DLF1TerminatingBHInfo
	UlF1TerminatingBHInfo *ULF1TerminatingBHInfo
	IEExtensions          []byte
}

func (ie *F1TerminatingBHInformationItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(f1TerminatingBHInformationItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.DlF1TerminatingBHInfo != nil, ie.UlF1TerminatingBHInfo != nil, false}); err != nil {
		return err
	}
	if err := ie.BHInfoIndex.Encode(e); err != nil {
		return err
	}
	if err := ie.DLTNLAddress.Encode(e); err != nil {
		return err
	}
	if ie.DlF1TerminatingBHInfo != nil {
		if err := ie.DlF1TerminatingBHInfo.Encode(e); err != nil {
			return err
		}
	}
	if ie.UlF1TerminatingBHInfo != nil {
		if err := ie.UlF1TerminatingBHInfo.Encode(e); err != nil {
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

func (ie *F1TerminatingBHInformationItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(f1TerminatingBHInformationItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.BHInfoIndex.Decode(d); err != nil {
		return err
	}
	if err := ie.DLTNLAddress.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(2) {
		ie.DlF1TerminatingBHInfo = new(DLF1TerminatingBHInfo)
		if err := ie.DlF1TerminatingBHInfo.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(3) {
		ie.UlF1TerminatingBHInfo = new(ULF1TerminatingBHInfo)
		if err := ie.UlF1TerminatingBHInfo.Decode(d); err != nil {
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
