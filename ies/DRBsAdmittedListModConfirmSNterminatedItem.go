package ies

import (
	"github.com/lvdund/asn1go/per"
)

var dRBsAdmittedListModConfirmSNterminatedItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "drb-ID"},
		{Name: "mN-DL-CG-UP-TNLInfo", Optional: true},
		{Name: "secondary-MN-DL-CG-UP-TNLInfo", Optional: true},
		{Name: "lCID", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type DRBsAdmittedListModConfirmSNterminatedItem struct {
	DrbID                    DRBID
	MNDLCGUPTNLInfo          *UPTransportParameters
	SecondaryMNDLCGUPTNLInfo *UPTransportParameters
	LCID                     *LCID
	IEExtensions             []byte
}

func (ie *DRBsAdmittedListModConfirmSNterminatedItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dRBsAdmittedListModConfirmSNterminatedItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.MNDLCGUPTNLInfo != nil, ie.SecondaryMNDLCGUPTNLInfo != nil, ie.LCID != nil, false}); err != nil {
		return err
	}
	if err := ie.DrbID.Encode(e); err != nil {
		return err
	}
	if ie.MNDLCGUPTNLInfo != nil {
		if err := ie.MNDLCGUPTNLInfo.Encode(e); err != nil {
			return err
		}
	}
	if ie.SecondaryMNDLCGUPTNLInfo != nil {
		if err := ie.SecondaryMNDLCGUPTNLInfo.Encode(e); err != nil {
			return err
		}
	}
	if ie.LCID != nil {
		if err := ie.LCID.Encode(e); err != nil {
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

func (ie *DRBsAdmittedListModConfirmSNterminatedItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dRBsAdmittedListModConfirmSNterminatedItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.DrbID.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.MNDLCGUPTNLInfo = new(UPTransportParameters)
		if err := ie.MNDLCGUPTNLInfo.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.SecondaryMNDLCGUPTNLInfo = new(UPTransportParameters)
		if err := ie.SecondaryMNDLCGUPTNLInfo.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(3) {
		ie.LCID = new(LCID)
		if err := ie.LCID.Decode(d); err != nil {
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
