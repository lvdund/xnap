package ies

import (
	"github.com/lvdund/asn1go/per"
)

var nonF1TerminatingBHInformationItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "bHInfoIndex"},
		{Name: "dlNon-F1TerminatingBHInfo", Optional: true},
		{Name: "ulNon-F1TerminatingBHInfo", Optional: true},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type NonF1TerminatingBHInformationItem struct {
	BHInfoIndex              BHInfoIndex
	DlNonF1TerminatingBHInfo *DLNonF1TerminatingBHInfo
	UlNonF1TerminatingBHInfo *ULNonF1TerminatingBHInfo
	IEExtensions             []byte
}

func (ie *NonF1TerminatingBHInformationItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nonF1TerminatingBHInformationItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.DlNonF1TerminatingBHInfo != nil, ie.UlNonF1TerminatingBHInfo != nil, false}); err != nil {
		return err
	}
	if err := ie.BHInfoIndex.Encode(e); err != nil {
		return err
	}
	if ie.DlNonF1TerminatingBHInfo != nil {
		if err := ie.DlNonF1TerminatingBHInfo.Encode(e); err != nil {
			return err
		}
	}
	if ie.UlNonF1TerminatingBHInfo != nil {
		if err := ie.UlNonF1TerminatingBHInfo.Encode(e); err != nil {
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

func (ie *NonF1TerminatingBHInformationItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nonF1TerminatingBHInformationItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.BHInfoIndex.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.DlNonF1TerminatingBHInfo = new(DLNonF1TerminatingBHInfo)
		if err := ie.DlNonF1TerminatingBHInfo.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.UlNonF1TerminatingBHInfo = new(ULNonF1TerminatingBHInfo)
		if err := ie.UlNonF1TerminatingBHInfo.Decode(d); err != nil {
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
