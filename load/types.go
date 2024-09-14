package load

import (
	"job-funnel/transform"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
)

// Gorm uses a different function to create the ID
type LoadDbInsertGorm struct {
	transform.JobPost
	ID        uint           `gorm:"primarykey"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// MongoDB uses a different function to create the ID
type LoadDbInsertMongoDb struct {
	transform.JobPost
	ID primitive.ObjectID `bson:"_id"`
}
