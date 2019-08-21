module streamConsole

go 1.12

require (
	github.com/astaxie/beego v1.11.1
	github.com/go-sql-driver/mysql v1.4.1
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/smartystreets/goconvey v0.0.0-20190731233626-505e41936337
	github.com/tidwall/gjson v1.3.2
	gopkg.in/mgo.v2 v2.0.0-20180705113604-9856a29383ce
)

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190820162420-60c769a6c586
	golang.org/x/net => github.com/golang/net v0.0.0-20190813141303-74dc4d7220e7
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190813064441-fde4db37ae7a
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190821162956-65e3620a7ae7
	golang.org/x/xerrors => github.com/golang/xerrors v0.0.0-20190717185122-a985d3407aa7

)
