package minio_driver

import (
	"bytes"
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	. "storage/pkg/types"
)

var bucketName = "zcbucket"

type minioStore struct {
	client *minio.Client
}

func NewMinioStore(ctx context.Context) (*minioStore, error) {
	// to-do read conf
	var endPoint = "127.0.0.1:9000"
	var accessKeyID = "testminio"
	var secretAccessKey = "testminio"
	var useSSL = false

	minioClient, err := minio.New(endPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		return nil, err
	}

	bucketExists, err := minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		return nil, err
	}

	if !bucketExists {
		err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return nil, err
		}
	}
	return &minioStore{
		client: minioClient,
	}, nil
}

func (s *minioStore) PUT(ctx context.Context, key Key, value Value) error {
	reader := bytes.NewReader(value)
	_, err := s.client.PutObject(ctx, bucketName, string(key), reader, int64(len(value)), minio.PutObjectOptions{})

	if err != nil {
		return err
	}

	return err
}

func (s *minioStore) GET(ctx context.Context, key Key) (Value, error) {
	object, err := s.client.GetObject(ctx, bucketName, string(key), minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	size := 256*1024
	buf := make([]byte, size)
	n, err := object.Read(buf)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return buf[:n], err
}

func (s *minioStore) GetByPrefix(ctx context.Context, prefix Key, keyOnly bool) ([]Key, []Value, error) {
	objects := s.client.ListObjects(ctx, bucketName, minio.ListObjectsOptions{Prefix: string(prefix)})

	var objectsKeys []Key
	var objectsValues []Value

	for object := range objects {
		objectsKeys = append(objectsKeys, []byte(object.Key))
		if !keyOnly{
			value, err := s.GET(ctx, []byte(object.Key))
			if err != nil{
				return nil, nil, err
			}
			objectsValues = append(objectsValues, value)
		}
	}

	return objectsKeys, objectsValues, nil

}

func (s *minioStore) Scan(ctx context.Context, keyStart Key, keyEnd Key, limit int, keyOnly bool) ([]Key, []Value, error){
	var keys []Key
	var values []Value
	limitCount := uint(limit)
	for object := range s.client.ListObjects(ctx, bucketName, minio.ListObjectsOptions{Prefix:  string(keyStart)}) {
		if object.Key <= string(keyEnd) {
			keys = append(keys, []byte(object.Key))
			if !keyOnly {
				value, err := s.GET(ctx, []byte(object.Key))
				if err != nil {
					return nil, nil, err
				}
				values = append(values, value)
			}
		}
		limitCount--;
		if limitCount <= 0{
			break
		}
	}

	return keys, values, nil
}

func (s *minioStore) Delete(ctx context.Context, key Key) error {
	err := s.client.RemoveObject(ctx, bucketName, string(key), minio.RemoveObjectOptions{})
	return err
}

func (s *minioStore) DeleteByPrefix(ctx context.Context, prefix Key) error{
	objectsCh := make(chan minio.ObjectInfo)

	go func() {
		defer close(objectsCh)

		for object := range s.client.ListObjects(ctx, bucketName, minio.ListObjectsOptions{Prefix:  string(prefix)}){
			objectsCh <- object
		}
	}()

	for rErr := range s.client.RemoveObjects(ctx, bucketName, objectsCh, minio.RemoveObjectsOptions{GovernanceBypass: true}){
		if rErr.Err != nil {
			return rErr.Err
		}
	}
	return nil
}

func (s *minioStore) DeleteRange(ctx context.Context, keyStart Key, keyEnd Key) error {
	objectsCh := make(chan minio.ObjectInfo)

	go func() {
		defer close(objectsCh)

		for object := range s.client.ListObjects(ctx, bucketName, minio.ListObjectsOptions{Prefix:  string(keyStart)}){
			if object.Key <= string(keyEnd) {
				objectsCh <- object
			}
		}
	}()

	for rErr := range s.client.RemoveObjects(ctx, bucketName, objectsCh, minio.RemoveObjectsOptions{GovernanceBypass: true}){
		if rErr.Err != nil {
			return rErr.Err
		}
	}
	return nil
}
