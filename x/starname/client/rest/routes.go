package rest

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/umma-chain/umma-core/pkg/queries"
	"github.com/umma-chain/umma-core/pkg/utils"
	"github.com/umma-chain/umma-core/x/starname/types"
)

// txRouteList clubs together all the transaction routes, which are the transactions
// // that return the bytes to sign to send a request that modifies state to the domain module
var txRoutesList = map[string]func(cliCtx client.Context) http.HandlerFunc{
	"registerDomain":          registerDomainHandler,
	"addAccountCertificates":  addAccountCertificatesHandler,
	"delAccountCertificates":  delAccountCertificateHandler,
	"deleteAccount":           deleteAccountHandler,
	"deleteDomain":            deleteDomainHandler,
	"registerAccount":         registerAccountHandler,
	"renewAccount":            renewAccountHandler,
	"renewDomain":             renewDomainHandler,
	"replaceAccountResources": replaceAccountResourcesHandler,
	"transferAccount":         transferAccountHandler,
	"transferDomain":          transferDomainHandler,
	"setAccountMetadata":      setAccountMetadataHandler,
}

// registerTxRoutes registers all the transaction routes to the router
// the route will be exposed to storeName/handler, the handler will
// accept only post request with json codec
func registerTxRoutes(cliCtx client.Context, r *mux.Router, storeName string) {
	for route, handler := range txRoutesList {
		path := fmt.Sprintf("/%s/tx/%s", storeName, route)
		r.HandleFunc(path, handler(cliCtx))
	}
}

func queryHandlerBuild(cliCtx client.Context, storeName string, queryType queries.QueryHandler) http.HandlerFunc {
	// get query type
	typ := utils.GetPtrType(queryType)
	// return function
	return func(writer http.ResponseWriter, request *http.Request) {
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(writer, cliCtx, request)
		if !ok {
			return
		}
		// clone queryType so we can unmarshal data to it
		query := utils.CloneFromType(typ).(queries.QueryHandler)
		// read request bytes
		b, err := ioutil.ReadAll(request.Body)
		if err != nil {
			rest.WriteErrorResponse(writer, http.StatusInternalServerError, err.Error())
			return
		}
		// unmarshal request from the client to the query handler
		err = queries.DefaultQueryDecode(b, query)
		if err != nil {
			rest.WriteErrorResponse(writer, http.StatusBadRequest, err.Error())
			return
		}
		// verify query correctness
		if err = query.Validate(); err != nil {
			rest.WriteErrorResponse(writer, http.StatusBadRequest, err.Error())
			return
		}
		// marshal request to bytes understandable to the app query processor
		requestBytes, err := queries.DefaultQueryEncode(query)
		if err != nil {
			// this is an internal server error if we're not able to marshal a request TODO log
			rest.WriteErrorResponse(writer, http.StatusInternalServerError, err.Error())
			return
		}
		// build query path
		queryPath := fmt.Sprintf("custom/%s/%s", storeName, query.QueryPath())
		// do query
		res, height, err := cliCtx.QueryWithData(queryPath, requestBytes)
		if err != nil {
			rest.WriteErrorResponse(writer, http.StatusBadRequest, err.Error())
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		// success
		rest.PostProcessResponse(writer, cliCtx, res)
	}
}

// registerQueryRoutes registers all the routes used to query
// the domain module's keeper
func registerQueryRoutes(cliCtx client.Context, r *mux.Router, queries []queries.QueryHandler) {
	for _, query := range queries {
		path := fmt.Sprintf("/%s/query/%s", types.ModuleName, query.QueryPath())
		r.HandleFunc(path, queryHandlerBuild(cliCtx, types.ModuleName, query)).Methods("POST")
	}
}

// RegisterRoutes clubs together the tx and query routes
func RegisterRoutes(cliCtx client.Context, r *mux.Router, queries []queries.QueryHandler) {
	// register tx routes
	registerTxRoutes(cliCtx, r, types.ModuleName)
	// register query routes
	registerQueryRoutes(cliCtx, r, queries)
}
