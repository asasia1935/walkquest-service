package explorer

import "time"

type ExplorerProfile struct {
	ID        int64
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
