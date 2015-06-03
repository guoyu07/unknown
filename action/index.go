package action

import(
	"fmt"
	"github.com/erzha/http/server"
)

type IndexAction struct {
	base
}

func (action *IndexAction) DoGet(ctx context.Context, sapi *server.Sapi) {

	name := sapi.Get.Get("name")

	//if $name is dir -> show $dir/README.md
	//if $name is path -> show $file.md
}
