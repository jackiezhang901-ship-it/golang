package config

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	config := getDbInfo()
	usermame := config.DB.UserName
	password := config.DB.Password
	url := config.DB.Url
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/test", usermame, password, url)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	DB = db
	// 获取通用数据库对象 sql.DB ，以便于显示地设置连接池参数
	sqlDB, err := db.DB()

	if err != nil {
		log.Fatalf("failed to get *sql.DB: %v", err)
	}

	// 设置最大连接数
	sqlDB.SetMaxOpenConns(25)

	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(25)

	// 设置连接的最大生命周期
	sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Println("Database connected and connection pool configured successfully!")
}

func getDbInfo() Config {
	viper.SetConfigName("db")        // 设置配置文件名（不带扩展名）
	viper.SetConfigType("yaml")      // 如果配置文件名中没有扩展名，则需要设置
	viper.AddConfigPath("./conf/db") // 添加搜索路径（当前目录）

	// 尝试读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	fmt.Println("Using config file:", viper.ConfigFileUsed())
	// 解析到结构体
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	fmt.Println("confi is:", config.DB.UserName)
	return config

}
