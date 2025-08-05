package pkg1

import (
	"fmt"

	"github.com/learn/init_order/pkg2"
)

const pkgName string = "pkg1"

var PkgNameVar string = getPkgName() + pkg2.PkgNameVar

func init() {
	fmt.Println("pkg1.init() method invoked")
}

func getPkgName() string {
	fmt.Println("pkg1.PkgNameVar has been initialized")
	return pkgName
}
