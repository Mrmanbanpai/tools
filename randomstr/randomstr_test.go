package randomstr

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomString(t *testing.T) {
	assert := assert.New(t)

	{
		randStr := RandomString(Number, 10)
		assert.Equal(len(randStr), 10)
		for _, v := range randStr {
			assert.True(strings.Contains(numberByte, string(v)))
		}
	}

	{
		randStr := RandomString(UpperByte, 10)
		assert.Equal(len(randStr), 10)
		for _, v := range randStr {
			assert.True(strings.Contains(upperByte, string(v)))
		}
	}

	{
		randStr := RandomString(LowerByte, 10)
		assert.Equal(len(randStr), 10)
		for _, v := range randStr {
			assert.True(strings.Contains(lowerByte, string(v)))
		}
	}

	{
		randStr := RandomString(UpperByteWithNumber, 10)
		assert.Equal(len(randStr), 10)
		for _, v := range randStr {
			assert.True(strings.Contains(upperByte, string(v)) || strings.Contains(numberByte, string(v)))
		}
	}

	{
		randStr := RandomString(LowerByteWithNumber, 10)
		assert.Equal(len(randStr), 10)
		for _, v := range randStr {
			assert.True(strings.Contains(lowerByte, string(v)) || strings.Contains(numberByte, string(v)))
		}
	}

	{
		randStr := RandomString(UpperAndLowerByte, 10)
		assert.Equal(len(randStr), 10)
		for _, v := range randStr {
			assert.True(strings.Contains(lowerByte, string(v)) || strings.Contains(upperByte, string(v)))
		}
	}

	{
		randStr := RandomString(AllByte, 10)
		assert.Equal(len(randStr), 10)
		for _, v := range randStr {
			assert.True(strings.Contains(numberByte, string(v)) || strings.Contains(lowerByte, string(v)) || strings.Contains(upperByte, string(v)))
		}
	}
}
