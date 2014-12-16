package config

import (
	"testing"
)

const filepath = "test.ini"

func TestConfig(t *testing.T) {
	conf := NewConfig(filepath)
	conf.SetValue("section", "key", "value")
	conf.SetValue("section", "key2", "value")
	conf.SetValue("section1", "key1", "value")
	conf.SetValue("section1", "key2", "value")
	conf.SetValue("section1", "key3", "value")
	conf.SetValue("section2", "key4", "value")
	conf.SetValue("section2", "key5", "123")
	conf.Save()

	conf2 := NewConfig(filepath)
	t.Log(conf2.data)
	if conf2.GetString("section", "key") != "value" {
		t.Fail()
	}
	if conf2.GetInt("section2", "key5") != 123 {
		t.Fail()
	}
}
