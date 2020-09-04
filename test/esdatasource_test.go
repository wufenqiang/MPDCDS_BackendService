package test

import (
	"MPDCDS_BackendService/src/conf"
	"MPDCDS_BackendService/src/datasource"
	"context"
	"fmt"
	"gitlab.weather.com.cn/wufenqiang/MPDCDSPro/src/models"

	//uuid "github.com/iris-contrib/go.uuid"
	"github.com/olivere/elastic/v7"
	"testing"
	"time"
)

//校验Index是否存在
func TestIndexExists(t *testing.T) {
	index := "Hf_platform_log"
	exists, err := datasource.ES.GetDB().IndexExists(index).Do(context.Background())
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("index %s exists %v", index, exists)
}

//创建Index
func TestIndexCreate(t *testing.T) {
	index := "orderinfo"
	result, err := datasource.ES.GetDB().CreateIndex(index).Do(context.Background())

	/*var mapping = `{
			"settings":{
				"number_of_shards": 15,
				"number_of_replicas": 1
			},
			"mappings":{
				"hf_platform_log":{
					"properties":{
						"username":{
							"type":"keyword"
						},
						"operate":{
							"type":"keyword"
						},
						"catalog":{
							"type":"keyword"
						},
						"createtime":{
							"type":"date"
						}
					}
				}
			}
		}`
	result, err := esclient.CreateIndex(index).BodyString(mapping).Do(context.Background())*/
	if err != nil {
		fmt.Printf("create index failed, err: %v\n", err)
	}

	fmt.Println("create index success", result.Acknowledged)
}

type Person struct {
	UserName   string    `json:"user_name"`
	Password   string    `json:"password"`
	Age        int       `json:"age"`
	CreateTime time.Time `json:"create_time"`
	ArrayTest  []string  `json:"array_test"`
	SonList    []Son     `json:"son_list"`
	Remark     string    `json:"remark"`
}
type Son struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestCreate(t *testing.T) {
	//使用结构体
	var e1 models.Api_user

	creatTime, _ := time.ParseInLocation(conf.Layout, "2006-01-02 15:04:05", conf.Loc)

	e1.CreateTime = creatTime.String()
	e1.Username = "test001"
	e1.Password = "test001"
	e1.Phone = "1234567890"
	e1.Remark = "test"
	e1.Status = "1"

	put1, err := datasource.ES.GetDB().Index().
		Index("megacorp").
		Id("1").
		BodyJson(e1).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put1.Id, put1.Index, put1.Type)

	//使用字符串
	e2 := `{"first_name":"John","last_name":"Smith","age":25,"about":"I love to go rock climbing","interests":["sports","music"]}`
	put2, err := datasource.ES.GetDB().Index().
		Index("megacorp").
		Id("2").
		BodyJson(e2).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put2.Id, put2.Index, put2.Type)

	e3 := `{"first_name":"Douglas","last_name":"Fir","age":35,"about":"I like to build cabinets","interests":["forestry"]}`
	put3, err := datasource.ES.GetDB().Index().
		Index("megacorp").
		Id("3").
		BodyJson(e3).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put3.Id, put3.Index, put3.Type)
}

func TestUpdate(t *testing.T) {
	res, err := datasource.ES.GetDB().Update().
		Index("megacorp").
		Id("2").
		Doc(map[string]interface{}{"age": 88}).
		Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	fmt.Printf("update age %s\n", res.Result)
}

func TestGetById(t *testing.T) {
	//通过id查找
	get1, err := datasource.ES.GetDB().Get().Index("megacorp").Id("2").Do(context.Background())
	if err != nil {
		panic(err)
	}
	if get1.Found {
		fmt.Printf("Got document %s in version %d from index %s, type %s, source %s\n", get1.Id, get1.Version, get1.Index, get1.Type, get1.Source)
	}
}

func TestQuery(t *testing.T) {
	var res *elastic.SearchResult
	var err error
	//取所有
	res, err = datasource.ES.GetDB().Search("megacorp").Do(context.Background())
	printEmployee(res, err)

	//字段相等
	q := elastic.NewQueryStringQuery("last_name:Smith")
	res, err = datasource.ES.GetDB().Search("megacorp").Query(q).Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	printEmployee(res, err)

	//条件查询
	//年龄大于30岁的
	boolQ := elastic.NewBoolQuery()
	boolQ.Must(elastic.NewMatchQuery("last_name", "smith"))
	boolQ.Filter(elastic.NewRangeQuery("age").Gt(30))
	res, err = datasource.ES.GetDB().Search("megacorp").Type("employee").Query(q).Do(context.Background())
	printEmployee(res, err)

	//短语搜索 搜索about字段中有 rock climbing
	matchPhraseQuery := elastic.NewMatchPhraseQuery("about", "rock climbing")
	res, err = datasource.ES.GetDB().Search("megacorp").Type("employee").Query(matchPhraseQuery).Do(context.Background())
	printEmployee(res, err)

}

func TestListWithPage(t *testing.T) {
	listEs(1, 1)
}

func TestDelete(t *testing.T) {
	res, err := datasource.ES.GetDB().Delete().Index("megacorp").
		Id("2").
		Do(context.Background())
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("delete result %s\n", res.Result)
}

//简单分页
func listEs(size, page int) {
	if size < 0 || page < 1 {
		fmt.Printf("param error")
		return
	}
	res, err := datasource.ES.GetDB().Search("megacorp").
		Size(size).
		From((page - 1) * size).
		Do(context.Background())
	printEmployee(res, err)

}

//打印查询到的Employee
func printEmployee(res *elastic.SearchResult, err error) {
	//if err != nil {
	//	print(err.Error())
	//	return
	//}
	//var typ Employee
	//for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
	//	t := item.(Employee)
	//	fmt.Printf("%#v\n", t)
	//}
}
