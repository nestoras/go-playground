package magic

import "testing"

//This is a unit test
func TestMagic(t *testing.T){
	got := Magic(1,2)
	if got != 9 {
		t.Errorf("Magic() = %v, want %v", got, 9)
	}
}
