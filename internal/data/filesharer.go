package data

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
	"sync"

	v1 "filesharer/api/file/v1"
	"fmt"

	"net/http"

	"filesharer/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type FilesharerRepo struct {
	data       *Data
	log        *log.Helper
	httpClient *http.Client
}

func getAllFiles(path string) []*os.FileInfo {
	files, err := filepath.Glob(path + "/*")
	if err != nil {
		return nil
	}
	resp := make([]*os.FileInfo, 0)
	wg := &sync.WaitGroup{}
	ch := make(chan *os.FileInfo, len(files))
	for _, v := range files {
		wg.Add(1)
		go func() {
			defer wg.Done()
			v := v
			info, err := os.Stat(v)
			if err != nil {
				return
			}
			if info.IsDir() {
				allFiles := getAllFiles(v)
				for _, f := range allFiles {
					ch <- f
				}
			}
			ch <- &info
		}()
	}

	chDone := make(chan struct{})
	go func() {
		for v := range ch {
			resp = append(resp, v)
		}

		chDone <- struct{}{}
	}()
	wg.Wait()
	close(ch)
	<-chDone
	return resp
}
func getAllFilesByWalk(path string) []*os.FileInfo {
	resp := make([]*os.FileInfo, 0)
	stat, err := os.Stat(path)
	if err != nil {
		return nil
	}
	if stat.IsDir() {
		open, err := os.Open(path)
		if err != nil {
			return nil
		}
		readdir, err := open.Readdir(-1)
		if err != nil {
			return nil
		}
		resp = append(resp, &stat)
		for _, v := range readdir {
			resp = append(resp, getAllFilesByWalk(filepath.Join(path, v.Name()))...)
		}

	} else {
		resp = append(resp, &stat)
	}
	return resp
}
func (f *FilesharerRepo) ListByAddr(ctx context.Context, req *v1.ListByAddrRequest) (*v1.ListByAddrReply, error) {
	info, err := os.Stat(req.Path)
	//if errors.Is(err,os.ErrNotExist) {

	//}
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return nil, errors.New("不是文件夹")
	}

	files, err := filepath.Glob(req.Path + "/*")
	if err != nil {
		return nil, err
	}
	resp := &v1.ListByAddrReply{
		Data: make([]*v1.ListByAddrReplyItem, 0, len(files)),
	}
	wg := &sync.WaitGroup{}
	ch := make(chan *v1.ListByAddrReplyItem, len(files))
	for _, v := range files {
		wg.Add(1)
		go func() {
			defer wg.Done()
			v := v
			info, err := os.Stat(v)
			if err != nil {
				return
			}
			ch <- &v1.ListByAddrReplyItem{
				Path:  v,
				IsDir: info.IsDir(),
				Size:  info.Size(),
			}
		}()
	}
	chDone := make(chan struct{})
	go func() {
		for v := range ch {
			resp.Data = append(resp.Data, v)
		}

		chDone <- struct{}{}
	}()
	wg.Wait()
	close(ch)
	<-chDone

	return resp, nil
}

func (f *FilesharerRepo) GetDetailByAddr(ctx context.Context, req *v1.GetDetailByAddrRequest) (*v1.GetDetailByAddrReply, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FilesharerRepo) DownloadByAddr(ctx context.Context, req *v1.DownloadByAddrRequest) (*v1.DownloadByAddrReply, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FilesharerRepo) ListNode(ctx context.Context, req *v1.ListNodeRequest) (*v1.ListNodeReply, error) {

	httreq, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("http://%s/v1/catalog/service/filesharer", f.data.getConsulInfoClient.ip), nil)
	if err != nil {
		return nil, err
	}

	// 发送 HTTP 请求
	httpresp, err := f.httpClient.Do(httreq)
	if err != nil {
		return nil, err
	}

	// 读取 HTTP 响应主体
	body, err := io.ReadAll(httpresp.Body)
	if err != nil {
		return nil, err
	}

	resp := &v1.ListNodeReply{
		Data: make([]*v1.ListNodeReplyItem, 0),
	}

	// 将 HTTP 响应主体解组到 Person 结构体中
	if err := json.Unmarshal(body, &resp.Data); err != nil {
		return nil, err
	}
	return resp, nil
}

// NewFilesharerRepo .
func NewFilesharerRepo(data *Data, logger log.Logger) biz.FilesharerRepo {
	return &FilesharerRepo{
		data:       data,
		log:        log.NewHelper(logger),
		httpClient: &http.Client{},
	}
}
