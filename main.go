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
)



func main() {
	fmt.Printf("Hello, %s\n", parser.Parse())
	u4, err := uuid.NewV4()
	if err != nil {
	    fmt.Println("error:", err)
	    return
	}
	fmt.Println(u4)

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



    b, err := ioutil.ReadFile("sample-services/northwind.metadata.edmx")
    if err != nil { panic(err) }
	
	var q Edmx
	err = xml.Unmarshal(b, &q)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Printf("Edmx XMLName: %#v\n", q.XMLName)
	fmt.Printf("DataServiceVersion: %#v\n", q.DataServices.DataServiceVersion)
	fmt.Printf("First EntityType Property Name: %#v\n", q.DataServices.Schema[0].EntityType[0].Property[0].Name)
	fmt.Printf("First EntityType Name: %#v\n", q.DataServices.Schema[0].EntityType[0].Name)
	fmt.Printf("Count Schema: %#v\n", len(q.DataServices.Schema))
	fmt.Printf("Count EntityType: %#v\n", len(q.DataServices.Schema[0].EntityType))
	fmt.Printf("Count Property: %#v\n", len(q.DataServices.Schema[0].EntityType[0].Property))
	fmt.Printf("Second Schema: %#v\n", q.DataServices.Schema[1])

	for _,element := range q.DataServices.Schema {
		fmt.Printf("Second Schema: %#v\n", element)	  
	}

	fmt.Printf("reflect: %v\n", runtime.FuncForPC(reflect.ValueOf(main).Pointer()).Name())



	type Email struct {
		Where string `xml:"where,attr"`
		Addr  string
	}
	type Address struct {
		City, State string
	}
	type Result struct {
		XMLName xml.Name `xml:"Person"`
		Name    string   `xml:"FullName"`
		Phone   string
		Email   []Email
		Groups  []string `xml:"Group>Value"`
		Address
	}
	v := Result{Name: "none", Phone: "none"}

	data := `
		<Person>
			<FullName>Grace R. Emlin</FullName>
			<Company>Example Inc.</Company>
			<Email where="home">
				<Addr>gre@example.com</Addr>
			</Email>
			<Email where='work'>
				<Addr>gre@work.com</Addr>
			</Email>
			<Group>
				<Value>Friends</Value>
				<Value>Squash</Value>
			</Group>
			<City>Hanga Roa</City>
			<State>Easter Island</State>
		</Person>
	`
	err = xml.Unmarshal([]byte(data), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Printf("XMLName: %#v\n", v.XMLName)
	fmt.Printf("Name: %q\n", v.Name)
	fmt.Printf("Phone: %q\n", v.Phone)
	fmt.Printf("Email: %v\n", v.Email)
	fmt.Printf("Groups: %v\n", v.Groups)
	fmt.Printf("Address: %v\n", v.Address)
}