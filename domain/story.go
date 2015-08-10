package domain

type StoryType string

const (
	StoryTypeLink = StoryType("link")
	StoryTypeAsk  = StoryType("ask")
)

type Story struct {
	TimeStamp

	Id        Id        `json:"id" gorethink:"id"`
	Url       string    `json:"url" gorethink:"url"`
	Slug      string    `json:"slug" gorethink:"slug"`
	UserId    Id        `json:"user_id" gorethink:"user_id"`
	Title     string    `json:"title" gorethink:"title"`
	Score     int       `json:"score" gorethink:"score"`
	Type      StoryType `json:"type" gorethink:"type"`
	Text      string    `json:"text" gorethink:"text"`
	IsDeleted bool      `json:"is_deleted" gorethink:"is_deleted"`
	IsDead    bool      `json:"is_dead" gorethink:"is_dead"`
}
