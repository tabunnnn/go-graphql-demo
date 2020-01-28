package provider

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/graphql-go/graphql"
	"tabu4n.me/graphql-demo/service"

	"github.com/gin-gonic/gin"
)

type GraphQLProvider struct{}

// Request a graphql request
type Request struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

var gqlProvider GraphQLProvider

// Get parse get request
func (p *GraphQLProvider) Get(c *gin.Context) {

	request := Request{
		Query:         c.Query("query"),
		OperationName: c.Query("operationName"),
	}

	result := p.execute(c, &request)

	c.JSON(http.StatusOK, result.Data)
}

// Post parse post request
func (p *GraphQLProvider) Post(c *gin.Context) {

	var err error

	defer func() {
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}()

	rb, err := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()

	if err != nil {
		return
	}

	var request Request

	if err = json.Unmarshal(rb, &request); err != nil {
		return
	}

	result := p.execute(c, &request)

	c.JSON(http.StatusOK, result.Data)
}

// execute query
func (p *GraphQLProvider) execute(ctx context.Context, request *Request) (result *graphql.Result) {
	query := graphql.Params{
		Context:        ctx,
		Schema:         service.Helloworld,
		RequestString:  request.Query,
		VariableValues: request.Variables,
		OperationName:  request.OperationName,
	}

	result = graphql.Do(query)

	return
}
