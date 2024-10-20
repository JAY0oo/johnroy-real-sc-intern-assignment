package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func main() {
	OrgID := uuid.FromStringOrNil(folder.DefaultOrgID)
	res := folder.GetAllFolders()
	folderDriver := folder.NewDriver(res)

	// show folder of orgID
	orgFolder := folderDriver.GetFoldersByOrgID(OrgID)
	fmt.Printf("Folders of orgID: %v:\n", OrgID)
	folder.PrettyPrint(orgFolder)
	println("\n")

	// show child folders
	parent := "alpha"
	child := folderDriver.GetAllChildFolders(OrgID, parent)
	fmt.Printf("Child folders of %v:\n", parent)
	folder.PrettyPrint(child)
	println("\n")

	// Showcase:
	// move folder alpha to folder golf
	newFolders, err := folderDriver.MoveFolder(OrgID, "alpha", "golf")
	if err != nil {
		fmt.Println("Error moving folder: ", err)
	} else {
		fmt.Println("New folder: ")
		
		res = newFolders
		folder.PrettyPrint(res)
		folderDriver = folder.NewDriver(res)
	}
	// list golf new subfolders
	fmt.Println("New subfolders of golf:")
	folder.PrettyPrint(folderDriver.GetAllChildFolders(OrgID, "golf"))
	println("\n")
}
