package nonil

import "testing"

func BenchmarkOption(b *testing.B) {
	b.Run("functional", func(b *testing.B) {
		var evenCount uint64
		var oddCount uint64
		for i := 0; i < b.N; i++ {
			op := toOption(i)
			op.Handle(
				func(_ int) {
					evenCount++
				},
				func() {
					oddCount++
				},
			)
		}
		if evenCount - oddCount > 1 {
			b.Fatalf("count mismatch on %d: even: %d; odd: %d", b.N, evenCount, oddCount)
		}
	})
	b.Run("imperative", func(b *testing.B) {
		var evenCount uint64
		var oddCount uint64
		for i := 0; i < b.N; i++ {
			op := toOption(i)
			if op.IsSome() {
				evenCount++
			} else {
				oddCount++
			}
		}
		if evenCount - oddCount > 1 {
			b.Fatalf("count mismatch on %d: even: %d; odd: %d", b.N, evenCount, oddCount)
		}
	})
}

func toOption(n int) Option[int] {
	if n&1 == 0 {
		return Some(n)
	}
	return None[int]()
}