package repositories

import (
	"reflect"
	"testing"

	"github.com/aditya-nagare/d2h/db"
	"github.com/aditya-nagare/d2h/services/pkg/models"
	"github.com/jinzhu/gorm"
)

func TestRepositoryStruct_Recharge(t *testing.T) {
	u := models.User{
		ID:   2,
		Name: "ganesh",
	}
	type fields struct {
		dbConn *gorm.DB
	}
	type args struct {
		user *models.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.User
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				dbConn: db.NewTestDbConnection(),
			},
			args: args{
				user: &u,
			},
			want:    &u,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RepositoryStruct{
				dbConn: tt.fields.dbConn,
			}
			got, err := r.Recharge(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("RepositoryStruct.Recharge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RepositoryStruct.Recharge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepositoryStruct_ViewPacks(t *testing.T) {
	type fields struct {
		dbConn *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				dbConn: db.NewTestDbConnection(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RepositoryStruct{
				dbConn: tt.fields.dbConn,
			}
			_, _, err := r.ViewPacks()
			if (err != nil) != tt.wantErr {
				t.Errorf("RepositoryStruct.ViewPacks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestRepositoryStruct_SubscribeBasePack(t *testing.T) {
	u := models.User{
		ID:   2,
		Name: "ganesh",
	}
	s := models.Subscription{
		UserID: 2,
		PackID: 1,
	}
	type fields struct {
		dbConn *gorm.DB
	}
	type args struct {
		user models.User
		subs models.Subscription
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.User
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				dbConn: db.NewTestDbConnection(),
			},
			args: args{
				user: u,
				subs: s,
			},
			want:    &u,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RepositoryStruct{
				dbConn: tt.fields.dbConn,
			}
			got, err := r.SubscribeBasePack(tt.args.user, tt.args.subs)
			if (err != nil) != tt.wantErr {
				t.Errorf("RepositoryStruct.SubscribeBasePack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RepositoryStruct.SubscribeBasePack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepositoryStruct_AddChannel(t *testing.T) {
	u := models.User{
		ID:   2,
		Name: "ganesh",
	}
	s := models.Subscription{
		UserID:    2,
		ChannelID: 1,
	}
	type fields struct {
		dbConn *gorm.DB
	}
	type args struct {
		user models.User
		subs models.Subscription
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.User
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				dbConn: db.NewTestDbConnection(),
			},
			args: args{
				user: u,
				subs: s,
			},
			want:    &u,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RepositoryStruct{
				dbConn: tt.fields.dbConn,
			}
			got, err := r.AddChannel(tt.args.user, tt.args.subs)
			if (err != nil) != tt.wantErr {
				t.Errorf("RepositoryStruct.AddChannel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RepositoryStruct.AddChannel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepositoryStruct_SubscribeSpecialService(t *testing.T) {
	u := models.User{
		ID:   2,
		Name: "ganesh",
	}
	s := models.Subscription{
		UserID:    2,
		ServiceID: 1,
	}
	type fields struct {
		dbConn *gorm.DB
	}
	type args struct {
		user models.User
		subs models.Subscription
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.User
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				dbConn: db.NewTestDbConnection(),
			},
			args: args{
				user: u,
				subs: s,
			},
			want:    &u,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RepositoryStruct{
				dbConn: tt.fields.dbConn,
			}
			got, err := r.SubscribeSpecialService(tt.args.user, tt.args.subs)
			if (err != nil) != tt.wantErr {
				t.Errorf("RepositoryStruct.SubscribeSpecialService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RepositoryStruct.SubscribeSpecialService() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestRepositoryStruct_ViewSubscription(t *testing.T) {
	u := models.User{
		ID:   2,
		Name: "ganesh",
	}
	type fields struct {
		dbConn *gorm.DB
	}
	type args struct {
		user models.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				dbConn: db.NewTestDbConnection(),
			},
			args: args{
				user: u,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RepositoryStruct{
				dbConn: tt.fields.dbConn,
			}
			_, err := r.ViewSubscription(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("RepositoryStruct.ViewSubscription() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestRepositoryStruct_UpdateInfo(t *testing.T) {
	u := models.User{
		ID:    2,
		Name:  "ganesh",
		Email: "ganesh@thorat.com",
	}
	type fields struct {
		dbConn *gorm.DB
	}
	type args struct {
		user models.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.User
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				dbConn: db.NewTestDbConnection(),
			},
			args: args{
				user: u,
			},
			want:    &u,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RepositoryStruct{
				dbConn: tt.fields.dbConn,
			}
			got, err := r.UpdateInfo(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("RepositoryStruct.UpdateInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RepositoryStruct.UpdateInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepositoryStruct_ListUsers(t *testing.T) {
	type fields struct {
		dbConn *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				dbConn: db.NewTestDbConnection(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RepositoryStruct{
				dbConn: tt.fields.dbConn,
			}
			_, err := r.ListUsers()
			if (err != nil) != tt.wantErr {
				t.Errorf("RepositoryStruct.ListUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestRepositoryStruct_SelectUser(t *testing.T) {
	u := models.User{
		ID:   2,
		Name: "ganesh",
	}
	type fields struct {
		dbConn *gorm.DB
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.User
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				dbConn: db.NewTestDbConnection(),
			},
			args: args{
				id: 2,
			},
			want:    &u,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RepositoryStruct{
				dbConn: tt.fields.dbConn,
			}
			got, err := r.SelectUser(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("RepositoryStruct.SelectUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.ID, tt.want.ID) && !reflect.DeepEqual(got.Name, tt.want.Name) {
				t.Errorf("RepositoryStruct.SelectUser() = %v,%v, want %v,%v", got.ID, got.Name, tt.want.ID, tt.want.Name)
			}
		})
	}
}

func TestRepositoryStruct_ListPacks(t *testing.T) {
	type fields struct {
		dbConn *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				dbConn: db.NewTestDbConnection(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RepositoryStruct{
				dbConn: tt.fields.dbConn,
			}
			_, err := r.ListPacks()
			if (err != nil) != tt.wantErr {
				t.Errorf("RepositoryStruct.ListPacks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestRepositoryStruct_ListServices(t *testing.T) {
	type fields struct {
		dbConn *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				dbConn: db.NewTestDbConnection(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RepositoryStruct{
				dbConn: tt.fields.dbConn,
			}
			_, err := r.ListServices()
			if (err != nil) != tt.wantErr {
				t.Errorf("RepositoryStruct.ListServices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestRepositoryStruct_ListChannels(t *testing.T) {
	type fields struct {
		dbConn *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				dbConn: db.NewTestDbConnection(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RepositoryStruct{
				dbConn: tt.fields.dbConn,
			}
			_, err := r.ListChannels()
			if (err != nil) != tt.wantErr {
				t.Errorf("RepositoryStruct.ListChannels() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
