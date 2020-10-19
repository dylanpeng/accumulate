package main

import (
	"fmt"
	"github.com/go-zookeeper/zk"
	"time"
)

var (
	hosts       = []string{"127.0.0.1:2181"}
	path        = "/wtzk"
	flags int32 = zk.FlagEphemeral
	data        = []byte("zk data 001")
	acls        = zk.WorldACL(zk.PermAll)
)

func main() {
	// 连接zk
	conn, _, err := zk.Connect(hosts, time.Second*5)
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 开始监听path
	_, _, event, err := conn.ExistsW(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 协程调用监听事件
	go watchZkEvent(event)

	// 触发创建数据操作
	create(conn, path, data)

}

// zk watch 回调函数
func callback(event zk.Event) {
	// zk.EventNodeCreated
	// zk.EventNodeDeleted
	fmt.Println("###########################")
	fmt.Println("path: ", event.Path)
	fmt.Println("type: ", event.Type.String())
	fmt.Println("state: ", event.State.String())
	fmt.Println("---------------------------")
}

// zk 回调函数
func watchZkEvent(e <-chan zk.Event) {
	event := <-e
	fmt.Println("###########################")
	fmt.Println("path: ", event.Path)
	fmt.Println("type: ", event.Type.String())
	fmt.Println("state: ", event.State.String())
	fmt.Println("---------------------------")
}

// 创建数据
func create(conn *zk.Conn, path string, data []byte) {
	_, err := conn.Create(path, data, flags, acls)
	if err != nil {
		fmt.Printf("创建数据失败: %v\n", err)
		return
	}
	fmt.Println("创建数据成功")
}

// 增
func add(conn *zk.Conn, path string) {
	var data = []byte("test value")
	// flags有4种取值：
	// 0:永久，除非手动删除
	// zk.FlagEphemeral = 1:短暂，session断开则该节点也被删除
	// zk.FlagSequence  = 2:会自动在节点后面添加序号
	// 3:Ephemeral和Sequence，即，短暂且自动添加序号
	var flags int32 = 0
	// 获取访问控制权限
	acls := zk.WorldACL(zk.PermAll)
	s, err := conn.Create(path, data, flags, acls)
	if err != nil {
		fmt.Printf("创建失败: %v\n", err)
		return
	}
	fmt.Printf("创建: %s 成功", s)
}

// 查
func get(conn *zk.Conn, path string) {
	data, _, err := conn.Get(path)
	if err != nil {
		fmt.Printf("查询%s失败, err: %v\n", path, err)
		return
	}
	fmt.Printf("%s 的值为 %s\n", path, string(data))
}

// 删改与增不同在于其函数中的version参数,其中version是用于 CAS支持
// 可以通过此种方式保证原子性
// 改
func modify(conn *zk.Conn, path string) {
	new_data := []byte("hello zookeeper")
	_, sate, _ := conn.Get(path)
	_, err := conn.Set(path, new_data, sate.Version)
	if err != nil {
		fmt.Printf("数据修改失败: %v\n", err)
		return
	}
	fmt.Println("数据修改成功")
}

// 删
func del(conn *zk.Conn, path string) {
	_, sate, _ := conn.Get(path)
	err := conn.Delete(path, sate.Version)
	if err != nil {
		fmt.Printf("数据删除失败: %v\n", err)
		return
	}
	fmt.Println("数据删除成功")
}
