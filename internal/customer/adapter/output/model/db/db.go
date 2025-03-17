package db

type CustomerDB struct {
	ID 				string `gorm:"primaryKey"`
	Name 			string
	Username  string
	Email 		string
	Password  string
}