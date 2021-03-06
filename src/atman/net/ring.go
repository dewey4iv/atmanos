package net

import (
	"atman/mm"
	"atman/xen"
	"unsafe"
)

type NetifTxRequest struct {
	Gref   xen.Gref
	Offset uint16
	Flags  uint16
	ID     uint16
	Size   uint16
}

type NetifTxResponse struct {
	ID     uint16
	Status int16
}

func newTxRing(r *xen.SharedRing) *xen.FrontendRing {
	return xen.NewFrontendRing(r, txRingEntrySize())
}

func txRingEntrySize() int {
	size := unsafe.Sizeof(NetifTxRequest{})

	if rspSize := unsafe.Sizeof(NetifTxResponse{}); rspSize > size {
		size = rspSize
	}

	return int(size)
}

type NetifRxRequest struct {
	ID   uint16
	_    uint16
	Gref xen.Gref
}

type NetifRxResponse struct {
	ID     uint16
	Offset uint16
	Flags  uint16
	Status int16
}

func newRxRing(r *xen.SharedRing) *xen.FrontendRing {
	return xen.NewFrontendRing(r, rxRingEntrySize())
}

func rxRingEntrySize() int {
	size := unsafe.Sizeof(NetifRxRequest{})

	if rspSize := unsafe.Sizeof(NetifRxResponse{}); rspSize > size {
		size = rspSize
	}

	return int(size)
}

func initSharedRing(page *mm.Page) *xen.SharedRing {
	ring := (*xen.SharedRing)(page.Ptr)
	ring.RequestEvent = 1
	ring.ResponseEvent = 1

	return ring
}
