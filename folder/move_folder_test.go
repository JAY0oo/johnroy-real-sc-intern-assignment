package folder_test

import (
	"errors"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_folder_MoveFolder(t *testing.T) {
	t.Parallel()
	
	testUUID := uuid.Must(uuid.NewV1())
	res := folder.GetAllFolders()
	OrgID := uuid.FromStringOrNil(folder.DefaultOrgID)

	test := [...]struct {
		name string
		orgID uuid.UUID
		folder []folder.Folder
		target string
		destination string

		want []folder.Folder
		err error
	}{
		{
			name: "Normal operation",
			orgID: testUUID,
			folder: []folder.Folder{
				{Name: "alpha", OrgId: testUUID, Paths: "alpha"},
				{Name: "bravo", OrgId: testUUID, Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: testUUID, Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: testUUID, Paths: "alpha.delta"},
				{Name: "echo", OrgId: testUUID, Paths: "alpha.delta.echo"},
				{Name: "foxtrot", OrgId: testUUID, Paths: "foxtrot"},
				{Name: "golf", OrgId: testUUID, Paths: "golf"},
			},
			target: "alpha",
			destination: "golf",
			want: []folder.Folder{
				{Name: "alpha", OrgId: testUUID, Paths: "golf.alpha"},
				{Name: "bravo", OrgId: testUUID, Paths: "golf.alpha.bravo"},
				{Name: "charlie", OrgId: testUUID, Paths: "golf.alpha.bravo.charlie"},
				{Name: "delta", OrgId: testUUID, Paths: "golf.alpha.delta"},
				{Name: "echo", OrgId: testUUID, Paths: "golf.alpha.delta.echo"},
				{Name: "foxtrot", OrgId: testUUID, Paths: "foxtrot"},
				{Name: "golf", OrgId: testUUID, Paths: "golf"},
			},
			err: nil,
		},
		{
			name: "Moving a folder to its own child folder",
			orgID: testUUID,
			folder: []folder.Folder{
				{Name: "alpha", OrgId: testUUID, Paths: "alpha"},
				{Name: "bravo", OrgId: testUUID, Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: testUUID, Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: testUUID, Paths: "alpha.delta"},
				{Name: "echo", OrgId: testUUID, Paths: "alpha.delta.echo"},
				{Name: "foxtrot", OrgId: testUUID, Paths: "foxtrot"},
				{Name: "golf", OrgId: testUUID, Paths: "golf"},
			},
			target: "alpha",
			destination: "bravo",
			want: nil,
			err: errors.New("can't move folder to a child of itself"),
		},
		{
			name: "Moving folder to itself",
			orgID: testUUID,
			folder: []folder.Folder{
				{Name: "alpha", OrgId: testUUID, Paths: "alpha"},
				{Name: "bravo", OrgId: testUUID, Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: testUUID, Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: testUUID, Paths: "alpha.delta"},
				{Name: "echo", OrgId: testUUID, Paths: "alpha.delta.echo"},
				{Name: "foxtrot", OrgId: testUUID, Paths: "foxtrot"},
				{Name: "golf", OrgId: testUUID, Paths: "golf"},
			},
			target: "golf",
			destination: "golf",
			want: nil,
			err: errors.New("cannot move folder to itself"),
		},
		{
			name: "moving folder to a different orgID",
			orgID: OrgID,
			folder: res,
			target: "alpha",
			destination: "hotel",
			want: nil,
			err: errors.New("cannot move folder to a different organization"),
		},
		{
			name: "Source folder does not exist",
			orgID: OrgID,
			folder: res,
			target: "peanut",
			destination: "alpha",
			want: nil,
			err: errors.New("source folder does not exist"),
		},
		{
			name: "destination folder does not exist",
			orgID: OrgID,
			folder: res,
			target: "alpha",
			destination: "alpaca",
			want: nil,
			err: errors.New("destination folder does not exist"),
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folder)
			get, err := f.MoveFolder(tt.orgID, tt.target, tt.destination)

			if tt.err == nil {
				assert.Nil(t, err)
				assert.ElementsMatch(t, tt.want, get, "Expected folder output: %v\nGot: %v\n", tt.want, get)
			} else {
				assert.Nil(t, get)
				assert.EqualError(t, err, tt.err.Error())
			}
		})
	}

}
