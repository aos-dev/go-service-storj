package storj

import (
	"context"
	"io"
	"strings"

	"github.com/beyondstorage/go-storage/v4/services"
	. "github.com/beyondstorage/go-storage/v4/types"
	"storj.io/uplink"
)

func (s *Storage) create(path string, opt pairStorageCreate) (o *Object) {
	if opt.HasObjectMode && opt.ObjectMode.IsDir() {
		path += "/"
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

func (s *Storage) delete(ctx context.Context, path string, opt pairStorageDelete) (err error) {
	rp := s.getAbsPath(path)

	_, err = s.project.DeleteObject(ctx, s.name, rp)
	return err
}

func (s *Storage) list(ctx context.Context, path string, opt pairStorageList) (oi *ObjectIterator, err error) {
	rp := s.getAbsPath(path)
	if !opt.HasListMode || opt.ListMode.IsDir() {
		nextFn := func(ctx context.Context, page *ObjectPage) error {
			options := uplink.ListObjectsOptions{Prefix: rp, System: true}
			dirObject := s.project.ListObjects(ctx, s.name, &options)
			for dirObject.Next() {
				if dirObject.Item().Key == rp {
					continue
				}
				o := NewObject(s, true)
				o.Path = dirObject.Item().Key[len(rp):]
				if dirObject.Item().IsPrefix {
					o.Mode |= ModeDir
				} else {
					o.Mode |= ModeRead
				}
				o.SetContentLength(dirObject.Item().System.ContentLength)
				page.Data = append(page.Data, o)
			}
			return IterateDone
		}
		oi = NewObjectIterator(ctx, nextFn, nil)
		return oi, err
	} else {
		return nil, services.ListModeInvalidError{Actual: opt.ListMode}
	}
}

func (s *Storage) metadata(opt pairStorageMetadata) (meta *StorageMeta) {
	meta = NewStorageMeta()
	meta.WorkDir = s.workDir
	meta.Name = s.name
	return meta
}

func (s *Storage) read(ctx context.Context, path string, w io.Writer, opt pairStorageRead) (n int64, err error) {
	rp := s.getAbsPath(path)
	download, err := s.project.DownloadObject(ctx, s.name, rp, nil)
	if err != nil {
		return 0, services.ErrObjectNotExist
	}
	n, err = io.Copy(w, download)
	return n, err
}

func (s *Storage) stat(ctx context.Context, path string, opt pairStorageStat) (o *Object, err error) {
	rp := s.getAbsPath(path)
	object, err := s.project.StatObject(ctx, s.name, rp)
	if err != nil {
		return nil, services.ErrObjectNotExist
	}
	o = NewObject(s, true)
	o.Path = path
	if object.IsPrefix {
		o.Mode |= ModeDir
	} else {
		o.Mode |= ModeRead
	}
	o.SetContentLength(object.System.ContentLength)
	o.SetSystemMetadata(object.System)
	o.SetUserMetadata(object.Custom)
	return o, nil
}

func (s *Storage) write(ctx context.Context, path string, r io.Reader, size int64, opt pairStorageWrite) (n int64, err error) {
	rp := s.getAbsPath(path)

	upload, err := s.project.UploadObject(ctx, s.name, rp, nil)
	if err != nil {
		return 0, err
	}
	if r == nil {
		if size == 0 {
			r = strings.NewReader("")
		} else {
			return 0, services.ServiceError{}
		}
	}

	n, err = io.CopyN(upload, r, size)
	if err != nil {
		return 0, err
	}
	err = upload.Commit()
	if err != nil {
		return 0, err
	}
	return n, err
}
