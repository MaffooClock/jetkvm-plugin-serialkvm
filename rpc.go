package main

import (
	"context"
	"encoding/json"

	"serialkvm/plugin"
)

func (p *PluginImpl) GetPluginSupportedMethods(ctx context.Context) (plugin.SupportedMethodsResponse, error) {
	return plugin.SupportedMethodsResponse{
		SupportedRpcMethods: []string{"getPluginSupportedMethods", "getPluginStatus"},
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

	inputNumber := 1

	return p.SwitchInput(inputNumber)
}
