/*
Package tukilastic is a utility library to help you write data to AWS Elasticsearch.
It will take care of signing your requests against AWS, ensuring the index you want to write
to exists and write data using a type which implements io.Writer.
*/
package tukilastic

import (
	"context"
	"gopkg.in/olivere/elastic.v5"
)

// Config holds configuration for initializing a new tukilastic.Writer
type Config struct {
	ElasticUrl   string
	Sniff        bool
	Index        string
	DocumentType string
	AwsRegion    string
}

// Writer implements io.Writer and will write a document to Client on the index Index and with
// a type of DocumentType
type Writer struct {
	Client       *elastic.Client
	Index        string
	DocumentType string
}


// InitIndex checks if the index exists on the elasticsearch cluster, if it doesn't it will
// create it.
func (e *Writer) InitIndex() error {
	exists, err := e.Client.IndexExists(e.Index).Do(context.Background())

	if err != nil {
		return err
	}

	if !exists {
		_, createErr := e.Client.CreateIndex(e.Index).Do(context.Background())
		if createErr != nil {
			return createErr
		}
	}

	return nil
}


// Write writes a document to the elasticsearch cluster
func (e *Writer) Write(p []byte) (n int, err error) {
	_, err = e.Client.Index().
		Index(e.Index).
		Type(e.DocumentType).
		BodyString(string(p)).
		Do(context.Background())

	if err != nil {
		n = 0
	} else {
		n = len(p)
	}

	return n, err

}

// New returns a new Writer
func New(conf Config) (*Writer, error) {
	client, err := newClient(conf.ElasticUrl, conf.Sniff, conf.AwsRegion)

	if err != nil {
		return nil, err
	}

	writer := &Writer{
		Client:       client,
		Index:        conf.Index,
		DocumentType: conf.DocumentType,
	}

	return writer, nil

}