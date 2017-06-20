// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=ethernet -id punt_packet_disposition -d PoolType=punt_packet_disposition_pool -d Type=punt_packet_disposition -d Data=dispositions github.com/platinasystems/go/elib/pool.tmpl]

// Copyright 2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ethernet

import (
	"github.com/platinasystems/go/elib"
)

type punt_packet_disposition_pool struct {
	elib.Pool
	dispositions []punt_packet_disposition
}

func (p *punt_packet_disposition_pool) GetIndex() (i uint) {
	l := uint(len(p.dispositions))
	i = p.Pool.GetIndex(l)
	if i >= l {
		p.Validate(i)
	}
	return i
}

func (p *punt_packet_disposition_pool) PutIndex(i uint) (ok bool) {
	return p.Pool.PutIndex(i)
}

func (p *punt_packet_disposition_pool) IsFree(i uint) (v bool) {
	v = i >= uint(len(p.dispositions))
	if !v {
		v = p.Pool.IsFree(i)
	}
	return
}

func (p *punt_packet_disposition_pool) Resize(n uint) {
	c := elib.Index(cap(p.dispositions))
	l := elib.Index(len(p.dispositions) + int(n))
	if l > c {
		c = elib.NextResizeCap(l)
		q := make([]punt_packet_disposition, l, c)
		copy(q, p.dispositions)
		p.dispositions = q
	}
	p.dispositions = p.dispositions[:l]
}

func (p *punt_packet_disposition_pool) Validate(i uint) {
	c := elib.Index(cap(p.dispositions))
	l := elib.Index(i) + 1
	if l > c {
		c = elib.NextResizeCap(l)
		q := make([]punt_packet_disposition, l, c)
		copy(q, p.dispositions)
		p.dispositions = q
	}
	if l > elib.Index(len(p.dispositions)) {
		p.dispositions = p.dispositions[:l]
	}
}

func (p *punt_packet_disposition_pool) Elts() uint {
	return uint(len(p.dispositions)) - p.FreeLen()
}

func (p *punt_packet_disposition_pool) Len() uint {
	return uint(len(p.dispositions))
}

func (p *punt_packet_disposition_pool) Foreach(f func(x punt_packet_disposition)) {
	for i := range p.dispositions {
		if !p.Pool.IsFree(uint(i)) {
			f(p.dispositions[i])
		}
	}
}

func (p *punt_packet_disposition_pool) ForeachIndex(f func(i uint)) {
	for i := range p.dispositions {
		if !p.Pool.IsFree(uint(i)) {
			f(uint(i))
		}
	}
}
