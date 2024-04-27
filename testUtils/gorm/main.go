package main

import (
	"fmt"
	"gin-one/testUtils/gorm/internal"
)

func main() {
	db := InitGorm()
	internal.CreateBelone(db)
	internal.CreateOne2One(db)
	internal.CreateOne2Many(db)

	internal.GetBelone(db)
	internal.GetOne2One(db)
	//internal.GetOne2Many(db)
	fmt.Println("db:", db)
}
