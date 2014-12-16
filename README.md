enorzw/config
===================================

##install 

> go get github.com/enorzw/config

##how to use

"` 
> conf := NewConfig(filepath)

>	conf.SetValue("database", "user", "xxx")

>	conf.SetValue("database", "port", "12345")

> conf.Save()
  
> conf.GetString("database", "user")

> conf.GetInt("database", "port")
"`
