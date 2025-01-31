package main

import (
	"context"
	"encoding/json"
	"fmt"

	"serialkvm/plugin"
)

type SwitchInputRequest struct {
	InputID int `json:"inputID"`
}

func (p *PluginImpl) GetPluginSupportedMethods(ctx context.Context) (plugin.SupportedMethodsResponse, error) {
	return plugin.SupportedMethodsResponse{
		SupportedRpcMethods: []string{"getPluginSupportedMethods", "getPluginStatus", "switchInput"},
	}, nil
}

func (p *PluginImpl) GetPluginStatus(ctx context.Context) (plugin.PluginStatus, error) {

	var status string
	if p.serialPort == nil {
		status = "pending-configuration"
	} else {
		status = "ready"
	}

	return plugin.PluginStatus{
		Status: status,
	}, nil
}

func (p *PluginImpl) DoSwitchInput(ctx context.Context, params *json.RawMessage) error {

	// So... Unmarshal the `params` so we can get the one that
	// has the selected input the user wants to switch to. But
	// before we can make that work, we need to setup a Struct
	// to define what a `param` looks like... I think?
	var request SwitchInputRequest

	err := json.Unmarshal(*params, &request)
	if err != nil {
		return fmt.Errorf("failed to read request parameters: %v", err)
	}

	return p.SwitchInput(request.InputID)
}
