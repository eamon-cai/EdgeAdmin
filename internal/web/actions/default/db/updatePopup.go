package db

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/maps"
)

type UpdatePopupAction struct {
	actionutils.ParentAction
}

func (this *UpdatePopupAction) Init() {
	this.Nav("", "", "")
}

func (this *UpdatePopupAction) RunGet(params struct {
	NodeId int64
}) {
	nodeResp, err := this.RPC().DBNodeRPC().FindEnabledDBNode(this.AdminContext(), &pb.FindEnabledDBNodeRequest{NodeId: params.NodeId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	node := nodeResp.Node
	if node == nil {
		this.NotFound("dbNode", params.NodeId)
		return
	}

	this.Data["node"] = maps.Map{
		"id":          node.Id,
		"isOn":        node.IsOn,
		"name":        node.Name,
		"description": node.Description,
		"host":        node.Host,
		"port":        node.Port,
		"username":    node.Username,
		"password":    node.Password,
		"database":    node.Database,
	}

	this.Show()
}

func (this *UpdatePopupAction) RunPost(params struct {
	NodeId int64

	Name     string
	Host     string
	Port     int32
	Database string
	Username string
	Password string

	Description string
	IsOn        bool

	Must *actions.Must
}) {
	params.Must.
		Field("name", params.Name).
		Require("请输入节点名称").
		Field("host", params.Host).
		Require("请输入主机地址").
		Field("port", params.Port).
		Gt(0, "请输入正确的数据库端口").
		Lt(65535, "请输入正确的数据库端口").
		Field("database", params.Database).
		Require("请输入数据库名称").
		Field("username", params.Username).
		Require("请输入连接数据库的用户名")

	_, err := this.RPC().DBNodeRPC().UpdateDBNode(this.AdminContext(), &pb.UpdateDBNodeRequest{
		NodeId:      params.NodeId,
		IsOn:        params.IsOn,
		Name:        params.Name,
		Description: params.Description,
		Host:        params.Host,
		Port:        params.Port,
		Database:    params.Database,
		Username:    params.Username,
		Password:    params.Password,
		Charset:     "", // 暂时不能修改
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}