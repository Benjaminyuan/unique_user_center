package conf

import (
	"github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type MysqlConf struct {
	MysqlAddr     string `yaml:"MysqlAddr"`
	//MysqlPassword string `yaml:"MysqlPassword"`
}
type RPCConf struct {
	Addr string `yaml:"addr"`
	Network string `yaml:"network"`
}
type Conf struct {
	MysqlConf MysqlConf `yaml:"MysqlConf"`
	RPCConf RPCConf `yaml:"RPCConf"`
}
var (
	Data = &Conf{}
)
func InitConf()error{
	fileName,_:= filepath.Abs("./conf/conf.yaml")
	yamlData,err := ioutil.ReadFile(fileName)
	if err != nil{
		logrus.Fatalf("fail to load file,err: %v",err)
		return err
	}
	err = yaml.Unmarshal(yamlData, Data)
	if err != nil{
		logrus.Fatalf("fail to unmarshal yamlData,err: %v",err)
		return err
	}
	logrus.Debugf("conf_data: %v", Data)
	return nil
}