package main

import (
	//"fmt"

	//"os"

	//"github.com/LIJUCHACKO/XmlDB"
	"github.com/LIJUCHACKO/odf_Report_Generator"
)

func main() {

	Note := odf.NewDatabase()

	Text := Note.NewTextSpan("21", "")
	Text.ToItalic()
	Text.ToBold()
	Note.CreateArticle(Text.NodeId, "!%Para1%!")
	Text2 := Note.NewTextSpan("23", "")
	Note.CreateArticle(Text2.NodeId, "!%Para2%!")

	Para := Note.NewParagraph("")
	Para.AddText("List of items", "")
	List1 := Note.NewNumberedList("")
	List1.AddItem("Item1")
	List1.AddItem("Item2")
	ART := Note.CreateArticle(Para.NodeId, "!%Listofitems%!")
	ART.AddContentArticle(List1.NodeId)

	odf.ProcessOdtfile("Report_Template.odt", "Report.odt", Note)
}
