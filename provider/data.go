package provider

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var Conn = ConnectRedis()

/**测试服务连接
 */
func ConnectRedis() redis.Conn {
	conn, err := redis.Dial("tcp", "192.168.238.131:6379")
	if err != nil {
		fmt.Println("连接失败", err)
		return nil
	}

	return conn
}

func CreateComplexData(count int) []TestStruct{
	data := make([]TestStruct,0,count)
	for i := 0;i<count;i++{
		data = append(data,*CreateTestData(i))
	}
	return data
}


func CreateTestData(id int) *TestStruct {
	return &TestStruct{
		Id:    id,
		Name:  "测试姓名",
		Sex:   "男",
		Desc:  "描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述",
		Desc1: "描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述",
		Desc2: "描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述",
		Desc3: "描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述",
		Desc4: "描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述",
		Desc5: "描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述",
		Desc6: "描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述",
		Desc7: "描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述",
		Desc8: "描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述",
	}
}

type TestStruct struct {
	Id    int    `redis:"id" json:"id"`
	Name  string `redis:"name" json:"name"`
	Sex   string `redis:"sex" json:"sex"`
	Desc  string `redis:"desc" json:"desc"`
	Desc1 string `redis:"desc1" json:"desc1"`
	Desc2 string `redis:"desc2" json:"desc2"`
	Desc3 string `redis:"desc3" json:"desc3"`
	Desc4 string `redis:"desc4" json:"desc4"`
	Desc5 string `redis:"desc5" json:"desc5"`
	Desc6 string `redis:"desc6" json:"desc6"`
	Desc7 string `redis:"desc7" json:"desc7"`
	Desc8 string `redis:"desc8" json:"desc8"`
}
