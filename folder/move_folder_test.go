package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_folder_MoveFolder(t *testing.T) {
	// TODO: your tests here
	t.Parallel()
	
	testUUID := uuid.Must(uuid.NewV1())

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
				{Name: "foxtrot", OrgId: uuid.Must(uuid.NewV1()), Paths: "foxtrot"},
				{Name: "golf", OrgId: testUUID, Paths: "golf"},
			},
			target: "alpha",
			destination: "golf",
			want: []folder.Folder{
				{Name: "alpha", OrgId: testUUID, Paths: "golf.alpha"},
				{Name: "bravo", OrgId: testUUID, Paths: "golf.alpha.bravo"},
				{Name: "charlie", OrgId: testUUID, Paths: "golf.alpha.bravo.charlie"},
				{Name: "delta", OrgId: testUUID, Paths: "golf.alpha.delta"},
				{Name: "echo", OrgId: testUUID, Paths: "golf.aplpha.delta.echo"},
				{Name: "foxtrot", OrgId: uuid.Must(uuid.NewV1()), Paths: "foxtrot"},
				{Name: "golf", OrgId: testUUID, Paths: "golf"},
			},
			err: nil,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folder)
			get, err := f.MoveFolder(tt.orgID, tt.target, tt.destination)

			assert.Equal(t, tt.want, get, "Expected folder output: %v\nGot: %v\n", tt.want, get)
			
		})
	}

}
