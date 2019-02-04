package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

type wlpName string

func (n *wlpName) Name() string {
	return "wlp_" + string(*n)
}

type protocol struct {
	Name wlpName `xml:"name,attr"`
	Itfs []*itf  `xml:"interface"`
}

func (p *protocol) gen(dirname string) error {
	tmpl, err := template.ParseFiles("client.gotmpl")
	if err != nil {
		return err
	}

	for _, itf := range p.Itfs {
		fn := filepath.Join(dirname, itf.Name+".go")

		f, err := os.Create(fn)
		if err != nil {
			return err
		}
		defer f.Close()

		if err = tmpl.Execute(f, itf); err != nil {
			return err
		}
	}
	return nil
}

type itf struct {
	Name    string    `xml:"name,attr"`
	Version int       `xml:"version,attr"`
	Desc    desc      `xml:"description"`
	Reqs    []request `xml:"request"`
	Events  []event   `xml:"event"`
	Enums   []enum    `xml:"enum"`
}

func (itf *itf) StructName() string {
	return strcase.ToCamel(itf.Name)
}

type request struct {
	Name string `xml:"name,attr"`
	Desc desc   `xml:"description"`
	Args []arg  `xml:"arg"`
}

func (r *request) FuncName() string {
	return strcase.ToCamel(r.Name)
}

type event struct {
	Name string `xml:"name,attr"`
	Desc desc   `xml:"description"`
	Args []arg  `xml:"arg"`
}

func (e *event) FuncName() string {
	return strcase.ToCamel(e.Name) + "Handler"
}

func (e *event) TypeName() string {
	return strcase.ToCamel(e.Name)
}

func (e *event) FieldName() string {
	return strcase.ToLowerCamel(e.Name) + "Handler"
}

type enum struct {
	Name    string  `xml:"name,attr"`
	Desc    desc    `xml:"description"`
	Entries []entry `xml:"entry"`
}

func (e *enum) TypeName() string {
	return strcase.ToCamel(e.Name)
}

type entry struct {
	Name    string `xml:"name,attr"`
	Value   string `xml:"value,attr"`
	Summary string `xml:"summary,attr"`
}

func (e *entry) VarName() string {
	return strcase.ToCamel(e.Name)
}

type desc struct {
	Summary string `xml:"summary,attr"`
	Desc    string `xml:",chardata"`
}

func (d *desc) DescString() string {
	var b strings.Builder
	for _, line := range strings.Split(strings.Trim(d.Desc, " \t\n"), "\n") {
		b.WriteString("//")
		line = strings.TrimLeft(line, " \t")
		if len(line) > 0 {
			b.WriteString(" ")
			b.WriteString(line)
		}
		b.WriteString("\n")
	}
	return b.String()
}

type arg struct {
	Name      string `xml:"name,attr"`
	Type      string `xml:"type,attr"`
	Summary   string `xml:"summary,attr"`
	Interface string `xml:"interface,attr"`
}

func (a *arg) Parameter() string {
	return a.VarName() + " " + a.typeName()
}

func (a *arg) ExprEvent() string {
	return a.VarName()
}

func (a *arg) ExprRequest() string {
	name := a.VarName()

	if len(a.Interface) > 0 {
		name += ".ID()"
	}
	return name
}

func (a *arg) IsRealNewID() bool {
	return a.Type == "new_id" && len(a.Interface) > 0
}

func (a *arg) VarName() string {
	if a.Name == "interface" {
		return "itf"
	}
	return strcase.ToLowerCamel(a.Name)
}

func (a *arg) typeName() string {
	if len(a.Interface) > 0 {
		return "*" + strcase.ToCamel(a.Interface)
	}
	if a.Type == "new_id" {
		return "wire.ID"
	}
	if a.Type == "object" {
		return "wire.Object"
	}

	tn := strcase.ToCamel(a.Type)
	if tn == "Fd" {
		tn = "FD"
	}
	return "wire." + tn
}

func (a *arg) TypeNameOnly() string {
	tn := strcase.ToCamel(a.Type)
	if tn == "Fd" {
		tn = "FD"
	}
	if a.Type == "new_id" || a.Type == "object" {
		tn = "ID"
	}
	return tn
}

func (a *arg) ItfTypeName() string {
	return strcase.ToCamel(a.Interface)
}

func usage() {
	fmt.Fprintf(os.Stderr, "%s: <type> <inupt xml file> <output path>\n", filepath.Base(os.Args[0]))
}

func main() {
	if len(os.Args) != 4 {
		usage()
		os.Exit(1)
	}
	typ := os.Args[1]
	switch typ {
	case "client":
		if err := generateClientCode(os.Args[2], os.Args[3]); err != nil {
			log.Fatal(err)
		}
	// case "server":
	// 	if err := generateServerCode(os.Args[1], os.Args[2]); err != nil {
	// 		log.Fatal(err)
	// 	}
	default:
		usage()
		os.Exit(1)
	}
}

func generateClientCode(filename, dirname string) error {
	xmldata, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	var p protocol
	if err := xml.Unmarshal(xmldata, &p); err != nil {
		return err
	}

	return p.gen(dirname)
}
