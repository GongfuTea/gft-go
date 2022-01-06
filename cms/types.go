package cms

type ContentState string

const (
	ContentState_Draft     ContentState = "Draft"
	ContentState_Published ContentState = "Published"
	ContentState_Archived  ContentState = "Archived"
)

var ContentStates []ContentState = []ContentState{ContentState_Draft, ContentState_Published, ContentState_Archived}
