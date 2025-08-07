package blog

type Blog struct {
	Id        uint `gorm:"primaryKey;default:auto_random()"`
	Title string
	Text string
	AuthorId uint
}