package benchmarks

import (
	"testing"

	"github.com/Oudwins/govalidbench/packages"
	z "github.com/Oudwins/zog"
	internals "github.com/Oudwins/zog/internals"
	"github.com/asaskevich/govalidator"
)

func BenchmarkStructComplex(b *testing.B) {
	// --- zog ---
	b.Run("zog/Success", func(b *testing.B) {
		internals.Clear()
		for i := 0; i < b.N; i++ {
			packages.StructComplexZog.Validate(&packages.StructComplexSuccessVal)
		}
	})
	b.Run("zog/Error", func(b *testing.B) {
		internals.Clear()
		for i := 0; i < b.N; i++ {
			errs := packages.StructComplexZog.Validate(&packages.StructComplexFailureVal)
			z.Issues.CollectMap(errs)
		}
	})
	b.Run("zog/SuccessParallel", func(b *testing.B) {
		internals.Clear()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.StructComplexZog.Validate(&packages.StructComplexSuccessVal)
			}
		})
	})
	b.Run("zog/ErrorParallel", func(b *testing.B) {
		internals.Clear()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				errs := packages.StructComplexZog.Validate(&packages.StructComplexFailureVal)
				z.Issues.CollectMap(errs)
			}
		})
	})

	// --- validator ---
	b.Run("validator/Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			playgroundValidate.Struct(&packages.StructComplexSuccessVal)
		}
	})
	b.Run("validator/Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			playgroundValidate.Struct(&packages.StructComplexFailureVal)
		}
	})
	b.Run("validator/SuccessParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				playgroundValidate.Struct(&packages.StructComplexSuccessVal)
			}
		})
	})
	b.Run("validator/ErrorParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				playgroundValidate.Struct(&packages.StructComplexFailureVal)
			}
		})
	})

	// --- ozzo ---
	b.Run("ozzo/Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			packages.StructComplexOzzo(&packages.StructComplexSuccessVal)
		}
	})
	b.Run("ozzo/Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			packages.StructComplexOzzo(&packages.StructComplexFailureVal)
		}
	})
	b.Run("ozzo/SuccessParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.StructComplexOzzo(&packages.StructComplexSuccessVal)
			}
		})
	})
	b.Run("ozzo/ErrorParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.StructComplexOzzo(&packages.StructComplexFailureVal)
			}
		})
	})

	// --- govalidator ---
	b.Run("govalidator/Success", func(b *testing.B) {
		_, err := govalidator.ValidateStruct(packages.StructComplexSuccessVal)
		if err != nil {
			b.Fatal(err)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			govalidator.ValidateStruct(packages.StructComplexSuccessVal)
		}
	})
	b.Run("govalidator/Error", func(b *testing.B) {
		_, err := govalidator.ValidateStruct(packages.StructComplexFailureVal)
		if err == nil {
			b.Fatal("expected error")
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			govalidator.ValidateStruct(packages.StructComplexFailureVal)
		}
	})
	b.Run("govalidator/SuccessParallel", func(b *testing.B) {
		_, err := govalidator.ValidateStruct(packages.StructComplexSuccessVal)
		if err != nil {
			b.Fatal(err)
		}
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				govalidator.ValidateStruct(packages.StructComplexSuccessVal)
			}
		})
	})
	b.Run("govalidator/ErrorParallel", func(b *testing.B) {
		_, err := govalidator.ValidateStruct(packages.StructComplexFailureVal)
		if err == nil {
			b.Fatal("expected error")
		}
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				govalidator.ValidateStruct(packages.StructComplexFailureVal)
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
			s := packages.StructComplexCreateZog()
			s.Validate(&packages.StructComplexSuccessVal)
		}
	})
	b.Run("zog/Error", func(b *testing.B) {
		internals.Clear()
		for i := 0; i < b.N; i++ {
			s := packages.StructComplexCreateZog()
			errs := s.Validate(&packages.StructComplexFailureVal)
			z.Issues.CollectMap(errs)
		}
	})
	b.Run("zog/SuccessParallel", func(b *testing.B) {
		internals.Clear()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				s := packages.StructComplexCreateZog()
				s.Validate(&packages.StructComplexSuccessVal)
			}
		})
	})
	b.Run("zog/ErrorParallel", func(b *testing.B) {
		internals.Clear()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				s := packages.StructComplexCreateZog()
				errs := s.Validate(&packages.StructComplexFailureVal)
				z.Issues.CollectMap(errs)
			}
		})
	})
}
