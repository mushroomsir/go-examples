package main

import (
	"log"
	"reflect"
)

// AppVerifyInfoTpl ...
type AppVerifyInfoTpl struct {
	Token    string `valid:"required;xx" json:"token" swaggo:"true,app token,"`
	Tenant   string `valid:"-" json:"tenant"`
	Resource string `valid:"required" json:"resource"`
	Action   int32  `valid:"required" json:"action"`
}

func main() {

	tpl := &AppVerifyInfoTpl{
		Token: "x",
	}
	v := reflect.ValueOf(tpl)
	log.Println(v.Kind())
	v = v.Elem()
	log.Println(v.Kind())
	//  v.Field(0)
	//log.Println(v.FieldByName("Token"))

	t := v.Type()
	log.Println(t.Field(0).Tag)
	opt, _ := t.Field(0).Tag.Lookup("valid")
	log.Println(opt)
}
