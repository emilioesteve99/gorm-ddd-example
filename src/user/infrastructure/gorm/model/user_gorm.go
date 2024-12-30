package userGormModels

type UserGorm struct {
	ID       string `gorm:"column:id;primaryKey;type:uuid;"`
	Email    string `gorm:"column:email;uniqueIndex;type:varchar;not null;"`
	Name     string `gorm:"column:name;type:varchar;not null;"`
	Password string `gorm:"column:password;type:varchar;not null;"`
}

func (UserGorm) TableName() string {
	return "user"
}
