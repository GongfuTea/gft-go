package pool

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/GongfuTea/gft-go/edu/gs/xj"
	"github.com/GongfuTea/gft-go/x/jsonx"
	"github.com/GongfuTea/gft-go/x/timex"
)

var zxsFiles []string

type ImportFromXxw struct {
	files []string
}

func NewImportFromXxw() *ImportFromXxw {
	x := &ImportFromXxw{}

	return x
}

func (x *ImportFromXxw) LoadAndValid() {
	files, err := ioutil.ReadDir("./data/gs/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
		if strings.HasPrefix(file.Name(), "zxs") {
			zxsFiles = append(zxsFiles, file.Name()[3:11])
		}
	}

	jsonx.PrintAsJson("files", zxsFiles)
	x.files = zxsFiles
	x.validFile()
}

func (x *ImportFromXxw) UpdateDb() int {
	total := 0
	for _, rq := range x.files {
		total += x.importFromXlsx(rq)
	}
	println("load files done! total:", total)
	return total
}

func (x *ImportFromXxw) validFile() {
	file := xj.NewXxwFile()
	for i, rq := range x.files {
		println(fmt.Sprintf("load file: %d  %s", i, rq))
		file.LoadFile(fmt.Sprintf("./data/gs/zxs%s.xlsx", rq))
	}
}

func (x *ImportFromXxw) importFromXlsx(rq string) int {
	file := xj.NewXxwFile()
	file.LoadFile(fmt.Sprintf("./data/gs/zxs%s.xlsx", rq))

	xwwrq := timex.ParseLocalUTC(rq)
	ls := file.XjList
	println(len(ls))

	items := make([]any, 0)

	for _, r := range ls {
		item := &GftGsXjPoolItem{
			GftGsXj: *r,
			Xxwrq:   &xwwrq,
		}
		item.Id = fmt.Sprintf("%s-%s", item.Xh, rq)
		item.CreatedAt = time.Now()
		items = append(items, item)
	}

	start := time.Now()

	if res, err := GsXjPoolItemRepo.Coll().InsertMany(context.Background(), items); err != nil {
		fmt.Printf("err:%s, %v\n", rq, err)
	} else {
		idCount := len(res.InsertedIDs)
		if idCount != len(ls) {
			println("插入数量不一致", idCount, len(ls))
		}
	}

	println(rq, len(ls), "time:", time.Since(start).String())

	return len(ls)
}
