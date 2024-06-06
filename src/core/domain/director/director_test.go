package director

import (
	"moviedle/src/core/domain/errors"
	"moviedle/src/core/messages"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDirect(t *testing.T) {
	testCases := directorTests{}
	t.Run("Test director with valid parameters", testCases.testWithValidParameters)
	t.Run("Test director ID setters", testCases.testIDSetters)
	t.Run("Test director name setters", testCases.testNameSetters)
}
type directorTests struct {}

func (*directorTests) testWithValidParameters (t *testing.T) {
	director := new(director)
	id := uuid.New()
    director.SetID(id)
    director.SetName("Name")
    assert.Equal(t, director.ID(), &id)
    assert.Equal(t, director.Name(), "Name")
}

func (*directorTests) testIDSetters (t *testing.T) {
	id := uuid.New()
    director := director{}
    err := director.SetID(id)
    assert.Nil(t, err)
    assert.NotNil(t, director)
    
    invalidID := uuid.Nil
    err = director.SetID(invalidID)
    assert.NotNil(t, err)
	assert.Equal(t, err, errors.NewValidationFromString(messages.DirectorIDInvalidIDMessage))
}

func (*directorTests) testNameSetters (t *testing.T) {
	name := "Test name"
    director := director{}
    err := director.SetName(name)
    assert.Nil(t, err)
    assert.NotNil(t, director)
    
    invalidName := ""
    err = director.SetName(invalidName)
    assert.NotNil(t, err)
	assert.Equal(t, err, errors.NewValidationFromString(messages.DirectorNameCannotBeEmptyMessage))
}