package explorer

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// 도메인 에러 정의
var (
	ErrProfileNotFound      = errors.New("explorer profile not found")
	ErrProfileAlreadyExists = errors.New("explorer profile already exists")
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, userID string) (ExplorerProfile, error) {
	// 첫번째 파라미터(userID)를 데이터로 넣고 DB가 생성한 컬럼들을 돌려받는다
	const query = `
		INSERT INTO explorer_profiles (user_id)
		VALUES ($1)
		RETURNING id, user_id, created_at, updated_at
	`

	// Scan() : 쿼리 결과 row의 컬럼 값을 Go 변수에 담는 함수
	var profile ExplorerProfile
	err := r.db.QueryRowContext(ctx, query, userID).Scan(
		&profile.ID,
		&profile.UserID,
		&profile.CreatedAt,
		&profile.UpdatedAt,
	)
	if err != nil {
		if isUniqueViolation(err) {
			return ExplorerProfile{}, ErrProfileAlreadyExists
		}

		return ExplorerProfile{}, err
	}

	return profile, nil
}

func (r *Repository) FindByUserID(ctx context.Context, userID string) (ExplorerProfile, error) {
	const query = `
		SELECT id, user_id, created_at, updated_at
		FROM explorer_profiles
		WHERE user_id = $1
	`

	var profile ExplorerProfile
	err := r.db.QueryRowContext(ctx, query, userID).Scan(
		&profile.ID,
		&profile.UserID,
		&profile.CreatedAt,
		&profile.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ExplorerProfile{}, ErrProfileNotFound
		}

		return ExplorerProfile{}, err
	}

	return profile, nil
}

// 에러가 PostgreSQL의 unique violation인지 확인하는 함수
func isUniqueViolation(err error) bool {
	var pqErr *pq.Error
	return errors.As(err, &pqErr) && pqErr.Code == "23505"
}
