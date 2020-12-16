package types

// 通用的状态
type ComStatus = int

const (
	ComStatusNormal   ComStatus = 1
	ComStatusDisabled ComStatus = 2
	ComStatusDeleted  ComStatus = -9
)

// 用户状态
type CustomerStatus = int
type CustomerLoginStatus = int
type CustomerAuthType = int

const (
	CustomerStatusNormal   = 1
	CustomerStatusDisabled = 2
	CustomerStatusDeleted  = -9
)

const (
	CustomerAuthTypeUsername = 1
	CustomerAuthTypeEmail    = 2
	CustomerAuthTypePhone    = 3
	CustomerAuthTypeWechat   = 4
	CustomerAuthTypeGoogle   = 5
	CustomerAuthTypeFacebook = 6
	CustomerAuthTypeAppleId  = 7
	CustomerAuthTypeWeibo    = 8
	CustomerAuthTypeQQ       = 9
)

const (
	CustomerLoginStatusNormal   = 1
	CustomerLoginStatusDisabled = 2
	CustomerLoginStatusDeleted  = -9
)

// 属性状态
type AttributeType = int

// product
type ProductType int

const (
	ProductItem        ProductType = 1
	ProductVirtualItem ProductType = 2
)
