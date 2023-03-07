package types

import "time"

type GftTimelineDiff = [2]any

type GftTimeline struct {
	TlStart   *time.Time                 `bson:"tlStart,omitempty" json:"tlStart,omitempty"`
	TlEnd     *time.Time                 `bson:"tlEnd,omitempty" json:"tlEnd,omitempty"`
	TlVersion int                        `bson:"tlVersion,omitempty" json:"tlVersion,omitempty"`
	TlNote    string                     `bson:"tlNote,omitempty" json:"tlNote,omitempty"`
	TlMeta    map[string]any             `bson:"tlMeta" json:"tlMeta,omitempty"`
	TlDiff    map[string]GftTimelineDiff `bson:"TlDiff" json:"TlDiff,omitempty"`
}
