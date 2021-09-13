package workflow_plugin

import (
	"fmt"
	"time"
)

type FormDataModelInterface interface {
	GetID() int
	GetSN() string
	GetTitle() string
	GetFormID() int
	GetFormVersion() int64
	GetProcessInstanceID() string
	GetProcessDefinitionID() int
	GetProcessDefinitionVersion() int64
	GetCreatorUid() string
	GetCreatorDepId() int64
	GetAllParticipantUids() []string
	GetFormValues() ValueInterface
	IsEnd() bool
	IsNew() bool
	GetCreateTime() time.Time
	SetTitle(title string)
	SetSN(sn string)
	SetFormID(id int)
	SetFormVersion(version int64)
	SetProcessInstanceID(id string)
	SetProcessDefinitionID(id int)
	SetProcessDefinitionVersion(version int64)
	SetCreatorUid(uid string)
	SetCreatorDepId(depId int64)
	SetParticipantUid(uid string)
	SetEnd(end bool)
	SetFormValues(value ValueInterface)
	GetSubFormData(paths ...string) FormDataModelsInterface
	Clone() FormDataModelInterface
	ToJson() ([]byte, error)
	ToValue() ValueInterface
}

type FormDataModelsInterface interface {
	Size() int
	Find(func(data FormDataModelInterface) bool) (int, FormDataModelInterface)
	Each(func(idx int, data FormDataModelInterface))
	Add(data ...FormDataModelInterface)
	Get(idx int) FormDataModelInterface
	Remove(idx int)
	List() []FormDataModelInterface
}

type ValueGetterInterface interface {
	fmt.Stringer
	ToJson() ([]byte, error)
	ToMap() map[string]interface{}
	Clone() ValueGetterInterface
	Keys() []string
	Values() []interface{}
	Get(key string) interface{}
	Size() int
}

type ValueSetterInterface interface {
	Merge(getter ValueGetterInterface)
	Set(key string, val interface{})
	Remove(key string)
}

type ValueInterface interface {
	ValueSetterInterface
	ValueGetterInterface
}
