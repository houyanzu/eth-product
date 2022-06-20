package loop

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func Loop(t int, fs ...func(wg *sync.WaitGroup)) {
	c := make(chan os.Signal)
	//监听指定信号 ctrl+c kill
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	fmt.Println("程序已启动...")
	wg := &sync.WaitGroup{}

	go func() {
		for {
			fmt.Println("start")
			for _, f := range fs {
				wg.Add(1)
				f(wg)
			}
			fmt.Println("ok")
			time.Sleep(time.Duration(t) * time.Second)
		}
	}()

	<-c
	fmt.Println("正在退出，请稍后...")
	wg.Wait()
	fmt.Println("退出成功！")
	os.Exit(0)
}
