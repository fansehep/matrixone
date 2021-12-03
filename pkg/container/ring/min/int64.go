// Copyright 2021 Matrix Origin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package min

import (
	"fmt"
	"math"
	"github.com/matrixorigin/matrixone/pkg/container/nulls"
	"github.com/matrixorigin/matrixone/pkg/container/ring"
	"github.com/matrixorigin/matrixone/pkg/container/types"
	"github.com/matrixorigin/matrixone/pkg/container/vector"
	"github.com/matrixorigin/matrixone/pkg/encoding"
	"github.com/matrixorigin/matrixone/pkg/vm/mheap"
)

func NewInt64(typ types.Type) *Int64Ring {
	return &Int64Ring{Typ: typ}
}

func (r *Int64Ring) String() string {
	return fmt.Sprintf("%v-%v", r.Vs, r.Ns)
}

func (r *Int64Ring) Free(m *mheap.Mheap) {
	if r.Da != nil {
		mheap.Free(m, r.Da)
		r.Da = nil
		r.Vs = nil
		r.Ns = nil
	}
}

func (r *Int64Ring) Count() int {
	return len(r.Vs)
}

func (r *Int64Ring) Size() int {
	return cap(r.Da)
}

func (r *Int64Ring) Dup() ring.Ring {
	return &Int64Ring{
		Typ: r.Typ,
	}
}

func (r *Int64Ring) Type() types.Type {
	return r.Typ
}

func (r *Int64Ring) SetLength(n int) {
	r.Vs = r.Vs[:n]
	r.Ns = r.Ns[:n]
}

func (r *Int64Ring) Shrink(sels []int64) {
	for i, sel := range sels {
		r.Vs[i] = r.Vs[sel]
		r.Ns[i] = r.Ns[sel]
	}
	r.Vs = r.Vs[:len(sels)]
	r.Ns = r.Ns[:len(sels)]
}

func (r *Int64Ring) Shuffle(_ []int64, _ *mheap.Mheap) error {
	return nil
}

func (r *Int64Ring) Grow(m *mheap.Mheap) error {
	n := len(r.Vs)
	if n == 0 {
		data, err := mheap.Alloc(m, 8*8)
		if err != nil {
			return err
		}
		r.Da = data
		r.Ns = make([]int64, 0, 8)
		r.Vs = encoding.DecodeInt64Slice(data)
	} else if n+1 >= cap(r.Vs) {
		r.Da = r.Da[:n*8]
		data, err := mheap.Grow(m, r.Da, int64(n+1)*8)
		if err != nil {
			return err
		}
		mheap.Free(m, r.Da)
		r.Da = data
		r.Vs = encoding.DecodeInt64Slice(data)
	}
	r.Vs = r.Vs[:n+1]
	r.Vs[n] = math.MaxInt64
	r.Ns = append(r.Ns, 0)
	return nil
}

func (r *Int64Ring) Grows(size int, m *mheap.Mheap) error {
	n := len(r.Vs)
	if n == 0 {
		data, err := mheap.Alloc(m, int64(size*8))
		if err != nil {
			return err
		}
		r.Da = data
		r.Ns = make([]int64, 0, size)
		r.Vs = encoding.DecodeInt64Slice(data)
	} else if n+size >= cap(r.Vs) {
		r.Da = r.Da[:n*8]
		data, err := mheap.Grow(m, r.Da, int64(n+size)*8)
		if err != nil {
			return err
		}
		mheap.Free(m, r.Da)
		r.Da = data
		r.Vs = encoding.DecodeInt64Slice(data)
	}
	r.Vs = r.Vs[:n+size]
	for i := 0; i < size; i++ {
		r.Ns = append(r.Ns, 0)
		r.Vs[i+n] = math.MaxInt64
	}
	return nil
}

func (r *Int64Ring) Fill(i int64, sel, z int64, vec *vector.Vector) {
	if v := vec.Col.([]int64)[sel]; v < r.Vs[i] {
		r.Vs[i] = v
	}
	if nulls.Contains(vec.Nsp, uint64(sel)) {
		r.Ns[i] += z
	}
}

func (r *Int64Ring) BatchFill(start int64, os []uint8, vps []*uint64, zs []int64, vec *vector.Vector) {
	vs := vec.Col.([]int64)
	for i, o := range os {
		if o == 1 {
			j := *vps[i]
			if vs[int64(i)+start] < r.Vs[j] {
				r.Vs[j] = vs[int64(i)+start]
			}
		}
	}
	if nulls.Any(vec.Nsp) {
		for i, o := range os {
			if o == 1 {
				if nulls.Contains(vec.Nsp, uint64(start)+uint64(i)) {
					r.Ns[*vps[i]] += zs[int64(i)+start]
				}
			}
		}
	}
}

func (r *Int64Ring) BulkFill(i int64, zs []int64, vec *vector.Vector) {
	vs := vec.Col.([]int64)
	for _, v := range vs {
		if v < r.Vs[i] {
			r.Vs[i] = v
		}
	}
	if nulls.Any(vec.Nsp) {
		for j := range vs {
			if nulls.Contains(vec.Nsp, uint64(j)) {
				r.Ns[i] += zs[j]
			}
		}
	}
}

func (r *Int64Ring) Add(a interface{}, x, y int64) {
	ar := a.(*Int64Ring)
	if ar.Vs[y] < r.Vs[x] {
		r.Vs[x] = ar.Vs[y]
	}
	r.Ns[x] += ar.Ns[y]
}

func (r *Int64Ring) BatchAdd(a interface{}, start int64, os []uint8, vps []*uint64) {
	ar := a.(*Int64Ring)
	for i, o := range os {
		if o == 1 {
			j := *vps[i]
			if ar.Vs[int64(i)+start] < r.Vs[j] {
				r.Vs[j] = ar.Vs[int64(i)+start]
			}
			r.Ns[j] += ar.Ns[int64(i)+start]
		}
	}
}

func (r *Int64Ring) Mul(x, z int64) {
	r.Ns[x] *= z
}

func (r *Int64Ring) Eval(zs []int64) *vector.Vector {
	defer func() {
		r.Da = nil
		r.Vs = nil
		r.Ns = nil
	}()
	nsp := new(nulls.Nulls)
	for i, z := range zs {
		if z-r.Ns[i] == 0 {
			nulls.Add(nsp, uint64(i))
		}
	}
	return &vector.Vector{
		Nsp:  nsp,
		Data: r.Da,
		Col:  r.Vs,
		Or:   false,
		Typ:  r.Typ,
	}
}