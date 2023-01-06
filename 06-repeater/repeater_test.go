package repeater

import "testing"

func TestRepeater(t *testing.T) {
	want := "aaaaa"
	got := repeater("a", 5)

	if want != got {
		t.Errorf("wanted %q but instead, got %q", want, got)
	}
}

func BenchmarkRepeater(b *testing.B) {
	for i := 0; i < b.N; i++ {
		repeater("a", 1)
	}
}
