//+build openbsd,386

// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs defs.go

package wgh

const SIOCGIFGMEMB = 0xc024698a

type Ifgroupreq struct {
	Name  [16]byte
	Len   uint32
	Ifgru [16]byte
}

type Ifgreq struct {
	Ifgrqu [16]byte
}

type Timespec struct {
	Sec  int64
	Nsec int32
}

const (
	SIOCGWGSERV = 0xc03c69c8
	SIOCGWGPEER = 0xc09469c9

	WGStateNoSession = 0x0
	WGStateInitiator = 0x1
	WGStateResponder = 0x2
)

type WGGetServ struct {
	Name      [16]byte
	Pubkey    [32]byte
	Port      uint16
	Num_peers uint32
	Peers     *[32]byte
}

type WGGetPeer struct {
	Name           [16]byte
	Pubkey         [32]byte
	Psk            [32]byte
	Tx_bytes       uint64
	Rx_bytes       uint64
	Ip             [28]byte
	State          uint32
	Last_handshake Timespec
	Num_aip        uint32
	Aip            *[28]byte
}

type WGIP [28]byte
