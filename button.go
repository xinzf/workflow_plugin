package workflow_plugin

// ButtonType 按钮类型
type ButtonType string

const (
	SUBMIT_BUTTON   ButtonType = "submit"
	STAGING_BUTTON  ButtonType = "staging"
	ROLLBACK_BUTTON ButtonType = "rollback"
	TRANSFER_BUTTON ButtonType = "transfer"
	KILL_BUTTON     ButtonType = "kill"
	PAUSE_BUTTON    ButtonType = "pause"
	AUTO_BUTTON     ButtonType = "auto"
	CANCEL_BUTTON   ButtonType = "cancel"
	RELAUNCH_BUTTON ButtonType = "relaunch"
)

func (static ButtonType) String() string {
	return string(static)
}

func (static ButtonType) Label() string {
	mp := map[ButtonType]string{
		SUBMIT_BUTTON:   "提交",
		STAGING_BUTTON:  "暂存",
		ROLLBACK_BUTTON: "回退",
		TRANSFER_BUTTON: "转交",
		KILL_BUTTON:     "结束流程",
		PAUSE_BUTTON:    "暂停流程",
		AUTO_BUTTON:     "自动触发",
		CANCEL_BUTTON:   "撤销",
		RELAUNCH_BUTTON: "重新发起",
	}

	return mp[static]
}

type ButtonInterface interface {
	GetID() string
	GetLabel() string
	GetButtonType() ButtonType
	IsEnabled() bool
}
