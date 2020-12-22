package bootstrap

import (
	_ "litshop/src/config"
	_ "litshop/src/config/mysql"
	_ "litshop/src/config/redis"
	_ "litshop/src/lvm/runtime"
	_ "litshop/src/pkg/logger"
)

func init() {

}
