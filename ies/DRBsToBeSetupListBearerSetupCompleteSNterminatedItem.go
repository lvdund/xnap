package ies

import (
	"github.com/lvdund/asn1go/per"
)

var dRBsToBeSetupListBearerSetupCompleteSNterminatedItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dRB-ID"},
		{Name: "mN-Xn-U-TNLInfoatM"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type DRBsToBeSetupListBearerSetupCompleteSNterminatedItem struct {
	DRBID           DRBID
	MNXnUTNLInfoatM UPTransportLayerInformation
	IEExtensions    []byte
}

func (ie *DRBsToBeSetupListBearerSetupCompleteSNterminatedItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dRBsToBeSetupListBearerSetupCompleteSNterminatedItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.DRBID.Encode(e); err != nil {
		return err
	}
	if err := ie.MNXnUTNLInfoatM.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *DRBsToBeSetupListBearerSetupCompleteSNterminatedItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dRBsToBeSetupListBearerSetupCompleteSNterminatedItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.DRBID.Decode(d); err != nil {
		return err
	}
	if err := ie.MNXnUTNLInfoatM.Decode(d); err != nil {
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
