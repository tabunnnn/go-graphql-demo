package service

import (
	"log"

	"github.com/graphql-go/graphql"
)

var Helloworld graphql.Schema

// init prepare Helloworld as a graphql Schema
// see [example in github.com/graphql-go/graphql](https://github.com/graphql-go/graphql/blob/master/examples/hello-world/main.go)
func init() {

	var err error

	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}

	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}

	if Helloworld, err = graphql.NewSchema(schemaConfig); err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
}
