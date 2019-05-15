package main

import (
	"GoTest/redis/provider"
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var testStruct = provider.CreateTestData(1111)






func main() {
	conn := provider.Conn
	DoHashStore(conn)
	DoJsonEncodingStore(conn)
	DoGobEncodingStore(conn)

}




//3ms
func DoHashStore(conn redis.Conn)  {

	//使用 以hash类型保存
	conn.Do("hmset",redis.Args{"struct1"}.AddFlat(testStruct)...)


	//获取缓存
	value, _ := redis.Values(conn.Do("hgetall",  "struct1"))


	//将values转成结构体
	object := &provider.TestStruct{}
	redis.ScanStruct(value, object)

}




func DoGobEncodingStore(conn redis.Conn)  {

	var buffer bytes.Buffer
	ecoder := gob.NewEncoder(&buffer)
	ecoder.Encode(testStruct)


	conn.Do("set","struct2",buffer.Bytes())

	rebytes,_ := redis.Bytes(conn.Do("get","struct2"))
	fmt.Println("gob","buffer.buytes:" ,len(buffer.Bytes()))


	//进行解码
	reader := bytes.NewReader(rebytes)
	dec := gob.NewDecoder(reader)
	object := &provider.TestStruct{}
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
	object := &provider.TestStruct{}
	json.Unmarshal(rebytes,object)

}


