package odata

type edmxUnmarshaler interface {
	UnmarshalEdmx(text []byte) error
}

type edmxParser struct {
	done bool // whether the parsing is finished (success or error)
}
