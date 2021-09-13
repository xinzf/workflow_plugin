package workflow_plugin

import (
	"time"
)

type UserModelInterface interface {
	GetID() string
	GetName() string
	GetMobile() string
	GetDDUid() string
	GetDDUnionID() string
	GetWechatID() string
	GetWechatUnionID() string
	GetDepIds() []int64
	GetRoleIds() []int64
	GetDepartments() []DepartmentModelInterface
	GetRoles() []RoleModelInterface
	IsAdmin() bool
	IsLeader() bool
	IsBoss() bool
	IsActivity() bool            // 账号是否正常
	GetInductionTime() time.Time // 入职日期
	GetDepartureTime() time.Time // 离职日期
}

type UserModelsInterface interface {
	Find(func(user UserModelInterface) bool) (UserModelInterface, int)
	Each(func(idx int, user UserModelInterface))
	Add(user UserModelInterface)
	Remove(uid string)
	Get(index int) UserModelInterface
	Uids() []string
	ToMap() map[string]UserModelInterface
	Size() int
}

//
//type UserQuery interface {
//    AfterFind(func(model UserModel, db *gorm.DB) error) UserQuery
//    Unscoped() UserQuery
//    WithUserID(id ...string) UserQuery
//    WithDDUnionID(id string) UserQuery
//    WithWechatOpenID(id string) UserQuery
//    WithWechatUnionID(id string) UserQuery
//    WithMobile(mobile string) UserQuery
//    WithName(name string) UserQuery
//    WithDepID(id ...int64) UserQuery
//    WithRoleID(id ...int64) UserQuery
//    WithTel(tel string) UserQuery
//    Activity() UserQuery
//    NotActivity() UserQuery
//    Admin() UserQuery
//    NotAdmin() UserQuery
//    Boss() UserQuery
//    NotBoss() UserQuery
//    Leader() UserQuery
//    NotLeader() UserQuery
//    LoadDepartments() UserQuery
//    LoadRoles() UserQuery
//    LoadLeaders(depId ...int64) UserQuery
//    First() (UserModel, bool, error)
//    List() (UserModels, error)
//    ListPage(pageNo, pageSize int) (UserModels, int, error)
//}

type DepartmentModelInterface interface {
	GetID() int64
	GetName() string
	GetParentDepID() int64
	GetUsers() UserModelsInterface
}

type DepartmentTreeModelInterface interface {
	DepartmentModelInterface
	GetSubDepartments() []DepartmentTreeModelInterface
}

//
//type DepartmentQuery interface {
//    AfterFind(func(model DepartmentModel, db *gorm.DB) error) DepartmentQuery
//    WithDepID(id ...int64) DepartmentQuery
//    WithName(name string) DepartmentQuery
//    WithParentDepID(id int64) DepartmentQuery
//    LoadUsers() DepartmentQuery
//    First() (DepartmentModel, bool, error)
//    List() ([]DepartmentModel, error)
//    ListPage(pageNo, pageSize int) ([]DepartmentModel, int, error)
//    Tree() (DepartmentTreeModel, error)
//}

type RoleModelInterface interface {
	GetID() int64
	GetName() string
	GetUsers() UserModelsInterface
}

//
//type RoleQuery interface {
//    AfterFind(func(model RoleModel, db *gorm.DB) error) RoleQuery
//    WithRoleID(id ...int64) RoleQuery
//    WithName(name string) RoleQuery
//    LoadUsers() RoleQuery
//    First() (RoleModel, bool, error)
//    List() ([]RoleModel, error)
//    ListPage(pageNo, pageSize int) ([]RoleModel, int, error)
//}
//

type SessionInterface interface {
	IsFromPC() bool
	IsFromMobile() bool
	IsLoginByMobile() bool
	IsLoginByDingDing() bool
	IsLoginByWechat() bool
	GetUser() UserModelInterface
	GetDepartment() DepartmentModelInterface
	GetRoles() []RoleModelInterface
}
