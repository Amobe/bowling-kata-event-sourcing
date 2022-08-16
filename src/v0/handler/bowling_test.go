package handler

import (
	"fmt"
	"os"
	"testing"

	mh "github.com/amobe/bowling-kata-event-sourcing/src/v0/handler/mocks"
	"github.com/amobe/bowling-kata-event-sourcing/src/v0/service"
	ms "github.com/amobe/bowling-kata-event-sourcing/src/v0/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewBowlingHandler(t *testing.T) {
	type args struct {
		bs service.Bowling
	}
	tests := []struct {
		name      string
		args      args
		want      Handler
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "new bowling handler",
			args: args{
				bs: &ms.MockBowling{},
			},
			want: &bowlingHandler{
				bs: &ms.MockBowling{},
			},
			assertion: assert.NoError,
		},
		{
			name: "bowling service is nill",
			args: args{
				bs: nil,
			},
			want:      nil,
			assertion: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBowlingHandler(tt.args.bs)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_bowlingHandler_Handle(t *testing.T) {
	type fields struct {
		getBS func(ctrl *gomock.Controller) service.Bowling
	}
	type args struct {
		getCtx func(ctrl *gomock.Controller) Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "throw the ball",
			fields: fields{
				getBS: func(ctrl *gomock.Controller) service.Bowling {
					s := ms.NewMockBowling(ctrl)
					s.EXPECT().Throw("0", uint32(1)).Return(nil)
					return s
				},
			},
			args: args{
				getCtx: func(ctrl *gomock.Controller) Context {
					ctx := mh.NewMockContext(ctrl)
					ctx.EXPECT().Query("hit").Return("1")
					return ctx
				},
			},
		},
		{
			name: "bowling service error",
			fields: fields{
				getBS: func(ctrl *gomock.Controller) service.Bowling {
					s := ms.NewMockBowling(ctrl)
					s.EXPECT().Throw(gomock.Any(), gomock.Any()).Return(fmt.Errorf("err"))
					return s
				},
			},
			args: args{
				getCtx: func(ctrl *gomock.Controller) Context {
					ctx := mh.NewMockContext(ctrl)
					ctx.EXPECT().Query("hit").Return("1")
					ctx.EXPECT().Writer().Return(os.Stdout)
					return ctx
				},
			},
		},
		{
			name: "hit value parsing error",
			fields: fields{
				getBS: func(ctrl *gomock.Controller) service.Bowling {
					s := ms.NewMockBowling(ctrl)
					return s
				},
			},
			args: args{
				getCtx: func(ctrl *gomock.Controller) Context {
					ctx := mh.NewMockContext(ctrl)
					ctx.EXPECT().Query("hit").Return("aaa")
					ctx.EXPECT().Writer().Return(os.Stdout)
					return ctx
				},
			},
		},
		{
			name: "hit param missing error",
			fields: fields{
				getBS: func(ctrl *gomock.Controller) service.Bowling {
					s := ms.NewMockBowling(ctrl)
					return s
				},
			},
			args: args{
				getCtx: func(ctrl *gomock.Controller) Context {
					ctx := mh.NewMockContext(ctrl)
					ctx.EXPECT().Query("hit").Return("")
					ctx.EXPECT().Writer().Return(os.Stdout)
					return ctx
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			h := &bowlingHandler{
				bs: tt.fields.getBS(ctrl),
			}
			h.Handle(tt.args.getCtx(ctrl))
		})
	}
}
