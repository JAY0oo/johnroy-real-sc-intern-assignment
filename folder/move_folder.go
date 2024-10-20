package folder

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gofrs/uuid"
)


func (f *driver) MoveFolder(orgId uuid.UUID, name string, dst string) ([]Folder, error) {
	
	fmt.Printf("Moving folder: %v to folder: %v ..... \n", name, dst)

	dstFolderSource := f.GetFoldersByOrgID(orgId)

	// check: moving folder to it self
	if name == dst {
		return nil, errors.New("cannot move folder to itself")
	}
	
	srcExist := false
	dstExist := false
	var srcFolder Folder
	var dstFolder Folder

	newFolder := []Folder{}
	for _, folder := range dstFolderSource {
		// if true: source exit
		if folder.Name == name {
			srcExist = true
			srcFolder = folder
		}
		// if true: destination exist
		if folder.Name == dst {
			dstExist = true
			dstFolder = folder
		}
		// if true: non child of target folder are excluded, but added to new folder
		if !strings.HasPrefix(folder.Paths, name) {
			newFolder = append(newFolder, folder)
		}
	}

	// check: source doesn't exist
	if !srcExist {
		return nil, errors.New("source folder does not exist")
	}

	// check: if destination doesn't exist its either in a different org or destination does not exist
	if !dstExist {
		allFolders := GetAllFolders()
		folderFound := false

		for _, folder := range allFolders {
			if folder.Name == dst {
				folderFound = true

				// check: if folder in a different orgID
				if folder.OrgId != orgId {
					return nil, errors.New("cannot move folder to a different organization")
				} 
			} 
		}
		// check: if destination doesn't exist
		if !folderFound {
			return nil, errors.New("destination folder does not exist")
		}
	}

	// check: moving folder to its own subfolder
	if strings.HasPrefix(dstFolder.Paths, srcFolder.Paths) {
		return nil, errors.New("cannot move folder to a child of itself")
	}

	// main move method
	children := f.GetAllChildFolders(orgId, name)
	folderToMove := Folder{Name: name, OrgId: orgId, Paths: name}
	children = append([]Folder{folderToMove}, children...)

	for i, child := range children {
		children[i].Paths = strings.Replace(child.Paths, name, dst + "." + name, 1)
	}
	
	newFolder = append(newFolder, children...)


	return newFolder, nil
}
