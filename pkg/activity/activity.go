package activity

import (
	"github.com/google/uuid"
	"github.com/redbow26/en-verre-et-contre-toux-api/pkg/database"
	"gorm.io/gorm"
)

type Activity struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"required"`
	MaxPlayer uint      `json:"max_player" gorm:"required"`
}

func (activity *Activity) BeforeCreate(tx *gorm.DB) (err error) {
	activity.ID = uuid.New()
	return
}

func FindAll() (activities []*Activity) {
	database.DB.Find(&activities)
	return activities
}

func FindOne(id string) (activity *Activity) {
	database.DB.Where("id = ?", id).First(&activity)
	return activity
}

func Create(activity *Activity) *Activity {
	database.DB.Create(&activity)
	return activity
}

func (activity *Activity) Update(newActivity *Activity) {
	database.DB.Model(&activity).Updates(&newActivity)
}

func Delete(id string) {
	database.DB.Where("id = ?", id).Delete(&Activity{})
}
