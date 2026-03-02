package benchmarks

import (
	"testing"

	"github.com/Oudwins/govalidbench/packages"
	z "github.com/Oudwins/zog"
	internals "github.com/Oudwins/zog/internals"
)

func BenchmarkLotsOfTests(b *testing.B) {
	// --- zog ---
	b.Run("zog/Success", func(b *testing.B) {
		internals.Clear()
		for i := 0; i < b.N; i++ {
			packages.LotsOfTestsZog.Validate(&packages.LotsOfTestsSuccessVal)
		}
	})
	b.Run("zog/Error", func(b *testing.B) {
		internals.Clear()
		for i := 0; i < b.N; i++ {
			errs := packages.LotsOfTestsZog.Validate(&packages.LotsOfTestsFailureVal)
			z.Issues.CollectList(errs)
		}
	})
	b.Run("zog/SuccessParallel", func(b *testing.B) {
		internals.Clear()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.LotsOfTestsZog.Validate(&packages.LotsOfTestsSuccessVal)
			}
		})
	})
	b.Run("zog/ErrorParallel", func(b *testing.B) {
		internals.Clear()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				errs := packages.LotsOfTestsZog.Validate(&packages.LotsOfTestsFailureVal)
				z.Issues.CollectList(errs)
			}
		})
	})

	// --- validator ---
	b.Run("validator/Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			playgroundValidate.Var(packages.LotsOfTestsSuccessVal, packages.LotsOfTestsValidator)
		}
	})
	b.Run("validator/Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			playgroundValidate.Var(packages.LotsOfTestsFailureVal, packages.LotsOfTestsValidator)
		}
	})
	b.Run("validator/SuccessParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				playgroundValidate.Var(packages.LotsOfTestsSuccessVal, packages.LotsOfTestsValidator)
			}
		})
	})
	b.Run("validator/ErrorParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				playgroundValidate.Var(packages.LotsOfTestsFailureVal, packages.LotsOfTestsValidator)
			}
		})
	})

	// --- ozzo ---
	b.Run("ozzo/Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			packages.LotsOfTestsOzzo(&packages.LotsOfTestsSuccessVal)
		}
	})
	b.Run("ozzo/Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			packages.LotsOfTestsOzzo(&packages.LotsOfTestsFailureVal)
		}
	})
	b.Run("ozzo/SuccessParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.LotsOfTestsOzzo(&packages.LotsOfTestsSuccessVal)
			}
		})
	})
	b.Run("ozzo/ErrorParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.LotsOfTestsOzzo(&packages.LotsOfTestsFailureVal)
			}
		})
	})

	// --- govalidator ---
	b.Run("govalidator/Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			packages.LotsOfTestsGoValidator(&packages.LotsOfTestsSuccessVal)
		}
	})
	b.Run("govalidator/Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			packages.LotsOfTestsGoValidator(&packages.LotsOfTestsFailureVal)
		}
	})
	b.Run("govalidator/SuccessParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.LotsOfTestsGoValidator(&packages.LotsOfTestsSuccessVal)
			}
		})
	})
	b.Run("govalidator/ErrorParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.LotsOfTestsGoValidator(&packages.LotsOfTestsFailureVal)
			}
		})
	})
}
