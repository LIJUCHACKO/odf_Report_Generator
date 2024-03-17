package odf

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
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
	OrginalPictureFiles   []string
	NewPictureFiles   []string
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
	xml_content.Libreofficemod = true
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

/**below function is not complete**/
func (Note *Notes) CreateArticleFromTemplate(Inputfile string, ArtName string) *Article {
	r, err := zip.OpenReader(Inputfile)
	if err != nil {
		fmt.Println(err)
	}
	defer r.Close()
	var Art *Article = new(Article)
	styletext := `<styles>`
	for _, f := range r.File {
		//fmt.Println(f.Name)
		fileToZip, _ := f.Open()
		//defer fileToZip.Close()
		//	header, _ := zip.FileInfoHeader(f.FileInfo())
		//header.Name = f.Name
		//header.Method = zip.Deflate
		//writer, _ := targetzipwriter.CreateHeader(header)
		if f.Name == "content.xml" {
			xmlfile, _ := ioutil.ReadAll(fileToZip)
			xmlstring := string(xmlfile)

			//xmlline = processString(xmlline)
			var DB *xmlDB.Database = new(xmlDB.Database)
			DB.MaxNooflines = 9999999
			xmllines := strings.Split(xmlstring, "\n")
			xmlDB.Load_dbcontent(DB, xmllines)
			DB.Debug_enabled = false
			var Doc *Odt = new(Odt)
			Doc.Content = DB
			Doc.ReplaceMarkers(Note)
			//MERGE STYLES
			Doc.ImportAutoStyles(Note)
			remNodeid, _ := xmlDB.GetNode(Doc.Content, 0, "office:body/office:forms")
			xmlDB.RemoveNode(Doc.Content, remNodeid[0])
			remNodeid, _ = xmlDB.GetNode(Doc.Content, 0, "office:body/text:sequence-decls")
			xmlDB.RemoveNode(Doc.Content, remNodeid[0])
			Nodeid, _ := xmlDB.GetNode(Doc.Content, 0, "office:body/office:text")

			xmlstring = "<article name=\"" + ArtName + "\">" + xmlDB.GetNodeContentRaw(Doc.Content, Nodeid[0]) + "</article>"
			id := Note.WritetoScratchpad(xmlstring)
			Art.Content = Note.Content
			Art.Nodeid = id
			Styleid, _ := xmlDB.GetNode(Doc.Content, 0, "office:automatic-styles")
			styletext = styletext + `<office:automatic-styles>` + xmlDB.GetNodeContentRaw(Doc.Content, Styleid[0]) + `</office:automatic-styles>`

			//_, _ = writer.Write([]byte(xmlDB.Dump_DB(Doc.Content)))
		} else if f.Name == "styles.xml" {
			xmlfile, _ := ioutil.ReadAll(fileToZip)
			xmlstring := string(xmlfile)

			//xmlline = processString(xmlline)
			var DB *xmlDB.Database = new(xmlDB.Database)
			DB.MaxNooflines = 9999999
			xmllines := strings.Split(xmlstring, "\n")
			xmlDB.Load_dbcontent(DB, xmllines)
			DB.Debug_enabled = false
			var Doc *Odt = new(Odt)
			Doc.Content = DB
			//MERGE STYLES
			Doc.ImportStyles(Note)
			Styleid, _ := xmlDB.GetNode(Doc.Content, 0, "office:styles")
			styletext = styletext + `<office:styles>` + xmlDB.GetNodeContentRaw(Doc.Content, Styleid[0]) + `</office:styles>`
			//_, _ = writer.Write([]byte(xmlDB.Dump_DB(Doc.Content)))
		} else {

			//io.Copy(writer, fileToZip)
		}

		fileToZip.Close()

	}
	styletext = styletext + `</styles>`
	id := Note.WritetoScratchpad(styletext)
	Note.IncludeStyle(id)

	return Art
}
