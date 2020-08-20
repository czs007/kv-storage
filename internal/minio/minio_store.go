package minio_driver

import (
	"context"
	. "storage/pkg/types"
	. "storage/internal/minio/codec"
)

type minioDriver struct {
	driver *minioStore
}

func (s *minioDriver) put(ctx context.Context, key Key, value Value, timestamp Timestamp, suffix string) error {
	minioKey := MvccEncode(key, timestamp)

	s.driver.PUT(ctx, minioKey, )
}
scanLE(ctx context.Context, key Key, timestamp Timestamp, withValue bool) ([]Timestamp, []Key, []Value, error)
scanGE(ctx context.Context, key Key, timestamp Timestamp, withValue bool) ([]Timestamp, []Key, []Value, error)
scan(ctx context.Context, key Key, start Timestamp, end Timestamp, withValue bool) ([]Timestamp, []Key, []Value, error)
deleteLE(ctx context.Context, key Key, timestamp Timestamp) error
deleteGE(ctx context.Context, key Key, timestamp Timestamp) error
deleteRange(ctx context.Context, key Key, start Timestamp, end Timestamp) error

GetRow(ctx context.Context, key Key, timestamp Timestamp) (Value, error)
GetRows(ctx context.Context, keys []Key, timestamp Timestamp) ([]Value, error)

AddRow(ctx context.Context, key Key, value Value, segment string, timestamp Timestamp) error
AddRows(ctx context.Context, keys []Key, values []Value, segments []string, timestamp Timestamp) error

DeleteRow(ctx context.Context, key Key, timestamp Timestamp) error
DeleteRows(ctx context.Context, keys []Key, timestamp Timestamp) error

PutLog(ctx context.Context, key Key, value Value, timestamp Timestamp, channel int) error
FetchLog(ctx context.Context, start Timestamp, end Timestamp, channels []int) error

GetSegmenIndex(ctx context.Context, segment string) (SegmentIndex, error)
PutSegmentIndex(ctx context.Context, segment string, index SegmentIndex) error
DeleteSegmentIndex(ctx context.Context, segment string) error

GetSegmentDL(ctx context.Context, segment string) (SegmentDL, error)
SetSegmentDL(ctx context.Context, segment string, log SegmentDL) error
DeleteSegmentDL(ctx context.Context, segment string) error
