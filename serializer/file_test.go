package serializer

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/leogsouza/pcbook-go/pb"

	"github.com/stretchr/testify/require"

	"github.com/leogsouza/pcbook-go/sample"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"

	laptop1 := sample.NewLaptop()
	err := WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))

	err = WriteProtocolbufToJSONFile(laptop1, jsonFile)
	require.NoError(t, err)
}
