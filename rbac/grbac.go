package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	// Initialize a Xorm adapter with MySQL database.
	a, err := xormadapter.NewAdapter("mysql", "root:root@tcp(127.0.0.1:3306)/casbin?charset=utf8", true)
	if err != nil {
		log.Fatalf("error: adapter: %s", err)
	}

	m, err := model.NewModelFromString(`
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[role_definition]
g = _, _

[matchers]
m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
`)
	if err != nil {
		log.Fatalf("error: model: %s", err)
	}

	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		log.Fatalf("error: enforcer: %s", err)
	}

	sub := "alice" // 想要访问资源的用户。
	obj := "data1" // 将被访问的资源。
	act := "read" // 用户对资源执行的操作。

	ok, err := e.Enforce(sub, obj, act)
	if err != nil {
		// 处理err
	}

	if ok == true {
		// 允许alice读取data1
		fmt.Println("success")
	} else {
		// 拒绝请求，抛出异常
		fmt.Println("fail")
	}
	roles, err := e.GetRolesForUser("tom")
	if err == nil {
		fmt.Println(roles)
	}

}
