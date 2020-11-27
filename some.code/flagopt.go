package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type StringArray []string

func (a *StringArray) Get() interface{} {
	fmt.Println("22222222")
	return []string(*a)
}

func (a *StringArray) Set(s string) error {
	fmt.Println("1111111111")
	*a = append(*a, s)
	return nil
}
func (a *StringArray) String() string {
	fmt.Println("**************")
	return strings.Join(*a, ",")

}

func flagSet() *flag.FlagSet {
	fs := flag.NewFlagSet("nsqd", flag.ExitOnError)

	fs.Bool("version", false, "print version string")
	fs.String("log-level", "info", "set log verbosity: debug, info, warn, error, or fatal")
	fs.String("log-prefix", "[nsq_to_file] ", "log message prefix")

	fs.String("channel", "nsq_to_file", "nsq channel")
	fs.Int("max-in-flight", 200, "max number of messages to allow in flight")

	fs.String("output-dir", "/tmp", "directory to write output files to")
	fs.String("work-dir", "", "directory for in-progress files before moving to output-dir")
	fs.String("datetime-format", "%Y-%m-%d_%H", "strftime compatible format for <DATETIME> in filename format")
	fs.String("filename-format", "<TOPIC>.<HOST><REV>.<DATETIME>.log", "output filename format (<TOPIC>, <HOST>, <PID>, <DATETIME>, <REV> are replaced. <REV> is increased when file already exists)")
	fs.String("host-identifier", "", "value to output in log filename in place of hostname. <SHORT_HOST> and <HOSTNAME> are valid replacement tokens")
	fs.Int("gzip-level", 6, "gzip compression level (1-9, 1=BestSpeed, 9=BestCompression)")
	fs.Bool("gzip", false, "gzip output files.")
	fs.Bool("skip-empty-files", false, "skip writing empty files")
	fs.Duration("topic-refresh", time.Minute, "how frequently the topic list should be refreshed")
	fs.String("topic-pattern", "", "only log topics matching the following pattern")

	fs.Int64("rotate-size", 0, "rotate the file when it grows bigger than `rotate-size` bytes")
	fs.Duration("rotate-interval", 0, "rotate the file every duration")
	fs.Duration("sync-interval", 30*time.Second, "sync file to disk every duration")

	fs.Duration("http-client-connect-timeout", 2*time.Second, "timeout for HTTP connect")
	fs.Duration("http-client-request-timeout", 5*time.Second, "timeout for HTTP request")

	nsqdTCPAddrs := StringArray{}
	lookupdHTTPAddrs := StringArray{}
	topics := StringArray{}
	consumerOpts := StringArray{}

	//自定义flag
	fs.Var(&nsqdTCPAddrs, "nsqd-tcp-address", "nsqd TCP address (may be given multiple times)")
	fs.Var(&lookupdHTTPAddrs, "lookupd-http-address", "lookupd HTTP address (may be given multiple times)")
	fs.Var(&topics, "topic", "nsq topic (may be given multiple times)")
	fs.Var(&consumerOpts, "consumer-opt", "option to passthrough to nsq.Consumer (may be given multiple times, http://godoc.org/github.com/nsqio/go-nsq#Config)")

	return fs
}

func main() {
	fs := flagSet()

	fs.Parse(os.Args[1:])

	if args := fs.Args(); len(args) > 0 {
		fmt.Println("44444:%s", args)
	}
	//实现了get函数
	ipstr := fs.Lookup("nsqd-tcp-address").Value.(flag.Getter).Get()
	fmt.Println(ipstr)
	fmt.Println(fs)
}
