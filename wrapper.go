package endpoint_wrapper

import (
	"github.com/ValGoldun/httperror"
	"github.com/ValGoldun/logger"
	"github.com/gin-gonic/gin"
)

type EndpointWrapper struct {
	problemWriter httperror.ProblemWriter
}

func NewEndpointWrapper(logger logger.Logger) EndpointWrapper {
	return EndpointWrapper{
		problemWriter: httperror.NewProblemWriter(logger),
	}
}

func (r *EndpointWrapper) Endpoint(handler func(ctx *gin.Context) error) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		err := handler(ctx)
		r.problemWriter.Problem(ctx, err)
	}
}
