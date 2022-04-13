package defertest

import (
	"testing"
)

func BenchmarkPut(b *testing.B) {
	q := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000; j++ {
			q.Put(j)
		}
	}
}

func BenchmarkPutDefer(b *testing.B) {
	q := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000; j++ {
			q.PutDefer(j)
		}
	}
}

func BenchmarkGet(b *testing.B) {
	q := New()
	for i := 0; i < 1000; i++ {
		q.Put(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 2000; j++ {
			q.Get()
		}
	}
}

func BenchmarkGetDefer(b *testing.B) {
	q := New()
	for i := 0; i < 1000; i++ {
		q.Put(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 2000; j++ {
			q.GetDefer()
		}
	}

}
