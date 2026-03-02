package benchmarks

import (
	"testing"

	"github.com/Oudwins/govalidbench/packages"
	z "github.com/Oudwins/zog"
	internals "github.com/Oudwins/zog/internals"
)

func BenchmarkSliceField(b *testing.B) {
	// --- zog ---
	b.Run("zog/Success", func(b *testing.B) {
		internals.Clear()
		for i := 0; i < b.N; i++ {
			packages.SliceFieldZog.Validate(&packages.SliceFieldSucessVal)
		}
	})
	b.Run("zog/Error", func(b *testing.B) {
		internals.Clear()
		for i := 0; i < b.N; i++ {
			errs := packages.SliceFieldZog.Validate(&packages.SliceFieldFailureVal)
			z.Issues.CollectMap(errs)
		}
	})
	b.Run("zog/SuccessParallel", func(b *testing.B) {
		internals.Clear()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.SliceFieldZog.Validate(&packages.SliceFieldSucessVal)
			}
		})
	})
	b.Run("zog/ErrorParallel", func(b *testing.B) {
		internals.Clear()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				errs := packages.SliceFieldZog.Validate(&packages.SliceFieldFailureVal)
				z.Issues.CollectMap(errs)
			}
		})
	})

	// --- validator ---
	b.Run("validator/Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			playgroundValidate.Var(packages.SliceFieldSucessVal, packages.SliceFieldValidator)
		}
	})
	b.Run("validator/Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			playgroundValidate.Var(packages.SliceFieldFailureVal, packages.SliceFieldValidator)
		}
	})
	b.Run("validator/SuccessParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				playgroundValidate.Var(packages.SliceFieldSucessVal, packages.SliceFieldValidator)
			}
		})
	})
	b.Run("validator/ErrorParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				playgroundValidate.Var(packages.SliceFieldFailureVal, packages.SliceFieldValidator)
			}
		})
	})

	// --- ozzo ---
	b.Run("ozzo/Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			packages.SliceFieldOzzo(&packages.SliceFieldSucessVal)
		}
	})
	b.Run("ozzo/Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			packages.SliceFieldOzzo(&packages.SliceFieldFailureVal)
		}
	})
	b.Run("ozzo/SuccessParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.SliceFieldOzzo(&packages.SliceFieldSucessVal)
			}
		})
	})
	b.Run("ozzo/ErrorParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.SliceFieldOzzo(&packages.SliceFieldFailureVal)
			}
		})
	})

	// --- govalidator ---
	b.Run("govalidator/Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			packages.SliceFieldGoValidator(&packages.SliceFieldSucessVal)
		}
	})
	b.Run("govalidator/Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			packages.SliceFieldGoValidator(&packages.SliceFieldFailureVal)
		}
	})
	b.Run("govalidator/SuccessParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.SliceFieldGoValidator(&packages.SliceFieldSucessVal)
			}
		})
	})
	b.Run("govalidator/ErrorParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.SliceFieldGoValidator(&packages.SliceFieldFailureVal)
			}
		})
	})
}
