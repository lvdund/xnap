package ies

import (
	"github.com/lvdund/asn1go/per"
)

var iABMTCellListItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nRCellIdentity"},
		{Name: "dU-RX-MT-RX"},
		{Name: "dU-TX-MT-TX"},
		{Name: "dU-RX-MT-TX"},
		{Name: "dU-TX-MT-RX"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type IABMTCellListItem struct {
	NRCellIdentity NRCellIdentity
	DURXMTRX       DURXMTRX
	DUTXMTTX       DUTXMTTX
	DURXMTTX       DURXMTTX
	DUTXMTRX       DUTXMTRX
	IEExtensions   []byte
}

func (ie *IABMTCellListItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(iABMTCellListItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.NRCellIdentity.Encode(e); err != nil {
		return err
	}
	if err := ie.DURXMTRX.Encode(e); err != nil {
		return err
	}
	if err := ie.DUTXMTTX.Encode(e); err != nil {
		return err
	}
	if err := ie.DURXMTTX.Encode(e); err != nil {
		return err
	}
	if err := ie.DUTXMTRX.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *IABMTCellListItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(iABMTCellListItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.NRCellIdentity.Decode(d); err != nil {
		return err
	}
	if err := ie.DURXMTRX.Decode(d); err != nil {
		return err
	}
	if err := ie.DUTXMTTX.Decode(d); err != nil {
		return err
	}
	if err := ie.DURXMTTX.Decode(d); err != nil {
		return err
	}
	if err := ie.DUTXMTRX.Decode(d); err != nil {
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
