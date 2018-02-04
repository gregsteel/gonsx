package ipset

import (
	"fmt"
	"log"
)

func (i List) String() string {
	return fmt.Sprintf("List object, contains ipset objects.")
}

func (i IPSet) String() string {
	return fmt.Sprintf("objectId: %-20s name: %-20s.", i.ObjectID, i.Name)
}

// FilterByName returns a single ipset object if it matches the name in List
func (i List) FilterByName(name string) *IPSet {
	var ipsetFound IPSet
	for _, ipset := range i.IPSets {
    log.Printf(fmt.Sprintf("[DEBUG] ipset.IPSet %s", ipset.Name))
		if ipset.Name == name {
			ipsetFound = ipset
			break
		}
	}
	return &ipsetFound
}

// CheckByName - Returns true or false depending if name is in ApplicationList
func (i List) CheckByName(name string) bool {
	for _, ipSet := range i.IPSets {
		if ipSet.Name == name {
			return true
		}
	}
	return false
}
