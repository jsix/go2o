/**
 * Copyright 2015 @ z3q.net.
 * name : personfinance_service
 * author : jarryliu
 * date : 2016-04-01 09:41
 * description :
 * history :
 */
package dps

import (
	"errors"
	"fmt"
	"go2o/src/core/domain/interface/member"
	"go2o/src/core/domain/interface/personfinance"
	"go2o/src/core/infrastructure/domain"
)

const (
	TransferFromBalance = 1 + iota //通过余额转入
	TransferFromPresent            //通过创业金转入
)

type personFinanceService struct {
	_rep    personfinance.IPersonFinanceRepository
	_accRep member.IMemberRep
}

func NewPersonFinanceService(rep personfinance.IPersonFinanceRepository,
	accRep member.IMemberRep) *personFinanceService {
	return &personFinanceService{
		_rep:    rep,
		_accRep: accRep,
	}
}

func (this *personFinanceService) GetRiseInfo(personId int) (personfinance.RiseInfoValue, error) {
	pf := this._rep.GetPersonFinance(personId)
	return pf.GetRiseInfo().Value()
}

// 开通增利服务
func (this *personFinanceService) OpenRiseService(personId int) error {
	pf := this._rep.GetPersonFinance(personId)
	return pf.CreateRiseInfo()
}

// 提交转入/转出日志
func (this *personFinanceService) CommitTransfer(personId, logId int) error {
	pf := this._rep.GetPersonFinance(personId)
	rs := pf.GetRiseInfo()
	if rs == nil {
		return personfinance.ErrNoSuchRiseInfo
	}
	return rs.CommitTransfer(logId)
}

// 转入
func (this *personFinanceService) RiseTransferIn(personId int, transferFrom int, amount float32) (err error) {
	pf := this._rep.GetPersonFinance(personId)
	r := pf.GetRiseInfo()
	if amount < personfinance.RiseMinTransferInAmount {
		//金额不足最低转入金额
		return errors.New(fmt.Sprintf(personfinance.ErrLessThanMinTransferIn.Error(),
			personfinance.RiseMinTransferInAmount))
	}
	m := this._accRep.GetMember(personId)
	if m == nil {
		return member.ErrNoSuchMember
	}
	acc := m.GetAccount()
	if transferFrom == TransferFromBalance { //从余额转入
		if err = acc.DiscountBalance("理财转入", domain.NewTradeNo(10000), amount); err != nil {
			return err
		}
		if err = r.TransferIn(amount); err != nil { //转入
			return err
		}
		return pf.SyncToAccount() //同步到会员账户
	}

	if transferFrom == TransferFromPresent { //从奖金转入
		if err := acc.DiscountPresent("理财转入", domain.NewTradeNo(10000), amount, true); err != nil {
			return err
		}
		if err = r.TransferIn(amount); err != nil { //转入
			return err
		}
		return pf.SyncToAccount() //同步到会员账户
	}

	return errors.New("未知的转入方式")
}

// 转出
func (this *personFinanceService) RiseTransferOut(personId int, amount float32) (err error) {
	pf := this._rep.GetPersonFinance(personId)
	r := pf.GetRiseInfo()
	if err = r.TransferOut(amount); err != nil {
		return err
	}
	return pf.SyncToAccount() //同步到会员账户
}

// 结算收益(按日期每天结息)
func (this *personFinanceService) RiseSettleByDay(personId int, settleUnix int64, dayRatio float32) (err error) {
	pf := this._rep.GetPersonFinance(personId)
	r := pf.GetRiseInfo()
	if err = r.RiseSettleByDay(settleUnix, dayRatio); err != nil {
		return err
	}
	return pf.SyncToAccount() //同步到会员账户
}