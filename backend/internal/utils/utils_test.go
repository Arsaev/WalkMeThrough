package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	// test Find with a slice of integers  and a function that returns true if the element is 5 and false otherwise
	// expected output is 5 and 4
	slice := []int{1, 2, 3, 4, 5, 6}
	f := func(v int) bool {
		return v == 5
	}
	v, i := Find(slice, f)
	assert.Equal(t, 5, v)
	assert.Equal(t, 4, i)

	// test Find with a slice of strings and a function that returns true if the element is "world" and false otherwise
	// expected output is "world" and 1
	slice2 := []string{"hello", "world", "!"}
	f2 := func(v string) bool {
		return v == "world"
	}

	v2, i2 := Find(slice2, f2)
	assert.Equal(t, "world", v2)
	assert.Equal(t, 1, i2)

	// test Find with a slice of structs and a function that returns true if the element is the struct with the name "world" and false otherwise
	// expected output is the struct with the name "world" and 1
	type Person struct {
		Name string
		Age  int
	}

	slice3 := []Person{
		{"hello", 20},
		{"world", 30},
		{"!", 40},
	}

	f3 := func(v Person) bool {
		return v.Name == "world"
	}

	v3, i3 := Find(slice3, f3)
	assert.Equal(t, Person{"world", 30}, v3)
	assert.Equal(t, 1, i3)

}
