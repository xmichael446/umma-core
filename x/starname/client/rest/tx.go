package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/umma-chain/umma-core/x/starname/types"
)

// handleTxRequest is a helper function that takes care of checking base requests, sdk messages, after verifying
// requests it forwards an error to the client in case of error, otherwise it will return a transaction to sign
// and send to the /tx endpoint to do a request
func handleTxRequest(cliCtx client.Context, baseReq rest.BaseReq, msg sdk.Msg, writer http.ResponseWriter) {
	baseReq = baseReq.Sanitize()
	if !baseReq.ValidateBasic(writer) {
		return
	}
	// validate request
	if err := msg.ValidateBasic(); err != nil {
		rest.WriteErrorResponse(writer, http.StatusBadRequest, err.Error())
	}
	// write tx
	tx.WriteGeneratedTxResponse(cliCtx, writer, baseReq, msg)
}

// registerDomain defines the request model used for registerDomainHandler
type registerDomain struct {
	BaseReq rest.BaseReq             `json:"base_req"`
	Message *types.MsgRegisterDomain `json:"message"`
}

// registerDomainHandler builds the transaction to sign
func registerDomainHandler(cliCtx client.Context) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var req registerDomain
		if !rest.ReadRESTReq(writer, request, cliCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(writer, http.StatusBadRequest, "failed to parse request")
		}
		handleTxRequest(cliCtx, req.BaseReq, req.Message, writer)
	}
}

// addAccountCertificates is the request model for addAccountCertificatesHandler
type addAccountCertificates struct {
	BaseReq rest.BaseReq                    `json:"base_req"`
	Message *types.MsgAddAccountCertificate `json:"message"`
}

// addAccountCertificatesHandler builds the transaction to sign to add account certificates
func addAccountCertificatesHandler(cliCtx client.Context) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var req addAccountCertificates
		if !rest.ReadRESTReq(writer, request, cliCtx.LegacyAmino, &req) {
			return
		}
		handleTxRequest(cliCtx, req.BaseReq, req.Message, writer)
	}
}

// delAccountCertificate is the request model for delAccountCertificateHandler
type delAccountCertificate struct {
	BaseReq rest.BaseReq                       `json:"base_req"`
	Message *types.MsgDeleteAccountCertificate `json:"message"`
}

// delAccountCertificateHandler builds the transaction to sign to delete account certificates
func delAccountCertificateHandler(cliCtx client.Context) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var req delAccountCertificate
		if !rest.ReadRESTReq(writer, request, cliCtx.LegacyAmino, &req) {
			return
		}
		handleTxRequest(cliCtx, req.BaseReq, req.Message, writer)
	}
}

// deleteAccount is the request
type deleteAccount struct {
	BaseReq rest.BaseReq            `json:"base_req"`
	Message *types.MsgDeleteAccount `json:"message"`
}

// deleteAccountHandler builds the transaction to sign to delete an account
func deleteAccountHandler(cliCtx client.Context) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var req deleteAccount
		if !rest.ReadRESTReq(writer, request, cliCtx.LegacyAmino, &req) {
			return
		}
		handleTxRequest(cliCtx, req.BaseReq, req.Message, writer)
	}
}

// deleteDomain is the request model for deleteDomainHandler
type deleteDomain struct {
	BaseReq rest.BaseReq           `json:"base_req"`
	Message *types.MsgDeleteDomain `json:"message"`
}

// deleteDomainHandler builds the transaction to sign to delete a domain
func deleteDomainHandler(cliCtx client.Context) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var req deleteDomain
		if !rest.ReadRESTReq(writer, request, cliCtx.LegacyAmino, &req) {
			return
		}
		handleTxRequest(cliCtx, req.BaseReq, req.Message, writer)
	}
}

// registerAccount is the request model used for registerAccountHandler
type registerAccount struct {
	BaseReq rest.BaseReq              `json:"base_req"`
	Message *types.MsgRegisterAccount `json:"message"`
}

// registerAccountHandler builds the transaction to sign to register an account
func registerAccountHandler(cliCtx client.Context) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var req registerAccount
		if !rest.ReadRESTReq(writer, request, cliCtx.LegacyAmino, &req) {
			return
		}
		handleTxRequest(cliCtx, req.BaseReq, req.Message, writer)
	}
}

// renewAccount is the request model for renewAccountHandler
type renewAccount struct {
	BaseReq rest.BaseReq           `json:"base_req"`
	Message *types.MsgRenewAccount `json:"message"`
}

// renewAccountHandler builds the transaction request to sign to renew a domain
func renewAccountHandler(cliCtx client.Context) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var req renewAccount
		if !rest.ReadRESTReq(writer, request, cliCtx.LegacyAmino, &req) {
			return
		}
		handleTxRequest(cliCtx, req.BaseReq, req.Message, writer)
	}
}

// renewDomain is the request model for renewDomainHandler
type renewDomain struct {
	BaseReq rest.BaseReq          `json:"base_req"`
	Message *types.MsgRenewDomain `json:"message"`
}

// renewDomainHandler builds the transaction to sign to renew a domain
func renewDomainHandler(cliCtx client.Context) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var req renewDomain
		if !rest.ReadRESTReq(writer, request, cliCtx.LegacyAmino, &req) {
			return
		}
		handleTxRequest(cliCtx, req.BaseReq, req.Message, writer)
	}
}

// replaceAccountResources is the request model for replaceAccountResources
type replaceAccountResources struct {
	BaseReq rest.BaseReq                      `json:"base_req"`
	Message *types.MsgReplaceAccountResources `json:"message"`
}

// replaceAccountResources builds the transaction to sign to replace account resources
func replaceAccountResourcesHandler(cliCtx client.Context) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var req replaceAccountResources
		if !rest.ReadRESTReq(writer, request, cliCtx.LegacyAmino, &req) {
			return
		}
		handleTxRequest(cliCtx, req.BaseReq, req.Message, writer)
	}
}

// transferAccount is the request model for transferAccountHandler
type transferAccount struct {
	BaseReq rest.BaseReq              `json:"base_req"`
	Message *types.MsgTransferAccount `json:"message"`
}

// transferAccountHandler builds the transaction to sign to transfer accounts
func transferAccountHandler(cliCtx client.Context) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var req transferAccount
		if !rest.ReadRESTReq(writer, request, cliCtx.LegacyAmino, &req) {
			return
		}
		handleTxRequest(cliCtx, req.BaseReq, req.Message, writer)
	}
}

// transferDomain is the request model for transferDomainHandler
type transferDomain struct {
	BaseReq rest.BaseReq             `json:"base_req"`
	Message *types.MsgTransferDomain `json:"message"`
}

// transferDomainHandler builds the transaction to sign to transfer domains
func transferDomainHandler(cliCtx client.Context) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var req transferDomain
		if !rest.ReadRESTReq(writer, request, cliCtx.LegacyAmino, &req) {
			return
		}
		handleTxRequest(cliCtx, req.BaseReq, req.Message, writer)
	}
}

// transferDomain is the request model for transferDomainHandler
type setAccountMetadata struct {
	BaseReq rest.BaseReq                     `json:"base_req"`
	Message *types.MsgReplaceAccountMetadata `json:"message"`
}

// transferDomainHandler builds the transaction to sign to transfer domains
func setAccountMetadataHandler(cliCtx client.Context) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var req setAccountMetadata
		if !rest.ReadRESTReq(writer, request, cliCtx.LegacyAmino, &req) {
			return
		}
		handleTxRequest(cliCtx, req.BaseReq, req.Message, writer)
	}
}
