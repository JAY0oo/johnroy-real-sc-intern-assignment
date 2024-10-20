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
	folder1 := "noble-vixen"
	folder2 := "free-contessa" 


	children := folderDriver.GetAllChildFolders(OrgID, folder2)
	fmt.Printf("\nChild folders of %v:\n", folder2)
	folder.PrettyPrint(children)
	println("\n")

	// Showcase:
	// move folder noble-vixen to folder free-contessa 
	newFolders, err := folderDriver.MoveFolder(OrgID, folder1, folder2)
	if err != nil {
		fmt.Println("Error moving folder: ", err)
	} else {
		fmt.Println("New folder: ")
		
		res = newFolders
		folder.PrettyPrint(res)
		folderDriver = folder.NewDriver(res)
	}
	// list free-contessa new subfolders
	fmt.Printf("New subfolders of %v:", folder2)
	folder.PrettyPrint(folderDriver.GetAllChildFolders(OrgID, folder2))
	println("\n")













}
