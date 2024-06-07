package oscar

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type oscarBuilderTests struct{}

func TestOscarBuilder(t *testing.T) {
	testcases := oscarBuilderTests{}
	t.Run("Test builder with valid params", testcases.testWithValidParams)
	t.Run("Test builder with invalid ID", testcases.testWithInvalidID)
	t.Run("Test builder with invalid name", testcases.testWithInvalidName)
	t.Run("Test builder with invalid code", testcases.testWithInvalidCode)
} 
func (bt *oscarBuilderTests) testWithValidParams(t *testing.T) {
	builder := NewBuilder()
    builder.WithID(uuid.New())
    builder.WithName("Name")
    builder.WithCode("CODE")
	oscar, err := builder.Build()
	assert.Nil(t, err)
	assert.NotNil(t, oscar)
}

func (bt *oscarBuilderTests) testWithInvalidID(t *testing.T) {
	builder := NewBuilder()
    builder.WithID(uuid.Nil)
    builder.WithName("Name")
    builder.WithCode("CODE")
    _, err := builder.Build()
    assert.NotNil(t, err)
}

func (bt *oscarBuilderTests) testWithInvalidName(t *testing.T) {
	builder := NewBuilder()
    builder.WithID(uuid.New())
    builder.WithName("")
    builder.WithCode("CODE")
    _, err := builder.Build()
    assert.NotNil(t, err)
}

func (bt *oscarBuilderTests) testWithInvalidCode(t *testing.T) {
	builder := NewBuilder()
    builder.WithID(uuid.New())
    builder.WithName("Name")
    builder.WithCode("")
    _, err := builder.Build()
    assert.NotNil(t, err)
}