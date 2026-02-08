package cache

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Redis RedisConfig `yaml:"redis"`
}

type RedisConfig struct {
	Addr         string `yaml:"addr"`
	Password     string `yaml:"password"`
	DB           int    `yaml:"db"`
	PoolSize     int    `yaml:"pool_size"`
	DialTimeout  int    `yaml:"dial_timeout"`
	ReadTimeout  int    `yaml:"read_timeout"`
	WriteTimeout int    `yaml:"write_timeout"`
}

func GetCacheInfo() RedisConfig {
	viper.SetConfigName("cache")        // 设置配置文件名（不带扩展名）
	viper.SetConfigType("yaml")         // 如果配置文件名中没有扩展名，则需要设置
	viper.AddConfigPath("./cache/conf") // 添加搜索路径（当前目录）

	// 尝试读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	fmt.Println("Using config file:", viper.ConfigFileUsed())
	// 解析到结构体
	var config RedisConfig
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	fmt.Println("confi is:", config.Addr)
	return config

}
