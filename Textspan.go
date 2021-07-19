package odf

import (
	//"fmt"

	"math/rand"
	"strconv"
	"strings"

	"github.com/LIJUCHACKO/XmlDB"
)

type TextSpan struct {
	NodeId  int
	StyleId int
	Note    *Notes
}

const (
	BOLD      = iota // BOLD = 0
	ITALIC           // ITALIC = 1
	UNDERLINE        // uNDERLINE = 2
)

func (Note1 *Notes) NewTextSpan(text string, Style string) *TextSpan {

	var Text *TextSpan = new(TextSpan)
	Text.Note = Note1
	Style_name := "TS" + strconv.Itoa(rand.Intn(100))
	if len(strings.TrimSpace(Style)) > 0 {
		Style_name = strings.TrimSpace(Style)
	}

	//office_Text
	Text.NodeId = Note1.WritetoScratchpad("<text:span text:style-name=\"" + Style_name + "\">" + text + "</text:span>")

	if len(strings.TrimSpace(Style)) == 0 {
		//office_style
		styletext := `<styles><office:automatic-styles>
				<style:style style:name="` + Style_name + `" style:family="text">
					<style:text-properties/>
        		</style:style>
				</office:automatic-styles></styles>`

		id := Note1.WritetoScratchpad(styletext)

		Note1.IncludeStyle(id)

	}
	StyleNodeid, _ := xmlDB.GetNode(Note1.Content, Note1.Officeautostyleid, "style:style[style:name=\""+Style_name+"\"]")
	Text.StyleId = StyleNodeid[0]
	return Text
}
func (Text *TextSpan) Style() string {
	Style_name := xmlDB.GetNodeAttribute(Text.Note.Content, Text.StyleId, "style:name")
	return Style_name
}
func (Text *TextSpan) ToBold() {
	styletextproperty, _ := xmlDB.GetNode(Text.Note.Content, Text.StyleId, "style:text-properties")
	xmlDB.UpdateAttributevalue(Text.Note.Content, styletextproperty[0], "fo:font-weight", "bold")
	xmlDB.UpdateAttributevalue(Text.Note.Content, styletextproperty[0], "style:font-weight-asian", "bold")
	xmlDB.UpdateAttributevalue(Text.Note.Content, styletextproperty[0], "style:font-weight-complex", "bold")
}
func (Text *TextSpan) ToItalic() {
	styletextproperty, _ := xmlDB.GetNode(Text.Note.Content, Text.StyleId, "style:text-properties")
	xmlDB.UpdateAttributevalue(Text.Note.Content, styletextproperty[0], "fo:font-style", "italic")
	xmlDB.UpdateAttributevalue(Text.Note.Content, styletextproperty[0], "style:font-style-asian", "italic")
	xmlDB.UpdateAttributevalue(Text.Note.Content, styletextproperty[0], "style:font-style-complex", "italic")

}
func (Text *TextSpan) ToUnderLine() {
	styletextproperty, _ := xmlDB.GetNode(Text.Note.Content, Text.StyleId, "style:text-properties")
	xmlDB.UpdateAttributevalue(Text.Note.Content, styletextproperty[0], "style:text-underline-style", "solid")
	xmlDB.UpdateAttributevalue(Text.Note.Content, styletextproperty[0], "style:text-underline-width", "auto")
	xmlDB.UpdateAttributevalue(Text.Note.Content, styletextproperty[0], "style:text-underline-color", "font-color")
}
func (Text *TextSpan) ToNormal() {
	styletextproperty, _ := xmlDB.GetNode(Text.Note.Content, Text.StyleId, "style:text-properties")
	xmlDB.UpdateAttributevalue(Text.Note.Content, styletextproperty[0], "fo:font-weight", "normal")
	xmlDB.UpdateAttributevalue(Text.Note.Content, styletextproperty[0], "style:font-weight-asian", "normal")
	xmlDB.UpdateAttributevalue(Text.Note.Content, styletextproperty[0], "style:font-weight-complex", "normal")
	xmlDB.UpdateAttributevalue(Text.Note.Content, styletextproperty[0], "style:text-underline-style", "")
	xmlDB.UpdateAttributevalue(Text.Note.Content, styletextproperty[0], "style:text-underline-width", "")
	xmlDB.UpdateAttributevalue(Text.Note.Content, styletextproperty[0], "style:text-underline-color", "")
	xmlDB.UpdateAttributevalue(Text.Note.Content, styletextproperty[0], "fo:font-style", "")
	xmlDB.UpdateAttributevalue(Text.Note.Content, styletextproperty[0], "style:font-style-asian", "")
	xmlDB.UpdateAttributevalue(Text.Note.Content, styletextproperty[0], "style:font-style-complex", "")
}
