package npncontroller

import (
	"fmt"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"golang.org/x/text/language"
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
		sections, completed := parseGanttRequest(r)
		var pc = func(n int) float64 {return (float64(n) / float64(completed)) * 100 }

		ret := make([]string, 0, len(sections) + 2)
		var ap = func(s string) {
			ret = append(ret, s)
		}

		totalHeight := len(sections) * rowHeight
		svgDecl := `<svg version="1.1" baseProfile="full" height="%v" width="100" preserveAspectRatio="none" viewBox="0 0 100 %v" xmlns="http://www.w3.org/2000/svg">`
		ap(fmt.Sprintf(svgDecl, totalHeight, totalHeight))

		lineMsg := `<line x1="%v" y1="0" x2="%v" y2="%v" stroke="#666" stroke-width="0.5" />`
		ap(fmt.Sprintf(lineMsg, 1, 1, totalHeight));
		for idx := 1; idx < 10; idx ++ {
			ap(fmt.Sprintf(lineMsg, idx * 10, idx * 10, totalHeight));
		}
		ap(fmt.Sprintf(lineMsg, 99, 99, totalHeight));

		for idx, section := range sections {
			cy := rowHeight * idx
			per := pc(section.End - section.Start)
			start := pc(section.Start)
			startTitle := npncore.MicrosToMillis(language.AmericanEnglish, section.Start)
			width := pc(section.End - section.Start)
			endTitle := npncore.MicrosToMillis(language.AmericanEnglish, section.End)
			color := colorForSection(section.Key)

			ap(fmt.Sprintf(`<rect x="0" y="%v" width="100" height="%v" fill="transparent" />`, cy, rowHeight))

			rectTitle := fmt.Sprintf("<title>%v: %v\n%v - %v</title>", section.Key, per, startTitle, endTitle)
			rect := `<rect x="%v" y="%v" width="%v" height="%v" style="fill: %v">%v</rect>`
			ap(fmt.Sprintf(rect, start, cy, width, rowHeight, color, rectTitle))
		}

		ap("</svg>")
		return RespondMIME("", "image/svg+xml", "svg", []byte(strings.Join(ret, "")), w)
	})
}

func parseGanttRequest(r *http.Request) ([]*ganttSection, int) {
	qps := QueryParamsFromRaw(r.URL.RawQuery)
	width := -1
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
	return ret, width
}

func colorForSection(key string) string {
	switch (key) {
	case "dns":
		return "#89b6cc";
	case "connect":
		return "#89b6cc";
	case "tls":
		return "#c96112";
	case "reqheaders":
		return "#177245";
	case "reqbody":
		return "#177245";
	case "rspwait":
		return "#397adb"
	case "rspheaders":
		return "#397adb"
	case "rspbody":
		return "#397adb"
	default:
		return "#397adb"
	}
}
