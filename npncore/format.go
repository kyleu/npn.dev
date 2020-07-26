package npncore

func OxfordComma(names []string, clause string) string {
	ret := ""
	for idx, name := range names {
		if idx > 0 {
			if len(clause) > 0 && idx == (len(names)-1) {
				if idx > 1 {
					ret += ","
				}
				ret += " " + clause + " "
			} else {
				ret += ", "
			}
		}
		ret += name
	}
	return ret
}
