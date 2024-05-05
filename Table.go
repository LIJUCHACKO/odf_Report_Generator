package odf

import (
	//"fmt"

	"math/rand"
	"strconv"
	"strings"

	"github.com/LIJUCHACKO/XmlDB"
)

type Table struct {
	cells         [][]int
	colms         int
	rows          int
	NodeId        int
	TableStyleId  int
	ColumnStyleId int
	CellStyleId   int
	Note          *Notes
}
var randnoTable int =0
func (Note *Notes) NewTable(noofcols int, noofrows int, Style string) *Table {
	var Table *Table = new(Table)
	//fmt.Println(rand.Intn(100))
	Style_name := "Tbl" + strconv.Itoa(rand.Intn(100)+randnoTable)
	randnoTable=randnoTable+1
	Table.Note = Note
	if len(strings.TrimSpace(Style)) > 0 {
		Style_name = strings.TrimSpace(Style)
	}
	Table.colms = noofcols
	Table.rows = noofrows
	//office_Text
	//Para.NodeId = Note.WritetoScratchpad("<text:p text:style-name=\"" + Style_name + "\"/>")
	Table.NodeId = Note.WritetoScratchpad(` 
    	<table:table table:name="` + Style_name + `" table:style-name="` + Style_name + `">
        <table:table-column table:style-name="` + Style_name + `_COL" table:number-columns-repeated="` + strconv.Itoa(noofcols) + `"/>
		</table:table>
	 `)

	for i := 0; i < noofrows; i++ {
		row, _ := xmlDB.InserSubNode(Table.Note.Content, Table.NodeId, `<table:table-row table:style-name="`+Style_name+`_ROW"/>`)
		rowcells := []int{}
		for j := 0; j < noofcols; j++ {
			cell, _ := xmlDB.InserSubNode(Table.Note.Content, row[0], `<table:table-cell table:style-name="`+Style_name+`_CELL" office:value-type="string"/>`)
			rowcells = append(rowcells, cell[0])
		}
		Table.cells = append(Table.cells, rowcells)
	}
	if len(strings.TrimSpace(Style)) == 0 {
		//office_style
		styletext := `<styles><office:styles>
				<style:default-style style:family="table">
					<style:table-properties table:border-model="collapsing"/>
				</style:default-style>
				<style:default-style style:family="table-row">
					<style:table-row-properties fo:keep-together="auto"/>
				</style:default-style>
			</office:styles>
			<office:automatic-styles>
				<style:style style:name="` + Style_name + `" style:family="table">
				    <style:table-properties style:width="17cm" table:align="margins"/>
				</style:style>
				<style:style style:name="` + Style_name + `_COL" style:family="table-column">
				    <style:table-column-properties style:column-width="2.835cm" style:rel-column-width="10925*"/>
				</style:style>
				<style:style style:name="` + Style_name + `_CELL" style:family="table-cell">
				    <style:table-cell-properties fo:padding="0.097cm" fo:border-left="0.05pt solid #000000" fo:border-right="none" fo:border-top="0.05pt solid #000000" fo:border-bottom="0.05pt solid #000000"/>
				</style:style>

			</office:automatic-styles></styles>`
		id := Note.WritetoScratchpad(styletext)
		Note.IncludeStyle(id)
	}
	return Table

}

func (Table *Table) AddItemPara(rowno int, colno int, style string) *Paragraph {
	para := Table.Note.NewParagraph(style)
	xmlDB.CutPasteAsSubNode(Table.Note.Content, Table.cells[rowno][colno], para.NodeId)
	return para
}
