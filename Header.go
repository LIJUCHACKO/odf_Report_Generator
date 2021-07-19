package odf

import (
	//"fmt"

	//	"math/rand"
	"strconv"
	"strings"

	"github.com/LIJUCHACKO/XmlDB"
)

type Header struct {
	NodeId  int
	StyleId int
	Note    *Notes
}

//various attributes for different levels
var display_bame = []string{"Heading 1", "Heading 2", "Heading 3", "Heading 4", "Heading 5", "Heading 6"}
var default_outline_level = []string{"1", "2", "3", "4", "5", "6"}
var font_size_asian = []string{"18pt", "16pt", "14pt", "12pt", "12pt", "12pt"}
var font_size_asian_perce = []string{"130%", "115%", "101%", "95%", "85%", "85%"}
var margin_top = []string{"0.423cm", "0.353cm", "0.247cm", "0.212cm", "0.212cm", "0.106cm"}
var margin_bottom = []string{"0.212cm", "0.212cm", "0.212cm", "0.212cm", "0.106cm", "0.106cm"}

func (Note *Notes) NewHeader1(level int, Style string) *Header {

	var Header *Header = new(Header)
	levelno := strconv.Itoa(level)
	Style_name := "PH" + levelno
	if len(strings.TrimSpace(Style)) > 0 {
		Style_name = strings.TrimSpace(Style)
	}
	//office_Text
	Header.NodeId = Note.WritetoScratchpad("<text:h text:style-name=\"" + Style_name + "\"/>")

	if len(strings.TrimSpace(Style)) == 0 {
		//fmt.Println("stylenew")
		//office_style
		styletext := `<styles><office:styles>
				 <style:style style:name="Heading_20_` + levelno + `" style:display-name="Heading ` + levelno + `" style:family="paragraph" style:parent-style-name="Heading" style:next-style-name="Text_20_body" style:default-outline-level="` + levelno + `" style:class="text">
		            <style:paragraph-properties fo:margin-top="` + margin_top[level] + `" fo:margin-bottom="` + margin_bottom[level] + `" loext:contextual-spacing="false"/>
		            <style:text-properties fo:font-size="` + font_size_asian_perce[level] + `" fo:font-weight="bold" style:font-size-asian="` + font_size_asian_perce[level] + `" style:font-weight-asian="bold" style:font-size-complex="` + font_size_asian_perce[level] + `" style:font-weight-complex="bold"/>
		        </style:style>
			</office:styles>
			<office:automatic-styles>
				<style:style style:name="` + Style_name + `" style:family="paragraph" style:parent-style-name="Heading_20_` + levelno + `" >
				<style:text-properties fo:font-size="` + font_size_asian[level] + `" style:font-size-asian="` + font_size_asian[level] + `" style:font-size-complex="` + font_size_asian[level] + `"/>
				</style:style>
			</office:automatic-styles></styles>`
		id := Note.WritetoScratchpad(styletext)
		Note.IncludeStyle(id)
	}
	StyleNodeid, _ := xmlDB.GetNode(Note.Content, Note.Officeautostyleid, "style:style[style:name=\""+Style_name+"\"]")
	Header.StyleId = StyleNodeid[0]

	//Para.Note.Content = Note.Content
	Header.Note = Note
	return Header
}
func (Head *Header) Style() string {
	Style_name := xmlDB.GetNodeAttribute(Head.Note.Content, Head.StyleId, "style:name")
	return Style_name
}
func (Head *Header) AddText(text string, textspanstyle string) *TextSpan {
	textspan := Head.Note.NewTextSpan(text, textspanstyle)
	xmlDB.CutPasteAsSubNode(Head.Note.Content, Head.NodeId, textspan.NodeId)
	//xmlDB.InserSubNode(Para.Note.Content, Para.NodeId, xmlDB.GetNodeContentRaw(Para.Note.Content, textspan.NodeId)) //<<
	return textspan
}
