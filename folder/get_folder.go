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
	// Your code here...

	// orgID error test
	if orgID == uuid.Nil {
		log.Printf("ERROR: Invalid orgID: %v", orgID)
		return nil
	}

	folderWithOrgID := f.GetFoldersByOrgID(orgID)
	if len(folderWithOrgID) == 0 {
		// check if folderWithOrgID is empty to avoid operating on empty slice
		fmt.Printf("Error: No folders found for orgID: %v\n", orgID)
		return nil
	}

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
	
	// if children slice is empty its either the folder has no children or the folder does not exist
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

