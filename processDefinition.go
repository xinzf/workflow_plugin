package workflow_plugin

import (
	"time"
)

type ProcessDefinitionModelInterface interface {
	GetID() int
	GetTitle() string
	GetVersion() int64
	GetTitleExpression() string
	GetFlowNodes() FlowNodesInterface
	GetMountForm() int
	IsMobileEnabled() bool
	GetVariables() VariablesInterface
	IsAbandon() bool
	IsRevoke() bool
	IsUrged() bool
	GetCompletionDuration() time.Duration
	GetWarnDuration() time.Duration
}
