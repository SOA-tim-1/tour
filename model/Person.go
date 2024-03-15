package model

type Person struct {
	ID        int64  `json:"id" gorm:"primaryKey:autoIncrement"`
	Firstname string `json:"firstname" gorm:"type:string"`
	Lastname  string `json:"lastname" gorm:"type:string"`
}

// func (person *Person) BeforeCreate(scope *gorm.DB) error {
// 	person.ID = uuid.New()
// 	return nil
// }
