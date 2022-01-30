package tcp

import (
	"encoding/json"
	"github.com/kanyuanzhi/tialloy/tiface"
	"github.com/kanyuanzhi/tialloy/tilog"
	"tialloy-demo/face"
	"tialloy-demo/global"
	model "tialloy-demo/model/terminal"
)

type TerminalBasicRouter struct {
	*BaseTcpRouter
}

func NewTerminalBasicRouter(trafficHub face.ITrafficHub) tiface.IRouter {
	return &TerminalBasicRouter{
		NewBaseTcpRouter(trafficHub),
	}
}

func (r *TerminalBasicRouter) PreHandle(request tiface.IRequest) {
	r.TrafficHub.SetTcpConnList(request)
}

func (r *TerminalBasicRouter) Handle(request tiface.IRequest) {
	var terminalBasicPack = model.TerminalBasicPack{}
	if err := json.Unmarshal(request.GetData(), &terminalBasicPack); err != nil {
		tilog.Log.Error(err)
		return
	}
	terminalBasic := terminalBasicPack.Data
	terminal := model.Terminal{TerminalBasic: terminalBasic}
	key := terminalBasicPack.Key
	var total int64
	global.DB.Model(&terminal).Where("net_mac = ?", key).Count(&total)
	// 查询数据库中是否存在key的记录
	if total == 0 {
		// 没有则添加
		result := global.DB.Create(&terminal)
		if result.Error != nil {
			tilog.Log.Error(result.Error)
			return
		}
		tilog.Log.Infof("terminal mac=%s has been created", key)
	} else {
		global.DB.Model(&terminal).Where("net_mac = ?", key).Updates(&terminal)
		tilog.Log.Infof("terminal mac=%s has been updated", key)
	}
}
