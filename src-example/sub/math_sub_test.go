package sub

import (
	"testing"
)

func TestHello(t *testing.T) {
	emptyNameResult := hello("")

	if emptyNameResult != "What is your name ?" {
		t.Errorf("Output expect What is your name ? instead of %v", emptyNameResult)
	}

	result := hello("Gopher")

	if result != "Hello Gopher" {
		t.Errorf("Output expect Hello Gopher instead of %v", result)
	}
}
