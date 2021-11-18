// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	"context"
	"encoding/json"

	consts "github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/common"
	"github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/models"
)

type CmpStatus struct {
	Client APIClientHandler
	Cfg    Configuration
}

func (a *CmpStatus) GetCmpVersion(ctx context.Context) (models.CmpVersionModel, error) {
	checkResp := models.CmpVersionModel{}

	allCloudDSAPI := &api{
		method: "GET",
		path:   consts.WhoamiPath,
		client: a.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &checkResp)
		},
	}
	err := allCloudDSAPI.do(ctx, nil, nil)

	return checkResp, err
}
