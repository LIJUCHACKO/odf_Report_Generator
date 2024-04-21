package odf

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	//"math/rand"
	"os"
	//"strconv"
	"archive/zip"
	"io"

	"github.com/LIJUCHACKO/XmlDB"
	"path/filepath"
)

type Odt struct {
	Content *xmlDB.Database
}

func (Doc *Odt) ImportStyles(Note *Notes) {
	style1 := Doc.Content
	style2 := Note.Content
	//import styles in style2 into style1
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
func (Doc *Odt) ImportAutoStyles(Note *Notes) {
	style1 := Doc.Content
	style2 := Note.Content
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

}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
func (Doc *Odt) CreateMarkerNodes() { //newmodified
	fmt.Printf("\nCreateMarkerNodes started\n")
	docxml := Doc.Content

	paraNodeids, _ := xmlDB.GetNode(docxml, 0, "office:body/office:text/../text:p")
	for _, NodeId := range paraNodeids {
		//fmt.Println(NodeId)
		//fmt.Printf("\n==============================")
		//fmt.Println(xmlDB.GetNodeContents(docxml, NodeId))
		//xmlDB.NodeDebug(docxml, NodeId)
		markerpresent := true
		for markerpresent {
			items := xmlDB.ChildNodes(docxml, NodeId)
			Output := ""
			children_boundaries := []int{}
			if xmlDB.IslowestNode(docxml, NodeId) {
				childcontent := xmlDB.GetNodeValue(docxml, NodeId)
				Output = childcontent
				children_boundaries = append(children_boundaries, len(childcontent))
			} else {
				//fmt.Println(items)
				for _, item := range items {
					if xmlDB.IslowestNode(docxml, item) {
						childcontent := xmlDB.GetNodeValue(docxml, item)
						Output = Output + childcontent // textspan
						children_boundaries = append(children_boundaries, len(childcontent))
					} else {
						items2 := xmlDB.ChildNodes(docxml, item)
						for _, item2 := range items2 {
							childcontent := xmlDB.GetNodeValue(docxml, item2)
							Output = Output + childcontent //child inside textspan
							children_boundaries = append(children_boundaries, len(childcontent))
						}
					}

				}
			}
			//fmt.Printf("\n Text- %s\n", Output)
			//fmt.Println(children_boundaries)
			re := regexp.MustCompile("(!%([A-Za-z0-9]{1,})%!)")
			res := re.FindAllStringSubmatch(Output, -1)
			if len(res) > 0 {
				//for i := range res {
				//like Java: match.group(1), match.gropu(2), etc
				marker := res[0][1]
				fmt.Printf("\n Identified marker- %s\n", marker)
				marker_start := strings.Index(Output, marker)

				if marker_start >= 0 {
					started := false
					StartCase := -1
					StopCase := -1
					startid := -1
					parentstartid := -1
					// startindex := -1
					// endindex := -1
					parentendid := -1
					endid := -1
					ind := 0
					start := 0
					items := xmlDB.ChildNodes(docxml, NodeId)
					if xmlDB.IslowestNode(docxml, NodeId) {
						//only one  node
						StartCase = 1
						StopCase = 1
						startid = NodeId
						endid = NodeId
						start = start + children_boundaries[ind]
						ind++
					} else {
						for _, item := range items {
							//fmt.Println(xmlDB.GetNodeContents(docxml, item))
							if xmlDB.IslowestNode(docxml, item) {
								childcontent := xmlDB.GetNodeValue(docxml, item)
								//fmt.Println("childcontent=")
								//fmt.Println(childcontent)
								if marker_start >= start && marker_start < start+children_boundaries[ind] && children_boundaries[ind] != 0 {
									// <> ffd!%marke..... </>
									//
									//fmt.Printf("\nmarker_start %d  start %d  children_boundaries[ind] %d ind %d\n", marker_start, start, children_boundaries[ind], ind)
									started = true
									startid = item
									StartCase = 2
									if start+children_boundaries[ind] >= marker_start+len(marker) {
										// <> ffd!%marker%!kij </>
										StopCase = 2
										//fmt.Println("\n inside if=node content="+xmlDB.GetNodeContentRaw(docxml, item))
										//fmt.Println("\ninside if="+childcontent[0:(marker_start-start)])
										xmlDB.UpdateNodevalue(docxml, item, childcontent[0:(marker_start-start)])//==>  <> ffd</>
										newNode:=xmlDB.GetNodeContentRaw(docxml, item)

										newids, _ := xmlDB.AppendAfterNode(docxml, item,newNode )//just to maintaine node name
										//==>  <> ffd</><>ffd!%marker%!kij</>
										xmlDB.UpdateNodevalue(docxml, newids[0], childcontent[(marker_start+len(marker)-start):])
										//==> <> ffd</><>kij</>
										endid = newids[0]
									} else {
										// <> ffd!%mark</>
										xmlDB.UpdateNodevalue(docxml, item, childcontent[0:(marker_start-start)])//==> <> ffd</>

									}

								} else if started {
									//fmt.Printf("\nmarker_start %d  start %d  children_boundaries[ind] %d ind %d\n", marker_start, start, children_boundaries[ind], ind)
									if start < marker_start+len(marker) {
										if start+children_boundaries[ind] >= marker_start+len(marker) {
											// <>er%!kij</>
											endid = item
											StopCase = 2
											xmlDB.UpdateNodevalue(docxml, item, childcontent[(marker_start+len(marker)-start):])
											// <>kij</>
										} else {
											// <>arke</>
											xmlDB.RemoveNode(docxml, item)
										}

									} else {
										started = false
									}
								}
								start = start + children_boundaries[ind]
								ind++
							} else {
								items2 := xmlDB.ChildNodes(docxml, item)
								for _, item2 := range items2 {
									//fmt.Println(xmlDB.GetNodeContents(docxml, item2))
									childcontent := xmlDB.GetNodeValue(docxml, item2)
									//fmt.Println(childcontent)
									if marker_start >= start && marker_start < start+children_boundaries[ind] && children_boundaries[ind] != 0 {
										//fmt.Printf("\nmarker_start %d  start %d  children_boundaries[ind] %d ind %d\n", marker_start, start, children_boundaries[ind], ind)
										started = true
										startid = item2
										parentstartid = item
										StartCase = 3
										if start+children_boundaries[ind] >= marker_start+len(marker) {
											StopCase = 3
											xmlDB.UpdateNodevalue(docxml, item2, childcontent[0:(marker_start-start)])
											newids, _ := xmlDB.AppendAfterNode(docxml, item2, xmlDB.GetNodeContentRaw(docxml, item2))
											xmlDB.UpdateNodevalue(docxml, newids[0], childcontent[(marker_start+len(marker)-start):])
											endid = newids[0]
										} else {
											xmlDB.UpdateNodevalue(docxml, item2, childcontent[0:(marker_start-start)])
										}

									} else if started {
										if start < marker_start+len(marker) {
											//fmt.Printf("\nmarker_start %d  start %d  children_boundaries[ind] %d ind %d\n", marker_start, start, children_boundaries[ind], ind)
											if start+children_boundaries[ind] >= marker_start+len(marker) {
												endid = item2
												parentendid = item
												StopCase = 3
												xmlDB.UpdateNodevalue(docxml, item2, childcontent[(marker_start+len(marker)-start):])
											} else {
												xmlDB.RemoveNode(docxml, item2)
											}
										} else {
											started = false
										}
									}
									start = start + children_boundaries[ind]
									ind++
								}

							}
						}
					}

					//if case 3 split into 2
					if StartCase == 3 && StopCase == 3 && parentstartid == parentendid {
						if startid == endid {

						} else {
							newids, _ := xmlDB.AppendAfterNode(docxml, parentstartid, xmlDB.GetNodeContentRaw(docxml, parentstartid))
							items2 := xmlDB.ChildNodes(docxml, parentstartid)
							newitems2 := xmlDB.ChildNodes(docxml, newids[0])
							started := false
							for i, item := range items2 {

								if !started {
									xmlDB.RemoveNode(docxml, newitems2[i])
								} else {
									xmlDB.RemoveNode(docxml, item)
								}
								if item == startid {
									started = true
								}
							}
						}

					} else if StartCase == 3 {

					} else if StopCase == 3 {

					}
					if strings.TrimSpace(Output) == marker {
						fmt.Println("content contains only marker")
						xmlDB.ReplaceNode(docxml, NodeId, "<marker type=\"text:p\" name=\""+marker+"\"/>")
					} else if StartCase == 3 || StartCase == 2 {
						if StartCase == 3 {
							startid = parentstartid
						}
						//insert marker after startid
						xmlDB.AppendAfterNode(docxml, startid, "<marker type=\"text:span\" name=\""+marker+"\"/>")
					} else if StartCase == 1 && StopCase == 1 && startid == endid {
						childcontent := xmlDB.GetNodeValue(docxml, startid)
						if strings.TrimSpace(childcontent) == marker {
							xmlDB.ReplaceNode(docxml, startid, "<marker type=\"text:p\" name=\""+marker+"\"/>")
						}
					}
					// fmt.Println("After subprocessing")
					// fmt.Println(xmlDB.GetNodeContents(docxml, NodeId))
				}
			} else {
				markerpresent = false
			}
		}
		//	fmt.Println("After processing")
		//	fmt.Println(xmlDB.GetNodeContents(docxml, NodeId))
	}
	fmt.Printf("\nCreateMarkerNodes over\n")

}

func (Doc *Odt) ReplaceMarkers(Note *Notes) {

	Doc.CreateMarkerNodes()

	articles, _ := xmlDB.GetNode(Note.Content, Note.scratchpadid, "article")
	fmt.Println("ReplaceMarkers ")
	fmt.Println(articles)
	for _, article := range articles {
		marker := xmlDB.GetNodeAttribute(Note.Content, article, "name")
		fmt.Printf("\n marker %s", marker)
		markerIds, _ := xmlDB.GetNode(Doc.Content, 0, "office:body/office:text/../marker[name=\""+marker+"\"]")
		fmt.Println(markerIds)
		for _, id := range markerIds {
			markertype := xmlDB.GetNodeAttribute(Doc.Content, id, "type")
			// fmt.Printf("\n%d\n", id)
			// xmlDB.NodeDebug(Doc.Content, id)
			items := xmlDB.ChildNodes(Note.Content, article)
			//fmt.Printf("\n Replacing %d", id)
			//fmt.Println(xmlDB.GetNodeContentRaw(Note.Content, article))
			fmt.Println(items)
			previousitemid := -1
			for index, item := range items {
				//fmt.Println(xmlDB.GetNodeContentRaw(Note.Content, item))
				if index == 0 {
					newids, _ := xmlDB.ReplaceNode(Doc.Content, id, xmlDB.GetNodeContentRaw(Note.Content, item))
					previousitemid = newids[0]
				} else {
					newids, _ := xmlDB.AppendAfterNode(Doc.Content, previousitemid, xmlDB.GetNodeContentRaw(Note.Content, item))
					previousitemid = newids[0]
				}

				nodeName := xmlDB.GetNodeName(Note.Content, item)
				if markertype == "text:span" {
					if nodeName != "text:span" {
						fmt.Printf("\n Article should be of type %s(not %s)\n", markertype, nodeName)
					}
				}

			}

		}
	}

}

func ProcessOdtfile(Inputfile string, Outputfile string, Note *Notes) {
	targetfile, err := os.Create(Outputfile)
	defer targetfile.Close()
	targetzipwriter := zip.NewWriter(targetfile)
	defer targetzipwriter.Close()
	r, err := zip.OpenReader(Inputfile)
	if err != nil {
		fmt.Println(err)
	}
	defer r.Close()

	//adding additional files
	for ind, file_path := range Note.OrginalPictureFiles {
		f1, err := os.Open(file_path)
		if err != nil {
			panic(err)
		}
		defer f1.Close()

		fmt.Println("writing first file to archive...")
		//file_name := filepath.Base(file_path)
		//extension := filepath.Ext(file_path)
		filenameinside := Note.NewPictureFiles[ind]//"Pictures/"+strconv.Itoa(rand.Intn(10000000))+extension
		w1, err := targetzipwriter.Create(filenameinside)
		if err != nil {
			panic(err)
		}
		if _, err := io.Copy(w1, f1); err != nil {
			panic(err)
		}
	}
	///

	for _, f := range r.File {
		//fmt.Println(f.Name)
		fileToZip, _ := f.Open()
		defer fileToZip.Close()
		header, _ := zip.FileInfoHeader(f.FileInfo())
		header.Name = f.Name
		header.Method = zip.Deflate
		writer, _ := targetzipwriter.CreateHeader(header)

		if f.Name == "content.xml" {
			xmlfile, _ := ioutil.ReadAll(fileToZip)
			xmlstring := string(xmlfile)

			//xmlline = processString(xmlline)
			var DB *xmlDB.Database = new(xmlDB.Database)
			DB.MaxNooflines = 9999999

			xmllines := strings.Split(xmlstring, "\n")
			xmlDB.Load_dbcontent(DB, xmllines)
			DB.Debug_enabled = false
			DB.Libreofficemod= true
			var Doc *Odt = new(Odt)
			Doc.Content = DB
			Doc.ReplaceMarkers(Note)
			//MERGE STYLES
			Doc.ImportAutoStyles(Note)
			_, _ = writer.Write([]byte(xmlDB.Dump_DB(Doc.Content)))
		} else if f.Name == "styles.xml" {
			xmlfile, _ := ioutil.ReadAll(fileToZip)
			xmlstring := string(xmlfile)

			//xmlline = processString(xmlline)
			var DB *xmlDB.Database = new(xmlDB.Database)
			DB.MaxNooflines = 9999999
			xmllines := strings.Split(xmlstring, "\n")
			xmlDB.Load_dbcontent(DB, xmllines)
			DB.Debug_enabled = false
			DB.Libreofficemod= true
			var Doc *Odt = new(Odt)
			Doc.Content = DB
			//MERGE STYLES
			Doc.ImportStyles(Note)
			_, _ = writer.Write([]byte(xmlDB.Dump_DB(Doc.Content)))
		} else if f.Name == "META-INF/manifest.xml" {
			xmlfile, _ := ioutil.ReadAll(fileToZip)
			xmlstring := string(xmlfile)
			var DB *xmlDB.Database = new(xmlDB.Database)
			DB.MaxNooflines = 999
			xmllines := strings.Split(xmlstring, "\n")
			xmlDB.Load_dbcontent(DB, xmllines)
			DB.Debug_enabled = false
			DB.Libreofficemod= false
			for _, file_path := range Note.NewPictureFiles {
				extension := filepath.Ext(file_path)
				extension=strings.ReplaceAll(extension, ".", "" )
				xmlDB.InserSubNode(DB, 0,`<manifest:file-entry manifest:full-path="`+file_path+`" manifest:media-type="image/`+ extension+`"/>`)
			}
			_, _ = writer.Write([]byte(xmlDB.Dump_DB(DB)))
	    } else {

			io.Copy(writer, fileToZip)
		}

		fileToZip.Close()

	}

	targetzipwriter.Close()
	targetfile.Close()
}
