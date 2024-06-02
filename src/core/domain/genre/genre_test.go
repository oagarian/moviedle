package genre

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type genreTests struct{}

func TestGenreSetters(t *testing.T) {
	tests := genreTests{}
	t.Run("Test with valid parameters", tests.testWithValidParameters)
	t.Run("Test the ID setters", tests.testIDSetters)
	t.Run("Test the name setters", tests.testNameSetters)
	t.Run("Test the code setters", tests.testCodeSetters)
}

func (*genreTests) testWithValidParameters (t *testing.T) {
	genre := new(genre)
	id := uuid.New()
	genre.SetID(id)
    genre.SetName("Name")
    genre.SetCode("Code")

	assert.NotNil(t, genre)
	assert.Equal(t, &id, genre.ID())
	assert.Equal(t, "Name", genre.Name())
	assert.Equal(t, "Code", genre.Code())
}

func (*genreTests) testIDSetters (t *testing.T) {
	id := uuid.New()
	genre := genre{}
	err := genre.SetID(id)
	assert.Nil(t, err)
	assert.NotNil(t, genre)
	
	invalidID := uuid.Nil
	err = genre.SetID(invalidID)
	assert.NotNil(t, err)
}

func (*genreTests) testNameSetters (t *testing.T) {
	name := "Test name"
	genre := genre{}
	err := genre.SetName(name)
	assert.Nil(t, err)
	assert.NotNil(t, genre)
	
	invalidName := ""
	err = genre.SetName(invalidName)
	assert.NotNil(t, err)
}

func (*genreTests) testCodeSetters (t *testing.T) {
	code := "Test code"
    genre := genre{}
    err := genre.SetCode(code)
    assert.Nil(t, err)
    assert.NotNil(t, genre)
    
    invalidCode := ""
    err = genre.SetCode(invalidCode)
    assert.NotNil(t, err)
}