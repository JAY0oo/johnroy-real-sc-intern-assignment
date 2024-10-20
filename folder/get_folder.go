package folder

import (
	"fmt"
	"log"
	"strings"

	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}
	return res

}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {

	// check: orgID validity
	if orgID == uuid.Nil {
		fmt.Printf("ERROR: Invalid orgID: %v", orgID)
		return nil
	}

	folderWithOrgID := f.GetFoldersByOrgID(orgID)
	// check: if folderWithOrgID is empty
	if len(folderWithOrgID) == 0 {
		fmt.Printf("Error: No folders found for orgID: %v\n", orgID)
		return nil
	}

	// main
	var nameExist bool = false
	children := []Folder{}

	for _, folder := range folderWithOrgID {
		if strings.HasPrefix(folder.Paths, strings.ToLower(name) + ".") {
			children = append(children, folder)
		}
		if folder.Name == strings.ToLower(name) {
			nameExist = true
		}
	} 
	
	// check: if children slice is empty its either the folder has no children or the folder does not exist
	if len(children) == 0 {
		if nameExist{
			fmt.Printf("Alert: Folder: " + name + " has no children folders")
			return []Folder{}
		} else {
			log.Printf("Folder not found!")
			return nil
		}
	}

	return children
}

