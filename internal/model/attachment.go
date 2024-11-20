package model

type Attachment struct {
	ID          int    `db:"id"`
	BusinessId  string `db:"business_id"`
	Title       string `db:"title"`
	Size        int    `db:"size"`
	Filename    string `db:"filename"`
	ContentType string `db:"content_type"`
	Url         string `db:"url"`
}
