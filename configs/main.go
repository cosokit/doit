package configs

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type DoitConf struct {
	App AppConf `mapstructure:"app"`
}

type AppConf struct {
	Debug bool `mapstructure:"debug"`
}

var Doit DoitConf
var VTool *viper.Viper

func InitDoit(path, name, configType string) *viper.Viper {
	v := viper.New()
	v.SetConfigName(name) //  设置配置文件名 (不带后缀)
	v.AddConfigPath(path) // 第一个搜索路径
	v.SetConfigType(configType)
	err := v.ReadInConfig() // 搜索路径，并读取配置数据
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		if err := v.Unmarshal(&Doit); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&Doit); err != nil {
		fmt.Println(err)
	}
	VTool = v

	return VTool
}
