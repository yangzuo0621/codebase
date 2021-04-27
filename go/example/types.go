package example

//go:generate mockgen -destination=./mock_$GOPACKAGE/types.go -package=mock_$GOPACKAGE github.com/yangzuo0621/codebase/go/example Foo
type Foo interface {
	Bar() (string, error)
}
