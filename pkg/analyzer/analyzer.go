package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "whyvarscope",
	Doc:      "Check for unnecessary variable scope",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	ins := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nf := []ast.Node{(*ast.IfStmt)(nil)}
	ins.Preorder(nf, func(node ast.Node) {
		ifDecl := node.(*ast.IfStmt)

		initStatement := ifDecl.Init
		if initStatement == nil {
			return
		}

		assignStatement, ok := initStatement.(*ast.AssignStmt)
		if !ok {
			return
		}

		condExpr := ifDecl.Cond
		if condExpr == nil {
			return
		}

		if len(assignStatement.Lhs) != len(assignStatement.Rhs) {
			return
		}

		for i := range len(assignStatement.Lhs) {
			v := assignStatement.Lhs[i]
			vIdent, ok := v.(*ast.Ident)
			if !ok {
				continue
			}
			if countIdentUsedInNode(vIdent, ifDecl) > 2 {
				continue
			}
			if countIdentUsedInNode(vIdent, condExpr) > 1 {
				continue
			}
			pass.Reportf(vIdent.Pos(), "variable %s can be removed and use assignee directly", vIdent.Name)
		}
		return
	})
	return nil, nil
}

func countIdentUsedInNode(v *ast.Ident, node ast.Node) int {
	count := 0
	ast.Inspect(node, func(n ast.Node) bool {
		ident, ok := n.(*ast.Ident)
		if !ok {
			return true
		}
		if ident.Name == v.Name {
			count++
		}
		return true
	})
	return count
}
