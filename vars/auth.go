package vars

var (
	JwtAuthExtend uint = 0 // 鉴权：继承自分组，如果路由无分组，则视同为无需鉴权
	JwtAuthNot    uint = 1 // 鉴权：无需鉴权
	JwtAuthMust   uint = 2 // 鉴权：必须鉴权
	JwtAuthShould uint = 3 // 鉴权：就当鉴权，可以无授权信息或鉴权失败（jwt过期等），但应带上用户信息
)
