package storj

import (
	"context"
	"io"

	. "github.com/beyondstorage/go-storage/v4/types"
)

func (s *Storage) copy(ctx context.Context, src string, dst string, opt pairStorageCopy) (err error) {
	panic("not implemented")
}

func (s *Storage) create(path string, opt pairStorageCreate) (o *Object) {
	if opt.HasObjectMode && opt.ObjectMode.IsDir() {
		path = path + "/"
		o = NewObject(s, true)
		o.Mode = ModeDir
	} else {
		o = NewObject(s, false)
		o.Mode = ModeRead
	}
	o.ID = s.getAbsPath(path)
	o.Path = path
	return o
}

func (s *Storage) createDir(ctx context.Context, path string, opt pairStorageCreateDir) (o *Object, err error) {
	panic("not implemented")
}

func (s *Storage) delete(ctx context.Context, path string, opt pairStorageDelete) (err error) {
	panic("not implemented")
}

func (s *Storage) list(ctx context.Context, path string, opt pairStorageList) (oi *ObjectIterator, err error) {
	panic("not implemented")
}

func (s *Storage) metadata(opt pairStorageMetadata) (meta *StorageMeta) {
	meta = NewStorageMeta()
	meta.WorkDir = s.workDir
	return meta
}

func (s *Storage) move(ctx context.Context, src string, dst string, opt pairStorageMove) (err error) {
	panic("not implemented")
}

func (s *Storage) read(ctx context.Context, path string, w io.Writer, opt pairStorageRead) (n int64, err error) {
	buckets := s.project.ListBuckets(ctx, nil)

	if err := buckets.Err(); err != nil {
		return 0, err
	}

	object, err := s.project.DownloadObject(ctx, s.bucket.Name, path, nil)
	if err != nil {
		return 0, err
	}
	defer object.Close()

	return io.Copy(w, object)
}

func (s *Storage) stat(ctx context.Context, path string, opt pairStorageStat) (o *Object, err error) {
	panic("not implemented")
}

func (s *Storage) write(ctx context.Context, path string, r io.Reader, size int64, opt pairStorageWrite) (n int64, err error) {
	upload, err := s.project.UploadObject(ctx, s.bucket.Name, path, nil)
	if err != nil {
		return 0, err
	}

	// Copy the data to the upload.
	_, err = io.Copy(upload, r)
	if err != nil {
		_ = upload.Abort()
		return 0, err
	}

	// Commit the uploaded object.
	err = upload.Commit()
	if err != nil {
		return 0, err
	}
}
