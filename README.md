enorzw/config
===================================

##install 

    > go get github.com/enorzw/config
  
###how to use
    conf := NewConfig("conf")
    conf.SetValue("database", "user", "xxx")
    conf.SetValue("database", "port", "12345")

    conf.Save()
  
    conf.GetString("database", "user")
    conf.GetInt("database", "port")
###
    
        > go run main.go
        xxx
        12345
        > cat conf
        [database]
        user = xxx
        port = 12345
