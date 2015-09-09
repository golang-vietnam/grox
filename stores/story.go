package stores

import (
	"github.com/golang-vietnam/grox/api/rethink"
	"github.com/golang-vietnam/grox/domain"
)

const (
	StoryTable = "story"
)

type StoryStore struct {
	re *rethink.Instance
}

func NewStoryStore(re *rethink.Instance) *StoryStore {
	return &StoryStore{
		re: re,
	}
}

func (this *StoryStore) ListAll() (stories []*domain.Story) {
	this.re.All(this.re.Table(StoryTable), &stories)
	return
}

func (this *StoryStore) GetById(id string) (story domain.Story, err error) {
	err = this.re.One(this.re.Table(StoryTable).Get(id), &story)
	return
}

func (this *StoryStore) Create(story *domain.Story) (err error) {
	_, err = this.re.RunWrite(this.re.Table(StoryTable).Insert(story))
	return
}

func (this *StoryStore) UpdateAll(query map[string]interface{}, data map[string]interface{}) (int, error) {
	if res, err := this.re.RunWrite(this.re.Table(StoryTable).Filter(query).Update(data)); err != nil {
		return 0, err
	} else {
		return res.Updated, err
	}
}
