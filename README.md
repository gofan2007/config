enorzw/config
===================================

##Install 

    > go get github.com/enorzw/config
  
###How to use
```go
package main

import(
	"fmt"
	"github.com/enorzw/config"
)

func main(){
	filepath:="config.ini"
	conf := config.NewConfig(filepath)
	conf.SetValue("database", "user", "xxx")
	conf.SetValue("database", "port", "12345")

	conf.Save()

	str:=conf.GetString("database", "user")
	num:=conf.GetInt("database", "port")
	fmt.Println(str)
	fmt.Println(num)
}
```
###Runing Result

    > go run main.go
    xxx
    12345
    > cat conf
    [database]
    user = xxx
    port = 12345
