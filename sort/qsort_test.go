package sort

import (
	"testing"
)

func TestSort(t *testing.T) {
	cases := map[string]sortFunc{
		// "quick":           quick,
		"quick parallel":      quickParallel,
		"quick parallel-slow": quickParallelMagic,
	}
	for name, f := range cases {
		src := randomSlice(10)
		t.Run(name, func(t *testing.T) {
			checkSort(t, src, f)
		})
	}
}

// benchmarks

func BenchmarkQuick(b *testing.B) {
	benchSort(b, "quick", quick, 15000)
	benchSort(b, "quick-parallel", quickParallel, 15000)
	// benchSort(b, "quick-parallel-slow", quickParallelMagic, 15000)

	// b.Run("no-gc", func(b *testing.B) {
	// 	debug.SetGCPercent(-1)
	// 	b.Cleanup(func() {
	// 		debug.SetGCPercent(100)
	// 	})
	// 	benchSort(b, "quick-parallel", quickParallel, benchSize)
	// })
	//
	// benchSort(b, "quick", quick, benchSize)
}
