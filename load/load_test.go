package load

import (
	"encoding/xml"
	"fmt"
	"testing"
)

// load toml
//
//
type tomlConfig struct {
	Title   string
	Owner   ownerInfo
	DB      database `toml:"database"`
	Servers map[string]server
	Clients clients
}

type ownerInfo struct {
	Name string
	Org  string `toml:"organization"`
}

type database struct {
	Server  string
	Ports   []int
	ConnMax int `toml:"connection_max"`
	Enabled bool
}

type server struct {
	IP string
	DC string
}

type clients struct {
	Data  [][]interface{}
	Hosts []string
}

func TestLoadToml(t *testing.T) {
	var tomlConf tomlConfig
	if loadErr:=LoadedToml("./toml_config.toml", &tomlConf);loadErr != nil {
		t.Log(loadErr)
	}
	t.Log(fmt.Sprintf("tomlConf: %+v", tomlConf))
}

// load yaml
//
//
type yamlConfig struct {
	AppName string    `yaml:"APP_NAME"`
	IP      string    `yaml:"IP"`
	Port    string    `yaml:"PORT"`
	Host    string    `yaml:"HOST"`
	DB      yamlDB    `yaml:"DB"`
	Redis   yamlRedis `yaml:"REDIS"`
}
type yamlDB struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type yamlRedis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	Timeout  string `yaml:"timeout"`
	Auth     string `yaml:"auth"`
}

func TestLoadYaml(t *testing.T) {
	var yamlConf yamlConfig
	if loadErr:=LoadedYaml("./yaml_config.yaml", &yamlConf);loadErr != nil {
		t.Log(loadErr)
	}
	t.Log(fmt.Sprintf("yamlConf: %+v", yamlConf))
}

// load json
//
//
type jsonConfig struct {
	Name     string        `json:"name"`
	Age      int64         `json:"age"`
	Database jsonDBCluster `json:"database"`
}

type jsonDBCluster struct {
	Read  jsonDB `json:"read"`
	Write jsonDB `json:"write"`
}

type jsonDB struct {
	Host     string `json:"host"`
	Port     string `json:"json"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func TestLoadJson(t *testing.T) {
	var jsonConf jsonConfig
	if loadErr := LoadedJson("./json_config.json", &jsonConf);loadErr != nil {
		t.Log(loadErr)
	}

	t.Log(fmt.Sprintf("jsonConf: %+v", jsonConf))
}

// load xml
//
//
type xmlConfig struct {
	XMLName   xml.Name   `xml:"config"` // 指定最外层的标签
	DbDriver  string     `xml:"dbDriver"`
	Db        string     `xml:"db"`
	DbHost    string     `xml:"dbHost"`
	DbPort    string     `xml:"dbPort"`
	DbUser    string     `xml:"dbUser"`
	DbPasswd  string     `xml:"dbPasswd"`
	Receivers SReceivers `xml:"receivers"`
}

type SReceivers struct {
	Flag string   `xml:"flag,attr"` // 读取flag属性
	Age  string   `xml:"age"`
	User []string `xml:"user"` // 读取user数组
}

func TestLoadXml(t *testing.T) {
	var xmlConf xmlConfig
	if loadErr := LoadedXml("./xml_config.xml", &xmlConf);loadErr != nil {
		t.Log(loadErr)
	}
	t.Log(fmt.Sprintf("xmlConf: %+v", xmlConf))
}
