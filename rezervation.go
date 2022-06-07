package main

import "fmt"

type Train struct {
	CarName  string
	Fullness int
	Capacity int
	difCar   bool
}

func (t *Train) RezervationFunc(person int) (int, error) {
	if person <= 0 {
		return 0, fmt.Errorf("number of people cannot be negative")
	}
	if t.difCar {
		k := int(float32(t.Capacity)*0.7 - float32(t.Fullness))
		x := k - person
		if k <= 0 {
			return person, fmt.Errorf("capacity is not enough")
		} else if x >= 0 {
			t.Fullness = t.Fullness + person
			return person, nil
		}
		t.Fullness = t.Fullness + k
		person = person - k
		return person, nil
	}
	k := int(float32(t.Capacity)*0.7 - float32(t.Fullness))
	if k < person {
		return person, fmt.Errorf("no availaible rezervation")
	}
	t.Fullness = t.Fullness + person

	return person, nil
}
