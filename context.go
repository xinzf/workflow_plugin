package workflow_plugin

import "gorm.io/gorm"

type ExecutionContextInterface interface {
	AddAlertMessage(msg string, code int)
	GetAlertMessages()
	GetData() FormDataModelInterface
	GetForm() FormModelInterface
	GetProcessDefinition() ProcessDefinitionModelInterface
	GetProcessInstance() ProcessInstanceModelInterface
	GetActivityInstance() TaskInstanceModelInterface
	GetSession() SessionInterface
	GetTX() *gorm.DB
	GetVariables() VariablesInterface
}

type ExpressionContextInterface interface {
	GetData() FormDataModelInterface
	GetProcessDefinition() ProcessDefinitionModelInterface
	GetProcessInstance() ProcessInstanceModelInterface
	GetActivityInstance() TaskInstanceModelInterface
	GetVariables() VariablesInterface
	GetSession() SessionInterface
}
