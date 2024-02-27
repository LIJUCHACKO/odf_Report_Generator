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
	Note    *Notes
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
					<style:paragraph-properties fo:margin-top="0cm" fo:margin-bottom="0cm" loext:contextual-spacing="false" fo:line-height="100%"/>
					<style:text-properties style:font-name="Liberation Sans" fo:font-size="12pt" style:font-size-asian="12pt" style:font-size-complex="12pt" fo:font-weight="normal" style:font-weight-asian="normal" style:font-weight-complex="normal"/>
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
func (Para *Paragraph) AddTextSpan(textspan *TextSpan)  {
	//textspan := Para.Note.NewTextSpan(text, textspanstyle)
	xmlDB.CutPasteAsSubNode(Para.Note.Content, Para.NodeId, textspan.NodeId)
	//xmlDB.InserSubNode(Para.Note.Content, Para.NodeId, xmlDB.GetNodeContentRaw(Para.Note.Content, textspan.NodeId)) //<<
}
func (Para *Paragraph) Wrap(pic *Picture)  {
	//textspan := Para.Note.NewTextSpan(text, textspanstyle)
	xmlDB.CutPasteAsSubNode(Para.Note.Content, Para.NodeId, pic.NodeId)
	//xmlDB.InserSubNode(Para.Note.Content, Para.NodeId, xmlDB.GetNodeContentRaw(Para.Note.Content, textspan.NodeId)) //<<
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

func (Para *Paragraph) SetFontType(fontname string) {
	styletextproperty, _ := xmlDB.GetNode(Para.Note.Content, Para.StyleId, "style:text-properties")
	xmlDB.UpdateAttributevalue(Para.Note.Content, styletextproperty[0], "style:font-name", fontname)
}

func (Para *Paragraph) SetFontSize(size int) {
	font := strconv.Itoa(size) + "pt"
	styletextproperty, _ := xmlDB.GetNode(Para.Note.Content, Para.StyleId, "style:text-properties")
	xmlDB.UpdateAttributevalue(Para.Note.Content, styletextproperty[0], "fo:font-size", font)
	xmlDB.UpdateAttributevalue(Para.Note.Content, styletextproperty[0], "style:font-size-asian", font)
	xmlDB.UpdateAttributevalue(Para.Note.Content, styletextproperty[0], "style:font-size-complexr", font)
}
func (Para *Paragraph) SetMargins(top int, bottom int) {
	styletextproperty, _ := xmlDB.GetNode(Para.Note.Content, Para.StyleId, "style:paragraph-properties")
	xmlDB.UpdateAttributevalue(Para.Note.Content, styletextproperty[0], "fo:margin-top", strconv.Itoa(top)+"cm")
	xmlDB.UpdateAttributevalue(Para.Note.Content, styletextproperty[0], "fo:margin-bottom", strconv.Itoa(bottom)+"cm")
}
func (Para *Paragraph) SetLineHeight(height int) {
	styletextproperty, _ := xmlDB.GetNode(Para.Note.Content, Para.StyleId, "style:paragraph-properties")
	xmlDB.UpdateAttributevalue(Para.Note.Content, styletextproperty[0], "fo:line-height", strconv.Itoa(height)+"%")
}

func (Para *Paragraph) ToBold() {
	styletextproperty, _ := xmlDB.GetNode(Para.Note.Content, Para.StyleId, "style:text-properties")
	xmlDB.UpdateAttributevalue(Para.Note.Content, styletextproperty[0], "fo:font-weight", "bold")
	xmlDB.UpdateAttributevalue(Para.Note.Content, styletextproperty[0], "style:font-weight-asian", "bold")
	xmlDB.UpdateAttributevalue(Para.Note.Content, styletextproperty[0], "style:font-weight-complex", "bold")
}
func (Para *Paragraph) ToItalic() {
	styletextproperty, _ := xmlDB.GetNode(Para.Note.Content, Para.StyleId, "style:text-properties")
	xmlDB.UpdateAttributevalue(Para.Note.Content, styletextproperty[0], "fo:font-style", "italic")
	xmlDB.UpdateAttributevalue(Para.Note.Content, styletextproperty[0], "style:font-style-asian", "italic")
	xmlDB.UpdateAttributevalue(Para.Note.Content, styletextproperty[0], "style:font-style-complex", "italic")

}
func (Para *Paragraph) ToUnderLine() {
	styletextproperty, _ := xmlDB.GetNode(Para.Note.Content, Para.StyleId, "style:text-properties")
	xmlDB.UpdateAttributevalue(Para.Note.Content, styletextproperty[0], "style:text-underline-style", "solid")
	xmlDB.UpdateAttributevalue(Para.Note.Content, styletextproperty[0], "style:text-underline-width", "auto")
	xmlDB.UpdateAttributevalue(Para.Note.Content, styletextproperty[0], "style:text-underline-color", "font-color")
}

func (Para *Paragraph) AddListStyleName(value string) {
	xmlDB.UpdateAttributevalue(Para.Note.Content, Para.StyleId, "style:list-style-name", value)
}
