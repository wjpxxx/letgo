package dcode

import (
	"go/ast"
	"go/token"
)

//路由检测
type DCodeRouter struct {
	fset    *token.FileSet
	mfile   *ast.File
	modName string
	cs      []ControllerInfo
}

//Visit 遍历抽象语法树
func (c *DCodeRouter) Visit(node ast.Node) ast.Visitor {
	switch node.(type) {
	case *ast.File:
		c.astFile(node.(*ast.File))
		break
	case *ast.GenDecl:
		c.genDecl(node.(*ast.GenDecl))
		break
	case *ast.FuncDecl:
		c.funcDecl(node.(*ast.FuncDecl))
	}
	return c
}

//astFile 根
func (c *DCodeRouter) astFile(file *ast.File) {
	if len(file.Decls) == 0 {
		c.createFunc(file)
	}
}

//genDecl 字节点
func (c *DCodeRouter) genDecl(decl *ast.GenDecl) {

}

//funcDecl 函数节点
func (c *DCodeRouter) funcDecl(fn *ast.FuncDecl) {

}

//创建run函数
func (c *DCodeRouter) createFunc(file *ast.File) {
	obj := &ast.Object{
		Name: "InitRouter",
		Kind: ast.Fun,
		Decl: &ast.FuncDecl{
			Type:&ast.FuncType{
				
			},
			Body: &ast.BlockStmt{},
		},
	}
	fn := &ast.FuncDecl{
		Name: &ast.Ident{
			Name: "InitRouter",
			Obj:  obj,
		},
		Type: &ast.FuncType{
			Params: &ast.FieldList{},
		},
		Body: &ast.BlockStmt{},
	}
	file.Decls = append(file.Decls, fn)
	file.Scope.Objects["InitRouter"] = obj
}
