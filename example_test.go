package tukilastic

import (
	"os"
	"log"
)

func Example() {
	conf := Config{
		ElasticUrl:   os.Getenv("ES_URL"),
		Sniff:        false,
		Index:        "index-2017-02-09",
		DocumentType: "Demo",
		AwsRegion:    "us-east-1",
	}

	writer, err := New(conf)

	err = writer.InitIndex()


	val := `{"a":1}`

	n, err := writer.Write([]byte(val))

	if err != nil {
		log.Println("Error:", err)
	}

	log.Printf("Successfully wrote %d bytes", n)

}
