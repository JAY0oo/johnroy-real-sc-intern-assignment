package folder

import (
	"errors"
	"strings"

	"github.com/gofrs/uuid"
)


func (f *driver) MoveFolder(orgId uuid.UUID, name string, dst string) ([]Folder, error) {
	
	dstFolderSource := f.GetFoldersByOrgID(orgId)

	// ERROR CHECKS
	// Move folder to itself 
	if name == dst {
		return nil, errors.New("cannot move folder to itself")
	}
	
	
	srcExist := false
	dstExist := false
	var srcFolder Folder
	var dstFolder Folder

	

	newDstFolder := []Folder{}
	for _, folder := range dstFolderSource {
		// target folder exist?
		if folder.Name == name {
			srcExist = true
			srcFolder = folder
		}

		// destination folder exist?
		if folder.Name == dst {
			dstExist = true
			dstFolder = folder
		}
		
		if !strings.HasPrefix(folder.Paths, name) {
			newDstFolder = append(newDstFolder, folder)
		}
	}

	if !srcExist {
		return nil, errors.New("source folder does not exist")
	}

	// check if destination doesn't exist its either in a different org or destination does not exist
	if !dstExist {
		allFolders := GetAllFolders()
		folderFound := false
		for _, folder := range allFolders {
			if folder.Name == dst {
				folderFound = true
				if folder.OrgId != orgId {
					return nil, errors.New("cannot move folder to a different organization")
				} 
			} 
		}
		if !folderFound {
			return nil, errors.New("destination folder does not exist")
		}
	}

	if strings.HasPrefix(dstFolder.Paths, srcFolder.Paths) {
		return nil, errors.New("can't move folder to a child of itself")
	}


	children := f.GetAllChildFolders(orgId, name)
	folderToMove := Folder{Name: name, OrgId: orgId, Paths: name}
	children = append([]Folder{folderToMove}, children...)

	for i, child := range children {
		children[i].Paths = strings.Replace(child.Paths, name, dst + "." + name, 1)
	}
	
	newDstFolder = append(newDstFolder, children...)


	return newDstFolder, nil
}
