package api

import (
	"context"
	"github.com/pennsieve/repo-service/internal/api/response"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type LambdaHandler func(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error)

func Handler() LambdaHandler {
	return RepoServiceAPIHandler()
}

func RepoServiceAPIHandler() LambdaHandler {
	return func(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
		//routeKey := request.RouteKey
		//logger := logging.Default.With(slog.String("routeKey", routeKey),
		//	slog.String("requestId", request.RequestContext.RequestID))
		//
		//claims := authorizer.ParseClaims(request.RequestContext.Authorizer.Lambda)

		response := events.APIGatewayV2HTTPResponse{
			StatusCode: http.StatusOK,
			Body:       "{}",
			Headers:    response.StandardResponseHeaders(nil),
		}
		return response, nil
	}
}
