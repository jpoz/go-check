package check

import (
	"fmt"
	"reflect"
)

// CheckEquality will check if two given interfaces are equal. Note it will
// only check the presense of object that are in the expected (e) interface.
// If the returned (r) interface has more object or keys no error will be
// returned
func CheckEquality(expected, given interface{}) (bool, error) {
	s := reflect.ValueOf(expected)
	switch s.Kind() {
	case reflect.Slice:
		return checkSliceEquality(expected, given)
	case reflect.Map:
		return checkMapEquality(expected, given)
	case reflect.String:
		return checkStringEquality(expected, given)
	case reflect.Struct:
		return checkStructEquality(expected, given)
	case reflect.Bool:
		return checkBoolEquality(expected, given)
	default:
		return false, fmt.Errorf("Can not check equality of type %v", s.Kind())
	}
}

func mapInterface(e interface{}) ([]interface{}, error) {
	s := reflect.ValueOf(e)
	if s.Kind() != reflect.Slice {
		empty := []interface{}{}
		return empty, fmt.Errorf("expected %q to be a slice", e)
	}

	slice := make([]interface{}, s.Len())
	for i := 0; i < s.Len(); i++ {
		slice[i] = s.Index(i).Interface()
	}

	return slice, nil
}

type equalityMatcher interface {
	CheckEquality(r interface{}) error
}

func checkSliceEquality(e, r interface{}) (bool, error) {
	expected, err := mapInterface(e)
	if err != nil {
		return false, err
	}
	returned, err := mapInterface(r)
	if err != nil {
		return false, err
	}

	if len(returned) != len(expected) {
		return false, fmt.Errorf("wanted %d objects got %d\n%q",
			len(expected), len(returned), returned)
	}

	for i := range expected {
		want := expected[i]
		got := returned[i]

		b, err := CheckEquality(want, got)
		if err != nil {
			return b, err
		}
	}

	return true, nil
}

func checkMapEquality(e, r interface{}) (bool, error) {
	s := reflect.ValueOf(e)
	t := reflect.ValueOf(r)
	if s.Kind() != t.Kind() {
		return false, fmt.Errorf("%v does not match %v", s.Kind(), t.Kind())
	}
	if s.Kind() != reflect.Map {
		return false, fmt.Errorf("%v is not a Map", s)
	}

	expected := e.(map[string]interface{})
	returned := r.(map[string]interface{})

	for k := range expected {
		want := expected[k]
		got := returned[k]

		b, err := CheckEquality(want, got)
		if err != nil {
			return b, err
		}
	}

	return true, nil
}

func checkStringEquality(e, r interface{}) (bool, error) {
	s := reflect.ValueOf(e)
	t := reflect.ValueOf(r)
	if s.Kind() != t.Kind() {
		return false, fmt.Errorf("%v does not match %v", s.Kind(), t.Kind())
	}
	if s.Kind() != reflect.String {
		return false, fmt.Errorf("%v is not a String", s)
	}
	expected := e.(string)
	returned := r.(string)

	if expected != returned {
		return false, fmt.Errorf("wanted %q\nbut got\n%q",
			expected, returned)
	}
	return true, nil
}

func checkBoolEquality(e, r interface{}) (bool, error) {
	s := reflect.ValueOf(e)
	t := reflect.ValueOf(r)
	if s.Kind() != t.Kind() {
		return false, fmt.Errorf("%v does not match %v", s.Kind(), t.Kind())
	}
	if s.Kind() != reflect.Bool {
		return false, fmt.Errorf("%v is not a String", s)
	}
	expected := e.(bool)
	returned := r.(bool)

	if expected != returned {
		return false, fmt.Errorf("wanted %t\nbut got\n%t",
			expected, returned)
	}
	return true, nil
}

func checkStructEquality(e, r interface{}) (bool, error) {
	s := reflect.ValueOf(e)
	if s.Kind() != reflect.Struct {
		return false, fmt.Errorf("%v is not a Struct", s)
	}

	if expected, ok := e.(equalityMatcher); ok {
		return false, expected.CheckEquality(r)
	}
	if !reflect.DeepEqual(e, r) {
		return false, fmt.Errorf("wanted %q\nbut got\n%q",
			e, r)
	}
	return true, nil
}
