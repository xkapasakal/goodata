package main

import (
	// "os"
	// "fmt"
	// "github.com/nu7hatch/gouuid"
	"github.com/xkapasakal/goodata/odata-gen-go/generator"
)

// func reduce(entityTypes []EntityType, key string) EntityType {
// 	for _, value := range entityTypes {
// 		if value.Name == key {
// 			return value
// 		}
// 	}
// 	return EntityType{Name: ""}
// }

func main() {
	// fmt.Printf("Hello, %s\n", parser.Parse())
	// u4, err := uuid.NewV4()
	// if err != nil {
	// 	fmt.Println("error:", err)
	// 	return
	// }
	// fmt.Println(u4)

	g := generator.New("../testdata/northwind.metadata.edmx")
	g.Generate()

	// b, err := ioutil.ReadFile("sample-services/northwind.metadata.edmx")
	// if err != nil {
	// 	panic(err)
	// }

	// var q Edmx
	// err = xml.Unmarshal(b, &q)
	// if err != nil {
	// 	fmt.Printf("error: %v", err)
	// 	return
	// }

	// fmt.Printf("Edmx XMLName: %#v\n", q.XMLName)
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

	// fmt.Printf("reflect: %v\n", runtime.FuncForPC(reflect.ValueOf(main).Pointer()).Name())
}
