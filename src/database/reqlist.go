package database

type ReqList struct {
	Email string `db:"email" json:"email,omitempty"`
}
