package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

// GetAllChildFolders()
// (GetFoldersByOrgID tests below)
func Test_folder_GetAllChildFolders(t *testing.T) {
	t.Parallel()

	testUUID := uuid.Must(uuid.NewV1()) // test orgID that needs to be the same

	test := [...]struct {
		name string
		orgID uuid.UUID
		folder []folder.Folder
		folderName string

		want []folder.Folder
	}{
		{
			// Normal operation no error expected
			name: "Normal operation",
			orgID: testUUID,
			folder: []folder.Folder{
				{Name: "alpha", OrgId: testUUID, Paths: "alpha"},
				{Name: "bravo", OrgId: testUUID, Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: testUUID, Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: testUUID, Paths: "alpha.delta"},
				{Name: "echo", OrgId: testUUID, Paths: "echo"},
				{Name: "foxtrot", OrgId: uuid.Must(uuid.NewV1()), Paths: "foxtrot"},
			},
			folderName: "alpha",
			want: []folder.Folder{
				{Name: "bravo", OrgId: testUUID, Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: testUUID, Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: testUUID, Paths: "alpha.delta"},
			},
		},
		{
			// Invalid ordID as input, expected to raise error and get nil
			name: "Invalid ordID",
			orgID: uuid.Must(uuid.NewV1()),
			folder: []folder.Folder{
				{Name: "alpha", OrgId: testUUID, Paths: "alpha"},
				{Name: "bravo", OrgId: testUUID, Paths: "alpha.bravo"},
			},
			folderName: "alpha",
			want: nil,
		},
		{
			// Folder doesn't exist, expect to return null and error raised
			name: "Invalid folder name",
			orgID: testUUID,
			folder: []folder.Folder{
				{Name: "alpha", OrgId: testUUID, Paths: "alpha"},
				{Name: "bravo", OrgId: testUUID, Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: testUUID, Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: testUUID, Paths: "alpha.delta"},
				{Name: "echo", OrgId: testUUID, Paths: "echo"},
				{Name: "foxtrot", OrgId: uuid.Must(uuid.NewV1()), Paths: "foxtrot"},
			},
			folderName: "hotel",
			want: nil,
		},
		{
			// Folder has no child, expect to return an empty list and a non error alert
			name: "Folder has no children folders",
			orgID: testUUID,
			folder: []folder.Folder{
				{Name: "alpha", OrgId: testUUID, Paths: "alpha"},
				{Name: "bravo", OrgId: testUUID, Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: testUUID, Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: testUUID, Paths: "alpha.delta"},
				{Name: "echo", OrgId: testUUID, Paths: "echo"},
				{Name: "foxtrot", OrgId: uuid.Must(uuid.NewV1()), Paths: "foxtrot"},
			},
			folderName: "echo",
			want: []folder.Folder{},
		},
		{
			// No case sensitivity, function is expected to operate normally regardless of the case of the folder name
			name: "Case sensitivity",
			orgID: testUUID,
			folder: []folder.Folder{
				{Name: "alpha", OrgId: testUUID, Paths: "alpha"},
				{Name: "bravo", OrgId: testUUID, Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: testUUID, Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: testUUID, Paths: "alpha.delta"},
				{Name: "echo", OrgId: testUUID, Paths: "echo"},
				{Name: "foxtrot", OrgId: uuid.Must(uuid.NewV1()), Paths: "foxtrot"},
			},
			folderName: "aLpHa",
			want: []folder.Folder{
				{Name: "bravo", OrgId: testUUID, Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: testUUID, Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: testUUID, Paths: "alpha.delta"},
			},
		},
		{
			// Folder empty return nil early
			name: "Empty folder",
			orgID: testUUID,
			folder: []folder.Folder{},
			folderName: "bravo",
			want: nil,
		},
	
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folder)
			get := f.GetAllChildFolders(tt.orgID, tt.folderName)
			assert.Equal(t, tt.want, get)
		})
	}
}

func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()

	testUUID1 := uuid.Must(uuid.NewV1())
	testUUID2 := uuid.Must(uuid.NewV1())

	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder

	}{
		{
			// Normal operation, return expected folder based on orgID
			name: "Normal operation",
			orgID: testUUID2,
			folders: []folder.Folder{
				{Name: "alpha", OrgId: testUUID1, Paths: "alpha"}, 
				{Name: "bravo", OrgId: testUUID2, Paths: "alpha.bravo"}, 
				{Name: "charlie", OrgId: testUUID2, Paths: "alpha.bravo.charlie"}, 
				{Name: "delta", OrgId: testUUID1, Paths: "alpha.delta"},
			},
			want: []folder.Folder{
				{Name: "bravo", OrgId: testUUID2, Paths: "alpha.bravo"}, 
				{Name: "charlie", OrgId: testUUID2, Paths: "alpha.bravo.charlie"},
			},
		},
		{
			// Invalid orgID no folder returned
			name: "Invalid orgID",
			orgID: uuid.Must(uuid.NewV1()),
			folders: []folder.Folder{
				{Name: "alpha", OrgId: testUUID1, Paths: "alpha"}, 
				{Name: "bravo", OrgId: testUUID2, Paths: "alpha.bravo"}, 
				{Name: "charlie", OrgId: testUUID2, Paths: "alpha.bravo.charlie"}, 
				{Name: "delta", OrgId: testUUID1, Paths: "alpha.delta"},
			},
			want: []folder.Folder{},
		},


	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get := f.GetFoldersByOrgID(tt.orgID)
			
			assert.Equal(t, tt.want, get)

		})
	}
}