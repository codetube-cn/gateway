module codetube.cn/gateway

go 1.18

require (
	codetube.cn/core v1.0.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	gorm.io/driver/mysql v1.3.3
	gorm.io/gorm v1.23.4
)

require (
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
)

replace (
	codetube.cn/core v1.0.0 => ../core
	codetube.cn/proto v1.0.0 => ../proto
)
