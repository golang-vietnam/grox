package domain

type StoryType string

const (
	StoryTypeLink = StoryType("link")
	StoryTypeAsk  = StoryType("ask")
)

type Story struct {
	Common
	Url    string    `json:"url" gorethink:"url"`
	Slug   string    `json:"slug" gorethink:"slug"`
	UserId string    `json:"user_id" gorethink:"user_id"`
	Title  string    `json:"title" gorethink:"title"`
	Score  int       `json:"score" gorethink:"score"`
	Type   StoryType `json:"type" gorethink:"type"`
	Text   string    `json:"text" gorethink:"text"`
	IsDead bool      `json:"is_dead" gorethink:"is_dead"`
}
