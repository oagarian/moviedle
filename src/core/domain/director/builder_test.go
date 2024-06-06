package director

import (
	"testing"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type genreBuilderTests struct{}

func TestDirectorBuilder(t *testing.T) {
	testcases := genreBuilderTests{}
	t.Run("Test builder with valid params", testcases.testWithValidParams)
	t.Run("Test builder with invalid ID", testcases.testWithInvalidID)
	t.Run("Test builder with invalid name", testcases.testWithInvalidName)
} 
func (bt *genreBuilderTests) testWithValidParams(t *testing.T) {
	builder := NewBuilder()
    builder.WithID(uuid.New())
    builder.WithName("Name")
	genre, err := builder.Build()
	assert.Nil(t, err)
	assert.NotNil(t, genre)
}

func (bt *genreBuilderTests) testWithInvalidID(t *testing.T) {
	builder := NewBuilder()
    builder.WithID(uuid.Nil)
    builder.WithName("Name")
    _, err := builder.Build()
    assert.NotNil(t, err)
}

func (bt *genreBuilderTests) testWithInvalidName(t *testing.T) {
	builder := NewBuilder()
    builder.WithID(uuid.New())
    builder.WithName("")
    _, err := builder.Build()
    assert.NotNil(t, err)
}
