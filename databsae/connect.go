package databsae

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)
var D *gorm.DB
func Connect()*gorm.DB {
	db, _ := gorm.Open(mysql.New(mysql.Config{
		DSN: "bugmaker:2982@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize: 191, // string 类型字段的默认长度 uft8为255 utf8mb4为191
		//DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		//DontSupportRenameIndex: true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		//DontSupportRenameColumn: true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: true, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		SkipDefaultTransaction:false,
		DisableForeignKeyConstraintWhenMigrating: true,//使用逻辑外键,不使用物理外键以增加性能
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "dy_",   // table name prefix, table for `User` would be `t_users`
			SingularTable: false, // use singular table name, table for `User` would be `user` with this option enabled
		},
	})
	sqlDB, _ := db.DB()
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(3)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(10)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	db.AutoMigrate(&User{},&Video{},&FavoriteRelationship{},&UserInfo{},&Comment{},&FollowRelationship{})

	return db

}
