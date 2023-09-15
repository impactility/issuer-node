package loader

import (
	"github.com/iden3/go-schema-processor/processor"
	"github.com/piprate/json-gold/ld"
)

// DocumentLoader is an alias for json-gold DocumentLoader
type DocumentLoader ld.DocumentLoader

// Loader defines a Loader interface
type Loader interface {
	processor.SchemaLoader
}

// Factory defines the interface that a loader constructor should satisfy
type Factory func(url string) Loader
