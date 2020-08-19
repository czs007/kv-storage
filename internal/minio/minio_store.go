package minio_driver

import (
	"bytes"
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"storage/internal/minio/codec"
	. "storage/pkg/types"
)

type minioStore struct {
	client *minio.Client
}

func RawBatchGet(ctx context.Context, key Key) []Value{

}

func (s *minioStore) RawGet(ctx context.Context, key Key) Value{
	key1 := string(key)
	object, err := s.client.GetObject(ctx, bucketName, key1, minio.GetObjectOptions{})

	if err != nil {
		return nil
	}

	size := 256 * 1024
	buf := make([]byte, size)
	n, err := object.Read(buf)
	if err != nil && err != io.EOF {
		return nil
	}
	return buf[:n]
}
func (s *minioStore) RawPut(ctx context.Context, key Key, value Value){

	reader := bytes.NewReader(value)
	s.client.PutObject(ctx, bucketName, string(key), reader, int64(len(value)), minio.PutObjectOptions{})
}

func (s *minioStore) RawDeleteAll(ctx context.Context, key Key){

	for i := 0; i < len(key); i++ {
		s.client.RemoveObjects()
	}

}
func (s *minioStore) RawDelete(ctx context.Context, key Key){

}

func (s *minioStore)	Get(ctx context.Context, key Key, timestamp Timestamp) Value{

}
func (s *minioStore)	BatchGet(ctx context.Context, keys []Key, timestamp Timestamp) []Value{

}

func (s *minioStore)	GetAll(ctx context.Context, key Key, withValue bool) ([]Timestamp, []Key, []Value){

}
func (s *minioStore)	ScanLE(ctx context.Context, key Key, timestamp Timestamp, withValue bool) ([]Timestamp, []Key, []Value){

}
func (s *minioStore)	ScanGE(ctx context.Context, key Key, timestamp Timestamp, withValue bool) ([]Timestamp, []Key, []Value){

}
func (s *minioStore)	ScanRange(ctx context.Context, key Key, start Timestamp, end Timestamp, withValue bool) ([]Timestamp, []Key, []Value){

}
func (s *minioStore)	PUT(ctx context.Context, key Key, value Value, timestamp Timestamp, suffix string){

}
func (s *minioStore)	DeleteLE(ctx context.Context, key Key, timestamp Timestamp){

}
func (s *minioStore)	DeleteGE(ctx context.Context, key Key, timestamp Timestamp){

}
func (s *minioStore)	RangeDelete(ctx context.Context, key Key, start Timestamp, end Timestamp){

}
func (s *minioStore)	LogPut(ctx context.Context, key Key, value Value, timestamp Timestamp, suffix string){

}
func (s *minioStore)	LogFetch(ctx context.Context, start Timestamp, end Timestamp, channels []int){

}