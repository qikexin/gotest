package example10

import "testing"

/*func TestFib(t *testing.T) {
	var (
		in = 7
		expected = 13
	)
	actual := Fib(in)
	if actual != expected{
		t.Errorf("Fib(%d) = %d;expected %d",in,actual,expected)
	}
}*/

func BenchmarkFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(10)
	}
}