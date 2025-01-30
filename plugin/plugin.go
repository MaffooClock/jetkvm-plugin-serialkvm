package plugin

import (
  "context"
  "github.com/sourcegraph/jsonrpc2"
)

type PluginStatus struct {
  Status  string  `json:"status" oneOf:"pending-configuration,ready,error"`
  Message *string `json:"message,omitempty"`
}

type SupportedMethodsResponse struct {
  SupportedRpcMethods []string `json:"supported_rpc_methods"`
}

type PluginHandler interface {
  GetPluginSupportedMethods(ctx context.Context) (SupportedMethodsResponse, error)
  GetPluginStatus(ctx context.Context) (PluginStatus, error)
  //// Is this how I call the SwitchInput function?
  //DoSwitchInput(ctx context.Context, params json.RawMessage) (???, error)
}

func HandleRPC(handler PluginHandler) jsonrpc2.Handler {
  return jsonrpc2.HandlerWithError(func(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) (interface{}, error) {

    switch req.Method {
    //case "switchInput":
    //// is this where the command from the UI to switch the input is dispatched?
    //	return handler.DoSwitchInput(ctx, req.Params)
    case "getPluginSupportedMethods":
      return handler.GetPluginSupportedMethods(ctx)
    case "getPluginStatus":
      return handler.GetPluginStatus(ctx)
    }
    return nil, nil
  })
}
