package config

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	isLoad   bool
	FilePath string
	data     map[string]map[string]string
}

func NewConfig(filepath string) *Config {
	conf := new(Config)
	conf.FilePath = filepath
	conf.loadConfig()
	return conf
}

func (this *Config) GetBool(section, name string) bool {
	if !this.isLoad {
		this.loadConfig()
	}
	result, _ := strconv.ParseBool(this.data[section][name])
	return result
}

func (this *Config) GetString(section, name string) string {
	if !this.isLoad {
		this.loadConfig()
	}
	return this.data[section][name]
}

func (this *Config) GetInt(section, name string) int {
	if !this.isLoad {
		this.loadConfig()
	}
	result, _ := strconv.Atoi(this.data[section][name])
	return result
}

func (this *Config) SetValue(section string, name string, value interface{}) {
	for sect, kvs := range this.data {
		if sect == section {
			kvs[name] = fmt.Sprintf("%v", value)
			return
		}
	}
	this.data[section] = make(map[string]string)
	this.data[section][name] = fmt.Sprintf("%v", value)
}

func (this *Config) DeleteSection(section string) bool {
	for sect, _ := range this.data {
		if sect == section {
			delete(this.data, sect)
			return true
		}
	}
	return false
}

func (this *Config) DeleteKey(section, key string) bool {
	for sect, kvs := range this.data {
		if sect == section {
			for k, _ := range kvs {
				if k == key {
					delete(kvs, k)
					return true
				}
			}
		}
	}
	return false
}

func (this *Config) Reload() {
	this.loadConfig()
}

func (this *Config) loadConfig() {
	this.isLoad = true
	file, err := os.OpenFile(this.FilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		fmt.Println("Error in Opening File:" + err.Error())
		return
	}
	defer file.Close()
	this.data = make(map[string]map[string]string)
	var section string
	buf := bufio.NewReader(file)
	for {
		l, err := buf.ReadString('\n')
		line := strings.TrimSpace(l)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error in Reading Config File:" + err.Error())
			}
			if len(line) == 0 {
				break
			}
		}
		switch {
		case len(line) == 0:
		case strings.HasPrefix(line, "#"):
		case line[0] == '[' && line[len(line)-1] == ']':
			section = strings.TrimSpace(line[1 : len(line)-1])
			this.data[section] = make(map[string]string)
		default:
			i := strings.IndexAny(line, "=")
			value := strings.TrimSpace(line[i+1 : len(line)])
			this.data[section][strings.TrimSpace(line[0:i])] = value
		}
	}
}

func (this *Config) Save() {
	file, err := os.Create(this.FilePath)
	if err != nil {
		fmt.Println("Error in Opening File:" + err.Error())
		return
	}
	defer file.Close()
	var buffer string
	for section, kvs := range this.data {
		buffer += "[" + section + "]\n"
		for key, value := range kvs {
			buffer += key + " = " + value + "\n"
		}
		buffer += "\n"
	}
	buffer += ""
	file.WriteString(buffer)
}
