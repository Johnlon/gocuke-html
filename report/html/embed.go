package html

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"strconv"
	"strings"

	"gitlab.com/rodrigoodhin/gocure/helpers"
	"gitlab.com/rodrigoodhin/gocure/models"
)

func (d *Data) embedFiles(embed models.Embed, isFirst bool, index int, modalID string) (embedHTML string, err error) {

	rowClass := ""
	if isFirst {
		rowClass = "embed-row-top"
	}

	mediaContent := ""
	copyOption := ""
	contentID := modalID + "-content-" + strconv.Itoa(index)
	printOption := `<div class="embed-options tooltip-embed top" data-tooltip="Print"><span class="material-icons" data-id="`+contentID+`" onclick="printDiv(this)">print</span></div>`

	if strings.HasPrefix(embed.Media.Type, "image/") {
		mediaContent = `<img src="data:` + embed.Media.Type + `;base64,` + embed.Data + `" />`

	} else if strings.HasPrefix(embed.Media.Type, "audio/") {
		mediaContent = `<audio autobuffer="autobuffer" controls><source src="data:` + embed.Media.Type + `;base64,` + embed.Data + `" /></audio>`
		
		printOption = ""

	} else if strings.HasPrefix(embed.Media.Type, "video/") {
		mediaContent = `<video controls><source type="` + embed.Media.Type + `" src="data:` + embed.Media.Type + `;base64,` + embed.Data + `" /></video>`
		
		printOption = ""

	} else if strings.HasPrefix(embed.Media.Type, "text/") {
		copyOption = `<div class="embed-options tooltip-embed top" data-tooltip="Copy"><span class="material-icons" id="` + modalID + `-` + strconv.Itoa(index) + `-copy-btn">content_copy</span></div>
<script>
document.getElementById("` + modalID + `-` + strconv.Itoa(index) + `-copy-btn").addEventListener("click", function() {
	copyjs("#` + modalID + `-` + strconv.Itoa(index) + `-copy-content", {
		copyFromSelector: true,
		html: true
	});
});
</script>`

		rawDecodedText, err := base64.StdEncoding.DecodeString(embed.Data)
		if err != nil {
			return "", fmt.Errorf("error decoding base64 text: %v", err)
		}

		rawText := strings.ReplaceAll(string(rawDecodedText), "<", "&lt;")
		rawText = strings.ReplaceAll(rawText, ">", "&gt;")

		mediaContent = `<pre id="` + modalID + `-` + strconv.Itoa(index) + `-copy-content">` + rawText + `</pre>`

	} else if strings.HasPrefix(embed.Media.Type, "application/pdf") {
		mediaContent = `<embed src="data:` + embed.Media.Type + `;base64,` + embed.Data + `" /></embed>`
		
		printOption = ""

	} else {

		mediaContent = `<span class="material-icons">description</span>`
	}

	modalFileHTML := ModalFileHTML{
		RowClass:       rowClass,
		MediaType:      embed.Media.Type,
		MediaSize:      helpers.ByteCountSI(helpers.GetBase64Size(embed.Data)),
		MediaData:      embed.Data,
		MediaExtension: helpers.GetFileExtension(embed.Media.Type),
		CopyOption:     template.HTML(copyOption),
		PrintOption:     template.HTML(printOption),
		MediaContent:   template.HTML(mediaContent),
		ContentID:      contentID,
	}

	content, err := d.tmplParse("tmpl/modalFile.html", &modalFileHTML)
	if err != nil {
		return "", fmt.Errorf("error trying to parse modalFile file : %v", err)
	}

	embedHTML = content

	return
}

func (d *Data) createEmbedModals(modalID, keyword, name string, timer int, embeddings []models.Embed) (embedHTML string, err error) {

	if len(embeddings) > 0 {

		isFirst := true
		for i, embedCode := range embeddings {
			if helpers.GetBase64Size(embedCode.Data) > 0 {
				embedHTMLNew, err := d.embedFiles(embedCode, isFirst, i, modalID)
				if err != nil {
					return "", fmt.Errorf("error embedding file")
				} else {
					embedHTML += embedHTMLNew
				}
				isFirst = false
			}
		}

		if embedHTML != "" {
			modalHTML := ModalHTML{
				ModalID:   modalID,
				Keyword:   keyword,
				Name:      name,
				EmbedHTML: template.HTML(embedHTML),
				Duration:  template.HTML(helpers.HumanizeExecution(timer)),
			}

			if strings.HasPrefix(modalID, "step") {
				modalHTML.Keyword = "Step"
				modalHTML.Name = keyword + " " + name
			}

			content, err := d.tmplParse("tmpl/modal.html", &modalHTML)
			if err != nil {
				return "", fmt.Errorf("error trying to parse modal file : %v", err)
			}

			embedHTML = content
		}
	}

	return
}
