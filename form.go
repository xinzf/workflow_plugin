package workflow_plugin

import (
	"github.com/xinzf/dataframe/series"
)

type FormModelInterface interface {
	GetID() int
	GetTitle() string
	GetDraftVersion() int64
	GetVersion() int64
	SetDraftVersion(millisecond int64)
	SetTitle(title string)
	SetVersion(millisecond int64)
	SetLayout(layout FormLayoutInterface)
	GetLayout() FormLayoutInterface
}

type FormLayoutInterface interface {
	SetForm(form FormModelInterface) FormLayoutInterface
	GetWidgets() (FormWidgetsInterface, error)
}

type DefaultFormLayout struct {
}

func (this *DefaultFormLayout) SetForm(form FormModelInterface) FormLayoutInterface {
	panic("implement me")
}

func (this *DefaultFormLayout) GetWidgets() (FormWidgetsInterface, error) {
	panic("implement me")
}

type FormWidgetInterface interface {
	GetID() string
	GetLabel() string
	GetDataType() DataType
	GetSeriesType() series.Type
	IsSubform() bool
	IsArray() bool
	IsObject() bool
	IsSystem() bool
	GetSubForms() FormWidgetsInterface
	Clone() FormWidgetInterface
}

type FormWidgetsInterface interface {
	Size() int
	Add(wgt ...FormWidgetInterface)
	Get(idx int) (FormWidgetInterface, bool)
	Find(func(widget FormWidgetInterface) bool) (int, FormWidgetInterface)
	Each(func(idx int, widget FormWidgetInterface))
	Widgets() []FormWidgetInterface
}
