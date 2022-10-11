package rdb

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	rdbLogger "gorm.io/gorm/logger"
)

type Connection struct {
	*gorm.DB
}
type User struct {
	gorm.Model
	EmUserID  int32  `gorm:"em_user_id"`
	CompanyID int32  `gorm:"company_id"`
	Name      string `gorm:"name"`
	Email     string `gorm:"email"`
	IconID    int32  `gorm:"icon_id"`
	UserType  int32  `gorm:"user_type"`
	CreatedBy int32  `gorm:"created_by"`
	UpdatedBy int32  `gorm:"updated_by"`
	DeletedBy int32  `gorm:"deleted_by"`
}

type PostInvitee struct {
	// 「->」mean readonly (disable write permission unless it configured)
	UUID      uuid.NullUUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;->"`
	PostID    uint32         `gorm:"post_id"`
	UserID    uint32         `gorm:"user_id"`
	Comment   *string        `gorm:"comment"`
	Passcode  *string        `gorm:"passcode"`
	CreatedAt time.Time      `gorm:"created_at;->"`
	CreatedBy int32          `gorm:"created_by"`
	UpdatedAt time.Time      `gorm:"updated_at;->"`
	UpdatedBy int32          `gorm:"updated_by"`
	DeletedAt gorm.DeletedAt `gorm:"deleted_at;->"`
	DeletedBy int32          `gorm:"deleted_by"`

	User User `gorm:"foreignkey:UserID"`
}

type Env struct {
	Host     string // EMZ_RDB_HOST
	User     string // EMZ_RDB_USER
	Password string // EMZ_RDB_PASSWORD
	Dbname   string // EMZ_RDB_DBNAME
	Driver   string // EMZ_RDB_DRIVER
}

const searchPath = "auth"

var JST = time.FixedZone("JST", 9*60*60)

func Connect() (conn Connection, err error) {
	// rdb env
	var rdbEnv Env
	_ = envconfig.Process("EMZ_RDB", &rdbEnv)

	// dns
	dsn := fmt.Sprintf("host=%s user=%s "+
		"password=%s dbname=%s search_path=%s sslmode=disable",
		rdbEnv.Host, rdbEnv.User, rdbEnv.Password, rdbEnv.Dbname, searchPath)

	// connect to db
	conn.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("connect db fail: %v", err)
	}

	// log sql executing
	conn.Logger = conn.DB.Logger.LogMode(rdbLogger.Info)

	return
}

func (conn *Connection) ListPostInvitee(
	ctx context.Context, postID uint32,
) ([]PostInvitee, error) {
	var postInvitees []PostInvitee
	err := conn.
		WithContext(ctx).
		Where("post_id = ?", postID).
		Preload("User").
		Order("updated_at DESC").
		Find(&postInvitees).Error

	if err != nil {
		log.Printf("get post invitee list error")
		return nil, err
	}
	return postInvitees, nil
}
