package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var testStruct = CreateTestData(1111)
var textComplexStruct = CreateComplexData(3)

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





type TestStruct struct{
	Id int `redis:"id" json:"id"`
	Name string `redis:"name" json:"name"`
	Sex string `redis:"sex" json:"sex"`
	Desc string `redis:"desc" json:"desc"`
	Desc1 string `redis:"desc1" json:"desc1"`
	Desc2 string `redis:"desc2" json:"desc2"`
	Desc3 string `redis:"desc3" json:"desc3"`
	Desc4 string `redis:"desc4" json:"desc4"`
	Desc5 string `redis:"desc5" json:"desc5"`
	Desc6 string `redis:"desc6" json:"desc6"`
	Desc7 string `redis:"desc7" json:"desc7"`
	Desc8 string `redis:"desc8" json:"desc8"`

}


func main() {
	conn := ConnectRedis()
	//DoHashStore(conn)
	//DoJsonEncodingStore(conn)
	//DoGobEncodingStore(conn)
	DoComplexHashStore(conn)

}

func DoComplexHashStore(conn redis.Conn)  {

	var args = redis.Args{"complex1"}.AddFlat(textComplexStruct)
	//使用 以hash类型保存
	_,err := conn.Do("hmset",args...)

	if err != nil{
		fmt.Println(err)
	}

	//获取缓存
	value, _ := redis.Values(conn.Do("hgetall",  "complex1"))


	//将values转成结构体
	var object []TestStruct
	redis.ScanStruct(value, &object)

	fmt.Println(object)

}


//3ms
func DoHashStore(conn redis.Conn)  {

	//使用 以hash类型保存
	conn.Do("hmset",redis.Args{"struct1"}.AddFlat(testStruct)...)


	//获取缓存
	value, _ := redis.Values(conn.Do("hgetall",  "struct1"))


	//将values转成结构体
	object := &TestStruct{}
	redis.ScanStruct(value, object)

}


func DoHashComplexStore(){

}


func DoGobEncodingStore(conn redis.Conn)  {

	var buffer bytes.Buffer
	ecoder := gob.NewEncoder(&buffer)
	ecoder.Encode(testStruct)


	conn.Do("set","struct2",)

	rebytes,_ := redis.Bytes(conn.Do("get","struct2",buffer.Bytes()))
	fmt.Println("gob","buffer.buytes:" ,len(buffer.Bytes()))


	//进行解码
	reader := bytes.NewReader(rebytes)
	dec := gob.NewDecoder(reader)
	object := &TestStruct{}
	dec.Decode(object)

}

func DoJsonEncodingStore(conn redis.Conn)  {
	//json序列化
	datas,_ := json.Marshal(testStruct)
	//缓存数据
	conn.Do("set","struct3",datas)
	//读取数据
	rebytes,_ := redis.Bytes(conn.Do("get","struct3"))
	fmt.Println("gob","json.buytes:" ,len(rebytes))

	//json反序列化
	object := &TestStruct{}
	json.Unmarshal(rebytes,object)

}


/**测试服务连接
 */
func ConnectRedis() redis.Conn{
	conn,err := redis.Dial("tcp", "192.168.11.51:6379")
	if err != nil{
		fmt.Println("连接失败",err)
		return nil
	}

	return conn
}


