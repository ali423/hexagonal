package db

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type Config struct {
	DBName     string
	DBUsername string
	DBPassword string
	DBHostname string
	DBPort     int
}

type DB struct {
	c  *Config
	DB *gorm.DB
}

func NewDB(c *Config) *DB {
	return &DB{
		c: c,
	}
}

func (d *DB) SetupPG() error {
	dsn := "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Tehran"
	dsn = fmt.Sprintf(dsn, d.c.DBHostname, d.c.DBUsername, d.c.DBPassword, d.c.DBName, d.c.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	d.DB = db
	return nil
}

func (d *DB) SetupMySQL() error {
	time.Local, _ = time.LoadLocation("Asia/Tehran")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		d.c.DBUsername,
		d.c.DBPassword,
		d.c.DBHostname,
		d.c.DBPort,
		d.c.DBName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	d.DB = db
	return nil
}

func (d *DB) SetupSQLite() error {
	db, err := gorm.Open(sqlite.Open(d.c.DBName), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return err
	}
	d.DB = db
	return nil
}

// Repository provides a basic CRUD functionality
type Repository interface {
	Get(id uint, out interface{})
	GetBy(by string, val interface{}, out interface{})
	List(out interface{})
	Add(interface{}) error
	Edit(id uint, model interface{}) error
	Delete(id uint, model interface{}) error
}

// DefaultRepository implements Repository with minimum functionality
type DefaultRepository struct {
	DB *DB
}

func NewDefaultRepository(db *DB) *DefaultRepository {
	return &DefaultRepository{
		DB: db,
	}
}

func (d *DefaultRepository) Get(id uint, out interface{}) {
	d.DB.DB.First(out, "id = ?", id)
}

func (d *DefaultRepository) GetBy(by string, val interface{}, out interface{}) {
	q := fmt.Sprintf("%s = ?", by)
	d.DB.DB.Find(out, q, val)
}

func (d *DefaultRepository) List(out interface{}) {
	d.DB.DB.Find(out)
}

func (d *DefaultRepository) Add(data interface{}) error {
	return d.DB.DB.Create(data).Error
}

func (d *DefaultRepository) Edit(id uint, data interface{}) error {
	return d.DB.DB.Where("id = ?", id).Updates(data).Error
}

func (d *DefaultRepository) Delete(id uint, model interface{}) error {
	return d.DB.DB.Delete(model, "id = ?", id).Error
}
