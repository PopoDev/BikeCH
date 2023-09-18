package utils_test

import (
	"BikeCH/src/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckArgumentSucceedsForTrue(t *testing.T) {
	assert := assert.New(t)
	assert.NotPanics(func() {
		utils.CheckArgument(true, "message")
	})
}

func TestCheckArgumentPanicsForFalse(t *testing.T) {
	assert := assert.New(t)
	assert.Panics(func() {
		utils.CheckArgument(false, "message")
	})
}
