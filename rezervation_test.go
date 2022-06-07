package main

import (
	"testing"
)

func TestRezervation(t *testing.T) {

	t.Run("Rezervation person test", func(t *testing.T) {
		dummy_rez := getEmptyRez("Car 1", 50, 70, true)

		_, got := dummy_rez.RezervationFunc(-5)

		if got == nil {
			t.Errorf("expected error got nil")
		}
	})

	t.Run("Rezervation fullness can not be higher than %70 capacity ", func(t *testing.T) {
		dummy_rez := getEmptyRez("Car 1", 100, 70, true)

		_, got := dummy_rez.RezervationFunc(5)

		if got == nil {
			t.Errorf("expected error got nil")
		}

	})

	t.Run("1st car full 2nd car available", func(t *testing.T) {
		var dummy_rez [2]Train
		dummy_rez[0] = getEmptyRez("Car 1", 50, 35, true)
		dummy_rez[1] = getEmptyRez("Car 2", 100, 60, true)
		person := 3
		person, got := dummy_rez[0].RezervationFunc(person)
		if got == nil {
			t.Errorf("expected error got nil")
		}
		person, got2 := dummy_rez[1].RezervationFunc(person)
		if got2 != nil {
			t.Errorf("expected nil got error")
		}
	})
	t.Run("1st car half and 2nd car half available", func(t *testing.T) {
		dummy_rez := getEmptyRez("Car 1", 50, 33, true)
		dummy_rez2 := getEmptyRez("Car 2", 100, 68, true)
		person := 4
		person, got := dummy_rez.RezervationFunc(person)
		if got != nil {
			t.Errorf("expected nil got error")
		}
		person, got2 := dummy_rez2.RezervationFunc(person)
		if got2 != nil {
			t.Errorf("expected nil got error")
		}
	})
	t.Run("Don't allow different car", func(t *testing.T) {
		dummy_rez := getEmptyRez("Car 1", 100, 69, false)
		person := 4
		person, got := dummy_rez.RezervationFunc(person)
		if got == nil {
			t.Errorf("expected nil got error")
		}

	})

}

func getEmptyRez(carName string, capacity, fullness int, difCar bool) Train {
	return Train{
		CarName:  carName,
		Fullness: fullness,
		Capacity: capacity,
		difCar:   difCar,
	}
}
