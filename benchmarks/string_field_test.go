package benchmarks

import (
	"testing"

	"github.com/Oudwins/govalidbench/packages/data"
	z "github.com/Oudwins/zog"
	internals "github.com/Oudwins/zog/internals"
)

func BenchmarkStringFieldSimple(b *testing.B) {
	// --- zog ---
	b.Run("zog/Success", func(b *testing.B) {
		internals.Clear()
		for i := 0; i < b.N; i++ {
			data.StringFieldSimpleZog.Validate(&data.StringFieldSimpleSuccessVal)
		}
	})
	b.Run("zog/Error", func(b *testing.B) {
		internals.Clear()
		for i := 0; i < b.N; i++ {
			errs := data.StringFieldSimpleZog.Validate(&data.StringFieldSimpleFailureVal)
			z.Issues.CollectList(errs)
		}
	})
	b.Run("zog/SuccessParallel", func(b *testing.B) {
		internals.Clear()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				data.StringFieldSimpleZog.Validate(&data.StringFieldSimpleSuccessVal)
			}
		})
	})
	b.Run("zog/ErrorParallel", func(b *testing.B) {
		internals.Clear()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				errs := data.StringFieldSimpleZog.Validate(&data.StringFieldSimpleFailureVal)
				z.Issues.CollectList(errs)
			}
		})
	})

	// --- validator ---
	b.Run("validator/Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			playgroundValidate.Var(data.StringFieldSimpleSuccessVal, data.StringFieldSimpleValidator)
		}
	})
	b.Run("validator/Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			playgroundValidate.Var(data.StringFieldSimpleFailureVal, data.StringFieldSimpleValidator)
		}
	})
	b.Run("validator/SuccessParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				playgroundValidate.Var(data.StringFieldSimpleSuccessVal, data.StringFieldSimpleValidator)
			}
		})
	})
	b.Run("validator/ErrorParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				playgroundValidate.Var(data.StringFieldSimpleFailureVal, data.StringFieldSimpleValidator)
			}
		})
	})

	// --- ozzo ---
	b.Run("ozzo/Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			data.StringFieldSimpleOzzo(&data.StringFieldSimpleSuccessVal)
		}
	})
	b.Run("ozzo/Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			data.StringFieldSimpleOzzo(&data.StringFieldSimpleFailureVal)
		}
	})
	b.Run("ozzo/SuccessParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				data.StringFieldSimpleOzzo(&data.StringFieldSimpleSuccessVal)
			}
		})
	})
	b.Run("ozzo/ErrorParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				data.StringFieldSimpleOzzo(&data.StringFieldSimpleFailureVal)
			}
		})
	})

	// --- govalidator ---
	b.Run("govalidator/Success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			data.StringFieldSimpleGoValidator(&data.StringFieldSimpleSuccessVal)
		}
	})
	b.Run("govalidator/Error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			data.StringFieldSimpleGoValidator(&data.StringFieldSimpleFailureVal)
		}
	})
	b.Run("govalidator/SuccessParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				data.StringFieldSimpleGoValidator(&data.StringFieldSimpleSuccessVal)
			}
		})
	})
	b.Run("govalidator/ErrorParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				data.StringFieldSimpleGoValidator(&data.StringFieldSimpleFailureVal)
			}
		})
	})
}
