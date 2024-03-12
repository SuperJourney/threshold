package threshold

import (
	"context"
)

const (
	DAY   = "day"
	WEEK  = "week"
	MONTH = "month"
)

const (
	RULE_TYPE_COUNT              = "count"
	RULE_TYPE_NATURE_DAY_COUNT   = "count_day"
	RULE_TYPE_NATURE_WEEK_COUNT  = "count_week"
	RULE_TYPE_NATRUE_MONTH_COUNT = "count_month"
)

type Rule struct {
	Code      uint32    `json:"code,omitempty"`
	Type      string    `json:"type,omitempty"`      // 规则类型
	RuleExpr  string    `json:"rule_expr,omitempty"` // 规则表达式
	Threshold Threshold `json:"threshold,omitempty"` //设置指标
}

type Threshold struct {
	Time  int64 `json:"time,omitempty"`  // 当规则类型为count时，可以自定义过期时间，0为不过期
	Limit int32 `json:"limit,omitempty"` // 限制次数
}

type ThresholdRuleIFace interface {
	GetThresholdRuleInfo(ctx context.Context, ruleID uint64) ([]*Rule, error)
	// ruleID 为0新增
	SetThresholdRule(ctx context.Context, ruleID uint64, rule []*Rule) (uint64, error)
	// 规则校验 (req map[string]string req匹配是否缺失字段)
	Match(ctx context.Context, ruleID uint64, req map[string]string) (logID uint64, err error)
	// 撤销
	Rollback(ctx context.Context, logID uint64) error
	// 获取当前校验值
	GetThresholdRuleData(ctx context.Context, ruleID uint64, req map[string]string) (map[uint64]int64, error)
}

/**
 *  TODO, 剥离原仓库repo
 */
