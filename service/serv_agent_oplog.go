package service

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joanbabyfet/letsgo/global"
	"github.com/joanbabyfet/letsgo/model"
	"github.com/spf13/viper"
)

// 写入代理用户操作日志
func AddAgentOplog(c *gin.Context, msg string, module_id int) int {
	db_name := viper.GetString("mongo.database")
	collection := global.MongoClient().Database(db_name).Collection(model.Col_agent_oplog)
	status := 1

	//组装数据

	//批量写入mongo
	_, err := collection.InsertMany(context.Background(), []interface{}{
		model.AgentOplog{
			Uid:        "9eff3e40b42fa665b18437d2e91a7b3c",
			Useranme:   "agent1",
			Session_id: "",
			Msg:        msg,
			Module_id:  module_id,
			Op_ip:      c.ClientIP(),
			Op_country: "",
			Op_time:    int(time.Now().Unix()),
		},
	})
	if err != nil {
		fmt.Println(err)
		status = -1
	}
	return status
}
