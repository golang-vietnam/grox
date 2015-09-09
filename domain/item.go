package domain

type ItemType string

const (
	ItemType_Story   = ItemType("story")
	ItemType_Comment = ItemType("comment")
)

type Item struct {
	Common
	By     string `json:"by" gorethink:"by"`
	Parent string `json:"parent,omitempty" gorethink:"parent"`

	Title string   `json:"title" gorethink:"title"`
	Score int      `json:"score" gorethink:"score"`
	Type  ItemType `json:"type" gorethink:"type"`
	Text  string   `json:"text" gorethink:"text"`

	Children []string `json:"children" gorethink:"children,omitempty"`
}
