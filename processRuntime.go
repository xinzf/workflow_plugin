package workflow_plugin

import "time"

type ProcessInstanceModelInterface interface {
	GetID() string
	GetDefinitionID() string
	GetDataID() int
	GetState() ProcessInstanceState
	GetParentProcessInstanceID() string
	GetParentActivityInstanceID() string
	IsParentProcess() bool
	HasSubProcess() bool
	GetCreateTime() time.Time
	GetStartTime() time.Time
	GetEndTime() time.Time
	GetCostDuration() time.Duration
	GetExpireDuration() time.Duration
	GetCreatorUid() string
	GetCreatorDepId() int64
	GetRemindTimes() int
	IsException() bool
	IsEnd() bool
	IsStart() bool
	SetDefinitionID(id int)
	SetState(state ProcessInstanceState)
	SetParentProcessInstanceID(id int)
	SetParentActivityInstanceID(id int)
	ToStart()
	ToEnd()
}

type ProcessInstanceState string

type TaskInstanceModelInterface interface {
	GetID() string
	GetTitle() string
	GetProcessInstanceID() string
	GetProcessDefinitionID() int
	GetGroupID() string // 任务组ID，如果是并签或者或签，一个节点会产生多个属于不同人的任务，这些任务同属一个任务组
	GetFlowNodeID() string
	GetFlowType() FlowNodeType
	GetReadState() bool
	GetCreateTime() time.Time
	GetReadTime() time.Time
	GetBeginTime() time.Time
	GetEndTime() time.Time
	GetCreateUser() string
	GetOwner() string
	GetOwnerDepID() int64
	GetParentTaskInstanceID() string
	GetClaimType() ClaimType
	GetState() TaskInstanceState
	GetDelayTimes() time.Duration
	GetDelegateUser() string // 委托人ID，如果该任务是别人委托（转交）过来的，返回这个委托人ID
	GetDueTime() time.Time   // 任务处理期限时间
	GetRemindTimes() int
	GetApprovalType() ApprovalType
	IsNew() bool
	IsHistoryTask() bool
	IsException() bool
	Clone() TaskInstanceModelInterface
	ToHistory() TaskInstanceModelInterface
}

type TaskInstanceState string
