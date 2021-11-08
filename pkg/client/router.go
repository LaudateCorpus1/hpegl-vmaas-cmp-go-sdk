// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	"context"
	"encoding/json"
	"fmt"

	consts "github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/common"
	"github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/models"
)

type RouterAPIService struct {
	Client APIClientHandler
	Cfg    Configuration
}

func (r *RouterAPIService) GetAllRouter(
	ctx context.Context,
	queryParams map[string]string,
) (models.GetAllNetworkRouter, error) {
	routerResp := models.GetAllNetworkRouter{}
	serverAPI := &api{
		method: "GET",
		path: fmt.Sprintf("%s/%s/%s/%s", r.Cfg.Host, consts.VmaasCmpAPIBasePath,
			consts.NetworksPath, consts.NetworkRouterPath),
		client: r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &routerResp)
		},
	}
	err := serverAPI.do(ctx, nil, queryParams)

	return routerResp, err
}

func (r *RouterAPIService) GetSpecificRouter(
	ctx context.Context,
	routerID int,
) (models.GetNetworkRouter, error) {
	routerResp := models.GetNetworkRouter{}
	serverAPI := &api{
		method: "GET",
		path: fmt.Sprintf("%s/%s/%s/%s/%d", r.Cfg.Host, consts.VmaasCmpAPIBasePath,
			consts.NetworksPath, consts.NetworkRouterPath, routerID),
		client: r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &routerResp)
		},
	}
	err := serverAPI.do(ctx, nil, nil)

	return routerResp, err
}

func (r *RouterAPIService) CreateRouter(
	ctx context.Context,
	request models.CreateRouterRequest,
) (models.CreateRouterResp, error) {
	routerResp := models.CreateRouterResp{}
	serverAPI := &api{
		method: "POST",
		path: fmt.Sprintf("%s/%s/%s/%s", r.Cfg.Host, consts.VmaasCmpAPIBasePath,
			consts.NetworksPath, consts.NetworkRouterPath),
		client: r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &routerResp)
		},
	}
	err := serverAPI.do(ctx, request, nil)

	return routerResp, err
}

func (r *RouterAPIService) UpdateRouter(
	ctx context.Context,
	routerID int,
	request models.CreateRouterRequest,
) (models.SuccessOrErrorMessage, error) {
	routerResp := models.SuccessOrErrorMessage{}
	serverAPI := &api{
		method: "PUT",
		path: fmt.Sprintf("%s/%s/%s/%s/%d", r.Cfg.Host, consts.VmaasCmpAPIBasePath,
			consts.NetworksPath, consts.NetworkRouterPath, routerID),
		client: r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &routerResp)
		},
	}
	err := serverAPI.do(ctx, request, nil)

	return routerResp, err
}

func (r *RouterAPIService) DeleteRouter(
	ctx context.Context,
	routerID int,
) (models.SuccessOrErrorMessage, error) {
	routerResp := models.SuccessOrErrorMessage{}
	serverAPI := &api{
		method: "DELETE",
		path: fmt.Sprintf("%s/%s/%s/%s/%d", r.Cfg.Host, consts.VmaasCmpAPIBasePath,
			consts.NetworksPath, consts.NetworkRouterPath, routerID),
		client: r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &routerResp)
		},
	}
	err := serverAPI.do(ctx, nil, nil)

	return routerResp, err
}

func (r *RouterAPIService) GetRouterTypes(
	ctx context.Context,
	queryParams map[string]string,
) (models.GetNetworlRouterTypes, error) {
	routerResp := models.GetNetworlRouterTypes{}
	serverAPI := &api{
		method: "GET",
		path: fmt.Sprintf("%s/%s/%s", r.Cfg.Host, consts.VmaasCmpAPIBasePath,
			consts.NetworkRouterTypePath),
		client: r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &routerResp)
		},
	}
	err := serverAPI.do(ctx, nil, queryParams)

	return routerResp, err
}

func (r *RouterAPIService) GetNetworkServices(
	ctx context.Context,
	queryParams map[string]string,
) (models.GetNetworkServicesResp, error) {
	routerResp := models.GetNetworkServicesResp{}
	serverAPI := &api{
		method: "GET",
		path: fmt.Sprintf("%s/%s/%s/%s", r.Cfg.Host, consts.VmaasCmpAPIBasePath, consts.NetworksPath,
			consts.NetworkServicePath),
		client: r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &routerResp)
		},
	}
	err := serverAPI.do(ctx, nil, queryParams)

	return routerResp, err
}

