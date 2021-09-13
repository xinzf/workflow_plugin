package workflow_plugin

type VariableInterface interface {
	GetName() string
	GetLabel() string
	GetValue() interface{}
	GetDataType() DataType
}

type VariablesInterface interface {
	Size() int
	Set(name, label string, val interface{}, dataType DataType)
	Get(name string) VariableInterface
	Remove(name string)
	Each(func(v VariableInterface))
	Merge(variables VariablesInterface)
	Clear()
}
