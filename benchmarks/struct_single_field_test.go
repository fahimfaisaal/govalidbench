package benchmarks

import (
	"testing"

	"github.com/Oudwins/govalidbench/packages"
	z "github.com/Oudwins/zog"
	internals "github.com/Oudwins/zog/internals"
	"github.com/asaskevich/govalidator"
)

func BenchmarkStructSingleField(b *testing.B) {
	// --- zog ---
	b.Run("zog/Success", func(b *testing.B) {
		internals.Clear()
		for i := 0; i < b.N; i++ {
			packages.StructSingleFieldZog.Validate(&packages.StructSingleFieldSuccessVal)
		}
	})
	b.Run("zog/Error", func(b *testing.B) {
		internals.Clear()
		for i := 0; i < b.N; i++ {
			errs := packages.StructSingleFieldZog.Validate(&packages.StructSingleFieldFailureVal)
			z.Issues.CollectMap(errs)
		}
	})
	b.Run("zog/SuccessParallel", func(b *testing.B) {
		internals.Clear()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.StructSingleFieldZog.Validate(&packages.StructSingleFieldSuccessVal)
			}
		})
	})
	b.Run("zog/ErrorParallel", func(b *testing.B) {
		internals.Clear()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				errs := packages.StructSingleFieldZog.Validate(&packages.StructSingleFieldFailureVal)
				z.Issues.CollectMap(errs)
			}
		})
	})

	// --- validator ---
	b.Run("validator/Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			playgroundValidate.Struct(&packages.StructSingleFieldSuccessVal)
		}
	})
	b.Run("validator/Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			playgroundValidate.Struct(&packages.StructSingleFieldFailureVal)
		}
	})
	b.Run("validator/SuccessParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				playgroundValidate.Struct(&packages.StructSingleFieldSuccessVal)
			}
		})
	})
	b.Run("validator/ErrorParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				playgroundValidate.Struct(&packages.StructSingleFieldFailureVal)
			}
		})
	})

	// --- ozzo ---
	b.Run("ozzo/Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			packages.StructSingleFieldOzzo(&packages.StructSingleFieldSuccessVal)
		}
	})
	b.Run("ozzo/Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			packages.StructSingleFieldOzzo(&packages.StructSingleFieldFailureVal)
		}
	})
	b.Run("ozzo/SuccessParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.StructSingleFieldOzzo(&packages.StructSingleFieldSuccessVal)
			}
		})
	})
	b.Run("ozzo/ErrorParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.StructSingleFieldOzzo(&packages.StructSingleFieldFailureVal)
			}
		})
	})

	// --- govalidator ---
	b.Run("govalidator/Success", func(b *testing.B) {
		_, err := govalidator.ValidateStruct(packages.StructSingleFieldSuccessVal)
		if err != nil {
			b.Fatal(err)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			govalidator.ValidateStruct(packages.StructSingleFieldSuccessVal)
		}
	})
	b.Run("govalidator/Error", func(b *testing.B) {
		_, err := govalidator.ValidateStruct(packages.StructSingleFieldFailureVal)
		if err == nil {
			b.Fatal("expected error")
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			govalidator.ValidateStruct(packages.StructSingleFieldFailureVal)
		}
	})
	b.Run("govalidator/SuccessParallel", func(b *testing.B) {
		_, err := govalidator.ValidateStruct(packages.StructSingleFieldSuccessVal)
		if err != nil {
			b.Fatal(err)
		}
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				govalidator.ValidateStruct(packages.StructSingleFieldSuccessVal)
			}
		})
	})
	b.Run("govalidator/ErrorParallel", func(b *testing.B) {
		_, err := govalidator.ValidateStruct(packages.StructSingleFieldFailureVal)
		if err == nil {
			b.Fatal("expected error")
		}
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				govalidator.ValidateStruct(packages.StructSingleFieldFailureVal)
			}
		})
	})
}
