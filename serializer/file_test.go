package serializer_test

import (
	"testing"

	"github.com/anselmussuryo/savedbillpayment/pb"
	"github.com/anselmussuryo/savedbillpayment/proto/sample"
	"github.com/anselmussuryo/savedbillpayment/serializer"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/savedBillPayment.bin"
	jsonFile := "../tmp/savedBillPayment.json"

	savedBillPayment1 := sample.NewSavedBillPayment()

	err := serializer.WriteProtobufToBinaryFile(savedBillPayment1, binaryFile)
	require.NoError(t, err)

	err = serializer.WriteProtobufToJSONFile(savedBillPayment1, jsonFile)
	require.NoError(t, err)

	savedBillPayment2 := &pb.SavedBillPayment{}
	err = serializer.ReadProtobufFromBinaryFile(binaryFile, savedBillPayment2)
	require.NoError(t, err)

	require.True(t, proto.Equal(savedBillPayment1, savedBillPayment2))
}
