// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"nRIC/internal/xapp/restapi/operations"
	"nRIC/internal/xapp/restapi/operations/common"
	"nRIC/internal/xapp/restapi/operations/policy"
	"nRIC/internal/xapp/restapi/operations/query"
	"nRIC/internal/xapp/restapi/operations/report"
)

//go:generate swagger generate server --target ../../pkg --name XappFramework --spec ../../api/xapp_rest_api.yaml --exclude-main

func configureFlags(api *operations.XappFrameworkAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.XappFrameworkAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.CommonUnsubscribeHandler == nil {
		api.CommonUnsubscribeHandler = common.UnsubscribeHandlerFunc(func(params common.UnsubscribeParams) middleware.Responder {
			return middleware.NotImplemented("operation common.Unsubscribe has not yet been implemented")
		})
	}
	if api.QueryGetAllSubscriptionsHandler == nil {
		api.QueryGetAllSubscriptionsHandler = query.GetAllSubscriptionsHandlerFunc(func(params query.GetAllSubscriptionsParams) middleware.Responder {
			return middleware.NotImplemented("operation query.GetAllSubscriptions has not yet been implemented")
		})
	}
	if api.PolicySubscribePolicyHandler == nil {
		api.PolicySubscribePolicyHandler = policy.SubscribePolicyHandlerFunc(func(params policy.SubscribePolicyParams) middleware.Responder {
			return middleware.NotImplemented("operation policy.SubscribePolicy has not yet been implemented")
		})
	}
	if api.ReportSubscribeReportHandler == nil {
		api.ReportSubscribeReportHandler = report.SubscribeReportHandlerFunc(func(params report.SubscribeReportParams) middleware.Responder {
			return middleware.NotImplemented("operation report.SubscribeReport has not yet been implemented")
		})
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
