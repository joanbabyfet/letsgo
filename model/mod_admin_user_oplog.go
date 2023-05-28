package model

type AdminUserOplog struct {
	Uid        string `bson:"uid"`
	Useranme   string `bson:"username"`
	Session_id string `bson:"session_id"`
	Msg        string `bson:"msg"`
	Module_id  int    `bson:"module_id"`
	Op_time    int    `bson:"op_time"`
	Op_ip      string `bson:"op_ip"`
	Op_country string `bson:"op_country"`
	Op_url     string `bson:"op_url"`
}

// collection名称
var Col_admin_user_oplog string = "admin_users_oplog"
