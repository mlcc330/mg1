package dbfunc

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Persion struct {
	Name string
	Age  int
}

func ConnMongoDB() {
	//1. 连接数据库
	fmt.Println("准备连接localhost mongodb services.!")
	sess, err := mgo.Dial("localhost:27017")

	if err != nil {
		fmt.Println("failed to connect", err)
		return
	}

	fmt.Println("connction is successful!!!, session is : ", sess)

	//2. 增加一个人
	tablename := sess.DB("mgdb").C("persion") //得到集合（表）
	//err = tablename.Insert(Persion{"peter",8})						// 插入一条数据

	// 判断数据库的连接
	//if err !=nil{
	//	fmt.Println("failed to connect",err)
	//	return
	//}

	//if err !=nil{
	//	fmt.Println("failed to insert db",err)
	//	return
	//}

	//3. 查看文文档
	var persion []Persion
	tablename.Find(nil).All(&persion)
	fmt.Println(persion)

	// 修改文档
	tablename.Update(bson.M{"name": "peter"}, Persion{"peter", 8})
	tablename.Find(nil).All(&persion)
	fmt.Println(persion)

	// 删除文档
	tablename.RemoveAll(nil)
	tablename.Find(nil).All(&persion)
	fmt.Println(persion)
}
