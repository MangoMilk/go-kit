# go-kit
Golang development kit.

## Support
* Load file
    * [load json](#Id-LoadJson)
    * [load yaml](#Id-LoadYaml)
    * [load xml](#Id-LoadXml)
    * [load toml](#Id-LoadToml)
* Encrypt alg
    * [rsa](#Id-RSA)
    * [aes](#Id-AES)
* Distributed lock
  * [redis](#DLock)
  * [zk]()
  * [consul]()
## Installation
```shell script
go get github.com/MangoMilk/go-kit
```

## Example
### load-kit

#### <a id="Id-loadJson">load .json file</a>
##### json_config.json
```json
{
  "name": "Dwarf",
  "age": 18,
  "database": {
    "read": {
      "host": "127.0.0.1",
      "port": "3306",
      "database": "test_read",
      "username": "root",
      "password": "root"
    },
    "write": {
      "host": "127.0.0.1",
      "port": "3306",
      "database": "test_write",
      "username": "root",
      "password": "root"
    }
  }
}
```
##### example.go
```go
import (
    "github.com/MangoMilk/go-kit/load"
)

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

func ExampleLoadJson()  {
    var jsonConf jsonConfig
    loadErr := load.LoadedJson("./json_config.json", &jsonConf)
    if loadErr != nil {
        panic(loadErr)
    }
    fmt.Println(fmt.Sprintf("jsonConf: %+v", jsonConf)))
}
```

#### <a id="Id-loadYaml">load .yaml file</a>
##### yaml_config.yaml
```yaml
APP_NAME: yaml_config
IP: 127.0.0.1
PORT: 80
HOST:

DB:
  host: 127.0.0.1
  port: 3306
  user: root
  password: root
  database: test

REDIS:
  host: 127.0.0.1
  port: 6379
  database: 0
  timeout: 2000
  auth: 111111
```
##### example.go
```go
import (
    "github.com/MangoMilk/go-kit/load"
)

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

func ExampleLoadYaml() {
    var yamlConf yamlConfig
    loadErr := load.LoadedYaml("./yaml_config.yaml", &yamlConf)
    if loadErr != nil {
        panic(loadErr)
    }
    fmt.Println(fmt.Sprintf("yamlConf: %+v", yamlConf))
}
```

#### <a id="Id-loadToml">load .toml file</a>
##### toml_config.toml
```toml
title = "TOML Example"

[owner]
name = "VV"
organization = "Dwarf"

[database]
server = "127.0.0.1"
ports = [ 8001, 8001, 8002 ]
connection_max = 5000
enabled = true

[servers]

  [servers.alpha]
  ip = "10.0.0.1"
  dc = "eqdc10"

  [servers.beta]
  ip = "10.0.0.2"
  dc = "eqdc10"

[clients]
data = [ ["gamma", "delta"], [1, 2] ] # just an update to make sure parsers support it

hosts = [
  "alpha",
  "omega"
]
```
##### example.go
```go
import (
    "github.com/MangoMilk/go-kit/load"
)

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

func ExampleLoadToml() {
    var tomlConf tomlConfig
    loadErr != load.LoadedToml("./toml_config.toml", &tomlConf)
    if loadErr != nil {
        panic(loadErr)
    }
    fmt.Println(fmt.Sprintf("tomlConf: %+v", tomlConf))
}
```

#### <a id="Id-loadXml">load .xml file</a>
##### xml_config.xml
```xml
<?xml version="1.0" encoding="UTF-8"?>
<config>
    <dbDriver>mysql</dbDriver>
    <dbHost>127.0.0.1</dbHost>
    <dbPort>3306</dbPort>
    <db>test</db>
    <dbUser>root</dbUser>
    <dbPasswd>root</dbPasswd>
    <receivers flag="true">
        <age>16</age>
        <user>U1</user>
        <user>U2</user>
    </receivers>
</config>
```
##### example.go
```go
import (
    "github.com/MangoMilk/go-kit/load"
)

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

func TestLoadXml() {
    var xmlConf xmlConfig
    loadErr := load.LoadedXml("./xml_config.xml", &xmlConf)
    if loadErr != nil {
        panic(loadErr)
    }
    fmt.Println(fmt.Sprintf("xmlConf: %+v", xmlConf))
}
```
### encrypt-kit

#### <a id="Id-RSA">RSA</a>
```go
import (
    "github.com/MangoMilk/go-kit/encrypt"
    "encoding/json"
    "fmt"
    "testing"
)

func ExampleEncrypt()  {
    r := encrypt.NewRSA()
    // Use the second 'filePath' param to save key into file.
    //pubKey, priKey, genPemErr := r.GeneratePem(2048,"./pem")
    pubKey, priKey, genPemErr := r.GeneratePem(2048)
    if genPemErr != nil {
        panic(genPemErr)
    }
    fmt.Println("rsa Public Key: ", pubKey)
    fmt.Println("rsa Private Key: ", priKey)
    
    dataByte, _ := json.Marshal(People{Name: "Dwarf"})
    cipherByte := r.Encrypt(dataByte, []byte(pubKey))
    fmt.Println("rsa Encrypt: ", string(cipherByte))
    fmt.Println("rsa Decrypt: ", r.Decrypt(cipherByte, []byte(priKey)))
}
```

#### <a id="Id-AES">AES</a>
```go
import (
    "github.com/MangoMilk/go-kit/encrypt"
    "encoding/json"
    "fmt"
)

func ExampleEncrypt()  {
    type People struct {
        Name string
    }
    
    secret := "asdfreqw34thv123"
    dataByte, _ := json.Marshal(People{Name: "Dwarf"})

    a := NewAES(ECB)
    // encrypt
    cipher, err := a.Encrypt(dataByte, secret)
    if err != nil {
        panic(err)
    }
    fmt.Println("AES Encrypt: ", string(cipher))
    
    // decrypt
    data, aErr := a.Decrypt(cipher, secret)
    if aErr != nil {
        panic(aErr)
    }
    fmt.Println("AES Decrypt: ", string(data))
}
```

### lock-kit

#### <a id="DLock">Distributed lock</a>
```go
import (
	"github.com/go-redis/redis"
	"time"
	"github.com/MangoMilk/go-kit/dlock"
	"fmt"
)

func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "123465",
		DB:       0,
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		painc(err)
	}

	lockKey := "Test"
	dLock := dlock.NewRedisDLock(rdb)
	res, err := dLock.Lock(lockKey, time.Second*3)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

```