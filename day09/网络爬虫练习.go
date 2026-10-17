package main

import (
	"fmt"
	"sync" // 引入 sync 包，这个包里装着我们需要的两件神器：Mutex（锁）和 WaitGroup（等待组）
)

// Fetcher 是一个接口，相当于一个“合同”。
// 只要谁实现了 Fetch 这个方法，谁就可以被当成爬虫去抓取数据。
type Fetcher interface {
	// Fetch 接收一个 url，返回网页内容(body)、网页里的所有链接(urls) 和 错误信息(err)
	Fetch(url string) (body string, urls []string, err error)
}

// ================= 核心改造部分开始 =================

// 1. 定义一个“安全”的缓存结构体
// 为什么叫 SafeCache？因为多个协程同时操作 Map 会崩溃，必须加锁保护。
type SafeCache struct {
	mu   sync.Mutex       // 这是一把锁。mu 是 mutex 的缩写
	seen map[string]bool  // 这是一个记录本。记录哪个 url 已经被抓过（true表示抓过）
}

// 创建一个全局的缓存实例（整个程序共用这一个记录本）
// 记得用 make 初始化 map，不然它是 nil（空指针），一写就报错
var cache = SafeCache{seen: make(map[string]bool)}

// 2. 检查 URL 是否访问过的方法
// 注意：这里有个小括号 (c *SafeCache)，代表这个方法属于 SafeCache，且是用指针接收（这样修改的记录本才会生效）
func (c *SafeCache) isVisited(url string) bool {
	
	// 【上锁】进门先锁门。防止 10 个协程同时来查记录本导致数据错乱
	c.mu.Lock()
	
	// 【延迟解锁】defer 意思是：当这个函数准备结束（return 之前）的那一刻，自动执行这行代码。
	// 这样写的好处是，不管中间发生什么错误，门一定会被打开，不会死锁。
	defer c.mu.Unlock()
	
	// 查记录本：看看这个 url 以前来过没有？
	if c.seen[url] {
		return true // 来过，返回 true
	}
	
	// 第一次来，在记录本上写上你的名字，标记为 true
	c.seen[url] = true 
	
	// 返回 false，告诉外面“这是个新地方，没来过”
	return false 
}

// 3. 核心的爬虫函数 Crawl
// 多了一个参数 wg *sync.WaitGroup（等待组），作用是指挥官，控制主程序什么时候下班。
func Crawl(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup) {
	
	// 【极其重要】告诉等待组：“我（当前这个协程）的任务干完了，你可以把计数器减 1 了”。
	// 这里用 defer 也是为了保证不管函数怎么结束，都会汇报任务完成。
	defer wg.Done() 

	// 如果深度小于等于 0，说明爬到底了，不爬了，直接返回
	if depth <= 0 {
		return
	}

	// 查一下记录本，如果这个 URL 已经被别的协程爬过了
	if cache.isVisited(url) {
		return // 直接掉头走人，避免重复劳动
	}

	// 真正开始去抓取网页（调用传进来的 fetcher）
	body, urls, err := fetcher.Fetch(url)
	
	// 如果抓取失败（比如网址不存在），打印错误，结束当前任务
	if err != nil {
		fmt.Println(err)
		return
	}
	
	// 成功抓到了！打印出来看看
	fmt.Printf("found: %s %q\n", url, body)

	// 遍历从当前网页里发现的所有新链接（比如页面里的 a 标签）
	for _, u := range urls {
		
		// 【指挥官记录】告诉等待组：“等一下，我又要派出去一个小弟（启动新协程）了，
		// 你的计数器得加 1，不然你提前跑了小弟就没命了”。
		// 注意：Add(1) 必须在 go 关键字前面！
		wg.Add(1) 
		
		// 【并发核心】启动一个全新的协程，去抓这个新链接。
		// 注意深度要 depth-1（比如规定爬4层，下一层就是3层）。
		// 这里加上 go 关键字，意味着这句话执行完，不等它返回，直接继续下一次循环。
		go Crawl(u, depth-1, fetcher, wg)
	}
}

// ================= 以下均为原题代码，无需修改 =================

func main() {
	// 创建一个等待组（相当于一个带有计数器的哨兵）
	var wg sync.WaitGroup
	
	// 即将启动第一个大任务，先告诉哨兵：计数器加 1
	wg.Add(1) 
	
	// 【注意】这里用了 go 关键字！让 Crawl 去后台默默干活。
	// 如果不加 go，主函数会卡在这里等它爬完。加了 go，主函数会继续往下走。
	// 注意：传给 wg 的是它的内存地址 &wg，因为多个协程都要修改同一个计数器。
	go Crawl("https://golang.org/", 4, fetcher, &wg)
	
	// 【主程序等待】主函数就像老板，在这里一直死等。
	// 直到 wg 里面的计数器变成 0（所有爬虫任务都干完了），才解除阻塞，继续往下走。
	wg.Wait() 
	
	// 所有协程都结束了，打印一句话
	fmt.Println("爬虫任务全部结束！")
}

// fakeFetcher is Fetcher that returns canned results.
// （下面是模拟的假网络，不用管它，只是用来模拟数据）
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}