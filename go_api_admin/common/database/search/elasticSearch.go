package search

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"project/common/global"
)

var ctx = context.Background()
var esUrl string = "http://106.13.73.165:9200"
func Init() error{
	//连接客户端
	client, err := elastic.NewClient(elastic.SetURL(esUrl), elastic.SetSniff(false))
	if err != nil {
		// Handle error
		return err
	}
	info, code, err := client.Ping(esUrl).Do(ctx)
	if err != nil {
		// Handle error
		return err
	}
	fmt.Printf("Elasticsearch returned with code>: %d and version %s\n", code, info.Version.Number)
	// 获取版本号的直接API
	esVersion, err := client.ElasticsearchVersion(esUrl)
	if err != nil {
		return err
	}
	fmt.Printf("es的版本为%s\n", esVersion)
	global.ElasticSearch = client
	return nil
}

//func Close() {
//	_ = global.ElasticSearch.Close()
//}