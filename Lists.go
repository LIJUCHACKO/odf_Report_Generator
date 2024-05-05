package odf

import (
	//"fmt"

	"math/rand"
	"strconv"
	"strings"

	"github.com/LIJUCHACKO/XmlDB"
)

type List struct {
	NodeId      int
	Note        *Notes
	ListStyleId int
}
var randnoList int=0
func (Note *Notes) NewNumberedList(Style string) *List {
	var List *List = new(List)
	List.Note = Note
	ListStyle_name := "PLS" + strconv.Itoa(rand.Intn(100)+randnoList)
	randnoList=randnoList+1
	if len(strings.TrimSpace(Style)) > 0 {
		ListStyle_name = strings.TrimSpace(Style)
	}
	//fmt.Println("Style name=" + ListStyle_name)

	//office_Text
	List.NodeId = Note.WritetoScratchpad("<text:list xml:id=\"list5252122" + strconv.Itoa(rand.Intn(100)+randnoList) + "\" text:style-name=\"" + ListStyle_name + "\"/>")
	//	List.Office_Text = xml_content
    randnoList=randnoList+1
	if len(strings.TrimSpace(Style)) == 0 {
		//office_style
		styletext := `<styles>
	         <office:styles>
				<style:style style:name="Standard" style:family="paragraph" style:class="text"/>
		        <style:style style:name="Text_20_body" style:display-name="Text body" style:family="paragraph" style:parent-style-name="Standard" style:class="text">
		            <style:paragraph-properties fo:margin-top="0cm" fo:margin-bottom="0.247cm" loext:contextual-spacing="false" fo:line-height="120%"/>
		        </style:style>
		        <style:style style:name="Numbering_20_Symbols" style:display-name="Numbering Symbols" style:family="text"/>
			</office:styles>
			<office:automatic-styles>
				<text:list-style style:name="` + ListStyle_name + `">
		            <text:list-level-style-number text:level="1" text:style-name="Numbering_20_Symbols" style:num-suffix="." style:num-format="1">
		                <style:list-level-properties text:list-level-position-and-space-mode="label-alignment">
		                    <style:list-level-label-alignment text:label-followed-by="listtab" text:list-tab-stop-position="1.27cm" fo:text-indent="-0.635cm" fo:margin-left="1.27cm"/>
		                </style:list-level-properties>
		            </text:list-level-style-number>
		            <text:list-level-style-number text:level="2" text:style-name="Numbering_20_Symbols" style:num-suffix="." style:num-format="1">
		                <style:list-level-properties text:list-level-position-and-space-mode="label-alignment">
		                    <style:list-level-label-alignment text:label-followed-by="listtab" text:list-tab-stop-position="1.905cm" fo:text-indent="-0.635cm" fo:margin-left="1.905cm"/>
		                </style:list-level-properties>
		            </text:list-level-style-number>
		            <text:list-level-style-number text:level="3" text:style-name="Numbering_20_Symbols" style:num-suffix="." style:num-format="1">
		                <style:list-level-properties text:list-level-position-and-space-mode="label-alignment">
		                    <style:list-level-label-alignment text:label-followed-by="listtab" text:list-tab-stop-position="2.54cm" fo:text-indent="-0.635cm" fo:margin-left="2.54cm"/>
		                </style:list-level-properties>
		            </text:list-level-style-number>
		            <text:list-level-style-number text:level="4" text:style-name="Numbering_20_Symbols" style:num-suffix="." style:num-format="1">
		                <style:list-level-properties text:list-level-position-and-space-mode="label-alignment">
		                    <style:list-level-label-alignment text:label-followed-by="listtab" text:list-tab-stop-position="3.175cm" fo:text-indent="-0.635cm" fo:margin-left="3.175cm"/>
		                </style:list-level-properties>
		            </text:list-level-style-number>
		            <text:list-level-style-number text:level="5" text:style-name="Numbering_20_Symbols" style:num-suffix="." style:num-format="1">
		                <style:list-level-properties text:list-level-position-and-space-mode="label-alignment">
		                    <style:list-level-label-alignment text:label-followed-by="listtab" text:list-tab-stop-position="3.81cm" fo:text-indent="-0.635cm" fo:margin-left="3.81cm"/>
		                </style:list-level-properties>
		            </text:list-level-style-number>
		            <text:list-level-style-number text:level="6" text:style-name="Numbering_20_Symbols" style:num-suffix="." style:num-format="1">
		                <style:list-level-properties text:list-level-position-and-space-mode="label-alignment">
		                    <style:list-level-label-alignment text:label-followed-by="listtab" text:list-tab-stop-position="4.445cm" fo:text-indent="-0.635cm" fo:margin-left="4.445cm"/>
		                </style:list-level-properties>
		            </text:list-level-style-number>
		            <text:list-level-style-number text:level="7" text:style-name="Numbering_20_Symbols" style:num-suffix="." style:num-format="1">
		                <style:list-level-properties text:list-level-position-and-space-mode="label-alignment">
		                    <style:list-level-label-alignment text:label-followed-by="listtab" text:list-tab-stop-position="5.08cm" fo:text-indent="-0.635cm" fo:margin-left="5.08cm"/>
		                </style:list-level-properties>
		            </text:list-level-style-number>
		            <text:list-level-style-number text:level="8" text:style-name="Numbering_20_Symbols" style:num-suffix="." style:num-format="1">
		                <style:list-level-properties text:list-level-position-and-space-mode="label-alignment">
		                    <style:list-level-label-alignment text:label-followed-by="listtab" text:list-tab-stop-position="5.715cm" fo:text-indent="-0.635cm" fo:margin-left="5.715cm"/>
		                </style:list-level-properties>
		            </text:list-level-style-number>
		            <text:list-level-style-number text:level="9" text:style-name="Numbering_20_Symbols" style:num-suffix="." style:num-format="1">
		                <style:list-level-properties text:list-level-position-and-space-mode="label-alignment">
		                    <style:list-level-label-alignment text:label-followed-by="listtab" text:list-tab-stop-position="6.35cm" fo:text-indent="-0.635cm" fo:margin-left="6.35cm"/>
		                </style:list-level-properties>
		            </text:list-level-style-number>
		            <text:list-level-style-number text:level="10" text:style-name="Numbering_20_Symbols" style:num-suffix="." style:num-format="1">
		                <style:list-level-properties text:list-level-position-and-space-mode="label-alignment">
		                    <style:list-level-label-alignment text:label-followed-by="listtab" text:list-tab-stop-position="6.985cm" fo:text-indent="-0.635cm" fo:margin-left="6.985cm"/>
		                </style:list-level-properties>
		            </text:list-level-style-number>
		        </text:list-style>

			</office:automatic-styles>
		   </styles>`

		id := Note.WritetoScratchpad(styletext)
		Note.IncludeStyle(id)
	}
	//fmt.Println(xmlDB.GetNodeContents(Note.Content, Note.Officeautostyleid))

	StyleNodeid, _ := xmlDB.GetNode(Note.Content, Note.Officeautostyleid, "text:list-style[style:name=\""+ListStyle_name+"\"]")
	List.ListStyleId = StyleNodeid[0]

	return List
}
func (Note *Notes) NewBulletinList(Style string) *List {
	var List *List = new(List)
	List.Note = Note
	ListStyle_name := "PLS" + strconv.Itoa(rand.Intn(100))
	if len(strings.TrimSpace(Style)) > 0 {
		ListStyle_name = strings.TrimSpace(Style)
	}
	//fmt.Println("Style name=" + ListStyle_name)

	//office_Text
	List.NodeId = Note.WritetoScratchpad("<text:list xml:id=\"list5252122" + strconv.Itoa(rand.Intn(100)) + "\" text:style-name=\"" + ListStyle_name + "\"/>")
	//	List.Office_Text = xml_content

	if len(strings.TrimSpace(Style)) == 0 {
		//office_style
		styletext := `<styles>
	         <office:styles>
				<style:style style:name="Standard" style:family="paragraph" style:class="text"/>
		        <style:style style:name="Text_20_body" style:display-name="Text body" style:family="paragraph" style:parent-style-name="Standard" style:class="text">
		            <style:paragraph-properties fo:margin-top="0cm" fo:margin-bottom="0.247cm" loext:contextual-spacing="false" fo:line-height="120%"/>
		        </style:style>
		        <style:style style:name="Numbering_20_Symbols" style:display-name="Numbering Symbols" style:family="text"/>
			</office:styles>
			<office:automatic-styles>
				<text:list-style style:name="` + ListStyle_name + `">
		            <text:list-level-style-bullet text:level="1" text:style-name="Bullet_20_Symbols" text:bullet-char="•">
		                <style:list-level-properties text:list-level-position-and-space-mode="label-alignment">
		                    <style:list-level-label-alignment text:label-followed-by="listtab" text:list-tab-stop-position="1.27cm" fo:text-indent="-0.635cm" fo:margin-left="1.27cm"/>
		                </style:list-level-properties>
		            </text:list-level-style-bullet>
		            <text:list-level-style-bullet text:level="2" text:style-name="Bullet_20_Symbols" text:bullet-char="◦">
		                <style:list-level-properties text:list-level-position-and-space-mode="label-alignment">
		                    <style:list-level-label-alignment text:label-followed-by="listtab" text:list-tab-stop-position="1.905cm" fo:text-indent="-0.635cm" fo:margin-left="1.905cm"/>
		                </style:list-level-properties>
		            </text:list-level-style-bullet>
		            <text:list-level-style-bullet text:level="3" text:style-name="Bullet_20_Symbols" text:bullet-char="▪">
		                <style:list-level-properties text:list-level-position-and-space-mode="label-alignment">
		                    <style:list-level-label-alignment text:label-followed-by="listtab" text:list-tab-stop-position="2.54cm" fo:text-indent="-0.635cm" fo:margin-left="2.54cm"/>
		                </style:list-level-properties>
		            </text:list-level-style-bullet>
		            <text:list-level-style-bullet text:level="4" text:style-name="Bullet_20_Symbols" text:bullet-char="•">
		                <style:list-level-properties text:list-level-position-and-space-mode="label-alignment">
		                    <style:list-level-label-alignment text:label-followed-by="listtab" text:list-tab-stop-position="3.175cm" fo:text-indent="-0.635cm" fo:margin-left="3.175cm"/>
		                </style:list-level-properties>
		            </text:list-level-style-bullet>
		            <text:list-level-style-bullet text:level="5" text:style-name="Bullet_20_Symbols" text:bullet-char="◦">
		                <style:list-level-properties text:list-level-position-and-space-mode="label-alignment">
		                    <style:list-level-label-alignment text:label-followed-by="listtab" text:list-tab-stop-position="3.81cm" fo:text-indent="-0.635cm" fo:margin-left="3.81cm"/>
		                </style:list-level-properties>
		            </text:list-level-style-bullet>
		            <text:list-level-style-bullet text:level="6" text:style-name="Bullet_20_Symbols" text:bullet-char="▪">
		                <style:list-level-properties text:list-level-position-and-space-mode="label-alignment">
		                    <style:list-level-label-alignment text:label-followed-by="listtab" text:list-tab-stop-position="4.445cm" fo:text-indent="-0.635cm" fo:margin-left="4.445cm"/>
		                </style:list-level-properties>
		            </text:list-level-style-bullet>
		            <text:list-level-style-bullet text:level="7" text:style-name="Bullet_20_Symbols" text:bullet-char="•">
		                <style:list-level-properties text:list-level-position-and-space-mode="label-alignment">
		                    <style:list-level-label-alignment text:label-followed-by="listtab" text:list-tab-stop-position="5.08cm" fo:text-indent="-0.635cm" fo:margin-left="5.08cm"/>
		                </style:list-level-properties>
		            </text:list-level-style-bullet>
		            <text:list-level-style-bullet text:level="8" text:style-name="Bullet_20_Symbols" text:bullet-char="◦">
		                <style:list-level-properties text:list-level-position-and-space-mode="label-alignment">
		                    <style:list-level-label-alignment text:label-followed-by="listtab" text:list-tab-stop-position="5.715cm" fo:text-indent="-0.635cm" fo:margin-left="5.715cm"/>
		                </style:list-level-properties>
		            </text:list-level-style-bullet>
		            <text:list-level-style-bullet text:level="9" text:style-name="Bullet_20_Symbols" text:bullet-char="▪">
		                <style:list-level-properties text:list-level-position-and-space-mode="label-alignment">
		                    <style:list-level-label-alignment text:label-followed-by="listtab" text:list-tab-stop-position="6.35cm" fo:text-indent="-0.635cm" fo:margin-left="6.35cm"/>
		                </style:list-level-properties>
		            </text:list-level-style-bullet>
		            <text:list-level-style-bullet text:level="10" text:style-name="Bullet_20_Symbols" text:bullet-char="•">
		                <style:list-level-properties text:list-level-position-and-space-mode="label-alignment">
		                    <style:list-level-label-alignment text:label-followed-by="listtab" text:list-tab-stop-position="6.985cm" fo:text-indent="-0.635cm" fo:margin-left="6.985cm"/>
		                </style:list-level-properties>
		            </text:list-level-style-bullet>
		        </text:list-style>
			</office:automatic-styles>
		   </styles>`

		id := Note.WritetoScratchpad(styletext)
		Note.IncludeStyle(id)
	}
	//fmt.Println(xmlDB.GetNodeContents(Note.Content, Note.Officeautostyleid))

	StyleNodeid, _ := xmlDB.GetNode(Note.Content, Note.Officeautostyleid, "text:list-style[style:name=\""+ListStyle_name+"\"]")
	List.ListStyleId = StyleNodeid[0]

	return List
}
func (List *List) Style() string {
	ListStyle_name := xmlDB.GetNodeAttribute(List.Note.Content, List.ListStyleId, "style:name")
	return ListStyle_name
}
func (List *List) toBulletinList() {
	ListStyle_name := List.Style()
	styletext := `<text:list-style style:name="` + ListStyle_name + `">
            <text:list-level-style-bullet text:level="1" text:style-name="Bullet_20_Symbols" text:bullet-char="•">
                <style:list-level-properties text:list-level-position-and-space-mode="label-alignment">
                    <style:list-level-label-alignment text:label-followed-by="listtab" text:list-tab-stop-position="1.27cm" fo:text-indent="-0.635cm" fo:margin-left="1.27cm"/>
                </style:list-level-properties>
            </text:list-level-style-bullet>
            <text:list-level-style-bullet text:level="2" text:style-name="Bullet_20_Symbols" text:bullet-char="◦">
                <style:list-level-properties text:list-level-position-and-space-mode="label-alignment">
                    <style:list-level-label-alignment text:label-followed-by="listtab" text:list-tab-stop-position="1.905cm" fo:text-indent="-0.635cm" fo:margin-left="1.905cm"/>
                </style:list-level-properties>
            </text:list-level-style-bullet>
            <text:list-level-style-bullet text:level="3" text:style-name="Bullet_20_Symbols" text:bullet-char="▪">
                <style:list-level-properties text:list-level-position-and-space-mode="label-alignment">
                    <style:list-level-label-alignment text:label-followed-by="listtab" text:list-tab-stop-position="2.54cm" fo:text-indent="-0.635cm" fo:margin-left="2.54cm"/>
                </style:list-level-properties>
            </text:list-level-style-bullet>
            <text:list-level-style-bullet text:level="4" text:style-name="Bullet_20_Symbols" text:bullet-char="•">
                <style:list-level-properties text:list-level-position-and-space-mode="label-alignment">
                    <style:list-level-label-alignment text:label-followed-by="listtab" text:list-tab-stop-position="3.175cm" fo:text-indent="-0.635cm" fo:margin-left="3.175cm"/>
                </style:list-level-properties>
            </text:list-level-style-bullet>
            <text:list-level-style-bullet text:level="5" text:style-name="Bullet_20_Symbols" text:bullet-char="◦">
                <style:list-level-properties text:list-level-position-and-space-mode="label-alignment">
                    <style:list-level-label-alignment text:label-followed-by="listtab" text:list-tab-stop-position="3.81cm" fo:text-indent="-0.635cm" fo:margin-left="3.81cm"/>
                </style:list-level-properties>
            </text:list-level-style-bullet>
            <text:list-level-style-bullet text:level="6" text:style-name="Bullet_20_Symbols" text:bullet-char="▪">
                <style:list-level-properties text:list-level-position-and-space-mode="label-alignment">
                    <style:list-level-label-alignment text:label-followed-by="listtab" text:list-tab-stop-position="4.445cm" fo:text-indent="-0.635cm" fo:margin-left="4.445cm"/>
                </style:list-level-properties>
            </text:list-level-style-bullet>
            <text:list-level-style-bullet text:level="7" text:style-name="Bullet_20_Symbols" text:bullet-char="•">
                <style:list-level-properties text:list-level-position-and-space-mode="label-alignment">
                    <style:list-level-label-alignment text:label-followed-by="listtab" text:list-tab-stop-position="5.08cm" fo:text-indent="-0.635cm" fo:margin-left="5.08cm"/>
                </style:list-level-properties>
            </text:list-level-style-bullet>
            <text:list-level-style-bullet text:level="8" text:style-name="Bullet_20_Symbols" text:bullet-char="◦">
                <style:list-level-properties text:list-level-position-and-space-mode="label-alignment">
                    <style:list-level-label-alignment text:label-followed-by="listtab" text:list-tab-stop-position="5.715cm" fo:text-indent="-0.635cm" fo:margin-left="5.715cm"/>
                </style:list-level-properties>
            </text:list-level-style-bullet>
            <text:list-level-style-bullet text:level="9" text:style-name="Bullet_20_Symbols" text:bullet-char="▪">
                <style:list-level-properties text:list-level-position-and-space-mode="label-alignment">
                    <style:list-level-label-alignment text:label-followed-by="listtab" text:list-tab-stop-position="6.35cm" fo:text-indent="-0.635cm" fo:margin-left="6.35cm"/>
                </style:list-level-properties>
            </text:list-level-style-bullet>
            <text:list-level-style-bullet text:level="10" text:style-name="Bullet_20_Symbols" text:bullet-char="•">
                <style:list-level-properties text:list-level-position-and-space-mode="label-alignment">
                    <style:list-level-label-alignment text:label-followed-by="listtab" text:list-tab-stop-position="6.985cm" fo:text-indent="-0.635cm" fo:margin-left="6.985cm"/>
                </style:list-level-properties>
            </text:list-level-style-bullet>
        </text:list-style>`

	xmlDB.ReplaceNode(List.Note.Content, List.ListStyleId, styletext)

}

func (List *List) AddItemPara(style string) *Paragraph {
	para := List.Note.NewParagraph(style)
	para.AddListStyleName(List.Style())
	newids, _ := xmlDB.InserSubNode(List.Note.Content, List.NodeId, "<text:list-item/>")
	xmlDB.CutPasteAsSubNode(List.Note.Content, newids[0], para.NodeId)
	return para
}
