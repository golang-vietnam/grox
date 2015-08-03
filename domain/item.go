package domain

type ItemId uint

type ItemType string

const (
	ItemType_Story   = ItemType("story")
	ItemType_Comment = ItemType("comment")
)

type Item struct {
	TimeStamp

	By      UserId `json:"by" gorethink:"by"`
	Parent  ItemId `json:"parent,omitempty" gorethink:"parent"`
	Id      ItemId `json:"id" gorethink:"id"`
	Deleted bool   `json:"-" gorethink:"deleted"`

	Title string   `json:"title" gorethink:"title"`
	Score int      `json:"score" gorethink:"score"`
	Type  ItemType `json:"type" gorethink:"type"`
	Text  string   `json:"text" gorethink:"text"`

	Children []ItemId `json:"children" gorethink:"children,omitempty"`
}
