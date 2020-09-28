package npncontroller

import (
	"fmt"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"golang.org/x/text/language"
	"math"
	"net/http"
	"strconv"
	"strings"
)

type ganttSection struct {
	Key   string  `json:"key"`
	Group string  `json:"group,omitempty"`
	Start int `json:"start"`
	End   int `json:"end"`
}

func Gantt(w http.ResponseWriter, r *http.Request) {
	Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		rowHeight := 24
		sections, completed, theme := parseGanttRequest(r)
		var pc = func(n int) float64 {return math.Floor((float64(n) / float64(completed)) * 10000)/100 }

		ret := make([]string, 0, len(sections) + 2)
		var ap = func(s string) {
			ret = append(ret, s)
		}

		totalHeight := len(sections) * rowHeight
		svgDecl := `<svg version="1.1" baseProfile="full" height="%v" width="100" preserveAspectRatio="none" viewBox="0 0 100 %v" xmlns="http://www.w3.org/2000/svg">`
		ap(fmt.Sprintf(svgDecl, totalHeight, totalHeight))

		lineMsg := `<line x1="%v" y1="0" x2="%v" y2="%v" stroke="#666" stroke-width="0.1" />`
		for idx := 0; idx < 11; idx ++ {
			ap(fmt.Sprintf(lineMsg, idx * 10, idx * 10, totalHeight));
		}

		bg := `<rect x="0" y="%v" width="100" height="%v" fill="transparent">%v</rect>`
		// bgLine := `<line x1="0" y1="%v" x2="100" y2="%v" stroke="#666" stroke-width="0.1" />`
		rect := `<rect x="%v" y="%v" width="%v" height="%v" style="fill: %v">%v</rect>`
		title := "<title>%v: %v%%\n%v - %v</title>"

		for idx, section := range sections {
			cy := rowHeight * idx
			per := pc(section.End - section.Start)
			start := pc(section.Start)
			startTitle := npncore.MicrosToMillis(language.AmericanEnglish, section.Start)
			width := pc(section.End - section.Start)
			endTitle := npncore.MicrosToMillis(language.AmericanEnglish, section.End)
			color := colorForSection(section.Key, theme)

			rectTitle := fmt.Sprintf(title, section.Key, per, startTitle, endTitle)
			ap(fmt.Sprintf(bg, cy, rowHeight, rectTitle))
			ap(fmt.Sprintf(rect, start, cy, width, rowHeight, color, rectTitle))
		}

		ap("</svg>")
		return RespondMIME("", "image/svg+xml", "svg", []byte(strings.Join(ret, "")), w)
	})
}

func parseGanttRequest(r *http.Request) ([]*ganttSection, int, string) {
	qps := QueryParamsFromRaw(r.URL.RawQuery)
	width := -1
	theme := "light"
	ret := make([]*ganttSection, 0)
	var get = func(k string) *ganttSection {
		for _, s := range ret {
			if s.Key == k {
				return s
			}
		}
		x := &ganttSection{Key: k}
		ret = append(ret, x)
		return x
	}
	for _, qp := range qps {
		if qp.Key == "w" {
			width, _ = strconv.Atoi(qp.Value)
		}
		if qp.Key == "t" {
			theme = qp.Value
		}
		if strings.Contains(qp.Key, ".") {
			k, t := npncore.SplitStringLast(qp.Key, '.', true)
			sec := get(k)
			switch t {
			case "g":
				sec.Group = qp.Value
			case "s":
				sec.Start, _ = strconv.Atoi(qp.Value)
			case "e":
				sec.End, _ = strconv.Atoi(qp.Value)
			}
		}
	}
	if width == -1 {
		for _, sec := range ret {
			if sec.End > width {
				width = sec.End
			}
		}
	}
	return ret, width, theme
}

func colorForSection(key string, theme string) string {
	switch (key) {
	case "dns":
		if theme == "dark" {
			return "#30444e";
		}
		return "#89b6cc";
	case "connect":
		if theme == "dark" {
			return "#30444e";
		}
		return "#89b6cc";
	case "tls":
		if theme == "dark" {
			return "#462206";
		}
		return "#c96112";
	case "reqheaders":
		if theme == "dark" {
			return "#072918";
		}
		return "#177245";
	case "reqbody":
		if theme == "dark" {
			return "#072918";
		}
		return "#177245";
	case "rspwait":
		if theme == "dark" {
			return "#101e33";
		}
		return "#397adb"
	case "rspheaders":
		if theme == "dark" {
			return "#101e33";
		}
		return "#397adb"
	case "rspbody":
		if theme == "dark" {
			return "#101e33";
		}
		return "#397adb"
	default:
		if theme == "dark" {
			return "#101e33";
		}
		return "#397adb"
	}
}
