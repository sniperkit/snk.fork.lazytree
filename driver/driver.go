/*
Sniperkit-Bot
- Status: analyzed
*/

package driver

import (
	"github.com/sniperkit/snk.fork.lazytree/object"
)

type Reference interface {
	Encode() string
}

// Common interfaces for all drivers to implement
type Driver interface {
	Name() string
	ResolveReference(objectType object.ObjectType, reference interface{}) (object.Object, error)
	DecodeReference(encodedReference string) (Reference, error)
}
