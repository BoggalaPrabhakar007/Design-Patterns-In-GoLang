/*Interface segregation principle*/
/*In an object-oriented class-based language, it states that if a class provides methods to multiple clients,
then rather than having a generic interface loaded with all methods, provide a separate interface for each
client and implement all of them in the class.*/

package main

type Document struct {
}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct {
}

func (m *MultiFunctionPrinter) Print(d Document) {

}

func (m *MultiFunctionPrinter) Fax(d Document) {

}

func (m *MultiFunctionPrinter) Scan(d Document) {

}

type OldFashionPrinter struct {
}

func (o *OldFashionPrinter) Print(d Document) {

}

//Fax Deprecated because old printer does not support
func (o *OldFashionPrinter) Fax(d Document) {

}

//Scan Deprecated because old printer does not support
func (o *OldFashionPrinter) Scan(d Document) {

}

//Interface segregation principle

type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type MyPrinter struct {
}
type PhotoScanner struct {
}

func (m *MyPrinter) Print(d Document) {

}

func (p *PhotoScanner) Scan(d Document) {

}

type MultiFunctionDevice interface {
	Printer
	Scanner
	Fax(d Document)
}

type multiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m *multiFunctionMachine) print(d Document) {
	m.printer.Print(d)
}

func main() {

}
