package entity

//Schedule represents users table in database
type Schedule struct {
	ID         uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Day        string `gorm:"type:varchar(20)" json:"day"`
	TimeStart  string `gorm:"type:varchar(20)" json:"time_start"`
	TimeEnd    string `gorm:"type:varchar(20)" json:"time_end"`
	HospitalID uint64 `gorm:"not null" json:"-"`
}
