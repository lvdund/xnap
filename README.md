# xnap

A Go library for the **Xn Application Protocol (XnAP)** defined by
**3GPP TS 38.423 V18.6.0 (2025-06)** - the signalling protocol used over the
**Xn interface** between NG-RAN nodes (gNB / ng-eNB) in 5G systems.

This package provides encode/decode support for the full set of XnAP elementary
procedures, Information Elements (IEs) and PDU structures, generated from the
ASN.1 definitions of TS 38.423 and encoded/decoded using ITU-T **APER**
(ALIGNED PER) as required by the specification.

## Features

- **Full TS 38.423 V18.6.0 (2025-06) coverage**: NG-RAN; Xn application protocol (XnAP)
- **APER encode / decode** based on [`github.com/lvdund/asn1go`](https://github.com/lvdund/asn1go)
- **Wire-ready**: simple `[]byte` ⇄ Go struct API — no extra dependencies for
  the radio network itself.
- **Converter helpers** in the [`utils`](./utils) package for translating between
  XnAP/ASN.1 representations and common Go/data-model types
  (SNSSAI, PLMN ID, IP addresses, bit strings).

## API reference

| Function / Type        | Location            | Description                                              |
| ---------------------- | ------------------- | -------------------------------------------------------- |
| `xnap.XnDecode`        | `decode.go`         | Decode an APER `[]byte` into an `XnApPDU`.               |
| `xnap.XnEncode`        | `encode.go`         | Encode an `XnApMessageEncoder` into an APER `[]byte`.    |
| `xnap.XnApPDU`         | `xnap.go`           | Top-level XnAP-PDU container (choice + procedure + msg). |
| `ies.*`                | `ies/`              | All Information Elements and procedure messages.         |
| `common.*`             | `common/`           | Procedure codes, IE IDs, criticality, choice constants.  |
| `utils.*`              | `utils/`            | ASN.1 ⇄ model converters.                                |

## License

Distributed under the [Apache License 2.0](./LICENSE).
