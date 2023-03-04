package dunmain

import (
	"context"
	"os"
	"time"

	"github.com/sagernet/sing-box/log"
)

func DisableColor() {
	disableColor = true
	log.SetStdLogger(log.NewDefaultFactory(context.Background(), log.Formatter{BaseTime: time.Now(), DisableColors: true}, os.Stderr, "", nil, false).Logger())
}
