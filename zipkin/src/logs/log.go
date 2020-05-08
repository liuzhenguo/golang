package logs

import (
	"fmt"

	"github.com/cihub/seelog"
)

var Logger seelog.LoggerInterface

func init() {
	disableLog()
	loadAppConfig()
}
func loadAppConfig() {
	appconfig := `<seelog >
    <outputs formatid="main">
        <filter levels="info,critical">
            <console />
        </filter>
        <filter levels="info,critical">
            <rollingfile formatid="info" type="size" filename="../log/zipkin_info.log" maxsize="100000000" maxrolls="100" />
        </filter>
        <filter levels="critical,error">
			<rollingfile formatid="critical" type="size" filename="../log/zipkin_critical.log" maxsize="100000000" maxrolls="100" />
        </filter>
    </outputs>
    <formats>
        <format id="main" format="%Date/%Time [%LEV] %Func[%Line] %Msg%n"/>
        <format id="info" format="%Date/%Time [%LEV] %File %Func[%Line] %Msg%n"/>
        <format id="critical" format="%Date/%Time [%LEV] %File %Func[%Line]  %Msg %n"/>
    </formats>
	</seelog>`
	logger, err := seelog.LoggerFromConfigAsBytes([]byte(appconfig))
	if err != nil {
		fmt.Println(err)
		return
	}
	userLogger(logger)
}
func userLogger(newLogger seelog.LoggerInterface) {
	Logger = newLogger
}
func disableLog() {
	Logger = seelog.Disabled
}
