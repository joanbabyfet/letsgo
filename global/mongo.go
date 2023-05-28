package global

import (
	"context"
	"fmt"
	"sync"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongo_connect struct {
	client *mongo.Client
}

var mongo_once = sync.Once{}
var _mongo_connect *mongo_connect

func connectMongo() {
	host := viper.GetString("mongo.host")
	port := viper.GetString("mongo.port")
	username := viper.GetString("mongo.username")
	password := viper.GetString("mongo.password")
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s",
		username,
		password,
		host,
		port,
	)
	opts := options.Client().ApplyURI(uri)

	//连接到mongo
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		fmt.Println(err)
	}
	//检查连接
	err = client.Ping(context.Background(), nil)
	if err != nil {
		fmt.Println(err)
	}

	_mongo_connect = &mongo_connect{
		client: client,
	}
}

// 使用单例模式, 提供一个外部访问方法 类似 GetInstance()
func MongoClient() *mongo.Client {
	if _mongo_connect == nil {
		mongo_once.Do(func() {
			connectMongo()
		})
	}
	return _mongo_connect.client
}
