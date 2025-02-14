package main

import (
	"context"
	"fmt"
)

import (
	"dubbo-go-pixiu-benchmark/dubbogo/pkg"
	"dubbo.apache.org/dubbo-go/v3/config"
	hessian "github.com/apache/dubbo-go-hessian2"
)

var (
	userProvider = &pkg.UserProvider1{}
)

func init() {
	config.SetConsumerService(userProvider)
	hessian.RegisterPOJO(&pkg.User{})
}

func main() {

	path := "dubbo-go-benchmark/3.0/dubbogo/stress/dubbogo.yml"
	err := config.Load(config.WithPath(path))

	if err != nil {

	}

	user, err := userProvider.GetUser(context.TODO(), &pkg.User{ID: "1113333", Name: "chengxingyuan", Code: 111, Age: 111})

	if err != nil {
		fmt.Println("Print the result of the current client call: %s", user)
	}

}
