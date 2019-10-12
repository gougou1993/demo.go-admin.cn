package main

import (
	"github.com/GoAdminGroup/demo/ecommerce"
	"github.com/GoAdminGroup/demo/login"
	"github.com/GoAdminGroup/demo/pages"
	_ "github.com/chenhg5/go-admin/adapter/gin"
	"github.com/chenhg5/go-admin/engine"
	"github.com/chenhg5/go-admin/examples/datamodel"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/plugins/admin"
	"github.com/chenhg5/go-admin/template"
	"github.com/chenhg5/go-admin/template/types"
	"github.com/gin-gonic/gin"
	template2 "html/template"
)

func main() {
	r := gin.Default()

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(datamodel.Generators)

	// add generator, first parameter is the url prefix of table when visit.
	// example:
	//
	// "user" => http://localhost:9033/admin/info/user
	//
	adminPlugin.AddGenerator("user", datamodel.GetUserTable)

	template.AddLoginComp(login.GetLoginComponent())

	rootPath := "/data/www/go-admin"
	//rootPath = "."

	cfg := config.ReadFromJson(rootPath + "/config.json")
	cfg.CustomFootHtml = template2.HTML(`<div style="display:none;">
    <script type="text/javascript" src="https://v1.cnzz.com/z_stat.php?id=1277862090&web_id=1277862090"></script>
</div>`)

	if err := eng.AddConfig(cfg).AddPlugins(adminPlugin).Use(r); err != nil {
		panic(err)
	}

	r.Static("/uploads", rootPath+"/uploads")

	// you can custom your pages like:

	r.GET("/admin", func(ctx *gin.Context) {
		engine.Content(ctx, func(ctx interface{}) (types.Panel, error) {
			return ecommerce.GetContent()
		})
	})

	r.GET("/admin/form1", func(ctx *gin.Context) {
		engine.Content(ctx, func(ctx interface{}) (types.Panel, error) {
			return pages.GetForm1Content()
		})
	})

	r.GET("/admin/e-commerce", func(ctx *gin.Context) {
		engine.Content(ctx, func(ctx interface{}) (types.Panel, error) {
			return ecommerce.GetContent()
		})
	})

	_ = r.Run(":9033")
}
