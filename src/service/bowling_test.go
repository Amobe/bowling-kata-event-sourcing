package service

import (
	"fmt"
	"testing"

	"github.com/amobe/bowling-kata-event-sourcing/src/event/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_bowlingService_Throw(t *testing.T) {
	type fields struct {
		getRepo func(*gomock.Controller) *mocks.MockRepository
	}
	type args struct {
		id  string
		hit uint32
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "thrown",
			fields: fields{
				getRepo: func(ctrl *gomock.Controller) *mocks.MockRepository {
					repo := mocks.NewMockRepository(ctrl)
					repo.EXPECT().
						Get("1").
						Return(nil, nil)
					repo.EXPECT().
						Append("1", gomock.Not(gomock.Nil())).
						Return(nil)
					return repo
				},
			},
			args: args{
				id:  "1",
				hit: 1,
			},
			assertion: assert.NoError,
		},
		{
			name: "fail to get events",
			fields: fields{
				getRepo: func(ctrl *gomock.Controller) *mocks.MockRepository {
					repo := mocks.NewMockRepository(ctrl)
					repo.EXPECT().
						Get(gomock.Any()).
						Return(nil, fmt.Errorf("failed to get"))
					return repo
				},
			},
			args: args{
				id:  "1",
				hit: 1,
			},
			assertion: assert.Error,
		},
		{
			name: "fail to append events",
			fields: fields{
				getRepo: func(ctrl *gomock.Controller) *mocks.MockRepository {
					repo := mocks.NewMockRepository(ctrl)
					repo.EXPECT().
						Get("1").
						Return(nil, nil)
					repo.EXPECT().
						Append(gomock.Any(), gomock.Any()).
						Return(fmt.Errorf("failed to append"))
					return repo
				},
			},
			args: args{
				id:  "1",
				hit: 1,
			},
			assertion: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &bowlingService{
				repo: tt.fields.getRepo(gomock.NewController(t)),
			}
			tt.assertion(t, s.Throw(tt.args.id, tt.args.hit))
		})
	}
}
