package control

import (
	"context"
	"fmt"
	"zipkin/src/envbuild"
	"zipkin/src/logs"
)

func GetDbData(ctx context.Context) error {
	//操作数据库获取数据
	db := envbuild.Getdb()
	err := make(chan error)
	var uname string
	go func() {

		sql := fmt.Sprintf("select username from wanbu_data_user where userid = 1")

		rows, err := db.Query(sql)
		if err != nil {
			fmt.Println(err)
			return
		}

		for rows.Next() {
			rows.Scan(&uname)
			logs.Logger.Info("获取userid:", uname)
		}
		defer rows.Close()
	}()

	select {

	case <-ctx.Done():
		logs.Logger.Info("其他调用失败", ctx.Err())
		return ctx.Err()

	case e := <-err:
		logs.Logger.Info("dbop获取数据库失败", e)
		return e

	}

}
