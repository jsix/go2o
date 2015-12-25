/**
 * Copyright 2014 @ z3q.net.
 * name :
 * author : jarryliu
 * date : 2014-01-08 21:35
 * description :
 * history :
 */

package daemon

import (
	"github.com/garyburd/redigo/redis"
	"go2o/src/core"
	"go2o/src/core/service/dps"
	"go2o/src/core/variable"
	"strconv"
	"strings"
)

// 监视新订单
func superviseOrder(ss []Service) {
	con := (*core.MainApp)(appCtx).Redis().Get()
	for {
		id, err := redis.Int(con.Do("RPOP", variable.KvOrderBusinessQueue))
		if err != nil {
			appCtx.Log().Println("[ DAEMON][ QUEUE][ NEW-ORDER] -", err.Error())
			continue
		}
		order := dps.ShoppingService.GetOrderById(id)
		for _, v := range ss {
			if !v.OrderObs(order) {
				break
			}
		}
	}
}

// 监视新订单
func superviseMemberUpdate(ss []Service) {
	con := (*core.MainApp)(appCtx).Redis().Get()
	for {
		s, err := redis.String(con.Do("RPOP", variable.KvMemberUpdateQueue))
		if err != nil {
			appCtx.Log().Println("[ DAEMON][ QUEUE][ MEMBER] -", err.Error())
			continue
		}

		arr := strings.Split(s, "-")
		id, err := strconv.Atoi(arr[0])
		if err == nil {
			m := dps.MemberService.GetMember(id)
			for _, v := range ss {
				if !v.MemberObs(m, arr[1] == "create") {
					break
				}
			}
		}
	}
}

//
//func confirmNewOrder(app gof.App, dfs []Func) {
//
//	if i, _ := appCtx.Storage().GetInt(variable.KvHaveNewCreatedOrder); i == enum.TRUE {
//		appCtx.Log().Printf("[ DAEMON][ ORDER][ CONFIRM] - begin invoke confirm handler.")
//		if dfs == nil || len(dfs) == 0 {
//			confirmOrderQueue(app)
//		} else {
//			for _, v := range dfs {
//				v(app)
//			}
//		}
//		appCtx.Storage().Set(variable.KvHaveNewCreatedOrder, enum.FALSE)
//	}
//}
//
//func completedOrderObs(app gof.App, dfs []Func) {
//	if len(dfs) < 0 {
//		return
//	}
//	if i, _ := appCtx.Storage().GetInt(variable.KvHaveNewCompletedOrder); i == enum.TRUE {
//		appCtx.Log().Printf("[ DAEMON][ ORDER][ FINISHED] - begin invoke finish handler.\n")
//		for _, v := range dfs {
//			v(app)
//		}
//		appCtx.Storage().Set(variable.KvHaveNewCompletedOrder, enum.FALSE)
//	}
//}
//
//type orderInfo struct {
//	PartnerId int
//	OrderNo   string
//}
//
//func confirmOrderQueue(app gof.App) {
//	var list []*orderInfo = []*orderInfo{}
//	appCtx.Db().GetOrm().SelectByQuery(&list, fmt.Sprintf("SELECT partner_id,order_no FROM pt_order WHERE status=%d",
//		enum.ORDER_WAIT_CONFIRM))
//	for _, v := range list {
//		err := dps.ShoppingService.ConfirmOrder(v.PartnerId, v.OrderNo)
//		if err != nil {
//			appCtx.Log().Printf("[ DAEMON][ ORDER][ ERROR] - %s\n", err.Error())
//		}
//	}
//}