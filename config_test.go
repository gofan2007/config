package config

import (
	"testing"
)

const filepath = "test.ini"

func TestConfig(t *testing.T) {

	conf := NewConfig(filepath)
	conf.SetValue("database", "user", "xxx")
	conf.SetValue("database", "port", 12345)
	conf.SetValue("database", "isSync", true)
	conf.SetValue("RemoveSection", "Remove", true)
	conf.SetValue("TestSection2", "Remove", true)
	conf.SetValue("TestSection2", "Remain", true)

	conf.DeleteSection("RemoveSection")
	conf.DeleteKey("TestSection2", "Remove")

	conf.Save()

	conf2 := NewConfig(filepath)
	str := conf2.GetString("database", "user")
	num := conf2.GetInt("database", "port")
	isSync := conf2.GetBool("database", "isSync")

	if str != "xxx" {
		t.Fail()
	}
	if num != 12345 {
		t.Fail()
	}
	if isSync != true {
		t.Fail()
	}

	for sec, _ := range conf2.data {
		if sec == "RemoveSection" {
			t.Fail()
		}
	}

	if conf2.GetBool("TestSection2", "Remove") {
		t.Fail()
	}

	if !conf2.GetBool("TestSection2", "Remain") {
		t.Fail()
	}
}
