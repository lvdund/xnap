package ies

import (
	"github.com/lvdund/asn1go/per"
)

var pDUSessionResourceModRqdInfoSNterminatedConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dL-NG-U-TNLatNG-RAN", Optional: true},
		{Name: "qoSFlowsToBeReleased-List", Optional: true},
		{Name: "dataforwardinginfofromSource", Optional: true},
		{Name: "drbsToBeSetupList", Optional: true},
		{Name: "drbsToBeModifiedList", Optional: true},
		{Name: "dRBsToBeReleased", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type PDUSessionResourceModRqdInfoSNterminated struct {
	DLNGUTNLatNGRAN              *UPTransportLayerInformation
	QoSFlowsToBeReleasedList     *QoSFlowsListWithCause
	DataforwardinginfofromSource *DataforwardingandOffloadingInfofromSource
	DrbsToBeSetupList            *DRBsToBeSetupListModRqdSNterminated
	DrbsToBeModifiedList         *DRBsToBeModifiedListModRqdSNterminated
	DRBsToBeReleased             *DRBListWithCause
	IEExtensions                 []byte
}

func (ie *PDUSessionResourceModRqdInfoSNterminated) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDUSessionResourceModRqdInfoSNterminatedConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.DLNGUTNLatNGRAN != nil, ie.QoSFlowsToBeReleasedList != nil, ie.DataforwardinginfofromSource != nil, ie.DrbsToBeSetupList != nil, ie.DrbsToBeModifiedList != nil, ie.DRBsToBeReleased != nil, false}); err != nil {
		return err
	}
	if ie.DLNGUTNLatNGRAN != nil {
		if err := ie.DLNGUTNLatNGRAN.Encode(e); err != nil {
			return err
		}
	}
	if ie.QoSFlowsToBeReleasedList != nil {
		if err := ie.QoSFlowsToBeReleasedList.Encode(e); err != nil {
			return err
		}
	}
	if ie.DataforwardinginfofromSource != nil {
		if err := ie.DataforwardinginfofromSource.Encode(e); err != nil {
			return err
		}
	}
	if ie.DrbsToBeSetupList != nil {
		if err := ie.DrbsToBeSetupList.Encode(e); err != nil {
			return err
		}
	}
	if ie.DrbsToBeModifiedList != nil {
		if err := ie.DrbsToBeModifiedList.Encode(e); err != nil {
			return err
		}
	}
	if ie.DRBsToBeReleased != nil {
		if err := ie.DRBsToBeReleased.Encode(e); err != nil {
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

func (ie *PDUSessionResourceModRqdInfoSNterminated) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDUSessionResourceModRqdInfoSNterminatedConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.DLNGUTNLatNGRAN = new(UPTransportLayerInformation)
		if err := ie.DLNGUTNLatNGRAN.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(1) {
		ie.QoSFlowsToBeReleasedList = new(QoSFlowsListWithCause)
		if err := ie.QoSFlowsToBeReleasedList.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.DataforwardinginfofromSource = new(DataforwardingandOffloadingInfofromSource)
		if err := ie.DataforwardinginfofromSource.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(3) {
		ie.DrbsToBeSetupList = new(DRBsToBeSetupListModRqdSNterminated)
		if err := ie.DrbsToBeSetupList.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(4) {
		ie.DrbsToBeModifiedList = new(DRBsToBeModifiedListModRqdSNterminated)
		if err := ie.DrbsToBeModifiedList.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(5) {
		ie.DRBsToBeReleased = new(DRBListWithCause)
		if err := ie.DRBsToBeReleased.Decode(d); err != nil {
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
