package tukilastic

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestWriter(t *testing.T) {
	conf := Config{
		ElasticUrl:   os.Getenv("ES_URL"),
		Sniff:        false,
		Index:        "serverless-logs-2017-02-09",
		DocumentType: "Demo",
		AwsRegion:    "us-east-1",
	}

	writer, err := New(conf)

	require.Nil(t, err)

	err = writer.InitIndex()

	require.Nil(t, err)

	val := `{"a":1}`

	n, err := writer.Write([]byte(val))

	require.Equal(t, n, len(val))

}
