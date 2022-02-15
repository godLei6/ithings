package tests

import (
	"context"
	"fmt"
	"github.com/go-things/things/shared/errors"
	"github.com/go-things/things/src/dmsvr/dm"
	"github.com/go-things/things/src/dmsvr/dmclient"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
	"testing"
)

func TestGetProductInfo(t *testing.T) {
	fmt.Println("GetProductInfo")
	client := dmclient.NewDm(zrpc.MustNewClient(zrpc.RpcClientConf{Etcd: discov.EtcdConf{
		Hosts: []string{"127.0.0.1:2379"},
		Key:   "dm.rpc",
	}}))
	ctx := context.Background()
	productID := "21CYs1k9YpG"
	{
		req := &dm.GetProductInfoReq{
			ProductID: productID, //产品id
			//Page       : productID,//分页信息 只获取一个则不填
		}
		info, err := client.GetProductInfo(ctx, req)
		t.Log(req, info)
		if err != nil {
			t.Errorf("%+v", errors.Fmt(err))
		}
	}
	{
		req := &dm.GetProductInfoReq{
			Page: &dm.PageInfo{
				Page:     1,
				PageSize: 20,
			}, //分页信息 只获取一个则不填
		}
		info, err := client.GetProductInfo(ctx, req)
		t.Log(req, info)
		if err != nil {
			t.Errorf("%+v", errors.Fmt(err))
		}
	}
	{
		req := &dm.GetProductInfoReq{
			ProductID: "123", //产品id
			Page: &dm.PageInfo{
				Page:     1,
				PageSize: 20,
			}, //分页信息 只获取一个则不填
		}
		info, err := client.GetProductInfo(ctx, req)
		t.Log(req, info)
		if err != nil {
			t.Errorf("%+v", errors.Fmt(err))
		}
	}
}
