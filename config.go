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
	data     map[string]map[string]string //节.key.value
	sections []string                     //节名称顺序
	keys     map[string][]string          //节内key顺序
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
	for _, sec := range this.sections {
		if sec == section {
			for _, key := range this.keys[sec] {
				if key == name {
					//有key修改
					this.data[sec][key] = fmt.Sprintf("%v", value)
					return
				}
			}
			//无key增加
			this.keys[sec] = append(this.keys[sec], name)
			this.data[sec][name] = fmt.Sprintf("%v", value)
			return
		}
	}
	//无section增加
	this.sections = append(this.sections, section)
	this.keys[section] = []string{name}
	this.data[section] = map[string]string{name: fmt.Sprintf("%v", value)}
}

func (this *Config) DeleteSection(section string) bool {
	for i, sect := range this.sections {
		if sect == section {
			this.sections = append(this.sections[:i], this.sections[i+1:]...)
			delete(this.keys, sect)
			delete(this.data, sect)
			return true
		}
	}
	return false
}

func (this *Config) DeleteKey(section, key string) bool {
	for _, sect := range this.sections {
		if sect == section {
			for j, k := range this.keys[sect] {
				if k == key {
					this.keys[sect] = append(this.keys[sect][:j], this.keys[sect][j+1:]...)
					delete(this.data[sect], key)
					return true
				}
			}
			break
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
	this.sections = []string{"default"}
	this.keys = make(map[string][]string)
	//this.keys["default"] = nil
	this.data["default"] = make(map[string]string)
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
			this.sections = append(this.sections, section)
			//this.keys[section] = nil
		default:
			i := strings.IndexAny(line, "=")
			value := strings.TrimSpace(line[i+1 : len(line)])
			key := strings.TrimSpace(line[0:i])
			if section == "" {
				this.data["default"][key] = value
				this.keys["default"] = append(this.keys["default"], key)
			} else {
				this.data[section][strings.TrimSpace(line[0:i])] = value
				this.keys[section] = append(this.keys[section], key)
			}
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
	for _, section := range this.sections {
		if section != "default" {
			buffer += "[" + section + "]\n"
		}
		for _, key := range this.keys[section] {
			buffer += key + "=" + this.data[section][key] + "\n"
		}
		buffer += "\n"
	}
	file.WriteString(buffer)
}
