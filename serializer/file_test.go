package serializer

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/leogsouza/pcbook-go/sample"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"

	laptop1 := sample.NewLaptop()
	err := WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)
}
