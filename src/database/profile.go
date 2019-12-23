package database

type Profile struct {
	ID                uint64 `db:"id" json:"id"`
	UUID              string `db:"uuid" json:"uuid,omitempty"`
	Experience        uint64 `db:"experience" json:"experience"`
	Level             uint64 `db:"level" json:"level"`
	TotalScore        uint64 `db:"total_score" json:"total_score"`
	PlayCount         uint64 `db:"play_count" json:"play_count"`
	Mastery           uint8  `db:"mastery" json:"mastery"`
	PerformanceRating uint64 `db:"performance_rating" json:"performance_rating"`
}

type Registration struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	InviteCode string `json:"invite_code,omitempty"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
