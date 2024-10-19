package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func main() {
	OrgID := uuid.FromStringOrNil(folder.DefaultOrgID)
	
	res := folder.GetAllFolders()
	
	// example usage
	folderDriver := folder.NewDriver(res)
	folder.PrettyPrint(folderDriver)
	orgFolder := folderDriver.GetFoldersByOrgID(OrgID)
	child := folderDriver.GetAllChildFolders(OrgID, "alpha")

	fmt.Println("OrgFolder Before")
	folder.PrettyPrint(orgFolder)

	// MOVE
	fmt.Println("MOVE")
	test, err := folderDriver.MoveFolder(OrgID, "alpha", "golf")
	if err != nil {
		fmt.Println("Error moving folder: ", err)
	} else {
		fmt.Println("OrgFolder After")
		folder.PrettyPrint(test)
		// orgFolder = append(orgFolder, test...)
	}

	// folder.PrettyPrint(test)

	
	//folder.PrettyPrint(orgFolder)

	//folder.PrettyPrint(res)
	//fmt.Printf("\n Folders for orgID: %s", orgID)
	

	fmt.Println("FOLDERS WITH NAME:")
	folder.PrettyPrint(child)
}
