package benchmarks

import (
	"testing"

	"github.com/Oudwins/govalidbench/packages"
	z "github.com/Oudwins/zog"
	internals "github.com/Oudwins/zog/internals"
)

func BenchmarkStringFieldSimple(b *testing.B) {
	// --- zog ---
	b.Run("zog/Success", func(b *testing.B) {
		internals.Clear()
		for i := 0; i < b.N; i++ {
			packages.StringFieldSimpleZog.Validate(&packages.StringFieldSimpleSuccessVal)
		}
	})
	b.Run("zog/Error", func(b *testing.B) {
		internals.Clear()
		for i := 0; i < b.N; i++ {
			errs := packages.StringFieldSimpleZog.Validate(&packages.StringFieldSimpleFailureVal)
			z.Issues.CollectList(errs)
		}
	})
	b.Run("zog/SuccessParallel", func(b *testing.B) {
		internals.Clear()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.StringFieldSimpleZog.Validate(&packages.StringFieldSimpleSuccessVal)
			}
		})
	})
	b.Run("zog/ErrorParallel", func(b *testing.B) {
		internals.Clear()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				errs := packages.StringFieldSimpleZog.Validate(&packages.StringFieldSimpleFailureVal)
				z.Issues.CollectList(errs)
			}
		})
	})

	// --- validator ---
	b.Run("validator/Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			playgroundValidate.Var(packages.StringFieldSimpleSuccessVal, packages.StringFieldSimpleValidator)
		}
	})
	b.Run("validator/Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			playgroundValidate.Var(packages.StringFieldSimpleFailureVal, packages.StringFieldSimpleValidator)
		}
	})
	b.Run("validator/SuccessParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				playgroundValidate.Var(packages.StringFieldSimpleSuccessVal, packages.StringFieldSimpleValidator)
			}
		})
	})
	b.Run("validator/ErrorParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				playgroundValidate.Var(packages.StringFieldSimpleFailureVal, packages.StringFieldSimpleValidator)
			}
		})
	})

	// --- ozzo ---
	b.Run("ozzo/Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			packages.StringFieldSimpleOzzo(&packages.StringFieldSimpleSuccessVal)
		}
	})
	b.Run("ozzo/Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			packages.StringFieldSimpleOzzo(&packages.StringFieldSimpleFailureVal)
		}
	})
	b.Run("ozzo/SuccessParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.StringFieldSimpleOzzo(&packages.StringFieldSimpleSuccessVal)
			}
		})
	})
	b.Run("ozzo/ErrorParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.StringFieldSimpleOzzo(&packages.StringFieldSimpleFailureVal)
			}
		})
	})

	// --- govalidator ---
	b.Run("govalidator/Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			packages.StringFieldSimpleGoValidator(&packages.StringFieldSimpleSuccessVal)
		}
	})
	b.Run("govalidator/Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			packages.StringFieldSimpleGoValidator(&packages.StringFieldSimpleFailureVal)
		}
	})
	b.Run("govalidator/SuccessParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.StringFieldSimpleGoValidator(&packages.StringFieldSimpleSuccessVal)
			}
		})
	})
	b.Run("govalidator/ErrorParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				packages.StringFieldSimpleGoValidator(&packages.StringFieldSimpleFailureVal)
			}
		})
	})
}
