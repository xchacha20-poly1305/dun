//go:build debug

package dunmain

import (
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"runtime/debug"

	"github.com/sagernet/sing-box/log"
	"github.com/sagernet/sing/common/json"
	"github.com/sagernet/sing/common/json/badjson"

	"github.com/dustin/go-humanize"
)

func init() {
	http.HandleFunc("/debug/gc", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusNoContent)
		go debug.FreeOSMemory()
	})
	http.HandleFunc("/debug/memory", func(writer http.ResponseWriter, request *http.Request) {
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)

		var memObject badjson.JSONObject
		memObject.Put("heap", humanize.IBytes(memStats.HeapInuse))
		memObject.Put("stack", humanize.IBytes(memStats.StackInuse))
		memObject.Put("idle", humanize.IBytes(memStats.HeapIdle-memStats.HeapReleased))
		memObject.Put("goroutines", runtime.NumGoroutine())
		memObject.Put("rss", rusageMaxRSS())

		encoder := json.NewEncoder(writer)
		encoder.SetIndent("", "  ")
		encoder.Encode(memObject)
	})
	go func() {
		err := http.ListenAndServe("0.0.0.0:8964", nil)
		if err != nil {
			log.Debug(err)
		}
	}()
}
