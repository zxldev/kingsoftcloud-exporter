package config

var AllSLBMetricConfigs = []KscMetricConfig{
	{
		Namespace:        "SLB",
		MetricName:       "slb.bps.in",
		MetricDesc:       "SLB入网流量",
		MetricType:       1,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "bps",
	}, {
		Namespace:        "SLB",
		MetricName:       "slb.bps.out",
		MetricDesc:       "SLB出网流量",
		MetricType:       1,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "bps",
	}, {
		Namespace:        "SLB",
		MetricName:       "slb.pps.in",
		MetricDesc:       "SLB每秒流入包数",
		MetricType:       1,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "个",
	}, {
		Namespace:        "SLB",
		MetricName:       "slb.pps.out",
		MetricDesc:       "SLB每秒流出包数",
		MetricType:       1,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "个",
	}, {
		Namespace:        "SLB",
		MetricName:       "slb.cps",
		MetricDesc:       "SLB每秒新建连接数",
		MetricType:       1,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "个",
	}, {
		Namespace:        "SLB",
		MetricName:       "slb.activeconn",
		MetricDesc:       "SLB当前活跃连接数",
		MetricType:       1,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "个",
	}, {
		Namespace:        "SLB",
		MetricName:       "slb.inactiveconn",
		MetricDesc:       "SLB当前不活跃连接数",
		MetricType:       1,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "个",
	}, {
		Namespace:        "SLB",
		MetricName:       "slb.concurrentconn",
		MetricDesc:       "SLB并发连接数",
		MetricType:       1,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "个",
	}, {
		Namespace:        "SLB",
		MetricName:       "slb.drop.bps.in",
		MetricDesc:       "丢弃入流量",
		MetricType:       1,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "bps",
	}, {
		Namespace:        "SLB",
		MetricName:       "slb.drop.bps.out",
		MetricDesc:       "丢弃出流量",
		MetricType:       1,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "bps",
	}, {
		Namespace:        "SLB",
		MetricName:       "slb.drop.bps.acl",
		MetricDesc:       "ACL丢弃流量",
		MetricType:       1,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "bps",
	}, {
		Namespace:        "SLB",
		MetricName:       "slb.drop.pps.in",
		MetricDesc:       "丢弃流入数据包",
		MetricType:       1,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "pps",
	}, {
		Namespace:        "SLB",
		MetricName:       "slb.drop.pps.out",
		MetricDesc:       "丢弃流出数据包",
		MetricType:       1,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "pps",
	}, {
		Namespace:        "SLB",
		MetricName:       "slb.drop.pps.acl",
		MetricDesc:       "ACL丢弃数据包",
		MetricType:       1,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "pps",
	}, {
		Namespace:        "SLB",
		MetricName:       "slb.req_rate",
		MetricDesc:       "7层协议QPS",
		MetricType:       1,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "qps",
	}, {
		Namespace:        "SLB",
		MetricName:       "slb.rs_disable",
		MetricDesc:       "健康检查未开启的rs数量",
		MetricType:       1,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "个",
	}, {
		Namespace:        "SLB",
		MetricName:       "slb.rs_up",
		MetricDesc:       "健康检查状态为健康的rs数量",
		MetricType:       1,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "个",
	}, {
		Namespace:        "SLB",
		MetricName:       "slb.rs_down",
		MetricDesc:       "健康检查状态为不健康的rs数量",
		MetricType:       1,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "个",
	}, {
		Namespace:        "SLB",
		MetricName:       "slb.rs_init",
		MetricDesc:       "健康检查未完成的rs数量",
		MetricType:       1,
		Labels:           []string{},
		Statistics:       []string{"avg"},
		MinPeriodSeconds: 60,
		Unit:             "个",
	},
}
