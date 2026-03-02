package benchmarks

import (
	"testing"

	"github.com/Oudwins/govalidbench/packages/data"
	z "github.com/Oudwins/zog"
	internals "github.com/Oudwins/zog/internals"
	"github.com/asaskevich/govalidator"
)

func BenchmarkStructComplex(b *testing.B) {
	// --- zog ---
	b.Run("zog/Success", func(b *testing.B) {
		internals.Clear()
		for i := 0; i < b.N; i++ {
			data.StructComplexZog.Validate(&data.StructComplexSuccessVal)
		}
	})
	b.Run("zog/Error", func(b *testing.B) {
		internals.Clear()
		for i := 0; i < b.N; i++ {
			errs := data.StructComplexZog.Validate(&data.StructComplexFailureVal)
			z.Issues.CollectMap(errs)
		}
	})
	b.Run("zog/SuccessParallel", func(b *testing.B) {
		internals.Clear()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				data.StructComplexZog.Validate(&data.StructComplexSuccessVal)
			}
		})
	})
	b.Run("zog/ErrorParallel", func(b *testing.B) {
		internals.Clear()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				errs := data.StructComplexZog.Validate(&data.StructComplexFailureVal)
				z.Issues.CollectMap(errs)
			}
		})
	})

	// --- validator ---
	b.Run("validator/Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			playgroundValidate.Struct(&data.StructComplexSuccessVal)
		}
	})
	b.Run("validator/Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			playgroundValidate.Struct(&data.StructComplexFailureVal)
		}
	})
	b.Run("validator/SuccessParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				playgroundValidate.Struct(&data.StructComplexSuccessVal)
			}
		})
	})
	b.Run("validator/ErrorParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				playgroundValidate.Struct(&data.StructComplexFailureVal)
			}
		})
	})

	// --- ozzo ---
	b.Run("ozzo/Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			data.StructComplexOzzo(&data.StructComplexSuccessVal)
		}
	})
	b.Run("ozzo/Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			data.StructComplexOzzo(&data.StructComplexFailureVal)
		}
	})
	b.Run("ozzo/SuccessParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				data.StructComplexOzzo(&data.StructComplexSuccessVal)
			}
		})
	})
	b.Run("ozzo/ErrorParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				data.StructComplexOzzo(&data.StructComplexFailureVal)
			}
		})
	})

	// --- govalidator ---
	b.Run("govalidator/Success", func(b *testing.B) {
		_, err := govalidator.ValidateStruct(data.StructComplexSuccessVal)
		if err != nil {
			b.Fatal(err)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			govalidator.ValidateStruct(data.StructComplexSuccessVal)
		}
	})
	b.Run("govalidator/Error", func(b *testing.B) {
		_, err := govalidator.ValidateStruct(data.StructComplexFailureVal)
		if err == nil {
			b.Fatal("expected error")
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			govalidator.ValidateStruct(data.StructComplexFailureVal)
		}
	})
	b.Run("govalidator/SuccessParallel", func(b *testing.B) {
		_, err := govalidator.ValidateStruct(data.StructComplexSuccessVal)
		if err != nil {
			b.Fatal(err)
		}
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				govalidator.ValidateStruct(data.StructComplexSuccessVal)
			}
		})
	})
	b.Run("govalidator/ErrorParallel", func(b *testing.B) {
		_, err := govalidator.ValidateStruct(data.StructComplexFailureVal)
		if err == nil {
			b.Fatal("expected error")
		}
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				govalidator.ValidateStruct(data.StructComplexFailureVal)
			}
		})
	})
}

//
// StructComplexCreate (zog-only: schema creation + validation)
//

func BenchmarkStructComplexCreate(b *testing.B) {
	b.Run("zog/Success", func(b *testing.B) {
		internals.Clear()
		for i := 0; i < b.N; i++ {
			s := data.StructComplexCreateZog()
			s.Validate(&data.StructComplexSuccessVal)
		}
	})
	b.Run("zog/Error", func(b *testing.B) {
		internals.Clear()
		for i := 0; i < b.N; i++ {
			s := data.StructComplexCreateZog()
			errs := s.Validate(&data.StructComplexFailureVal)
			z.Issues.CollectMap(errs)
		}
	})
	b.Run("zog/SuccessParallel", func(b *testing.B) {
		internals.Clear()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				s := data.StructComplexCreateZog()
				s.Validate(&data.StructComplexSuccessVal)
			}
		})
	})
	b.Run("zog/ErrorParallel", func(b *testing.B) {
		internals.Clear()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				s := data.StructComplexCreateZog()
				errs := s.Validate(&data.StructComplexFailureVal)
				z.Issues.CollectMap(errs)
			}
		})
	})
}
