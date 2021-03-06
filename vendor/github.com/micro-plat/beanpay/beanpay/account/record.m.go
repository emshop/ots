package account

import (
	"github.com/micro-plat/lib4go/types"
)

// AccountRecord 账户余额变动记录
type AccountRecord struct {
	AccountID   int     `json:"account_id" m2s:"account_id"`
	RecordID    int     `json:"record_id" m2s:"record_id"`
	TradeNo     string  `json:"trade_no" m2s:"trade_no"`
	EID         string  `json:"eid" m2s:"eid"`
	Group       string  `json:"groups" m2s:"groups"`
	AccountName string  `json:"account_name" m2s:"account_name"`
	ChangeType  int     `json:"change_type" m2s:"change_type"`
	TradeType   int     `json:"trade_type" m2s:"trade_type"`
	Amount      float64 `json:"amount" m2s:"amount"`
	Balance     float64 `json:"balance" m2s:"balance"`
	Memo        string  `json:"memo" m2s:"memo"`
	CreateTime  string  `json:"create_time" m2s:"create_time"`
}

// RecordResult 记录结果
type RecordResult struct {
	*AccountRecord
	code int
}

// NewRecordResult map转RecordResult
func NewRecordResult(code int, m types.XMap) *RecordResult {
	var account AccountRecord
	m.ToAnyStruct(&account)
	return &RecordResult{AccountRecord: &account, code: code}
}

// GetResult 获取账户记录
func (r *RecordResult) GetResult() interface{} {
	return r.AccountRecord
}

// GetCode 获取结果码
func (r *RecordResult) GetCode() int {
	return r.code
}

// RecordResults 查询账户记录
type RecordResults struct {
	List  []*AccountRecord `json:"data,omitempty"`
	code  int
	Count int `json:"count" m2s:"count"`
}

// NewRecordResults maps转RecordResults
func NewRecordResults(code, count int, m types.XMaps) *RecordResults {
	var accounts []*AccountRecord
	m.ToAnyStructs(&accounts)
	return &RecordResults{List: accounts, code: code, Count: count}
}

// GetResult 获取账户记录
func (r *RecordResults) GetResult() interface{} {
	return r.List
}

// GetCode 获取状态码
func (r *RecordResults) GetCode() int {
	return r.code
}
