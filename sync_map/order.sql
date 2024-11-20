-- 2024-11-20 面试
-- 业务场景：
-- 假设你有一个电商平台的数据库，其中包含一个订单表（orders），
-- 该表记录了订单的详细信息，包括订单ID、用户ID、订单状态、订单总金额、订单创建时间等字段。
-- 你需要编写查询语句来分析过去30天内每天的订单数量和总销售额。
-- 要求：
-- 1. 使用MySQL编写一个SQL查询，返回过去30天每天的日期、订单数量和总销售额。
-- 2. 使用Elasticsearch编写一个DSL查询，实现相同的功能。

-- mysql
select date(created_at, 'Y-m-d') d, sum(amount), count(id) from orders where created_at < ?  group by d
order by created_at desc;

-- elasticsearch
{
  "query": {
    "bool": {
      "must": [
        { "range": { "@createdAt": { "gt": "x","lt": "y" } }  }
      ]
    }
  },
  "sort": {  "@createdAt": "desc"  },
  "aggs": {
    "groupDate": {
      "date_histogram": {
        "field": "create_date",
        "interval": "day",
        "format": "yyyy-MM-dd"
      }
    },
    "sum": {
        "field": "sum"
    },
    "count": {
        "field": "id"
    }
  }
}



