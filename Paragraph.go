///////////////////////
//			<text:p text:style-name="P7">
//                <text:span text:style-name="T5">hghgy</text:span>
//           </text:p>
//
/////////////////////
package odf

import (
	//"fmt"

	"math/rand"
	"strconv"
	"strings"

	"github.com/LIJUCHACKO/XmlDB"
)

type Paragraph struct {
	NodeId  int
	StyleId int
	//xmlDBptr *xmlDB.Database
	Note *Notes
}

func (Note *Notes) NewParagraph(Style string) *Paragraph {

	var Para *Paragraph = new(Paragraph)
	Style_name := "PS" + strconv.Itoa(rand.Intn(100))
	if len(strings.TrimSpace(Style)) > 0 {
		Style_name = strings.TrimSpace(Style)
	}
	//office_Text
	Para.NodeId = Note.WritetoScratchpad("<text:p text:style-name=\"" + Style_name + "\"/>")

	if len(strings.TrimSpace(Style)) == 0 {
		//office_style
		styletext := `<styles><office:styles>
				<style:style style:name="Standard" style:family="paragraph" style:class="text"/>
		        <style:style style:name="Text_20_body" style:display-name="Text body" style:family="paragraph" style:parent-style-name="Standard" style:class="text">
		            <style:paragraph-properties fo:margin-top="0cm" fo:margin-bottom="0.247cm" loext:contextual-spacing="false" fo:line-height="120%"/>
		        </style:style>
			</office:styles>
			<office:automatic-styles>
				<style:style style:name="` + Style_name + `" style:family="paragraph" style:parent-style-name="Text_20_body">
	            	<style:text-properties/>
	            </style:style>
			</office:automatic-styles></styles>`
		id := Note.WritetoScratchpad(styletext)
		Note.IncludeStyle(id)
	}
	StyleNodeid, _ := xmlDB.GetNode(Note.Content, Note.Officeautostyleid, "style:style[style:name=\""+Style_name+"\"]")
	Para.StyleId = StyleNodeid[0]

	//Para.Note.Content = Note.Content
	Para.Note = Note
	return Para
}
func (Para *Paragraph) Style() string {
	Style_name := xmlDB.GetNodeAttribute(Para.Note.Content, Para.StyleId, "style:name")
	return Style_name
}
func (Para *Paragraph) AddText(text string, textspanstyle string) *TextSpan {
	textspan := Para.Note.NewTextSpan(text, textspanstyle)
	xmlDB.CutPasteAsSubNode(Para.Note.Content, Para.NodeId, textspan.NodeId)
	//xmlDB.InserSubNode(Para.Note.Content, Para.NodeId, xmlDB.GetNodeContentRaw(Para.Note.Content, textspan.NodeId)) //<<
	return textspan
}
func (Para *Paragraph) GetPlainText() string {
	if xmlDB.IslowestNode(Para.Note.Content, Para.NodeId) {
		return xmlDB.GetNodeValue(Para.Note.Content, Para.NodeId)
	}
	items := xmlDB.ChildNodes(Para.Note.Content, Para.NodeId)
	Output := ""
	for _, item := range items {
		Output = Output + xmlDB.GetNodeValue(Para.Note.Content, item)
	}
	return Output
}
