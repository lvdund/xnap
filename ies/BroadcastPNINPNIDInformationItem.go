package ies

import (
	"github.com/lvdund/asn1go/per"
)

var broadcastPNINPNIDInformationItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "plmn-id"},
		{Name: "broadcastCAG-Identifier-List"},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type BroadcastPNINPNIDInformationItem struct {
	PlmnId                     PLMNIdentity
	BroadcastCAGIdentifierList BroadcastCAGIdentifierList
	IEExtensions               []byte
}

func (ie *BroadcastPNINPNIDInformationItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(broadcastPNINPNIDInformationItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.PlmnId.Encode(e); err != nil {
		return err
	}
	if err := ie.BroadcastCAGIdentifierList.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *BroadcastPNINPNIDInformationItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(broadcastPNINPNIDInformationItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.PlmnId.Decode(d); err != nil {
		return err
	}
	if err := ie.BroadcastCAGIdentifierList.Decode(d); err != nil {
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
