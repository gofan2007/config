gofan2007/config
===================================

fork from [enorzw/config](https://github.com/enorzw/config)

##修改点

1. 支持缺省Section
2. 存储时保持Section顺序
3. 存储时保持Key顺序
4. key=value,等号两边无空格

##Install 

    > go get github.com/gofan2007/config
  
###How to use
```go
package main

import(
	"fmt"
	"github.com/gofan2007/config"
)

func main(){
	filepath:="config.ini"
	conf := config.NewConfig(filepath)
	conf.SetValue("default", "user", "aaa")
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
	user=aaa
	
    [database]
    user=xxx
    port=12345

###GODOC
```go
package config 

TYPES

type Config struct {
    FilePath string 
}


func NewConfig(filepath string) *Config  

func (this *Config) DeleteKey(section, key string) bool

func (this *Config) DeleteSection(section string) bool

func (this *Config) GetBool(section, name string) bool

func (this *Config) GetInt(section, name string) int

func (this *Config) GetString(section, name string) string

func (this *Config) Reload()

func (this *Config) Save()

func (this *Config) SetValue(section string, name string, value interface{})

```