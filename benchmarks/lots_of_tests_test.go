package benchmarks

import (
	"testing"

	"github.com/Oudwins/govalidbench/packages/data"
	z "github.com/Oudwins/zog"
	internals "github.com/Oudwins/zog/internals"
)

func BenchmarkLotsOfTests(b *testing.B) {
	// --- zog ---
	b.Run("zog/Success", func(b *testing.B) {
		internals.Clear()
		for i := 0; i < b.N; i++ {
			data.LotsOfTestsZog.Validate(&data.LotsOfTestsSuccessVal)
		}
	})
	b.Run("zog/Error", func(b *testing.B) {
		internals.Clear()
		for i := 0; i < b.N; i++ {
			errs := data.LotsOfTestsZog.Validate(&data.LotsOfTestsFailureVal)
			z.Issues.CollectList(errs)
		}
	})
	b.Run("zog/SuccessParallel", func(b *testing.B) {
		internals.Clear()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				data.LotsOfTestsZog.Validate(&data.LotsOfTestsSuccessVal)
			}
		})
	})
	b.Run("zog/ErrorParallel", func(b *testing.B) {
		internals.Clear()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				errs := data.LotsOfTestsZog.Validate(&data.LotsOfTestsFailureVal)
				z.Issues.CollectList(errs)
			}
		})
	})

	// --- validator ---
	b.Run("validator/Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			playgroundValidate.Var(data.LotsOfTestsSuccessVal, data.LotsOfTestsValidator)
		}
	})
	b.Run("validator/Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			playgroundValidate.Var(data.LotsOfTestsFailureVal, data.LotsOfTestsValidator)
		}
	})
	b.Run("validator/SuccessParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				playgroundValidate.Var(data.LotsOfTestsSuccessVal, data.LotsOfTestsValidator)
			}
		})
	})
	b.Run("validator/ErrorParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				playgroundValidate.Var(data.LotsOfTestsFailureVal, data.LotsOfTestsValidator)
			}
		})
	})

	// --- ozzo ---
	b.Run("ozzo/Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			data.LotsOfTestsOzzo(&data.LotsOfTestsSuccessVal)
		}
	})
	b.Run("ozzo/Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			data.LotsOfTestsOzzo(&data.LotsOfTestsFailureVal)
		}
	})
	b.Run("ozzo/SuccessParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				data.LotsOfTestsOzzo(&data.LotsOfTestsSuccessVal)
			}
		})
	})
	b.Run("ozzo/ErrorParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				data.LotsOfTestsOzzo(&data.LotsOfTestsFailureVal)
			}
		})
	})

	// --- govalidator ---
	b.Run("govalidator/Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			data.LotsOfTestsGoValidator(&data.LotsOfTestsSuccessVal)
		}
	})
	b.Run("govalidator/Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			data.LotsOfTestsGoValidator(&data.LotsOfTestsFailureVal)
		}
	})
	b.Run("govalidator/SuccessParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				data.LotsOfTestsGoValidator(&data.LotsOfTestsSuccessVal)
			}
		})
	})
	b.Run("govalidator/ErrorParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				data.LotsOfTestsGoValidator(&data.LotsOfTestsFailureVal)
			}
		})
	})
}
