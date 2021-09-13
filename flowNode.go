package workflow_plugin

import "fmt"

// FlowNodeType 节点类型
type FlowNodeType string

const (
	// 空白节点，用于默认初始值
	//EMPTY_NODE FlowNodeType = "empty"
	// StartNode 申请节点
	START_NODE FlowNodeType = "startNode"
	// AuditNode 审核节点
	AUDIT_NODE FlowNodeType = "auditNode"
	// CopyNode 抄送节点
	COPY_NODE FlowNodeType = "copyNode"
	// ConditionNoe 条件节点
	CONDITION_NODE FlowNodeType = "conditionNode"
	// BranchNode 分支节点
	BRANCH_NODE FlowNodeType = "branchNode"
	// CreatorNode Robot.写入数据节点
	CREATE_NODE FlowNodeType = "creatorNode"
	// UpdaterNode Robot.更新数据节点
	UPDATE_NODE FlowNodeType = "updaterNode"
	// CallNode 调用子流程节点
	CALL_NODE FlowNodeType = "callNode"
)

func (static FlowNodeType) String() string {
	return string(static)
}

func (static FlowNodeType) IsUserNode() bool {
	if static == START_NODE || static == AUDIT_NODE {
		return true
	}
	return false
}

func (static FlowNodeType) Text() string {
	m := map[FlowNodeType]string{
		START_NODE:     "开始节点",
		AUDIT_NODE:     "审批节点",
		COPY_NODE:      "抄送节点",
		CONDITION_NODE: "条件节点",
		BRANCH_NODE:    "分支节点",
		CREATE_NODE:    "添加数据",
		UPDATE_NODE:    "更新数据",
		CALL_NODE:      "调用子流程",
	}
	return m[static]
}

// ApprovalType 审批类型
type ApprovalType string

const (
	AND_AUDIT ApprovalType = "AND"
	OR_AUDIT  ApprovalType = "OR"
)

func (static ApprovalType) String() string {
	return string(static)
}

func (static ApprovalType) Label() string {
	mp := map[ApprovalType]string{
		AND_AUDIT: "会签",
		OR_AUDIT:  "或签",
	}
	return mp[static]
}

// ClaimType 任务分配模式
type ClaimType int

const (
	SEIZE ClaimType = iota
	RANDOM
	MIN_TASK
)

func (static ClaimType) Int() int {
	return int(static)
}

func (static ClaimType) Label() string {
	mp := map[ClaimType]string{
		RANDOM:   "随机模式",
		SEIZE:    "抢占模式",
		MIN_TASK: "任务量最少",
	}

	return mp[static]
}

type FlowNodeInterface interface {
	GetID() string
	GetTitle() string
	GetFormAuth() map[string]interface{} // todo
	GetNodeType() FlowNodeType
	GetOpinion() int
	GetButtons() []ButtonInterface
	GetBranches() []FlowNodesInterface
	GetConditionExpression() string
	GetApprovalType() ApprovalType
	GetClaimType() ClaimType
	GetNext() (FlowNodeInterface, bool)
	GetPrev() (FlowNodeInterface, bool)
	GetParent(FlowNodeInterface, bool)
}

type FlowNodesInterface interface {
	fmt.Stringer
	Size() int
	Get(index int) FlowNodeInterface
	Find(func(node FlowNodeInterface) bool) (int, FlowNodeInterface)
	Each(func(idx int, node FlowNodeInterface))
}
