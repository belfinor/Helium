package categorizer

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2018-12-11

import (
	"container/list"
	"fmt"
	"io"

	"github.com/belfinor/Helium/log"
	"github.com/belfinor/Helium/text"
	"github.com/belfinor/Helium/text/corpus/categorizer/statements"
	"github.com/belfinor/Helium/text/corpus/index"
	"github.com/belfinor/Helium/text/corpus/opts"
	"github.com/belfinor/Helium/text/corpus/tags"
	"github.com/belfinor/Helium/time/timer"
)

const (
	bufSize int = 32
)

type Engine interface {
	Proc(io.RuneReader) ([]string, bool)
}

type engine struct {
	st    *statements.Statements
	slang int
	buf   *list.List
}

func New() Engine {
	return &engine{
		st:  statements.New(),
		buf: list.New(),
	}
}

func (eng *engine) bufProc() {

	buf := eng.buf
	st := eng.st

	if buf.Len() == 0 {
		return
	}

	v := buf.Front().Value
	buf.Remove(buf.Front())

	if v == nil {
		return
	}

	ws := v.(*index.Record)

	if ws == nil {
		return
	}

	fmt.Println(ws.Name)

	if ws.HasOpt(opts.Opt(opts.OPT_EOS)) {
		st.Tact()
		return
	}

	addTag := func(t uint16) {
		st.Add(t)
	}

	ws.ForEachTags(addTag)
}

func (eng *engine) Proc(rh io.RuneReader) ([]string, bool) {

	// init calc process time
	tm := timer.New()
	defer func() {
		log.Debug(fmt.Sprintf("categorizer.Proc time %.4fs", tm.DeltaFloat()))
	}()

	// make word stream
	wordStream := text.WordStream(rh, text.WSO_ENDS, text.WSO_NO_URLS, text.WSO_HASHTAG, text.WSO_NO_XXX_COLON)

	buf := eng.buf
	st := eng.st

	// read forms stream
	for ws := range wsStream(wordStream, &eng.slang) {

		if buf.Len() >= bufSize {
			eng.bufProc()
		}

		buf.PushBack(ws)
	}

	for buf.Len() > 0 {
		eng.bufProc()
	}

	result := st.Finish()
	res := []string{}

	for _, v := range result {
		val := tags.FromCode(v)
		if val != "" {
			res = append(res, val)
		}
	}

	return res, eng.slang > 0
}
