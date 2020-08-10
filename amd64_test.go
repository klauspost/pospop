// +build amd64

package pospop

import "golang.org/x/sys/cpu"
import "testing"

// test count8avx2
func TestCount8Avx2(t *testing.T) {
	if !cpu.X86.HasAVX2 || !cpu.X86.HasPOPCNT {
		t.SkipNow()
	}

	testCount8(t, count8avx2)
}
