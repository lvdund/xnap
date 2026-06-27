package ies

import (
	"github.com/lvdund/asn1go/per"
)

var fiveGProSeAuthorizedConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "fiveGproSeDirectDiscovery", Optional: true},
		{Name: "fiveGproSeDirectCommunication", Optional: true},
		{Name: "fiveGnrProSeLayer2UEtoNetworkRelay", Optional: true},
		{Name: "fiveGnrProSeLayer3UEtoNetworkRelay", Optional: true},
		{Name: "fiveGnrProSeLayer2RemoteUE", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type FiveGProSeAuthorized struct {
	FiveGproSeDirectDiscovery          *FiveGProSeDirectDiscovery
	FiveGproSeDirectCommunication      *FiveGProSeDirectCommunication
	FiveGnrProSeLayer2UEtoNetworkRelay *FiveGProSeLayer2UEtoNetworkRelay
	FiveGnrProSeLayer3UEtoNetworkRelay *FiveGProSeLayer3UEtoNetworkRelay
	FiveGnrProSeLayer2RemoteUE         *FiveGProSeLayer2RemoteUE
	IEExtensions                       []byte
}

func (ie *FiveGProSeAuthorized) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(fiveGProSeAuthorizedConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.FiveGproSeDirectDiscovery != nil, ie.FiveGproSeDirectCommunication != nil, ie.FiveGnrProSeLayer2UEtoNetworkRelay != nil, ie.FiveGnrProSeLayer3UEtoNetworkRelay != nil, ie.FiveGnrProSeLayer2RemoteUE != nil, false}); err != nil {
		return err
	}
	if ie.FiveGproSeDirectDiscovery != nil {
		if err := ie.FiveGproSeDirectDiscovery.Encode(e); err != nil {
			return err
		}
	}
	if ie.FiveGproSeDirectCommunication != nil {
		if err := ie.FiveGproSeDirectCommunication.Encode(e); err != nil {
			return err
		}
	}
	if ie.FiveGnrProSeLayer2UEtoNetworkRelay != nil {
		if err := ie.FiveGnrProSeLayer2UEtoNetworkRelay.Encode(e); err != nil {
			return err
		}
	}
	if ie.FiveGnrProSeLayer3UEtoNetworkRelay != nil {
		if err := ie.FiveGnrProSeLayer3UEtoNetworkRelay.Encode(e); err != nil {
			return err
		}
	}
	if ie.FiveGnrProSeLayer2RemoteUE != nil {
		if err := ie.FiveGnrProSeLayer2RemoteUE.Encode(e); err != nil {
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

func (ie *FiveGProSeAuthorized) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(fiveGProSeAuthorizedConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.FiveGproSeDirectDiscovery = new(FiveGProSeDirectDiscovery)
		if err := ie.FiveGproSeDirectDiscovery.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(1) {
		ie.FiveGproSeDirectCommunication = new(FiveGProSeDirectCommunication)
		if err := ie.FiveGproSeDirectCommunication.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.FiveGnrProSeLayer2UEtoNetworkRelay = new(FiveGProSeLayer2UEtoNetworkRelay)
		if err := ie.FiveGnrProSeLayer2UEtoNetworkRelay.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(3) {
		ie.FiveGnrProSeLayer3UEtoNetworkRelay = new(FiveGProSeLayer3UEtoNetworkRelay)
		if err := ie.FiveGnrProSeLayer3UEtoNetworkRelay.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(4) {
		ie.FiveGnrProSeLayer2RemoteUE = new(FiveGProSeLayer2RemoteUE)
		if err := ie.FiveGnrProSeLayer2RemoteUE.Decode(d); err != nil {
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
