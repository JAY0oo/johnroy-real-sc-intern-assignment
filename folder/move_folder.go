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
		return nil, errors.New("Error: cannot move folder to itself!")
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
		return nil, errors.New("Error: source folder does not exist!")
	}
	if !dstExist {
		return nil, errors.New("Error: destination folder does not exist!")
	}
	if strings.HasPrefix(dstFolder.Paths, srcFolder.Paths) {
		return nil, errors.New("Error: can't move folder to a child of itself!")
	}
	if srcFolder.OrgId != dstFolder.OrgId {
		return nil, errors.New("Error: cannot move folder to a different organization!")
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
