package models

import "strings"

func databaseAssociation(load string, property string, preloads *[]string, loads ...*[]string) {
	lp := len(property)
	*preloads = append(*preloads, property)
	if len(load) > lp && len(loads) == 1 {
		*loads[0] = append(*loads[0], load[lp+1:])
	}
}

func abstractAssociation(load string, property string, setFlag *bool, loads ...*[]string) {
	lp := len(property)
	*setFlag = true
	if len(load) > lp && len(loads) == 1 {
		*loads[0] = append(*loads[0], load[lp+1:])
	}
}

func primaryLoad(load string) (primary string) {
	primary = strings.Split(load, ".")[0]
	return
}

func primaryLoads(loads []string) (primaries []string) {
	primaries = []string{}
	for _, load := range loads {
		primaries = append(primaries, primaryLoad(load))
	}
	return
}
