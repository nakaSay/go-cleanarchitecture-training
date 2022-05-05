package entities

type Post struct {
	ID int64 `json:"id" datastore:"-"`
	Text string `json:"text" datastore:"text"`
}

type Posts struct {
	Posts []Post `json:"posts"`
}