// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql/driver"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Priority string

const (
	PriorityLow    Priority = "low"
	PriorityMedium Priority = "medium"
	PriorityHigh   Priority = "high"
)

func (e *Priority) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Priority(s)
	case string:
		*e = Priority(s)
	default:
		return fmt.Errorf("unsupported scan type for Priority: %T", src)
	}
	return nil
}

type NullPriority struct {
	Priority Priority `json:"priority"`
	Valid    bool     `json:"valid"` // Valid is true if Priority is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullPriority) Scan(value interface{}) error {
	if value == nil {
		ns.Priority, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Priority.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullPriority) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Priority), nil
}

type Status string

const (
	StatusPending    Status = "pending"
	StatusInprogress Status = "in progress"
	StatusCompleted  Status = "completed"
)

func (e *Status) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Status(s)
	case string:
		*e = Status(s)
	default:
		return fmt.Errorf("unsupported scan type for Status: %T", src)
	}
	return nil
}

type NullStatus struct {
	Status Status `json:"status"`
	Valid  bool   `json:"valid"` // Valid is true if Status is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullStatus) Scan(value interface{}) error {
	if value == nil {
		ns.Status, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Status.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Status), nil
}

type UserRole string

const (
	UserRoleAdmin UserRole = "admin"
	UserRoleUser  UserRole = "user"
)

func (e *UserRole) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserRole(s)
	case string:
		*e = UserRole(s)
	default:
		return fmt.Errorf("unsupported scan type for UserRole: %T", src)
	}
	return nil
}

type NullUserRole struct {
	UserRole UserRole `json:"user_role"`
	Valid    bool     `json:"valid"` // Valid is true if UserRole is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUserRole) Scan(value interface{}) error {
	if value == nil {
		ns.UserRole, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UserRole.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUserRole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UserRole), nil
}

type Comment struct {
	ID          int32              `json:"id"`
	TaskID      uuid.UUID          `json:"task_id"`
	CommentText string             `json:"comment_text"`
	CreatedBy   uuid.UUID          `json:"created_by"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
}

type Task struct {
	Taskid      uuid.UUID          `json:"taskid"`
	Title       string             `json:"title"`
	Description pgtype.Text        `json:"description"`
	AssignedTo  uuid.UUID          `json:"assigned_to"`
	Priority    Priority           `json:"priority"`
	Status      Status             `json:"status"`
	DueDate     pgtype.Date        `json:"due_date"`
	CreatedBy   pgtype.UUID        `json:"created_by"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
}

type User struct {
	Userid       uuid.UUID          `json:"userid"`
	Name         string             `json:"name"`
	Email        string             `json:"email"`
	Password     string             `json:"password"`
	Role         NullUserRole       `json:"role"`
	Profilephoto string             `json:"profilephoto"`
	Createdat    pgtype.Timestamptz `json:"createdat"`
}