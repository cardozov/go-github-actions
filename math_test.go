import "testing"

func TestAdd(t *testing.T) {
	total := Soma(2, 3)
	if total != 5 {
		t.Errorf("Soma was incorrect, got: %d, want: %d.", total, 5)
	}
}