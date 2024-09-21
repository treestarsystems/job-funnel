package utils

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

// MongoDB variables
var ClientMongo *mongo.Client
var CollectionMongo *mongo.Collection
var CtxMongo = context.TODO()

// SQLite variables
var DB *gorm.DB
var TableName *string
var SqliteDbName *string
