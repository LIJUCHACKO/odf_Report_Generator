package odf

import (
	//"fmt"
	"strings"

	"github.com/LIJUCHACKO/XmlDB"
)

type Article struct {
	Content *xmlDB.Database
	Nodeid  int
}
type Notes struct {
	Content           *xmlDB.Database
	scratchpadid      int
	Officeautostyleid int
	Officestyleid     int
}

func NewDatabase() *Notes {
	var Note *Notes = new(Notes)
	text := `<contents>
				<text/>
				<office:styles/>
				<office:automatic-styles/>
				<scratchpad/>
			</contents>`
	xmllines := strings.Split(text, "\n")
	var xml_content *xmlDB.Database = new(xmlDB.Database)
	///xml_content.Debug_enabled = true
	xmlDB.Load_dbcontent(xml_content, xmllines)
	Note.Content = xml_content
	scratchpadids, _ := xmlDB.GetNode(xml_content, 0, "scratchpad")
	Note.scratchpadid = scratchpadids[0]

	officeautostyleids, _ := xmlDB.GetNode(xml_content, 0, "office:automatic-styles")
	Note.Officeautostyleid = officeautostyleids[0]

	officestyleids, _ := xmlDB.GetNode(xml_content, 0, "office:styles")
	Note.Officestyleid = officestyleids[0]

	return Note
}

func (Note *Notes) InsertSubNode(NodeId int, xmlstring string) int {
	newnode, _ := xmlDB.InserSubNode(Note.Content, NodeId, xmlstring)
	return newnode[0]
}
func (Note *Notes) WritetoScratchpad(xmlstring string) int {
	//fmt.Println(xmlstring)
	newnode, _ := xmlDB.InserSubNode(Note.Content, Note.scratchpadid, xmlstring)
	return newnode[0]
}
func (Note *Notes) IncludeStyle(Nodeid int) {
	//import styles from scratchpad

	office_automatic_styles, _ := xmlDB.GetNode(Note.Content, Nodeid, "office:automatic-styles")
	office_automatic_styleEach := xmlDB.ChildNodes(Note.Content, office_automatic_styles[0])
	for _, style := range office_automatic_styleEach {
		identified, _ := xmlDB.GetNode(Note.Content, Note.Officeautostyleid, "<x>[style:name=\""+xmlDB.GetNodeAttribute(Note.Content, style, "style:name")+"\"]")
		//Will not replace if the style is already present
		if len(identified) == 0 {
			xmlDB.InserSubNode(Note.Content, Note.Officeautostyleid, xmlDB.GetNodeContentRaw(Note.Content, style))
		}
	}

	office_styles, _ := xmlDB.GetNode(Note.Content, Nodeid, "office:styles")
	if len(office_styles) > 0 {
		office_stylesEach := xmlDB.ChildNodes(Note.Content, office_styles[0])
		for _, style := range office_stylesEach {
			identified, _ := xmlDB.GetNode(Note.Content, Note.Officestyleid, "<x>[style:name=\""+xmlDB.GetNodeAttribute(Note.Content, style, "style:name")+"\"]")
			//Will not replace if the style is already present
			if len(identified) == 0 {
				xmlDB.InserSubNode(Note.Content, Note.Officestyleid, xmlDB.GetNodeContentRaw(Note.Content, style))
			}
		}
	}

}

func (Note *Notes) CreateArticle(Nodeid int, name string) *Article {
	var Art *Article = new(Article)
	xmlstring := "<article name=\"" + name + "\">" + xmlDB.GetNodeContentRaw(Note.Content, Nodeid) + "</article>"
	id := Note.WritetoScratchpad(xmlstring)
	Art.Content = Note.Content
	Art.Nodeid = id
	return Art
}
func (Art *Article) AddContentArticle(Nodeid int) {
	xmlDB.InserSubNode(Art.Content, Art.Nodeid, xmlDB.GetNodeContentRaw(Art.Content, Nodeid))
}
