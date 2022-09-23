package counting

import "testing"

func TestCount(t *testing.T) {
	arr, err := Count(1, 10)
	if arr == nil {
		t.Error("Arr is empty but should not")
	}
	if err != nil {
		t.Error("Err was thrown but should not")
	}

}

func TestCountWrongStart(t *testing.T) {
	arr, err := Count(10, 1)
	if err == nil {
		t.Error("Start argument is higher than end argument. This should have thrown an error.")
	}
	if arr != nil {
		t.Error("Arr is not empty but should be")
	}
}
