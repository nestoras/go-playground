package bench

import (
	"fmt"
	"testing"
)


// go test -bench .
func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}