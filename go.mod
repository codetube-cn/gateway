module codetube.cn/gateway

go 1.17

require (
	codetube.cn/core v1.0.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	gorm.io/driver/mysql v1.2.1
	gorm.io/gorm v1.22.4
)

require (
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.3 // indirect
)

replace (
	codetube.cn/core v1.0.0 => ../core
	codetube.cn/proto v1.0.0 => ../proto
)
