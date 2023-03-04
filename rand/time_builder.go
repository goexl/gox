package rand

import "time"

type timeBuilder struct {
	params *timeParams
}

func newTimeBuilder() *timeBuilder {
	return &timeBuilder{
		params: newTimeParams(),
	}
}

func (tb *timeBuilder) From(from time.Time) *timeBuilder {
	tb.params.from = from

	return tb
}

func (db *timeBuilder) To(to time.Time) *timeBuilder {
	db.params.to = to

	return db
}

func (db *timeBuilder) Between(from time.Time, to time.Time) *timeBuilder {
	db.params.from = from
	db.params.to = to

	return db
}

func (tb *timeBuilder) Build() *_time {
	return newTime(tb.params)
}
