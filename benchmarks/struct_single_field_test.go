package benchmarks

import (
	"testing"

	"github.com/Oudwins/govalidbench/packages/data"
	z "github.com/Oudwins/zog"
	internals "github.com/Oudwins/zog/internals"
	"github.com/asaskevich/govalidator"
)

func BenchmarkStructSingleField(b *testing.B) {
	// --- zog ---
	b.Run("zog/Success", func(b *testing.B) {
		internals.Clear()
		for i := 0; i < b.N; i++ {
			data.StructSingleFieldZog.Validate(&data.StructSingleFieldSuccessVal)
		}
	})
	b.Run("zog/Error", func(b *testing.B) {
		internals.Clear()
		for i := 0; i < b.N; i++ {
			errs := data.StructSingleFieldZog.Validate(&data.StructSingleFieldFailureVal)
			z.Issues.CollectMap(errs)
		}
	})
	b.Run("zog/SuccessParallel", func(b *testing.B) {
		internals.Clear()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				data.StructSingleFieldZog.Validate(&data.StructSingleFieldSuccessVal)
			}
		})
	})
	b.Run("zog/ErrorParallel", func(b *testing.B) {
		internals.Clear()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				errs := data.StructSingleFieldZog.Validate(&data.StructSingleFieldFailureVal)
				z.Issues.CollectMap(errs)
			}
		})
	})

	// --- validator ---
	b.Run("validator/Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			playgroundValidate.Struct(&data.StructSingleFieldSuccessVal)
		}
	})
	b.Run("validator/Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			playgroundValidate.Struct(&data.StructSingleFieldFailureVal)
		}
	})
	b.Run("validator/SuccessParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				playgroundValidate.Struct(&data.StructSingleFieldSuccessVal)
			}
		})
	})
	b.Run("validator/ErrorParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				playgroundValidate.Struct(&data.StructSingleFieldFailureVal)
			}
		})
	})

	// --- ozzo ---
	b.Run("ozzo/Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			data.StructSingleFieldOzzo(&data.StructSingleFieldSuccessVal)
		}
	})
	b.Run("ozzo/Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			data.StructSingleFieldOzzo(&data.StructSingleFieldFailureVal)
		}
	})
	b.Run("ozzo/SuccessParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				data.StructSingleFieldOzzo(&data.StructSingleFieldSuccessVal)
			}
		})
	})
	b.Run("ozzo/ErrorParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				data.StructSingleFieldOzzo(&data.StructSingleFieldFailureVal)
			}
		})
	})

	// --- govalidator ---
	b.Run("govalidator/Success", func(b *testing.B) {
		_, err := govalidator.ValidateStruct(data.StructSingleFieldSuccessVal)
		if err != nil {
			b.Fatal(err)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			govalidator.ValidateStruct(data.StructSingleFieldSuccessVal)
		}
	})
	b.Run("govalidator/Error", func(b *testing.B) {
		_, err := govalidator.ValidateStruct(data.StructSingleFieldFailureVal)
		if err == nil {
			b.Fatal("expected error")
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			govalidator.ValidateStruct(data.StructSingleFieldFailureVal)
		}
	})
	b.Run("govalidator/SuccessParallel", func(b *testing.B) {
		_, err := govalidator.ValidateStruct(data.StructSingleFieldSuccessVal)
		if err != nil {
			b.Fatal(err)
		}
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				govalidator.ValidateStruct(data.StructSingleFieldSuccessVal)
			}
		})
	})
	b.Run("govalidator/ErrorParallel", func(b *testing.B) {
		_, err := govalidator.ValidateStruct(data.StructSingleFieldFailureVal)
		if err == nil {
			b.Fatal("expected error")
		}
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				govalidator.ValidateStruct(data.StructSingleFieldFailureVal)
			}
		})
	})
}
