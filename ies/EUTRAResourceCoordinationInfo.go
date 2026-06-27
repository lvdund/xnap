package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var eUTRAResourceCoordinationInfoConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "e-utra-cell"},
		{Name: "ul-coordination-info"},
		{Name: "dl-coordination-info", Optional: true},
		{Name: "nr-cell", Optional: true},
		{Name: "e-utra-coordination-assistance-info", Optional: true},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type EUTRAResourceCoordinationInfo struct {
	EUtraCell                       EUTRACGI
	UlCoordinationInfo              per.BitString
	DlCoordinationInfo              *per.BitString
	NrCell                          *NRCGI
	EUtraCoordinationAssistanceInfo *EUTRACoordinationAssistanceInfo
	IEExtensions                    []byte
}

func (ie *EUTRAResourceCoordinationInfo) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(eUTRAResourceCoordinationInfoConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.DlCoordinationInfo != nil, ie.NrCell != nil, ie.EUtraCoordinationAssistanceInfo != nil, false}); err != nil {
		return err
	}
	if err := ie.EUtraCell.Encode(e); err != nil {
		return err
	}
	if err := e.EncodeBitString(ie.UlCoordinationInfo, per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(6)),
		Max:        common.Ptr(int64(4400)),
	}); err != nil {
		return err
	}
	if ie.DlCoordinationInfo != nil {
		if err := e.EncodeBitString(*ie.DlCoordinationInfo, per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(6)),
			Max:        common.Ptr(int64(4400)),
		}); err != nil {
			return err
		}
	}
	if ie.NrCell != nil {
		if err := ie.NrCell.Encode(e); err != nil {
			return err
		}
	}
	if ie.EUtraCoordinationAssistanceInfo != nil {
		if err := ie.EUtraCoordinationAssistanceInfo.Encode(e); err != nil {
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

func (ie *EUTRAResourceCoordinationInfo) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(eUTRAResourceCoordinationInfoConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.EUtraCell.Decode(d); err != nil {
		return err
	}
	{
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(6)),
			Max:        common.Ptr(int64(4400)),
		})
		if err != nil {
			return err
		}
		ie.UlCoordinationInfo = val
	}
	if seq.IsComponentPresent(2) {
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(6)),
			Max:        common.Ptr(int64(4400)),
		})
		if err != nil {
			return err
		}
		ie.DlCoordinationInfo = &val
	}
	if seq.IsComponentPresent(3) {
		ie.NrCell = new(NRCGI)
		if err := ie.NrCell.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(4) {
		ie.EUtraCoordinationAssistanceInfo = new(EUTRACoordinationAssistanceInfo)
		if err := ie.EUtraCoordinationAssistanceInfo.Decode(d); err != nil {
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
