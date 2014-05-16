package main

import (
	// "os"
	"fmt"
	"encoding/xml"
	"github.com/xkapasakal/go4OData/parser"
	"io/ioutil"
	"reflect"
    "runtime"
    "github.com/nu7hatch/gouuid"
    "log"
	"os"
	"text/template"
)

type PropertyRef struct {
	Name string `xml:"Name,attr"`
}

type Key struct {
	PropertyRef []PropertyRef
}

type NavigationProperty struct {
	Name string `xml:"Name,attr"`
	Relationship string `xml:"Relationship,attr"`
	ToRole string `xml:"ToRole,attr"`
	FromRole string `xml:"FromRole,attr"`
}

type Property struct {
	Name string `xml:"Name,attr"`
	Type string `xml:"Type,attr"`
	Nullable bool `xml:"Nullable,attr"`
	MaxLength int32 `xml:"MaxLength,attr"`
	FixedLength bool `xml:"FixedLength,attr"`
	Unicode bool `xml:"Unicode,attr"`
}   

type EntityType struct {
	XMLName xml.Name `xml:"http://schemas.microsoft.com/ado/2008/09/edm EntityType"`
	Name string `xml:"Name,attr"`
	Property []Property
	NavigationProperty NavigationProperty
	// Have to specify where to find episodes since this
	// doesn't match the xml tags of the data that needs to go into it
	// EpisodeList []Episode `xml:"Episode>"`
}

func (p *Property) ConvertTypes() string {
	var typesMap = map[string]string{
		"Edm.Binary": "[]byte",
		"Edm.Boolean": "bool",
		"Edm.Byte": "byte",
		"Edm.DateTime": "time.Time",
		"Edm.Decimal": "float32",
		"Edm.Double": "float64",
		"Edm.Single": "float32",
		"Edm.Guid": "gouuid.UUID",
		"Edm.Int16": "int16",
		"Edm.Int32": "int32",
		"Edm.Int64": "int64",
		"Edm.SByte": "",
		"Edm.String": "string",
		"Edm.Time": "time.Time",
		"Edm.DateTimeOffset": "",
		"Edm.Geography": "",
		"Edm.GeographyPoint": "",
		"Edm.GeographyLineString": "",
		"Edm.GeographyPolygon": "",
		"Edm.GeographyMultiPoint": "",
		"Edm.GeographyMultiLineString": "",
		"Edm.GeographyMultiPolygon": "",
		"Edm.GeographyCollection": "",
		"Edm.Geometry": "",
		"Edm.GeometryPoint": "",
		"Edm.GeometryLineString": "",
		"Edm.GeometryPolygon": "",
		"Edm.GeometryMultiPoint": "",
		"Edm.GeometryMultiLineString": "",
		"Edm.GeometryMultiPolygon": "",
		"Edm.GeometryCollection": "",
		"Edm.Stream": "",
	}
	return typesMap[p.Type]
}

type End struct {
	Multiplicity string `xml:"Multiplicity,attr"`
	Type string `xml:"Type,attr"`
	Role string `xml:"Role,attr"`
	EntitySet string `xml:"EntitySet,attr"`
}

type Dependent struct {
	Role string `xml:"Role,attr"`
	PropertyRef PropertyRef
}

type Principal struct {
	Role string `xml:"Role,attr"`
	PropertyRef PropertyRef
}

type ReferentialConstraint struct {
	Principal Principal
	Dependent Dependent
}

type Association struct {
	Name string `xml:"Name,attr"`
	End []End
	ReferentialConstraint ReferentialConstraint
}

type AssociationSet struct {
	Name string `xml:"Name,attr"`
	Association string `xml:"Association,attr"`
	End []End
}

type EntitySet struct {
	Name string `xml:"Name,attr"`
	EntityType string `xml:"EntityType,attr"`
}

type EntityContainer struct {
	Name string `xml:"Name,attr"`
	IsDefaultEntityContainer bool `xml:"IsDefaultEntityContainer,attr"`
	LazyLoadingEnabled bool `xml:"LazyLoadingEnabled,attr"`
	EntitySet []EntitySet
	AssociationSet []AssociationSet
}

type Schema struct {
	Namespace string `xml:"Namespace,attr"`
	EntityType   []EntityType
	EntityContainer EntityContainer
}

type DataServices struct {
	DataServiceVersion string `xml:"DataServiceVersion,attr"`
	MaxDataServiceVersion float32 `xml:"MaxDataServiceVersion,attr"`
	Schema []Schema
}

type Edmx struct {
	XMLName xml.Name `xml:"http://schemas.microsoft.com/ado/2007/06/edmx Edmx"`
	DataServices DataServices
}

func reduce(entityTypes []EntityType, key string) EntityType {
	for _, value := range entityTypes {
		if value.Name == key {
			return value
		}
	}	
	return EntityType{Name:""}
}

func main() {
	fmt.Printf("Hello, %s\n", parser.Parse())
	u4, err := uuid.NewV4()
	if err != nil {
	    fmt.Println("error:", err)
	    return
	}
	fmt.Println(u4)


    b, err := ioutil.ReadFile("sample-services/northwind.metadata.edmx")
    if err != nil { panic(err) }
	
	var q Edmx
	err = xml.Unmarshal(b, &q)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Printf("Edmx XMLName: %#v\n", q.XMLName)
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
	var entityType = reduce(q.DataServices.Schema[0].EntityType, "Category")
	fmt.Printf("EntityType: %#v\n", entityType)


	file, err := os.Create("generated/odata.go");
	if err != nil {
		log.Fatal(err)
	}
	t, _ := template.ParseFiles("templates/data_context.tmpl")
	err = t.Execute(file, q.DataServices)
	if err != nil { fmt.Printf("error: %v", err) }

	fmt.Printf("reflect: %v\n", runtime.FuncForPC(reflect.ValueOf(main).Pointer()).Name())



	// type Email struct {
	// 	Where string `xml:"where,attr"`
	// 	Addr  string
	// }
	// type Address struct {
	// 	City, State string
	// }
	// type Result struct {
	// 	XMLName xml.Name `xml:"Person"`
	// 	Name    string   `xml:"FullName"`
	// 	Phone   string
	// 	Email   []Email
	// 	Groups  []string `xml:"Group>Value"`
	// 	Address
	// }
	// v := Result{Name: "none", Phone: "none"}

	// data := `
	// 	<Person>
	// 		<FullName>Grace R. Emlin</FullName>
	// 		<Company>Example Inc.</Company>
	// 		<Email where="home">
	// 			<Addr>gre@example.com</Addr>
	// 		</Email>
	// 		<Email where='work'>
	// 			<Addr>gre@work.com</Addr>
	// 		</Email>
	// 		<Group>
	// 			<Value>Friends</Value>
	// 			<Value>Squash</Value>
	// 		</Group>
	// 		<City>Hanga Roa</City>
	// 		<State>Easter Island</State>
	// 	</Person>
	// `
	// err = xml.Unmarshal([]byte(data), &v)
	// if err != nil {
	// 	fmt.Printf("error: %v", err)
	// 	return
	// }
	// fmt.Printf("XMLName: %#v\n", v.XMLName)
	// fmt.Printf("Name: %q\n", v.Name)
	// fmt.Printf("Phone: %q\n", v.Phone)
	// fmt.Printf("Email: %v\n", v.Email)
	// fmt.Printf("Groups: %v\n", v.Groups)
	// fmt.Printf("Address: %v\n", v.Address)
}