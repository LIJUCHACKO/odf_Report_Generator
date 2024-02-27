package main

import (
	//"fmt"

	//"os"

	//"github.com/LIJUCHACKO/XmlDB"
	"github.com/LIJUCHACKO/odf_Report_Generator"
)

func main() {

	Note := odf.NewDatabase()

	///First Article
	Text := Note.NewTextSpan("21", "")
	Text.ToItalic()
	Text.ToBold()
	Note.CreateArticle(Text.NodeId, "!%Para1%!")
	Text2 := Note.NewTextSpan("23", "")
	Note.CreateArticle(Text2.NodeId, "!%Para2%!")

	///Second Article
	Header := Note.NewHeader1(1, "")
	Header.AddText("List of items", "")

	List1 := Note.NewBulletinList("")
	item1 := List1.AddItemPara("")
	item1.AddText("item1", "")
	item2 := List1.AddItemPara("")
	item2.AddText("item2", "")
	item1.ToUnderLine()
	item2.ToBold()
	ART := Note.CreateArticle(Header.NodeId, "!%Listofitems%!")
	ART.AddContentArticle(List1.NodeId)

	//Third Article
	Table := Note.NewTable(2, 2, "")
	C00 := Table.AddItemPara(0, 0, "")
	C00.AddText("A", "")
	C01 := Table.AddItemPara(0, 1, "")
	C01.AddText("B", "")
	C10 := Table.AddItemPara(1, 0, "")
	C10.AddText("C", "")
	C11 := Table.AddItemPara(1, 1, "")
	C11.AddText("D", "")
	Note.CreateArticle(Table.NodeId, "!%Table%!")

	Picture := Note.NewPicture("","periodic-table-bw-download.jpg")
	Note.CreateArticle(Picture.NodeId, "!%Picture1%!")


	Parag := Note.NewParagraph("")
	Parag.AddText("kjjjj","")
	Parag.AddTextSpan(Text2)
	Note.CreateArticle(Parag.NodeId, "!%Para3%!")

	//Processing template to generate final report
	odf.ProcessOdtfile("Report_Template.odt", "Report.odt", Note)
}
