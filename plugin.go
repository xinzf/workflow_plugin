package workflow_plugin

import "fmt"

type PluginRegister func(cfg interface{}) PluginInterface

type PluginInterface interface {
	Name() string
	Label() string
	Describe() string
}

type ExecutePluginInterface interface {
	PluginInterface
	Execute(ctx ExecutionContextInterface) error
}

type InterruptPluginInterface interface {
	PluginInterface
	Execute(ctx ExecutionContextInterface) (isContinue bool, err error)
}

type ExpressionResult interface {
	fmt.Stringer
	Int() (int, error)
	Float() (float64, error)
	Bool() (bool, error)
	Slice() ([]interface{}, error)
	Bind(obj interface{}) error
}

type ExpressionPluginInterface interface {
	PluginInterface
	Execution(ctx ExpressionContextInterface, expression string) ExpressionResult
}

//type ListenerRegisterModelInterface interface {
//    GetName() string
//    GetConfig() interface{}
//}
//
//type ListenerRegisterQueryInterface interface {
//    WithID(id ...int) ListenerRegisterQueryInterface
//    WithName(name string) ListenerRegisterQueryInterface
//    WithLabel(label string) ListenerRegisterQueryInterface
//    WithType(t ListenerType) ListenerRegisterQueryInterface
//    WithTarget(targetType ListenerType, targetId int) ListenerRegisterQueryInterface
//    First() (ListenerRegisterModelInterface, bool, error)
//    List() ([]ListenerRegisterModelInterface, error)
//}
