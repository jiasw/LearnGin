package repositories

import (
	"errors"

	"gorm.io/gorm"
)

type BaseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{db: db}
}

// 基础 CRUD 方法 ---------------------------------------------------

// Create 创建记录
func (r *BaseRepository[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

// GetByID 根据主键查询
func (r *BaseRepository[T]) GetByID(id uint) (*T, error) {
	var entity T
	err := r.db.First(&entity, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &entity, err
}

// Update 更新整个实体
func (r *BaseRepository[T]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

// Delete 删除记录（硬删除）
func (r *BaseRepository[T]) Delete(id uint) error {
	return r.db.Delete(new(T), id).Error
}

// 高级查询方法 ---------------------------------------------------

// Where 条件查询
func (r *BaseRepository[T]) Where(query interface{}, args ...interface{}) *BaseRepository[T] {
	return &BaseRepository[T]{db: r.db.Where(query, args...)}
}

// First 查询第一条记录
func (r *BaseRepository[T]) First() (*T, error) {
	var entity T
	err := r.db.First(&entity).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &entity, err
}

// Paginate 分页
func (r *BaseRepository[T]) Paginate(page, pageSize int) ([]T, int64, error) {
	var (
		entities []T
		total    int64
	)

	if err := r.db.Model(new(T)).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&entities).Error

	return entities, total, err
}

// 事务支持 ---------------------------------------------------

// BeginTx 开启事务
func (r *BaseRepository[T]) BeginTx() *gorm.DB {
	return r.db.Begin()
}

// Transaction 自动事务处理
func (r *BaseRepository[T]) Transaction(fn func(txRepo *BaseRepository[T]) error) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		txRepo := NewBaseRepository[T](tx)
		return fn(txRepo)
	})
}

// 工具方法 ---------------------------------------------------

// Exists 判断记录是否存在
func (r *BaseRepository[T]) Exists(conditions map[string]interface{}) (bool, error) {
	var count int64
	query := r.db.Model(new(T))
	for field, value := range conditions {
		query = query.Where(field+" = ?", value)
	}
	err := query.Count(&count).Error
	return count > 0, err
}

// UpdateFields 部分字段更新
func (r *BaseRepository[T]) UpdateFields(id uint, fields map[string]interface{}) error {
	return r.db.Model(new(T)).Where("id = ?", id).Updates(fields).Error
}
