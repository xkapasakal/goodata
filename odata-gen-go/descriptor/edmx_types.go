package descriptor

import "encoding/xml"

type PropertyRef struct {
	Name string `xml:"Name,attr"`
}

type Key struct {
	PropertyRef []PropertyRef
}

type NavigationProperty struct {
	Name         string `xml:"Name,attr"`
	Relationship string `xml:"Relationship,attr"`
	ToRole       string `xml:"ToRole,attr"`
	FromRole     string `xml:"FromRole,attr"`
}

type Property struct {
	Name        string `xml:"Name,attr"`
	Type        string `xml:"Type,attr"`
	Nullable    bool   `xml:"Nullable,attr"`
	MaxLength   int32  `xml:"MaxLength,attr"`
	FixedLength bool   `xml:"FixedLength,attr"`
	Unicode     bool   `xml:"Unicode,attr"`
}

type EntityType struct {
	XMLName            xml.Name `xml:"http://schemas.microsoft.com/ado/2008/09/edm EntityType"`
	Name               string   `xml:"Name,attr"`
	Property           []Property
	NavigationProperty NavigationProperty
	// Have to specify where to find episodes since this
	// doesn't match the xml tags of the data that needs to go into it
	// EpisodeList []Episode `xml:"Episode>"`
}

type End struct {
	Multiplicity string `xml:"Multiplicity,attr"`
	Type         string `xml:"Type,attr"`
	Role         string `xml:"Role,attr"`
	EntitySet    string `xml:"EntitySet,attr"`
}

type Dependent struct {
	Role        string `xml:"Role,attr"`
	PropertyRef PropertyRef
}

type Principal struct {
	Role        string `xml:"Role,attr"`
	PropertyRef PropertyRef
}

type ReferentialConstraint struct {
	Principal Principal
	Dependent Dependent
}

type Association struct {
	Name                  string `xml:"Name,attr"`
	End                   []End
	ReferentialConstraint ReferentialConstraint
}

type AssociationSet struct {
	Name        string `xml:"Name,attr"`
	Association string `xml:"Association,attr"`
	End         []End
}

type EntitySet struct {
	Name       string `xml:"Name,attr"`
	EntityType string `xml:"EntityType,attr"`
}

type EntityContainer struct {
	Name                     string `xml:"Name,attr"`
	IsDefaultEntityContainer bool   `xml:"IsDefaultEntityContainer,attr"`
	LazyLoadingEnabled       bool   `xml:"LazyLoadingEnabled,attr"`
	EntitySet                []EntitySet
	AssociationSet           []AssociationSet
}

type Schema struct {
	Namespace       string `xml:"Namespace,attr"`
	EntityType      []EntityType
	EntityContainer EntityContainer
}

type DataServices struct {
	DataServiceVersion    string  `xml:"DataServiceVersion,attr"`
	MaxDataServiceVersion float32 `xml:"MaxDataServiceVersion,attr"`
	Schema                []Schema
}

type Edmx struct {
	XMLName      xml.Name `xml:"http://schemas.microsoft.com/ado/2007/06/edmx Edmx"`
	DataServices DataServices
}

func (p *Property) ConvertTypes() string {
	var typesMap = map[string]string{
		"Edm.Binary":                   "[]byte",
		"Edm.Boolean":                  "bool",
		"Edm.Byte":                     "byte",
		"Edm.DateTime":                 "time.Time",
		"Edm.Decimal":                  "float32",
		"Edm.Double":                   "float64",
		"Edm.Single":                   "float32",
		"Edm.Guid":                     "gouuid.UUID",
		"Edm.Int16":                    "int16",
		"Edm.Int32":                    "int32",
		"Edm.Int64":                    "int64",
		"Edm.SByte":                    "",
		"Edm.String":                   "string",
		"Edm.Time":                     "time.Time",
		"Edm.DateTimeOffset":           "",
		"Edm.Geography":                "",
		"Edm.GeographyPoint":           "",
		"Edm.GeographyLineString":      "",
		"Edm.GeographyPolygon":         "",
		"Edm.GeographyMultiPoint":      "",
		"Edm.GeographyMultiLineString": "",
		"Edm.GeographyMultiPolygon":    "",
		"Edm.GeographyCollection":      "",
		"Edm.Geometry":                 "",
		"Edm.GeometryPoint":            "",
		"Edm.GeometryLineString":       "",
		"Edm.GeometryPolygon":          "",
		"Edm.GeometryMultiPoint":       "",
		"Edm.GeometryMultiLineString":  "",
		"Edm.GeometryMultiPolygon":     "",
		"Edm.GeometryCollection":       "",
		"Edm.Stream":                   "",
	}
	return typesMap[p.Type]
}
