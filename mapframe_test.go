/*
Copyright 2018 Iguazio Systems Ltd.

Licensed under the Apache License, Version 2.0 (the "License") with
an addition restriction as set forth herein. You may not use this
file except in compliance with the License. You may obtain a copy of
the License at http://www.apache.org/licenses/LICENSE-2.0.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
implied. See the License for the specific language governing
permissions and limitations under the License.

In addition, you may not use the software for any purposes that are
illegal under applicable law, and the grant of the foregoing license
under the Apache 2.0 license is conditioned upon your compliance with
such restriction.
*/

package frames

import (
	"fmt"
	"testing"
)

func TestMapFrameNew(t *testing.T) {
	val0, val1, size := 7, "n", 10
	col0, _ := NewLabelColumn("col0", val0, size)
	col1, _ := NewLabelColumn("col1", val1, size)
	cols := []Column{col0, col1}

	frame, err := NewMapFrame(cols)
	if err != nil {
		t.Fatalf("can't create frame - %s", err)
	}

	names := frame.Columns()
	if len(names) != len(cols) {
		t.Fatalf("# of columns mismatch - %d != %d", len(names), len(cols))
	}

	for i, name := range names {
		col := cols[i]
		if col.Name() != name {
			t.Fatalf("%d: name mismatch - %q != %q", i, col.Name(), name)
		}

		if col.Len() != size {
			t.Fatalf("%d: length mismatch - %d != %d", i, col.Len(), size)
		}

		switch i {
		case 0:
			val := col.IntAt(0)
			if val != val0 {
				t.Fatalf("%d: value mismatch - %d != %d", i, val, val0)
			}
		case 1:
			val := col.StringAt(0)
			if val != val1 {
				t.Fatalf("%d: value mismatch - %q != %q", i, val, val1)
			}
		}

	}
}

func TestMapFrameSlice(t *testing.T) {
	size := 10
	nCols := 7
	var cols []Column

	for i := 0; i < nCols; i++ {
		col, _ := NewLabelColumn(fmt.Sprintf("col%d", i), i, size)
		cols = append(cols, col)
	}

	frame, err := NewMapFrame(cols)
	if err != nil {
		t.Fatalf("can't create frame - %s", err)
	}

	names := frame.Columns()
	if len(names) != nCols {
		t.Fatalf("# of columns mismatch - %d != %d", len(names), nCols)
	}

	start, end := 2, 7
	frame2, err := frame.Slice(start, end)
	if err != nil {
		t.Fatalf("can't create slice - %s", err)
	}

	if frame2.Len() != end-start {
		t.Fatalf("bad # of rows in slice - %d != %d", frame2.Len(), end-start)
	}

	names2 := frame2.Columns()
	if len(names2) != nCols {
		t.Fatalf("# of columns mismatch - %d != %d", len(names2), nCols)
	}
}