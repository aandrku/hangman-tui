package words

import (
	"testing"
)

func TestStore(t *testing.T) {
	s := GetStore()
	if s == nil {
		t.Errorf("received nil from GetStore")
	}

	s.ReadFromFile("./sample_data/words1.txt")

	if s.words[0] != "hello" {
		t.Errorf("wrong first word expected %s, got %s", "hello", s.words[0])
	}

	expectedWordQuantity := 3
	expectedMinLength := 2
	expectedMaxLength := 10
	expectedLength := expectedMinLength

	gotWordQuantity := len(s.words)
	if expectedWordQuantity != gotWordQuantity {
		t.Errorf("exptected word quantity does not match expected %d, got %d", expectedWordQuantity, gotWordQuantity)
	}

	gotMinLength := s.minLength
	if gotMinLength != expectedMinLength {
		t.Errorf("expected minLength to be %d, got %d", expectedMinLength, gotMinLength)
	}

	gotMaxLength := s.maxLength
	if gotMaxLength != expectedMaxLength {
		t.Errorf("expected maxLength to be %d, got %d", expectedMaxLength, gotMaxLength)
	}

	gotLength := s.length
	if gotLength != expectedLength {
		t.Errorf("expected length to be %d, got %d", expectedLength, gotLength)
	}

	t.Log(s.availableLength)

	// test IncreaseLength function
	expectedLength = 2
	gotLength = s.length
	if gotLength != expectedLength {
		t.Errorf("expected length to be %d, got %d", expectedLength, gotLength)
	}

	s.IncreaseLength()
	expectedLength = 5
	gotLength = s.length
	if gotLength != expectedLength {
		t.Errorf("expected length to be %d, got %d", expectedLength, gotLength)
	}

	s.IncreaseLength()
	expectedLength = 10
	gotLength = s.length
	if gotLength != expectedLength {
		t.Errorf("expected length to be %d, got %d", expectedLength, gotLength)
	}

	s.IncreaseLength()
	expectedLength = 10
	gotLength = s.length
	if gotLength != expectedLength {
		t.Errorf("expected length to be %d, got %d", expectedLength, gotLength)
	}

	s.DecreaseLength()
	expectedLength = 5
	gotLength = s.length
	if gotLength != expectedLength {
		t.Errorf("expected length to be %d, got %d", expectedLength, gotLength)
	}

	s.DecreaseLength()
	expectedLength = 2
	gotLength = s.length
	if gotLength != expectedLength {
		t.Errorf("expected length to be %d, got %d", expectedLength, gotLength)
	}

}
