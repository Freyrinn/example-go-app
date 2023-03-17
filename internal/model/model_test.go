package model

import "testing"

func TestPerson(t *testing.T) {
	p := Person{
		Name:  "John Doe",
		Phone: "123-456-7890",
	}

	if p.Name != "John Doe" {
		t.Errorf("unexpected name: expected=%s, got=%s", "John Doe", p.Name)
	}

	if p.Phone != "123-456-7890" {
		t.Errorf("unexpected phone number: expected=%s, got=%s", "123-456-7890", p.Phone)
	}
}

func TestData(t *testing.T) {
	p1 := Person{Name: "John Doe", Phone: "123-456-7890"}
	p2 := Person{Name: "Jane Doe", Phone: "987-654-3210"}

	d := Data{People: []Person{p1, p2}}

	if len(d.People) != 2 {
		t.Errorf("unexpected number of people: expected=%d, got=%d", 2, len(d.People))
	}

	if d.People[0].Name != "John Doe" {
		t.Errorf("unexpected name for first person: expected=%s, got=%s", "John Doe", d.People[0].Name)
	}

	if d.People[0].Phone != "123-456-7890" {
		t.Errorf("unexpected phone number for first person: expected=%s, got=%s", "123-456-7890", d.People[0].Phone)
	}

	if d.People[1].Name != "Jane Doe" {
		t.Errorf("unexpected name for second person: expected=%s, got=%s", "Jane Doe", d.People[1].Name)
	}

	if d.People[1].Phone != "987-654-3210" {
		t.Errorf("unexpected phone number for second person: expected=%s, got=%s", "987-654-3210", d.People[1].Phone)
	}
}
