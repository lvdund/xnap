package ies

import (
	"github.com/lvdund/asn1go/per"
)

var userPlaneFailureIndicationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "userPlaneFailureType"},
		{Name: "dL-NG-U-TNLatNG-RAN"},
		{Name: "uL-NG-U-TNLatNG-RAN"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type UserPlaneFailureIndication struct {
	UserPlaneFailureType UserPlaneFailureType
	DLNGUTNLatNGRAN      UPTransportLayerInformation
	ULNGUTNLatNGRAN      UPTransportLayerInformation
	IEExtensions         []byte
}

func (ie *UserPlaneFailureIndication) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(userPlaneFailureIndicationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.UserPlaneFailureType.Encode(e); err != nil {
		return err
	}
	if err := ie.DLNGUTNLatNGRAN.Encode(e); err != nil {
		return err
	}
	if err := ie.ULNGUTNLatNGRAN.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *UserPlaneFailureIndication) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(userPlaneFailureIndicationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.UserPlaneFailureType.Decode(d); err != nil {
		return err
	}
	if err := ie.DLNGUTNLatNGRAN.Decode(d); err != nil {
		return err
	}
	if err := ie.ULNGUTNLatNGRAN.Decode(d); err != nil {
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
