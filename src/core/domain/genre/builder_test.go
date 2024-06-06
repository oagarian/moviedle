package genre

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type genreBuilderTests struct{}

func TestGenreBuilder(t *testing.T) {
	testcases := genreBuilderTests{}
	t.Run("Test builder with valid params", testcases.testWithValidParams)
	t.Run("Test builder with invalid ID", testcases.testWithInvalidID)
	t.Run("Test builder with invalid name", testcases.testWithInvalidName)
	t.Run("Test builder with invalid code", testcases.testWithInvalidCode)
} 
func (bt *genreBuilderTests) testWithValidParams(t *testing.T) {
	builder := NewBuilder()
    builder.WithID(uuid.New())
    builder.WithName("Name")
    builder.WithCode("CODE")
	genre, err := builder.Build()
	assert.Nil(t, err)
	assert.NotNil(t, genre)
}

func (bt *genreBuilderTests) testWithInvalidID(t *testing.T) {
	builder := NewBuilder()
    builder.WithID(uuid.Nil)
    builder.WithName("Name")
    builder.WithCode("CODE")
    _, err := builder.Build()
    assert.NotNil(t, err)
}

func (bt *genreBuilderTests) testWithInvalidName(t *testing.T) {
	builder := NewBuilder()
    builder.WithID(uuid.New())
    builder.WithName("")
    builder.WithCode("CODE")
    _, err := builder.Build()
    assert.NotNil(t, err)
}

func (bt *genreBuilderTests) testWithInvalidCode(t *testing.T) {
	builder := NewBuilder()
    builder.WithID(uuid.New())
    builder.WithName("Name")
    builder.WithCode("")
    _, err := builder.Build()
    assert.NotNil(t, err)
}