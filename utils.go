package storj

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/beyondstorage/go-endpoint"
	"github.com/beyondstorage/go-storage/v4/services"
	"github.com/beyondstorage/go-storage/v4/types"

	"storj.io/uplink"
)

// Storage is the example client.
type Storage struct {
	project      uplink.Project
	bucket       uplink.Bucket
	defaultPairs DefaultStoragePairs
	features     StorageFeatures

	workDir string
	types.UnimplementedStorager
	types.UnimplementedDirer
}

// String implements Storager.String
func (s *Storage) String() string {
	return fmt.Sprintf("Storager storj {WorkDir: %s}", s.workDir)
}

func NewStorager(pairs ...types.Pair) (types.Storager, error) {
	opt, err := parsePairStorageNew(pairs)
	if err != nil {
		return nil, err
	}

	st := &Storage{
		workDir: "/",
	}
	if opt.HasWorkDir {
		st.workDir = opt.WorkDir
	}

	_, err = endpoint.Parse(opt.Endpoint)
	if err != nil {
		return nil, err
	}

	//TODO
	return nil, nil
}

func (s *Storage) formatError(op string, err error, path ...string) error {
	if err == nil {
		return nil
	}
	return services.StorageError{
		Op:       op,
		Err:      formatError(err),
		Storager: s,
		Path:     path,
	}
}

func formatError(err error) error {
	if _, ok := err.(services.InternalError); ok {
		return err
	}

	switch {
	case errors.Is(err, os.ErrNotExist):
		return fmt.Errorf("%w: %v", services.ErrObjectNotExist, err)
	case errors.Is(err, os.ErrPermission):
		return fmt.Errorf("%w: %v", services.ErrPermissionDenied, err)
	default:
		return fmt.Errorf("%w: %v", services.ErrUnexpected, err)
	}
}

func (s *Storage) getAbsPath(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return s.workDir + path
}
