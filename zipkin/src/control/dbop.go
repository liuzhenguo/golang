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
	chstr := make(chan string)
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
		}
		chstr <- uname
		defer rows.Close()
	}()
	select {

	case <-ctx.Done():
		logs.Logger.Info("dbop其他调用失败", ctx.Err())
		return ctx.Err()

	case e := <-err:
		logs.Logger.Info("dbop自己调用失败", e)
		return e
	case u := <-chstr:
		fmt.Println("dbop data :", u)
		return nil

	}

}
