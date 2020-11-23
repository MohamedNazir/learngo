package printer

import "fmt"

func unExportedPrinter() {
	fmt.Println("This wil not be printed if called from other package, as it was not exported")
}

// ExportedPrinter method to print something.
func ExportedPrinter() {
	fmt.Println("This is a print statement from the exported package")
}
