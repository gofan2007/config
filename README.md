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

###GODOC
```go
package config 

TYPES

type Config struct {
    FilePath string
    // contains filtered or unexported fields
}


func NewConfig(filepath string) *Config
    Create an empty configuration file


func (this *Config) DeleteKey(section, key string) bool

func (this *Config) DeleteSection(section string) bool

func (this *Config) GetBool(section, name string) bool

func (this *Config) GetInt(section, name string) int

func (this *Config) GetString(section, name string) string

func (this *Config) Reload()

func (this *Config) Save()

func (this *Config) SetValue(section string, name string, value interface{})

```