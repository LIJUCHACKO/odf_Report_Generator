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
	"path/filepath"
)

type Picture struct {
	NodeId  int
	StyleId int
	Note    *Notes

}

func (Note *Notes) NewPicture(Style string,PicFile string) *Picture {

	var Pict *Picture = new(Picture)
	Pict.Note = Note
	Style_name := "Pic" + strconv.Itoa(rand.Intn(100))
	if len(strings.TrimSpace(Style)) > 0 {
		Style_name = strings.TrimSpace(Style)
	}
	extension := filepath.Ext(PicFile)
	Note.OrginalPictureFiles=append(Note.OrginalPictureFiles,PicFile)

	filenameinside := "Pictures/"+strconv.Itoa(rand.Intn(10000000))+extension
	Note.NewPictureFiles=append(Note.NewPictureFiles,filenameinside)
	//Picture.FileNameInside=filenameinside;
	//Picture.FileName=PicFile;
	//office_Text
	//fmt.println()
	Pict.NodeId = Note.WritetoScratchpad(`<draw:frame draw:style-name="`+Style_name+`" draw:name="Image`+strconv.Itoa(rand.Intn(100))+`" text:anchor-type="char" svg:x="0cm" svg:y="0.229cm" svg:width="17cm" svg:height="11.137cm" draw:z-index="0">
	<draw:image xlink:href="`+filenameinside+`" xlink:type="simple" xlink:show="embed" xlink:actuate="onLoad" draw:mime-type="image/jpeg"/>
	</draw:frame>`)

	NewPara := Note.NewParagraph("")

	NewPara.Wrap(Pict)
	Pict.NodeId = NewPara.NodeId

	if len(strings.TrimSpace(Style)) == 0 {
		//office_style
		styletext := `<styles><office:styles>
				<style:style style:name="Graphics" style:family="graphic">
			<style:graphic-properties text:anchor-type="paragraph" svg:x="0cm" svg:y="0cm" style:wrap="dynamic" style:number-wrapped-paragraphs="no-limit" style:wrap-contour="false" style:vertical-pos="top" style:vertical-rel="paragraph" style:horizontal-pos="center" style:horizontal-rel="paragraph"/>
			</style:style>
			</office:styles>
			<office:automatic-styles>
				<style:style style:name="`+Style_name+`" style:family="graphic" style:parent-style-name="Graphics">
			<style:graphic-properties style:wrap="none" style:vertical-pos="from-top" style:vertical-rel="paragraph" style:horizontal-pos="from-left" style:horizontal-rel="paragraph" style:mirror="none" fo:clip="rect(0cm, 0cm, 0cm, 0cm)" draw:luminance="0%" draw:contrast="0%" draw:red="0%" draw:green="0%" draw:blue="0%" draw:gamma="100%" draw:color-inversion="false" draw:image-opacity="100%" draw:color-mode="standard" draw:wrap-influence-on-position="once-concurrent" loext:allow-overlap="false"/>
			</style:style>
			</office:automatic-styles></styles>`
		id := Note.WritetoScratchpad(styletext)
		Note.IncludeStyle(id)
	}
	StyleNodeid, _ := xmlDB.GetNode(Note.Content, Note.Officeautostyleid, "style:style[style:name=\""+Style_name+"\"]")
	Pict.StyleId = StyleNodeid[0]

	//Pict.Note.Content = Note.Content
	Pict.Note = Note
	return Pict
}
func (Pict *Picture) Style() string {
	Style_name := xmlDB.GetNodeAttribute(Pict.Note.Content, Pict.StyleId, "style:name")
	return Style_name
}

func (Pict *Picture) SetSize(ht int, wdth int) {
	styletextproperty, _ := xmlDB.GetNode(Pict.Note.Content, Pict.NodeId, "draw:frame")
	xmlDB.UpdateAttributevalue(Pict.Note.Content, styletextproperty[0], "svg:height", strconv.Itoa(ht)+"cm")
	xmlDB.UpdateAttributevalue(Pict.Note.Content, styletextproperty[0], "svg:width", strconv.Itoa(wdth)+"cm")
}


//func (Pict *Picture) AddListStyleName(value string) {
//	xmlDB.UpdateAttributevalue(Pict.Note.Content, Pict.StyleId, "style:list-style-name", value)
//}