func (r *RouterAPIService) CreateRouterNat(
	ctx context.Context,
	routerID int,
	request models.CreateRouterNatRequest,
) (models.CreateRouterNatResponse, error) {
	natResp := models.CreateRouterNatResponse{}
	serverAPI := &api{
		method: "POST",
		path: fmt.Sprintf("%s/%s/%s/%s/%d/%s", r.Cfg.Host, consts.VmaasCmpAPIBasePath, consts.NetworksPath,
			consts.NetworkRouterPath, routerID, consts.RoutersNatPath),
		client: r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &natResp)
		},
	}
	err := serverAPI.do(ctx, request, nil)

	return natResp, err
}

func (r *RouterAPIService) GetSpecificRouterNat(
	ctx context.Context,
	routerID, natID int,
) (models.GetSpecificRouterNatResponse, error) {
	natResp := models.GetSpecificRouterNatResponse{}
	serverAPI := &api{
		method: "GET",
		path: fmt.Sprintf("%s/%s/%s/%s/%d/%s/%d",
			r.Cfg.Host, consts.VmaasCmpAPIBasePath, consts.NetworksPath,
			consts.NetworkRouterPath, routerID, consts.RoutersNatPath, natID),
		client: r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &natResp)
		},
	}
	err := serverAPI.do(ctx, nil, nil)

	return natResp, err
}

func (r *RouterAPIService) UpdateRouterNat(
	ctx context.Context,
	routerID, natID int,
	req models.CreateRouterNatRequest,
) (models.CreateRouterNatResponse, error) {
	natResp := models.CreateRouterNatResponse{}
	serverAPI := &api{
		method: "PUT",
		path: fmt.Sprintf("%s/%s/%s/%s/%d/%s/%d",
			r.Cfg.Host, consts.VmaasCmpAPIBasePath, consts.NetworksPath,
			consts.NetworkRouterPath, routerID, consts.RoutersNatPath, natID),
		client: r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &natResp)
		},
	}
	err := serverAPI.do(ctx, req, nil)

	return natResp, err
}

func (r *RouterAPIService) DeleteRouterNat(
	ctx context.Context,
	routerID, natID int,
) (models.SuccessOrErrorMessage, error) {
	natResp := models.SuccessOrErrorMessage{}
	serverAPI := &api{
		method: "DELETE",
		path: fmt.Sprintf("%s/%s/%s/%s/%d/%s/%d",
			r.Cfg.Host, consts.VmaasCmpAPIBasePath, consts.NetworksPath,
			consts.NetworkRouterPath, routerID, consts.RoutersNatPath, natID),
		client: r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &natResp)
		},
	}
	err := serverAPI.do(ctx, nil, nil)

	return natResp, err
}

func (r *RouterAPIService) CreateRouterFirewallRuleGroup(
	ctx context.Context,
	routerID int,
	request models.CreateRouterFirewallRuleGroupRequest,
) (models.CreateRouterFirewallRuleGroupResponse, error) {
	firewallGroupResp := models.CreateRouterFirewallRuleGroupResponse{}
	serverAPI := &api{
		method: "POST",
		path: fmt.Sprintf("%s/%s/%s/%s/%d/%s", r.Cfg.Host, consts.VmaasCmpAPIBasePath, consts.NetworksPath,
			consts.NetworkRouterPath, routerID, consts.RoutersFirewallRuleGroupPath),
		client: r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &firewallGroupResp)
		},
	}
	err := serverAPI.do(ctx, request, nil)

	return firewallGroupResp, err
}

func (r *RouterAPIService) GetSpecificRouterFirewallRuleGroup(
	ctx context.Context,
	routerID, firewallGroupID int,
) (models.GetSpecificRouterFirewallRuleGroupResponse, error) {
	firewallGroupResp := models.GetSpecificRouterFirewallRuleGroupResponse{}
	serverAPI := &api{
		method: "GET",
		path: fmt.Sprintf("%s/%s/%s/%s/%d/%s/%d",
			r.Cfg.Host, consts.VmaasCmpAPIBasePath, consts.NetworksPath,
			consts.NetworkRouterPath, routerID, consts.RoutersFirewallRuleGroupPath, firewallGroupID),
		client: r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &firewallGroupResp)
		},
	}
	err := serverAPI.do(ctx, nil, nil)

	return firewallGroupResp, err
}

func (r *RouterAPIService) DeleteRouterFirewallRuleGroup(
	ctx context.Context,
	routerID, firewallGroupID int,
) (models.SuccessOrErrorMessage, error) {
	firewallGroupResp := models.SuccessOrErrorMessage{}
	serverAPI := &api{
		method: "DELETE",
		path: fmt.Sprintf("%s/%s/%s/%s/%d/%s/%d",
			r.Cfg.Host, consts.VmaasCmpAPIBasePath, consts.NetworksPath,
			consts.NetworkRouterPath, routerID, consts.RoutersFirewallRuleGroupPath, firewallGroupID),
		client: r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &firewallGroupResp)
		},
	}
	err := serverAPI.do(ctx, nil, nil)

	return firewallGroupResp, err
}
