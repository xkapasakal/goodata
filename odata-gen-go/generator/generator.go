package generator

import (
	// "os"
	"encoding/xml"
	"fmt"
	// "github.com/nu7hatch/gouuid"
	"io/ioutil"
	// "log"
	// "os"
	// "reflect"
	// "runtime"
	// "text/template"
	"github.com/xkapasakal/goodata/odata-gen-go/descriptor"
	// edmx "github.com/metaleap/go-xsd-pkg/docs.oasis-open.org/odata/odata/v4.0/os/schemas/edmx.xsd_go"
)

type Generator struct {
	EdmxData []byte
}

// TODO path could be URL, file path ...
func New(path string) *Generator {
	g := new(Generator)
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	g.EdmxData = b
	return g
}

func (g *Generator) Generate() {
	var q descriptor.Edmx
	err := xml.Unmarshal(g.EdmxData, &q)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	// fmt.Printf("Edmx XMLName: %#v\n", q)
	for _,schema := range q.DataServices.Schema {
		fmt.Printf("Schema Namespace: %#v\n", schema.Namespace)
		fmt.Printf("Number of EntityTypes: %#v\n", len(schema.EntityType))
	}

	// fmt.Printf("DataServiceVersion: %#v\n", q.DataServices.DataServiceVersion)
	// fmt.Printf("First EntityType Property Name: %#v\n", q.DataServices.Schema[0].EntityType[0].Property[0].Name)
	// fmt.Printf("First EntityType Name: %#v\n", q.DataServices.Schema[0].EntityType[0].Name)
	// fmt.Printf("Count Schema: %#v\n", len(q.DataServices.Schema))
	// fmt.Printf("Count EntityType: %#v\n", len(q.DataServices.Schema[0].EntityType))
	// fmt.Printf("Count Property: %#v\n", len(q.DataServices.Schema[0].EntityType[0].Property))
	// fmt.Printf("Second Schema: %#v\n", q.DataServices.Schema[1])

	// for _,element := range q.DataServices.Schema {
	// 	fmt.Printf("Second Schema: %#v\n", element)
	// }
	// var entityType = reduce(q.DataServices.Schema[0].EntityType, "Category")
	// fmt.Printf("EntityType: %#v\n", entityType)

	// file, err := os.Create("generated/odata.go")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// t, _ := template.ParseFiles("templates/data_context.tmpl")
	// err = t.Execute(file, q.DataServices)
	// if err != nil {
	// 	fmt.Printf("error: %v", err)
	// }
}
