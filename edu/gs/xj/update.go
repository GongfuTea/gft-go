package xj

import (
	"fmt"
	"time"

	"github.com/GongfuTea/gft-go/base"
	"github.com/GongfuTea/gft-go/types"
	"github.com/GongfuTea/gft-go/x/jsonx"
	"github.com/GongfuTea/gft-go/x/timex"
)

func Set转专业(xh string, diff map[string]types.GftTimelineDiff, rq time.Time, note string) {

	ls, _ := GsXjRepo.FindByXh(xh)
	println("ss", xh, len(ls))
	if s := ls.FindByXhLastVersion(xh); s != nil {
		jsonx.PrintAsJson("Set转专业", s)

		// 修改旧记录
		end1 := rq.Add(-time.Second)
		s.TlEnd = &end1
		GsXjRepo.Save(s)

		s.UpdateByDiff(diff)
		s.GftTimeline = types.GftTimeline{
			TlStart:   &rq,
			TlVersion: s.TlVersion + 1,
			TlDiff:    diff,
			TlNote:    note,
		}
		s.Id = fmt.Sprintf("%s-%d", s.Xh, s.TlVersion)
		s.CreatedAt = time.Now()
		GsXjRepo.Save(s)
	}
}

func Set退学(xh string, txrq time.Time, note string) {
	dqzt, _ := base.DictItemRepo.FindByItemName("JY.XSDQZT", "退学")

	if dqzt == nil {
		panic("获取学生当前状态字典错误")
	}

	ls, _ := GsXjRepo.FindByXh(xh)
	println("ss", xh, len(ls))
	if s := ls.FindByXhLastVersion(xh); s != nil {
		jsonx.PrintAsJson("Set退学", s)

		// 修改旧记录
		end1 := txrq.Add(-time.Second)
		s.TlEnd = &end1
		GsXjRepo.Save(s)

		// 添加新记录

		diff := make(map[string]types.GftTimelineDiff, 0)
		diff["Xsdqztm"] = types.GftTimelineDiff{s.Xsdqztm, dqzt.Code}
		diff["Xsdqzt"] = types.GftTimelineDiff{s.Xsdqzt, dqzt.Name}
		s.Xsdqztm = dqzt.Code
		s.Xsdqzt = dqzt.Name

		tlStart := txrq
		tlEnd := timex.DateEnd(time.Now().AddDate(100, 0, 0))

		s.GftTimeline = types.GftTimeline{
			TlStart:   &tlStart,
			TlEnd:     &tlEnd,
			TlVersion: s.TlVersion + 1,
			TlDiff:    diff,
			TlNote:    note,
		}
		s.Id = fmt.Sprintf("%s-%d", s.Xh, s.TlVersion)
		s.CreatedAt = time.Now()
		GsXjRepo.Save(s)
	}

}
