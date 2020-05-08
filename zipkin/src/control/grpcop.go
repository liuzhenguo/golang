package control

import (
	"Backend/Stat/chufangrefresh/logs"
	"context"

	"fmt"
	//"zipkin/src/envbuild"
)

func GetRpcData(ctx context.Context) error {

	//操作RPC获取数据

	err := make(chan error)

	go func() {
		fmt.Println("获取GRPC中的数据")
	}()

	select {
	case <-ctx.Done():
		logs.Logger.Info("别的程序出错了")
		return ctx.Err()
	case e := <-err:
		logs.Logger.Info("grpc出错了")
		return e

	}

}
