package main

import (
	"GoTest/redis/provider"
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var testComplexStruct =provider.CreateComplexData(2000)



func main() {
	conn := provider.Conn
	DoComplexJSONStore(conn)
	DoComplexGobEncodingStore(conn)

}






/**经测试是无效方法，hash无法支持复杂结构**/
func DoComplexHashStore(conn redis.Conn)  {

	var args = redis.Args{"complex1"}.AddFlat(testComplexStruct)
	//使用 以hash类型保存
	_,err := conn.Do("hmset",args...)

	if err != nil{
		fmt.Println(err)
	}

	//获取缓存
	value, _ := redis.Values(conn.Do("hgetall",  "complex1"))


	//将values转成结构体
	var object []provider.TestStruct
	redis.ScanStruct(value, &object)

	fmt.Println(object)

}



func DoComplexJSONStore(conn redis.Conn){
	//序列化数组
	datas,_ := json.Marshal(testComplexStruct)
	//缓存数据
	conn.Do("set","complex2",datas)
	//读取数据
	rebytes,_ := redis.Bytes(conn.Do("get","complex2"))
	//fmt.Println("json",len(rebytes))

	//json反序列化
	var object []provider.TestStruct
	json.Unmarshal(rebytes,&object)

}


func DoComplexGobEncodingStore(conn redis.Conn)  {
	//序列化数组
	var buffer bytes.Buffer
	ecoder := gob.NewEncoder(&buffer)
	ecoder.Encode(testComplexStruct)

	//缓存数据
	conn.Do("set","complex3",buffer.Bytes())

	//读取数据
	rebytes,_ := redis.Bytes(conn.Do("get","complex3"))
	//fmt.Println("gob",len(buffer.Bytes()))

	//反序列化
	reader := bytes.NewReader(rebytes)
	dec := gob.NewDecoder(reader)
	var object []provider.TestStruct
	dec.Decode(&object)

}