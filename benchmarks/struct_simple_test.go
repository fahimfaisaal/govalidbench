package benchmarks

import (
	"testing"

	"github.com/Oudwins/govalidbench/packages/data"
	z "github.com/Oudwins/zog"
	internals "github.com/Oudwins/zog/internals"
	"github.com/asaskevich/govalidator"
)

func BenchmarkStructSimple(b *testing.B) {
	// --- zog ---
	b.Run("zog/Success", func(b *testing.B) {
		internals.Clear()
		for i := 0; i < b.N; i++ {
			data.StructSimpleZog.Validate(&data.StructSimpleSuccessVal)
		}
	})
	b.Run("zog/Error", func(b *testing.B) {
		internals.Clear()
		for i := 0; i < b.N; i++ {
			errs := data.StructSimpleZog.Validate(&data.StructSimpleFailureVal)
			z.Issues.CollectMap(errs)
		}
	})
	b.Run("zog/SuccessParallel", func(b *testing.B) {
		internals.Clear()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				data.StructSimpleZog.Validate(&data.StructSimpleSuccessVal)
			}
		})
	})
	b.Run("zog/ErrorParallel", func(b *testing.B) {
		internals.Clear()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				errs := data.StructSimpleZog.Validate(&data.StructSimpleFailureVal)
				z.Issues.CollectMap(errs)
			}
		})
	})

	// --- validator ---
	b.Run("validator/Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			playgroundValidate.Struct(&data.StructSimpleSuccessVal)
		}
	})
	b.Run("validator/Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			playgroundValidate.Struct(&data.StructSimpleFailureVal)
		}
	})
	b.Run("validator/SuccessParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				playgroundValidate.Struct(&data.StructSimpleSuccessVal)
			}
		})
	})
	b.Run("validator/ErrorParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				playgroundValidate.Struct(&data.StructSimpleFailureVal)
			}
		})
	})

	// --- ozzo ---
	b.Run("ozzo/Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			data.StructSimpleOzzo(&data.StructSimpleSuccessVal)
		}
	})
	b.Run("ozzo/Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			data.StructSimpleOzzo(&data.StructSimpleFailureVal)
		}
	})
	b.Run("ozzo/SuccessParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				data.StructSimpleOzzo(&data.StructSimpleSuccessVal)
			}
		})
	})
	b.Run("ozzo/ErrorParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				data.StructSimpleOzzo(&data.StructSimpleFailureVal)
			}
		})
	})

	// --- govalidator ---
	b.Run("govalidator/Success", func(b *testing.B) {
		_, err := govalidator.ValidateStruct(data.StructSimpleSuccessVal)
		if err != nil {
			b.Fatal(err)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			govalidator.ValidateStruct(data.StructSimpleSuccessVal)
		}
	})
	b.Run("govalidator/Error", func(b *testing.B) {
		_, err := govalidator.ValidateStruct(data.StructSimpleFailureVal)
		if err == nil {
			b.Fatal("expected error")
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			govalidator.ValidateStruct(data.StructSimpleFailureVal)
		}
	})
	b.Run("govalidator/SuccessParallel", func(b *testing.B) {
		_, err := govalidator.ValidateStruct(data.StructSimpleSuccessVal)
		if err != nil {
			b.Fatal(err)
		}
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				govalidator.ValidateStruct(data.StructSimpleSuccessVal)
			}
		})
	})
	b.Run("govalidator/ErrorParallel", func(b *testing.B) {
		_, err := govalidator.ValidateStruct(data.StructSimpleFailureVal)
		if err == nil {
			b.Fatal("expected error")
		}
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				govalidator.ValidateStruct(data.StructSimpleFailureVal)
			}
		})
	})
}
