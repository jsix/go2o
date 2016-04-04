/**
 * Copyright 2015 @ z3q.net.
 * name : person_finance
 * author : jarryliu
 * date : 2016-03-31 10:46
 * description :
 * history :
 */
package personfinance

import (
	"go2o/src/core/domain/interface/member"
	"go2o/src/core/infrastructure/domain"
)

var (
	RiseMinTransferInAmount  float32 = 100.00       //最低转入金额为100
	RiseMinTransferOutAmount float32 = 0.00         //最低转出金额
	RiseSettleTValue         int     = 2            //T+? 开使计算收益
	RiseNormalDayRatio       float32 = 0.0001369863 // 年化5%,按365天计算
	// 比例提供者,默认为:personfinance.RiseNormalDayRatio
	RiseDayRatioProvider func(personId int) float32 = func(personId int) float32 {
		return RiseNormalDayRatio
	}
)

type (
	// 在此聚合下, 会员抽象为Person, PersonId 对应 MemberId
	IPersonFinance interface {
		// 获取聚合根
		GetAggregateRootId() int
		// 获取账号
		GetMemberAccount() member.IAccount
		// 获取增利账户信息(类:余额宝)
		GetRiseInfo() IRiseInfo
		// 创建增利账户信息
		CreateRiseInfo() error
		// 同步到会员账户理财数据
		SyncToAccount() error
	}

	// 现金增利
	IRiseInfo interface {
		GetDomainId() int

		// 获取值
		Value() (RiseInfoValue, error)

		// 转入
		TransferIn(amount float32) error

		// 转出
		TransferOut(amount float32) error

		// 根据日志记录提交转入转出,如果已经确认操作,则返回错误
		// 通常是由系统计划任务来完成此操作,转入和转出必须经过提交!
		CommitTransfer(logId int) error

		// 结算收益(按天结息),settleUnix:结算日期的时间戳(不含时间),
		// dayRatio 为每天的收益比率
		RiseSettleByDay(settleUnix int64, dayRatio float32) error

		// 获取时间段内的增利信息
		GetRiseByTime(begin, end int64) []*RiseDayInfo

		// 保存
		Save() error
	}

	// 收益总记录
	RiseInfoValue struct {
		//Id  int `db:"id" pk:"yes" auto:"no"`
		PersonId    int     `db:"person_id" pk:"yes" auto:"no"` //人员编号
		Balance     float32 `db:"balance"`                      //本金及收益的余额
		Rise        float32 `db:"rise"`                         //当前的收益
		TransferIn  float32 `db:"transfer_in"`                  //今日转入
		TotalAmount float32 `db:"total_amount"`                 //总金额
		TotalRise   float32 `db:"total_rise"`                   //总收益
		SettledDate int64   `db:"settled_date"`                 //结算日期,用于筛选需要结算的数据
		UpdateTime  int64   `db:"update_time"`
	}

	// 收益每日结算数据
	RiseDayInfo struct {
		Id         int     `db:"id" pk:"yes" auto:"yes"`
		PersonId   int     `db:"person_id"`
		Date       string  `db:"date"`
		BaseAmount float32 `db:"base_amount"` //本金
		RiseAmount float32 `db:"rise_amount"` //增加金额
		UnixDate   int64   `db:"unix_date"`
		UpdateTime int64   `db:"update_time"`
	}

	// 收益日志
	RiseLog struct {
		Id         int     `db:"id" pk:"yes" auto:"yes"`
		PersonId   int     `db:"person_id"`   //会员编号
		Amount     float32 `db:"amount"`      //金额
		Type       int     `db:"type"`        //类型
		State      int     `db:"state"`       //状态
		UnixDate   int64   `db:"unix_date"`   // 日期
		LogTime    int64   `db:"log_time"`    //日志时间
		UpdateTime int64   `db:"update_time"` //更新时间
	}

	IPersonFinanceRepository interface {
		// 获取个人财富聚合根
		GetPersonFinance(personId int) IPersonFinance

		// 根据时间获取收益情况
		GetRiseByTime(personId int, begin, end int64) []*RiseDayInfo

		// 根据人员编号获取收益
		GetRiseValueByPersonId(id int) (v *RiseInfoValue, err error)

		// 保存收益信息
		SaveRiseInfo(*RiseInfoValue) (id int, err error)

		// 获取日志
		GetRiseLog(personId, logId int) *RiseLog

		// 保存日志
		SaveRiseLog(*RiseLog) (int, error)

		// 获取日志
		GetRiseLogs(personId int, date int64, riseType int) []*RiseLog

		// 保存每日收益
		SaveRiseDayInfo(*RiseDayInfo) (int, error)
	}
)

var (
	RiseStateOk      int = 1 //OK
	RiseStateDefault int = 0 //默认状态

	ErrIncorrectAmount *domain.DomainError = domain.NewDomainError(
		"err_balance_amount", "金额错误!")
	ErrIncorrectTransfer *domain.DomainError = domain.NewDomainError(
		"err_incorrent_transfer", "已确认或非转入(转出)操作!")
	ErrNoSuchRiseInfo *domain.DomainError = domain.NewDomainError(
		"err_no_such_rise_info", "未开通该功能!")
	ErrHasSettled *domain.DomainError = domain.NewDomainError(
		"err_has_settled", "已经结算!")
	ErrUnixDate *domain.DomainError = domain.NewDomainError(
		"err_unix_date", "错误的日期时间戳!")
	ErrRatio *domain.DomainError = domain.NewDomainError(
		"err_ratio", "利率不正确!")

	ErrLessThanMinTransferIn *domain.DomainError = domain.NewDomainError(
		"err_less_than_min_transfer_in", "转入金额最低%d!")

	ErrLessThanMinTransferOut *domain.DomainError = domain.NewDomainError(
		"err_less_than_min_transfer_out", "转出金额最低%d!")

	ErrOutOfBalance *domain.DomainError = domain.NewDomainError(
		"err_out_of_balance", "超出帐户最大金额!")
)

const (
	RiseTypeTransferIn  int = 1 + iota //转入
	RiseTypeTransferOut                //转出
	RiseTypeSettle                     //结算
	RiseTypeAdjust                     //人工调整
)