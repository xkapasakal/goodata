package odata

import (
	"net/url"
)

type DataServiceContext struct {
	BaseUri *url.URL
}

func New(rawurl string) *DataServiceContext {
	context := new(DataServiceContext)
	context.BaseUri, err := url.Parse(rawurl)
	if err != nil {
		panic(err)
	}
	return context
}
