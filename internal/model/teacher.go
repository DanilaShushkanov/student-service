package model

type Teacher struct {
	ID           int64  `db:"id" goqu:"skipinsert,skipupdate"`
	PositionType int64  `db:"position_type"`
	FullName     string `db:"full_name"`
	StudentID    int64  `db:"student_id" goqu:"skipupdate"`
}
