package commands

import "testing"

func TestOptionValueExtractBoolNotFound(t *testing.T) {
	t.Log("ensure that no error is returned when value is not found")
	optval := &OptionValue{found: false}
	_, _, err := optval.Bool()
	if err != nil {
		t.Fatal("Found was false. Err should have been nil")
	}
}

func TestOptionValueExtractWrongType(t *testing.T) {

	t.Log("ensure that error is returned when value if of wrong type")

	optval := &OptionValue{value: "wrong type: a string", found: true}
	_, _, err := optval.Bool()
	if err == nil {
		t.Fatal("No error returned. Failure.")
	}

	optval = &OptionValue{value: "wrong type: a string", found: true}
	_, _, err = optval.Int()
	if err == nil {
		t.Fatal("No error returned. Failure.")
	}
}
