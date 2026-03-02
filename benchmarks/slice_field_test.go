package benchmarks

import (
	"testing"

	"github.com/Oudwins/govalidbench/packages/data"
	z "github.com/Oudwins/zog"
	internals "github.com/Oudwins/zog/internals"
)

func BenchmarkSliceField(b *testing.B) {
	// --- zog ---
	b.Run("zog/Success", func(b *testing.B) {
		internals.Clear()
		for i := 0; i < b.N; i++ {
			data.SliceFieldZog.Validate(&data.SliceFieldSucessVal)
		}
	})
	b.Run("zog/Error", func(b *testing.B) {
		internals.Clear()
		for i := 0; i < b.N; i++ {
			errs := data.SliceFieldZog.Validate(&data.SliceFieldFailureVal)
			z.Issues.CollectMap(errs)
		}
	})
	b.Run("zog/SuccessParallel", func(b *testing.B) {
		internals.Clear()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				data.SliceFieldZog.Validate(&data.SliceFieldSucessVal)
			}
		})
	})
	b.Run("zog/ErrorParallel", func(b *testing.B) {
		internals.Clear()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				errs := data.SliceFieldZog.Validate(&data.SliceFieldFailureVal)
				z.Issues.CollectMap(errs)
			}
		})
	})

	// --- validator ---
	b.Run("validator/Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			playgroundValidate.Var(data.SliceFieldSucessVal, data.SliceFieldValidator)
		}
	})
	b.Run("validator/Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			playgroundValidate.Var(data.SliceFieldFailureVal, data.SliceFieldValidator)
		}
	})
	b.Run("validator/SuccessParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				playgroundValidate.Var(data.SliceFieldSucessVal, data.SliceFieldValidator)
			}
		})
	})
	b.Run("validator/ErrorParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				playgroundValidate.Var(data.SliceFieldFailureVal, data.SliceFieldValidator)
			}
		})
	})

	// --- ozzo ---
	b.Run("ozzo/Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			data.SliceFieldOzzo(&data.SliceFieldSucessVal)
		}
	})
	b.Run("ozzo/Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			data.SliceFieldOzzo(&data.SliceFieldFailureVal)
		}
	})
	b.Run("ozzo/SuccessParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				data.SliceFieldOzzo(&data.SliceFieldSucessVal)
			}
		})
	})
	b.Run("ozzo/ErrorParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				data.SliceFieldOzzo(&data.SliceFieldFailureVal)
			}
		})
	})

	// --- govalidator ---
	b.Run("govalidator/Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			data.SliceFieldGoValidator(&data.SliceFieldSucessVal)
		}
	})
	b.Run("govalidator/Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			data.SliceFieldGoValidator(&data.SliceFieldFailureVal)
		}
	})
	b.Run("govalidator/SuccessParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				data.SliceFieldGoValidator(&data.SliceFieldSucessVal)
			}
		})
	})
	b.Run("govalidator/ErrorParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				data.SliceFieldGoValidator(&data.SliceFieldFailureVal)
			}
		})
	})
}
