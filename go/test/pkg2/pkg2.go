package pkg2

import "fmt"

const pkgName string = "pkg2"

var PkgNameVar string = getPkgName()

func init() {
	fmt.Println("pkg2.init() method invoked")
}

func getPkgName() string {
	fmt.Println("pkg2.PkgNameVar has been initialized")
	return pkgName
}
