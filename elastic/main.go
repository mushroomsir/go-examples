package main

import (
	"log"

	elastic "gopkg.in/olivere/elastic.v5"
)

type Permission int

const (
	Get           Permission                                        = 1 << iota // 1
	Execute                                                                     // 2
	Update                                                                      // 4
	List                                                                        // 8
	Create                                                                      // 16
	Delete                                                                      // 32
	AllPermission = Get | Execute | Update | List | Create | Delete             // 0x3f, 63
	noPermission  = Permission(0)
)

func main() {
	log.Println(AllPermission & Execute)
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	log.Println(client)
	log.Println(err)
}
