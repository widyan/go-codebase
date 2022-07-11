package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestConvertTzToNormalError(t *testing.T) {
// 	// tms := time.Time
// 	dateStr := "2009-01-02 15:04:05"
// 	_, err := ConvertTzToNormal(dateStr)
// 	if err != nil {
// 		t.Error(err)
// 	}

// }

func TestUniqueString(t *testing.T) {
	distict := UniqueString([]string{"a", "a", "b", "b", "c", "c"})
	assert.Equal(t, distict, []string{"a", "b", "c"})
}

func TestUniqueInt(t *testing.T) {
	distict := UniqueInt([]int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 6})
	assert.Equal(t, distict, []int{1, 2, 3, 4, 5, 6})
}

func TestStringInSlice(t *testing.T) {
	slice := StringInSlice("a", []string{"a", "a", "b", "b", "c", "c"})
	assert.Equal(t, slice, true)
}

func TestStringInNotSlice(t *testing.T) {
	slice := StringInSlice("d", []string{"a", "a", "b", "b", "c", "c"})
	assert.Equal(t, slice, false)
}
