// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package metrics

import (
	"errors"
	"github.com/TeaOSLab/EdgeAdmin/internal/rpc"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/iwind/TeaGo/maps"
)

// InitItem 初始化指标信息
func InitItem(parent *actionutils.ParentAction, itemId int64) (*pb.MetricItem, error) {
	client, err := rpc.SharedRPC()
	if err != nil {
		return nil, err
	}
	resp, err := client.MetricItemRPC().FindEnabledMetricItem(parent.AdminContext(), &pb.FindEnabledMetricItemRequest{MetricItemId: itemId})
	if err != nil {
		return nil, err
	}
	var item = resp.MetricItem
	if item == nil {
		return nil, errors.New("not found")
	}
	parent.Data["item"] = maps.Map{
		"id":         item.Id,
		"name":       item.Name,
		"isOn":       item.IsOn,
		"keys":       item.Keys,
		"value":      item.Value,
		"period":     item.Period,
		"periodUnit": item.PeriodUnit,
		"category":   item.Category,
	}
	return item, nil
}