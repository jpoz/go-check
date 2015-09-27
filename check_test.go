package check

import "testing"

type Foo struct {
	Str string
	B   Boo
	I   int
}

type Boo struct {
	Str *string
	I   int
}

func TestCheckEquality(t *testing.T) {
	str := "string"
	str2 := "anotherstring"
	foo := Foo{
		Str: "string",
		B: Boo{
			Str: &str,
			I:   0,
		},
		I: 1,
	}

	// Complex structs should equal themselves
	b, err := CheckEquality(foo, foo)
	if !b {
		t.Errorf("CheckEquality(%#v,%#v)\nreturned: %#v\nwanted: true", foo, foo, b)
	}
	if err != nil {
		t.Errorf("CheckEquality(%#v,%#v)\nreturned: %#v, %#v\nwanted: true, nil",
			foo, foo, b, err)
	}

	// Strings should equal themselves
	b, err = CheckEquality(str, str)
	if !b {
		t.Errorf("CheckEquality(%#v,%#v)\nreturned: false\nwanted: true", str, str)
	}
	if err != nil {
		t.Errorf("CheckEquality(%#v,%#v)\nreturned: %#v, %#v\nwanted: true, nil",
			str, str, b, err)
	}
	// Strings should equal themselves
	b, err = CheckEquality(str, str2)
	if b {
		t.Errorf("CheckEquality(%q,%q)\nreturned: %v\nwanted: true", str, str2, b)
	}
	if err.Error() != "wanted \"string\"\nbut got\n\"anotherstring\"" {
		t.Errorf("CheckEquality(%q,%q)\nreturned: %#v\nwanted: wanted \"string\"\nbut got\n\"anotherstring\"", str, str2, err)
	}
}
