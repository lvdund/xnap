package ies

import (
	"github.com/lvdund/asn1go/per"
)

var mBSSessionInformationItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "mBS-Session-ID"},
		{Name: "mBS-Area-Session-ID", Optional: true},
		{Name: "active-MBS-SessioInformation", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type MBSSessionInformationItem struct {
	MBSSessionID               MBSSessionID
	MBSAreaSessionID           *MBSAreaSessionID
	ActiveMBSSessioInformation *ActiveMBSSessionInformation
	IEExtensions               []byte
}

func (ie *MBSSessionInformationItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mBSSessionInformationItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.MBSAreaSessionID != nil, ie.ActiveMBSSessioInformation != nil, false}); err != nil {
		return err
	}
	if err := ie.MBSSessionID.Encode(e); err != nil {
		return err
	}
	if ie.MBSAreaSessionID != nil {
		if err := ie.MBSAreaSessionID.Encode(e); err != nil {
			return err
		}
	}
	if ie.ActiveMBSSessioInformation != nil {
		if err := ie.ActiveMBSSessioInformation.Encode(e); err != nil {
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

func (ie *MBSSessionInformationItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mBSSessionInformationItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.MBSSessionID.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.MBSAreaSessionID = new(MBSAreaSessionID)
		if err := ie.MBSAreaSessionID.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.ActiveMBSSessioInformation = new(ActiveMBSSessionInformation)
		if err := ie.ActiveMBSSessioInformation.Decode(d); err != nil {
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
