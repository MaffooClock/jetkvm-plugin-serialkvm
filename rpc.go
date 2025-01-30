package main

import (
	"context"

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
