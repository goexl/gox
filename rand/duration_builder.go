package rand

import "time"

type durationBuilder struct {
	params *durationParams
}

func newDurationBuilder() *durationBuilder {
	return &durationBuilder{
		params: newDurationParams(),
	}
}

func (db *durationBuilder) To(to time.Duration) *durationBuilder {
	db.params.to = to

	return db
}

func (db *durationBuilder) Between(from time.Duration, to time.Duration) *durationBuilder {
	db.params.from = from
	db.params.to = to

	return db
}

func (db *durationBuilder) Build() *duration {
	return newDuration(db.params)
}
