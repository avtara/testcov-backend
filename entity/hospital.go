package entity

//Hospital represents users table in database
type Hospital struct {
	ID        uint64  `gorm:"primary_key:auto_increment" json:"id"`
	Name      string  `gorm:"type:varchar(255)" json:"name"`
	Address   string  `gorm:"type:varchar(255)" json:"address"`
	Facilies  string  `gorm:"type:varchar(255)" json:"facilies"`
	Telp      string  `gorm:"type:varchar(14)" json:"telp"`
	Thumbnail string  `gorm:"type:varchar(255)" json:"thumbnail"`
	Longitude string  `gorm:"type:varchar(255)" json:"longitude"`
	Latitude  string  `gorm:"type:varchar(255)" json:"latitude"`
	Distance  float64 `json:"distance"`
}
