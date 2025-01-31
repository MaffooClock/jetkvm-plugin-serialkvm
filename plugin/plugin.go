package plugin

import (
	"context"
	"encoding/json"
	"github.com/sourcegraph/jsonrpc2"
)

type PluginStatus struct {
	Status  string  `json:"status" oneOf:"pending-configuration,running,error"`
	Message *string `json:"message,omitempty"`
}

type SupportedMethodsResponse struct {
	SupportedRpcMethods []string `json:"supported_rpc_methods"`
}

type PluginHandler interface {
	GetPluginSupportedMethods(ctx context.Context) (SupportedMethodsResponse, error)
	GetPluginStatus(ctx context.Context) (PluginStatus, error)
	DoSwitchInput(ctx context.Context, params *json.RawMessage) error
}

func HandleRPC(handler PluginHandler) jsonrpc2.Handler {
	return jsonrpc2.HandlerWithError(func(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) (interface{}, error) {

		//b, _ := json.Marshal(req)
		//log.Printf("Received request: %s", string(b))

		switch req.Method {
		case "switchInput":
			return nil, handler.DoSwitchInput(ctx, req.Params)

		case "getPluginSupportedMethods":
			return handler.GetPluginSupportedMethods(ctx)

		case "getPluginStatus":
			return handler.GetPluginStatus(ctx)
		}
		return nil, nil
	})
}
