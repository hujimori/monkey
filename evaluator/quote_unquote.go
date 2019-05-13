package evaluator

import (
	"github.com/monkey/ast"
	"github.com/monkey/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}
