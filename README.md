# tukilastic
--
    import "github.com/rotemtam/tukilastic"

Package tukilastic is a utility library to help you write data to AWS
Elasticsearch. It will take care of signing your requests against AWS, ensuring
the index you want to write to exists and write data using a type which
implements io.Writer.

## Example

```go
package your_package

import (
	"os"
	"log"
	"github.com/rotemtam/tukilastic"
)

func Example() {
	conf := tukilastic.Config{
		ElasticUrl:   os.Getenv("ES_URL"),
		Sniff:        false,
		Index:        "index-2017-02-09",
		DocumentType: "Demo",
		AwsRegion:    "us-east-1",
	}

	writer, err := tukilastic.New(conf)

	err = writer.InitIndex()


	val := `{"a":1}`

	n, err := writer.Write([]byte(val))

	if err != nil {
		log.Println("Error:", err)
	}

	log.Printf("Successfully wrote %d bytes", n)

}

```

## Usage

#### type Config

```go
type Config struct {
	ElasticUrl   string
	Sniff        bool
	Index        string
	DocumentType string
	AwsRegion    string
}
```

Config holds configuration for initializing a new tukilastic.Writer

#### type Writer

```go
type Writer struct {
	Client       *elastic.Client
	Index        string
	DocumentType string
}
```

Writer implements io.Writer and will write a document to Client on the index
Index and with a type of DocumentType

#### func  New

```go
func New(conf Config) (*Writer, error)
```
New returns a new Writer

#### func (*Writer) InitIndex

```go
func (e *Writer) InitIndex() error
```
InitIndex checks if the index exists on the elasticsearch cluster, if it doesn't
it will create it.

#### func (*Writer) Write

```go
func (e *Writer) Write(p []byte) (n int, err error)
```
Write writes a document to the elasticsearch cluster
