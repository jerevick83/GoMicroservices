package data

import "testing"

func TestValidation(t *testing.T) {
	p := &Product{
		Name:  "Jeremiah",
		Price: 5.0,
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
