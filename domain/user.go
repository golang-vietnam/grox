package domain

type UserId string

type User struct {
	TimeStamp

	Id    UserId `json:"id" gorethink:"id"`
	About string `json:"about" gorethink:"about"`

	Submitted []Id `json:"submitted" gorethink:"-"`
}
