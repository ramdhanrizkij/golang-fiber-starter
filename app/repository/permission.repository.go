package repository

import (
	"golang-fiber/app/models"
	"golang-fiber/app/models/scopes"
	"golang-fiber/helpers"

	"gorm.io/gorm"
)

type PermissionRepo interface {
	GetAll() ([]models.Permission, error)
	GetAllPagination(pagination helpers.Pagination) *helpers.Pagination
	SavePermission(data models.Permission) (models.Permission, error)
	GetById(id int) (models.Permission, error)
	UpdatePermission(data models.Permission) (models.Permission, error)
	DeletePermission(data models.Permission) (models.Permission, error)
}

type permissionRepo struct {
	db *gorm.DB
}

func NewPermissionRepo(conn *gorm.DB) PermissionRepo {
	return &permissionRepo{db: conn}
}

func (r *permissionRepo) GetAll() ([]models.Permission, error) {
	var data []models.Permission
	err := r.db.Find(&data).Error
	if err != nil {
		return data, err
	}
	return data, nil
}

func (r *permissionRepo) GetAllPagination(pagination helpers.Pagination) *helpers.Pagination {
	var data []*models.Permission

	r.db.Scopes(scopes.Paginate(data, &pagination, r.db)).Find(&data)
	pagination.Rows = data
	return &pagination
}

func (r *permissionRepo) SavePermission(data models.Permission) (models.Permission, error) {
	err := r.db.Save(&data).Error
	if err != nil {
		return data, err
	}
	return data, nil
}

func (r *permissionRepo) GetById(id int) (models.Permission, error) {
	var permission models.Permission
	err := r.db.Where("id=?", id).Find(&permission).Error
	if err != nil {
		return permission, err
	}
	return permission, nil
}

func (r *permissionRepo) UpdatePermission(data models.Permission) (models.Permission, error) {
	err := r.db.Save(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *permissionRepo) DeletePermission(data models.Permission) (models.Permission, error) {
	err := r.db.Delete(&data).Error
	if err != nil {
		return data, err
	}
	return data, nil
}
