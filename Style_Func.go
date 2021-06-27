package odf

import (
	//"fmt"

	"github.com/LIJUCHACKO/XmlDB"
)

func MergeStyles(style1 *xmlDB.Database, style2 *xmlDB.Database) {
	//import styles in style2 into style1
	office_automatic_styles1, _ := xmlDB.GetNode(style1, 0, "office:automatic-styles")

	office_automatic_styles, _ := xmlDB.GetNode(style2, 0, "office:automatic-styles")
	office_automatic_styleEach := xmlDB.ChildNodes(style2, office_automatic_styles[0])
	for _, style := range office_automatic_styleEach {
		identified, _ := xmlDB.GetNode(style1, office_automatic_styles1[0], "<x>[style:name=\""+xmlDB.GetNodeAttribute(style2, style, "style:name")+"\"]")
		//Will not replace if the style is already present
		if len(identified) == 0 {
			xmlDB.InserSubNode(style1, office_automatic_styles1[0], xmlDB.GetNodeContentRaw(style2, style))
		}
	}

	office_styles1, _ := xmlDB.GetNode(style1, 0, "office:styles")

	office_styles, _ := xmlDB.GetNode(style2, 0, "office:styles")
	if len(office_styles) > 0 {
		office_stylesEach := xmlDB.ChildNodes(style2, office_styles[0])
		for _, style := range office_stylesEach {
			identified, _ := xmlDB.GetNode(style1, office_styles1[0], "<x>[style:name=\""+xmlDB.GetNodeAttribute(style2, style, "style:name")+"\"]")
			//Will not replace if the style is already present
			if len(identified) == 0 {
				xmlDB.InserSubNode(style1, office_styles1[0], xmlDB.GetNodeContentRaw(style2, style))
			}
		}
	}

}
