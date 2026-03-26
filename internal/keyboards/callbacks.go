package keyboards

import "github.com/erfjab/egobot/tools"

type UserSections string

const (
	UserSectionInfo   = "Info"
	UserSectionMenu   = "Menu"
	UserSectionUpdate = "Update"
	UserSectionCreate = "Create"
)

type ItemCreateCB struct {
	tools.CallbackData `prefix:"create"`
}

type ItemInfoCB struct {
	tools.CallbackData `prefix:"items"`
	Id				 int	
}

type ItemUpdateCB struct {
	tools.CallbackData `prefix:"update"`
	Id				 int
	Title			 bool
	Content			 bool
	Remove			 bool
}

type ItemRemoveCB struct {
	tools.CallbackData `prefix:"remove"`
	Id				 int
	Confirm			 bool
}
