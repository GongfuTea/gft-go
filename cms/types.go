package cms

type ActiveState string

const (
	ActiveState_Draft     ActiveState = "Draft"
	ActiveState_Published ActiveState = "Published"
	ActiveState_Archived  ActiveState = "Archived"
)

var ActiveStates []ActiveState = []ActiveState{ActiveState_Draft, ActiveState_Published, ActiveState_Archived}

type PostType string

const (
	PostType_Post PostType = "Post"
	PostType_Link PostType = "Link"
	PostType_File PostType = "File"
)

var PostTypes []PostType = []PostType{PostType_Post, PostType_Link, PostType_File}

type PostFormat string

const (
	PostFormat_Html     PostFormat = "Html"
	PostFormat_Markdown PostFormat = "Markdown"
)

var PostFormats []PostFormat = []PostFormat{PostFormat_Html, PostFormat_Markdown}

type MenuType string

const (
	MenuType_Node     MenuType = "Node"
	MenuType_Category MenuType = "Category"
	MenuType_Post     MenuType = "Post"
	MenuType_Link     MenuType = "Link"
	MenuType_File     MenuType = "File"
)

var MenuTypes []MenuType = []MenuType{MenuType_Node, MenuType_Category, MenuType_Post, MenuType_Link, MenuType_File}

type AccessLevel string

const (
	AccessLevel_Public    AccessLevel = "Public"
	AccessLevel_Protected AccessLevel = "Protected"
	AccessLevel_Private   AccessLevel = "Private"
	AccessLevel_None      AccessLevel = "None"
)

var AccessLevels []AccessLevel = []AccessLevel{AccessLevel_None, AccessLevel_Private, AccessLevel_Protected, AccessLevel_Public}
