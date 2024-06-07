package oscar

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type oscarTests struct{}

func TestGenreSetters(t *testing.T) {
	tests := oscarTests{}
	t.Run("Test with valid parameters", tests.testWithValidParameters)
	t.Run("Test the ID setters", tests.testIDSetters)
	t.Run("Test the name setters", tests.testNameSetters)
	t.Run("Test the code setters", tests.testCodeSetters)
}

func (*oscarTests) testWithValidParameters (t *testing.T) {
	oscar := new(oscar)
	id := uuid.New()
	oscar.SetID(id)
    oscar.SetName("Name")
    oscar.SetCode("Code")

	assert.NotNil(t, oscar)
	assert.Equal(t, &id, oscar.ID())
	assert.Equal(t, "Name", oscar.Name())
	assert.Equal(t, "Code", oscar.Code())
}

func (*oscarTests) testIDSetters (t *testing.T) {
	id := uuid.New()
	oscar := oscar{}
	err := oscar.SetID(id)
	assert.Nil(t, err)
	assert.NotNil(t, oscar)
	
	invalidID := uuid.Nil
	err = oscar.SetID(invalidID)
	assert.NotNil(t, err)
}

func (*oscarTests) testNameSetters (t *testing.T) {
	name := "Test name"
	oscar := oscar{}
	err := oscar.SetName(name)
	assert.Nil(t, err)
	assert.NotNil(t, oscar)
	
	invalidName := ""
	err = oscar.SetName(invalidName)
	assert.NotNil(t, err)
}

func (*oscarTests) testCodeSetters (t *testing.T) {
	code := "Test code"
    oscar := oscar{}
    err := oscar.SetCode(code)
    assert.Nil(t, err)
    assert.NotNil(t, oscar)
    
    invalidCode := ""
    err = oscar.SetCode(invalidCode)
    assert.NotNil(t, err)
}