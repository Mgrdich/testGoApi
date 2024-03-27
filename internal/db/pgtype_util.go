package db

import (
	"math/big"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func ToText(text string) pgtype.Text {
	return pgtype.Text{String: text, Valid: true}
}

func ToDate(date time.Time) pgtype.Date {
	return pgtype.Date{Time: date, Valid: true}
}

func ToNumeric(bigInt *big.Int) pgtype.Numeric {
	return pgtype.Numeric{Int: bigInt, Valid: true}
}

func ToTimeStamp(time time.Time) pgtype.Timestamp {
	return pgtype.Timestamp{Time: time, Valid: true}
}

func ToUUID(id uuid.UUID) pgtype.UUID {
	return pgtype.UUID{Bytes: id, Valid: true}
}
