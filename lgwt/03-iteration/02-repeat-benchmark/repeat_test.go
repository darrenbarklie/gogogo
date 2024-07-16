package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

// => goos: darwin
// => goarch: amd64
// => pkg: learn-go-with-tests/03-iteration/02-repeat-benchmark
// => cpu: VirtualApple @ 2.50GHz
// => BenchmarkRepeat-8   	13129598	        91.63 ns/op

// 		What 91.63 ns/op means is the function takes on average
// 		91 nanoseconds to run. To test, it ran it 10_000_000 times.
