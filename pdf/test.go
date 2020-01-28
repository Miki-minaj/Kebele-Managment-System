package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	err := GeneratePdf("c25.pdf")
	fmt.Println("waiting... from")
	if err != nil {
		panic(err)
	}
}

// GeneratePdf generates our pdf by adding text and images to the page
// then saving it to a file (name specified in params).
func GeneratePdf(filename string) error {

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	//pdf.SetFont("Arial", "B", 16)
	pdf.SetFont("Arial", "", 12)

	// CellFormat(width, height, text, border, position after, align, fill, link, linkStr)

	//pdf.CellFormat(100, 7, "ID CARD", "0", 0, "CM", false, 0, "")
	pdf.MultiCell(120, 4, "\t \t \n \n \n \n Name: \n Age: \n Sex: \n Mother's Name: \n Phone no: \n Nationality: \nRelegion: \n Ocuupation: \n Emergency Contanct Name: \n Emergency Contac Phone: ", "0", "L", false)
	//ImageOptions(src, x, y, width, height, flow, options, link, linkStr)
	pdf.ImageOptions(
		"cover.jpg",
		90, 30,
		20, 20,
		false,
		gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true},
		0,
		"",
	)
	fmt.Println("waiting...")
	return pdf.OutputFileAndClose(filename)
}
