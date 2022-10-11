package rdb_test

import (
	"context"
	"grpc/test/rdb"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestListInvitee(t *testing.T) {
	//connect db
	conn, err := ConnectForTest()
	// check error
	require.NoError(t, err)

	ctx := context.Background()

	var testcasesListPostInvitee = []struct {
		id     int
		name   string
		postID uint32
		want   []rdb.PostInvitee
		setup  func(ctx context.Context, t *testing.T, conn rdb.Connection)
	}{{
		id:     1,
		name:   "normal get",
		postID: 1,
		want: []rdb.PostInvitee{
			{
				UUID: uuid.NullUUID{
					UUID:  [16]byte{0x83, 0xc8, 0xa0, 0x39, 0x65, 0xd1, 0x46, 0x80, 0x98, 0x05, 0x9b, 0x34, 0xcf, 0x15, 0xb0, 0x95},
					Valid: true,
				},
				PostID:    1,
				UserID:    1,
				Comment:   stringToPtr("comment"),
				Passcode:  stringToPtr("passcode                                                                                            "),
				CreatedBy: 1,
				CreatedAt: time.Date(2017, 1, 1, 0, 0, 0, 0, JST),
				UpdatedBy: 1,
				UpdatedAt: time.Date(2017, 1, 1, 0, 0, 0, 0, JST),
				DeletedBy: 0,
				DeletedAt: gorm.DeletedAt{
					Time:  time.Time{},
					Valid: false,
				},
				User: rdb.User{
					Model: gorm.Model{
						ID:        1,
						CreatedAt: time.Date(2017, 1, 1, 0, 0, 0, 0, JST),
						UpdatedAt: time.Date(2016, 1, 1, 0, 0, 0, 0, JST),
						DeletedAt: gorm.DeletedAt{
							Time: time.Time{},
						},
					},
					EmUserID:  1000,
					CompanyID: 1,
					Name:      "name",
					Email:     "name@gmail.com",
					IconID:    1,
					UserType:  1,
					CreatedBy: 1,
					UpdatedBy: 1,
					DeletedBy: 0,
				},
			}},
		setup: func(ctx context.Context, t *testing.T, conn rdb.Connection) {
			require.NoError(t, conn.WithContext(ctx).Exec(`INSERT INTO auth.post_invitees (uuid, post_id, user_id, comment, passcode, created_by, created_at, updated_by, updated_at, deleted_by, deleted_at) VALUES ('83c8a039-65d1-4680-9805-9b34cf15b095', 1, 1, 'comment', 'passcode', 1, '2017-01-01T00:00:00+09:00',1, '2017-01-01T00:00:00+09:00', 0, null);`).Error)
			require.NoError(t, conn.WithContext(ctx).Exec(`INSERT INTO auth.users (id, em_user_id, company_id, name, email, icon_id, user_type, created_by, created_at, updated_by, updated_at, deleted_by, deleted_at) VALUES (1, 1000, 1, 'name', 'name@gmail.com', 1,1,1, '2017-01-01T00:00:00+09:00',1, '2016-01-01T00:00:00+09:00', 0, null);`).Error)
		},
	}}

	for _, tt := range testcasesListPostInvitee {
		t.Run(tt.name, func(t *testing.T) {
			initDBForTests(context.Background(), t, conn)
			if tt.setup != nil {
				tt.setup(ctx, t, conn)
			}
			// execute
			got, err := conn.ListPostInvitee(ctx, tt.postID)
			require.NoError(t, err)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("diff %s", cmp.Diff(got, tt.want))
			}
		},
		)

	}

}

func stringToPtr(s string) *string {
	return &s
}
