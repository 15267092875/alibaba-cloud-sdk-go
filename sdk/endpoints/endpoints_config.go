package endpoints

import (
	"encoding/json"
	"fmt"
	"sync"
)

const endpointsJson = `{
	"products": [
		{
			"code": "emr",
			"document_id": "28140",
			"location_service_code": "emr",
			"regional_endpoints": [
				{
					"region": "cn-qingdao",
					"endpoint": "emr.cn-qingdao.chenxuewen.me"
				},
				{
					"region": "cn-beijing",
					"endpoint": "emr.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "emr.chenxuewen.me"
				},
				{
					"region": "eu-west-1",
					"endpoint": "emr.eu-west-1.chenxuewen.me"
				},
				{
					"region": "us-west-1",
					"endpoint": "emr.chenxuewen.me"
				},
				{
					"region": "me-east-1",
					"endpoint": "emr.me-east-1.chenxuewen.me"
				},
				{
					"region": "ap-northeast-1",
					"endpoint": "emr.ap-northeast-1.chenxuewen.me"
				},
				{
					"region": "eu-central-1",
					"endpoint": "emr.eu-central-1.chenxuewen.me"
				},
				{
					"region": "cn-huhehaote",
					"endpoint": "emr.cn-huhehaote.chenxuewen.me"
				},
				{
					"region": "ap-south-1",
					"endpoint": "emr.ap-south-1.chenxuewen.me"
				},
				{
					"region": "us-east-1",
					"endpoint": "emr.us-east-1.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "emr.chenxuewen.me"
				},
				{
					"region": "cn-hongkong",
					"endpoint": "emr.cn-hongkong.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "emr.chenxuewen.me"
				},
				{
					"region": "ap-southeast-2",
					"endpoint": "emr.ap-southeast-2.chenxuewen.me"
				},
				{
					"region": "ap-southeast-3",
					"endpoint": "emr.ap-southeast-3.chenxuewen.me"
				},
				{
					"region": "cn-zhangjiakou",
					"endpoint": "emr.cn-zhangjiakou.chenxuewen.me"
				},
				{
					"region": "ap-southeast-5",
					"endpoint": "emr.ap-southeast-5.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "emr.chenxuewen.me"
				}
			],
			"global_endpoint": "emr.chenxuewen.me",
			"regional_endpoint_pattern": "emr.[RegionId].chenxuewen.me"
		},
		{
			"code": "petadata",
			"document_id": "",
			"location_service_code": "petadata",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "petadata.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "petadata.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "petadata.chenxuewen.me"
				},
				{
					"region": "us-east-1",
					"endpoint": "petadata.chenxuewen.me"
				},
				{
					"region": "me-east-1",
					"endpoint": "petadata.me-east-1.chenxuewen.me"
				},
				{
					"region": "ap-southeast-2",
					"endpoint": "petadata.ap-southeast-2.chenxuewen.me"
				},
				{
					"region": "eu-central-1",
					"endpoint": "petadata.eu-central-1.chenxuewen.me"
				},
				{
					"region": "cn-beijing",
					"endpoint": "petadata.chenxuewen.me"
				},
				{
					"region": "cn-hongkong",
					"endpoint": "petadata.chenxuewen.me"
				},
				{
					"region": "ap-southeast-5",
					"endpoint": "petadata.ap-southeast-5.chenxuewen.me"
				},
				{
					"region": "us-west-1",
					"endpoint": "petadata.chenxuewen.me"
				},
				{
					"region": "cn-huhehaote",
					"endpoint": "petadata.cn-huhehaote.chenxuewen.me"
				},
				{
					"region": "cn-zhangjiakou",
					"endpoint": "petadata.cn-zhangjiakou.chenxuewen.me"
				},
				{
					"region": "cn-qingdao",
					"endpoint": "petadata.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "petadata.chenxuewen.me"
				}
			],
			"global_endpoint": "petadata.chenxuewen.me",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "dbs",
			"document_id": "",
			"location_service_code": "dbs",
			"regional_endpoints": [
				{
					"region": "cn-shenzhen",
					"endpoint": "dbs-api.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "cn-hongkong",
					"endpoint": "dbs-api.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "dbs-api.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "ap-northeast-1",
					"endpoint": "dbs-api.ap-northeast-1.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "dbs-api.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "dbs-api.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "cn-qingdao",
					"endpoint": "dbs-api.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "cn-beijing",
					"endpoint": "dbs-api.cn-hangzhou.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "alidnsgtm",
			"document_id": "",
			"location_service_code": "alidnsgtm",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "alidns.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "alidns.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "elasticsearch",
			"document_id": "",
			"location_service_code": "elasticsearch",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "elasticsearch.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "elasticsearch.cn-shenzhen.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "elasticsearch.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "cn-beijing",
					"endpoint": "elasticsearch.cn-beijing.chenxuewen.me"
				},
				{
					"region": "cn-hongkong",
					"endpoint": "elasticsearch.cn-hongkong.chenxuewen.me"
				},
				{
					"region": "ap-southeast-3",
					"endpoint": "elasticsearch.ap-southeast-3.chenxuewen.me"
				},
				{
					"region": "us-west-1",
					"endpoint": "elasticsearch.us-west-1.chenxuewen.me"
				},
				{
					"region": "ap-southeast-2",
					"endpoint": "elasticsearch.ap-southeast-2.chenxuewen.me"
				},
				{
					"region": "ap-southeast-5",
					"endpoint": "elasticsearch.ap-southeast-5.chenxuewen.me"
				},
				{
					"region": "eu-central-1",
					"endpoint": "elasticsearch.eu-central-1.chenxuewen.me"
				},
				{
					"region": "ap-south-1",
					"endpoint": "elasticsearch.ap-south-1.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "elasticsearch.cn-shanghai.chenxuewen.me"
				},
				{
					"region": "cn-qingdao",
					"endpoint": "elasticsearch.cn-qingdao.chenxuewen.me"
				},
				{
					"region": "cn-zhangjiakou",
					"endpoint": "elasticsearch.cn-zhangjiakou.chenxuewen.me"
				},
				{
					"region": "ap-northeast-1",
					"endpoint": "elasticsearch.ap-northeast-1.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "baas",
			"document_id": "",
			"location_service_code": "baas",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "baas.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "ap-northeast-1",
					"endpoint": "baas.ap-northeast-1.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "baas.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "cn-beijing",
					"endpoint": "baas.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "baas.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "baas.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "cr",
			"document_id": "60716",
			"location_service_code": "cr",
			"regional_endpoints": null,
			"global_endpoint": "cr.chenxuewen.me",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "cloudap",
			"document_id": "",
			"location_service_code": "cloudap",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "cloudwf.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "imagesearch",
			"document_id": "",
			"location_service_code": "imagesearch",
			"regional_endpoints": [
				{
					"region": "ap-southeast-2",
					"endpoint": "imagesearch.ap-southeast-2.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "imagesearch.cn-shanghai.chenxuewen.me"
				},
				{
					"region": "ap-northeast-1",
					"endpoint": "imagesearch.ap-northeast-1.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "imagesearch.ap-southeast-1.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "pts",
			"document_id": "",
			"location_service_code": "pts",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "pts.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "ehs",
			"document_id": "",
			"location_service_code": "ehs",
			"regional_endpoints": [
				{
					"region": "cn-beijing",
					"endpoint": "ehpc.cn-beijing.chenxuewen.me"
				},
				{
					"region": "cn-hongkong",
					"endpoint": "ehpc.cn-hongkong.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "ehpc.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "cn-qingdao",
					"endpoint": "ehpc.cn-qingdao.chenxuewen.me"
				},
				{
					"region": "eu-central-1",
					"endpoint": "ehpc.eu-central-1.chenxuewen.me"
				},
				{
					"region": "cn-zhangjiakou",
					"endpoint": "ehpc.cn-zhangjiakou.chenxuewen.me"
				},
				{
					"region": "cn-huhehaote",
					"endpoint": "ehpc.cn-huhehaote.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "ehpc.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "ehpc.cn-shanghai.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "ehpc.cn-shenzhen.chenxuewen.me"
				},
				{
					"region": "ap-southeast-2",
					"endpoint": "ehpc.ap-southeast-2.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "polardb",
			"document_id": "58764",
			"location_service_code": "polardb",
			"regional_endpoints": [
				{
					"region": "ap-south-1",
					"endpoint": "polardb.ap-south-1.chenxuewen.me"
				},
				{
					"region": "cn-hongkong",
					"endpoint": "polardb.chenxuewen.me"
				},
				{
					"region": "cn-huhehaote",
					"endpoint": "polardb.cn-huhehaote.chenxuewen.me"
				},
				{
					"region": "cn-qingdao",
					"endpoint": "polardb.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "polardb.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "polardb.chenxuewen.me"
				},
				{
					"region": "cn-beijing",
					"endpoint": "polardb.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "polardb.chenxuewen.me"
				},
				{
					"region": "ap-southeast-5",
					"endpoint": "polardb.ap-southeast-5.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": "polardb.chenxuewen.me"
		},
		{
			"code": "r-kvstore",
			"document_id": "60831",
			"location_service_code": "redisa",
			"regional_endpoints": [
				{
					"region": "cn-shenzhen",
					"endpoint": "r-kvstore.chenxuewen.me"
				},
				{
					"region": "cn-huhehaote",
					"endpoint": "r-kvstore.cn-huhehaote.chenxuewen.me"
				},
				{
					"region": "ap-southeast-3",
					"endpoint": "r-kvstore.ap-southeast-3.chenxuewen.me"
				},
				{
					"region": "ap-southeast-2",
					"endpoint": "r-kvstore.ap-southeast-2.chenxuewen.me"
				},
				{
					"region": "cn-zhangjiakou",
					"endpoint": "r-kvstore.cn-zhangjiakou.chenxuewen.me"
				},
				{
					"region": "eu-central-1",
					"endpoint": "r-kvstore.eu-central-1.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "r-kvstore.chenxuewen.me"
				},
				{
					"region": "us-west-1",
					"endpoint": "r-kvstore.chenxuewen.me"
				},
				{
					"region": "cn-hongkong",
					"endpoint": "r-kvstore.cn-hongkong.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "r-kvstore.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "ap-south-1",
					"endpoint": "r-kvstore.ap-south-1.chenxuewen.me"
				},
				{
					"region": "cn-qingdao",
					"endpoint": "r-kvstore.chenxuewen.me"
				},
				{
					"region": "cn-beijing",
					"endpoint": "r-kvstore.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "r-kvstore.chenxuewen.me"
				},
				{
					"region": "eu-west-1",
					"endpoint": "r-kvstore.eu-west-1.chenxuewen.me"
				},
				{
					"region": "us-east-1",
					"endpoint": "r-kvstore.chenxuewen.me"
				},
				{
					"region": "me-east-1",
					"endpoint": "r-kvstore.me-east-1.chenxuewen.me"
				},
				{
					"region": "ap-northeast-1",
					"endpoint": "r-kvstore.ap-northeast-1.chenxuewen.me"
				},
				{
					"region": "ap-southeast-5",
					"endpoint": "r-kvstore.ap-southeast-5.chenxuewen.me"
				}
			],
			"global_endpoint": "r-kvstore.chenxuewen.me",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "xianzhi",
			"document_id": "",
			"location_service_code": "xianzhi",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "xianzhi.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "pcdn",
			"document_id": "",
			"location_service_code": "pcdn",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "pcdn.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "cdn",
			"document_id": "27148",
			"location_service_code": "cdn",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "cdn.chenxuewen.me"
				}
			],
			"global_endpoint": "cdn.chenxuewen.me",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "cloudauth",
			"document_id": "60687",
			"location_service_code": "cloudauth",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "cloudauth.chenxuewen.me"
				}
			],
			"global_endpoint": "cloudauth.chenxuewen.me",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "nas",
			"document_id": "62598",
			"location_service_code": "nas",
			"regional_endpoints": [
				{
					"region": "ap-southeast-2",
					"endpoint": "nas.ap-southeast-2.chenxuewen.me"
				},
				{
					"region": "ap-south-1",
					"endpoint": "nas.ap-south-1.chenxuewen.me"
				},
				{
					"region": "eu-central-1",
					"endpoint": "nas.eu-central-1.chenxuewen.me"
				},
				{
					"region": "us-west-1",
					"endpoint": "nas.us-west-1.chenxuewen.me"
				},
				{
					"region": "cn-huhehaote",
					"endpoint": "nas.cn-huhehaote.chenxuewen.me"
				},
				{
					"region": "cn-qingdao",
					"endpoint": "nas.cn-qingdao.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "nas.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "cn-beijing",
					"endpoint": "nas.cn-beijing.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "nas.cn-shanghai.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "nas.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "cn-zhangjiakou",
					"endpoint": "nas.cn-zhangjiakou.chenxuewen.me"
				},
				{
					"region": "ap-northeast-1",
					"endpoint": "nas.ap-northeast-1.chenxuewen.me"
				},
				{
					"region": "us-east-1",
					"endpoint": "nas.us-east-1.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "nas.cn-shenzhen.chenxuewen.me"
				},
				{
					"region": "ap-southeast-5",
					"endpoint": "nas.ap-southeast-5.chenxuewen.me"
				},
				{
					"region": "cn-hongkong",
					"endpoint": "nas.cn-hongkong.chenxuewen.me"
				},
				{
					"region": "ap-southeast-3",
					"endpoint": "nas.ap-southeast-3.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "alidns",
			"document_id": "29739",
			"location_service_code": "alidns",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "alidns.chenxuewen.me"
				}
			],
			"global_endpoint": "alidns.chenxuewen.me",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "dts",
			"document_id": "",
			"location_service_code": "dts",
			"regional_endpoints": [
				{
					"region": "cn-beijing",
					"endpoint": "dts.chenxuewen.me"
				},
				{
					"region": "cn-zhangjiakou",
					"endpoint": "dts.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "dts.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "dts.chenxuewen.me"
				},
				{
					"region": "cn-hongkong",
					"endpoint": "dts.chenxuewen.me"
				},
				{
					"region": "cn-qingdao",
					"endpoint": "dts.chenxuewen.me"
				},
				{
					"region": "cn-huhehaote",
					"endpoint": "dts.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "dts.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "dts.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "emas",
			"document_id": "",
			"location_service_code": "emas",
			"regional_endpoints": [
				{
					"region": "cn-shanghai",
					"endpoint": "mhub.cn-shanghai.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "mhub.cn-hangzhou.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "dysmsapi",
			"document_id": "",
			"location_service_code": "dysmsapi",
			"regional_endpoints": [
				{
					"region": "ap-southeast-3",
					"endpoint": "dysmsapi.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "cn-qingdao",
					"endpoint": "dysmsapi.chenxuewen.me"
				},
				{
					"region": "cn-chengdu",
					"endpoint": "dysmsapi.chenxuewen.me"
				},
				{
					"region": "me-east-1",
					"endpoint": "dysmsapi.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "ap-southeast-5",
					"endpoint": "dysmsapi.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "ap-southeast-2",
					"endpoint": "dysmsapi.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "cn-zhangjiakou",
					"endpoint": "dysmsapi.chenxuewen.me"
				},
				{
					"region": "cn-hongkong",
					"endpoint": "dysmsapi.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "us-east-1",
					"endpoint": "dysmsapi.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "cn-beijing",
					"endpoint": "dysmsapi.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "dysmsapi.chenxuewen.me"
				},
				{
					"region": "ap-northeast-1",
					"endpoint": "dysmsapi.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "ap-south-1",
					"endpoint": "dysmsapi.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "dysmsapi.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "us-west-1",
					"endpoint": "dysmsapi.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "eu-central-1",
					"endpoint": "dysmsapi.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "dysmsapi.chenxuewen.me"
				},
				{
					"region": "cn-huhehaote",
					"endpoint": "dysmsapi.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "dysmsapi.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "cloudwf",
			"document_id": "58111",
			"location_service_code": "cloudwf",
			"regional_endpoints": null,
			"global_endpoint": "cloudwf.chenxuewen.me",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "fc",
			"document_id": "",
			"location_service_code": "fc",
			"regional_endpoints": [
				{
					"region": "cn-beijing",
					"endpoint": "cn-beijing.fc.chenxuewen.me"
				},
				{
					"region": "ap-southeast-2",
					"endpoint": "ap-southeast-2.fc.chenxuewen.me"
				},
				{
					"region": "cn-huhehaote",
					"endpoint": "cn-huhehaote.fc.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "cn-shanghai.fc.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "cn-hangzhou.fc.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "cn-shenzhen.fc.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "saf",
			"document_id": "",
			"location_service_code": "saf",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "saf.cn-shanghai.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "saf.cn-shanghai.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "saf.cn-shenzhen.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "rds",
			"document_id": "26223",
			"location_service_code": "rds",
			"regional_endpoints": [
				{
					"region": "ap-northeast-1",
					"endpoint": "rds.ap-northeast-1.chenxuewen.me"
				},
				{
					"region": "ap-south-1",
					"endpoint": "rds.ap-south-1.chenxuewen.me"
				},
				{
					"region": "cn-zhangjiakou",
					"endpoint": "rds.cn-zhangjiakou.chenxuewen.me"
				},
				{
					"region": "cn-beijing",
					"endpoint": "rds.chenxuewen.me"
				},
				{
					"region": "cn-hongkong",
					"endpoint": "rds.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "rds.chenxuewen.me"
				},
				{
					"region": "eu-west-1",
					"endpoint": "rds.eu-west-1.chenxuewen.me"
				},
				{
					"region": "us-east-1",
					"endpoint": "rds.chenxuewen.me"
				},
				{
					"region": "ap-southeast-3",
					"endpoint": "rds.ap-southeast-3.chenxuewen.me"
				},
				{
					"region": "ap-southeast-2",
					"endpoint": "rds.ap-southeast-2.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "rds.chenxuewen.me"
				},
				{
					"region": "cn-huhehaote",
					"endpoint": "rds.cn-huhehaote.chenxuewen.me"
				},
				{
					"region": "ap-southeast-5",
					"endpoint": "rds.ap-southeast-5.chenxuewen.me"
				},
				{
					"region": "eu-central-1",
					"endpoint": "rds.eu-central-1.chenxuewen.me"
				},
				{
					"region": "cn-qingdao",
					"endpoint": "rds.chenxuewen.me"
				},
				{
					"region": "me-east-1",
					"endpoint": "rds.me-east-1.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "rds.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "rds.chenxuewen.me"
				},
				{
					"region": "us-west-1",
					"endpoint": "rds.chenxuewen.me"
				}
			],
			"global_endpoint": "rds.chenxuewen.me",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "vpc",
			"document_id": "34962",
			"location_service_code": "vpc",
			"regional_endpoints": [
				{
					"region": "ap-south-1",
					"endpoint": "vpc.ap-south-1.chenxuewen.me"
				},
				{
					"region": "cn-beijing",
					"endpoint": "vpc.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "vpc.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "vpc.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "vpc.chenxuewen.me"
				},
				{
					"region": "ap-southeast-2",
					"endpoint": "vpc.ap-southeast-2.chenxuewen.me"
				},
				{
					"region": "ap-northeast-1",
					"endpoint": "vpc.ap-northeast-1.chenxuewen.me"
				},
				{
					"region": "cn-qingdao",
					"endpoint": "vpc.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "vpc.chenxuewen.me"
				},
				{
					"region": "cn-hongkong",
					"endpoint": "vpc.chenxuewen.me"
				},
				{
					"region": "cn-huhehaote",
					"endpoint": "vpc.cn-huhehaote.chenxuewen.me"
				},
				{
					"region": "me-east-1",
					"endpoint": "vpc.me-east-1.chenxuewen.me"
				},
				{
					"region": "ap-southeast-5",
					"endpoint": "vpc.ap-southeast-5.chenxuewen.me"
				},
				{
					"region": "ap-southeast-3",
					"endpoint": "vpc.ap-southeast-3.chenxuewen.me"
				},
				{
					"region": "eu-central-1",
					"endpoint": "vpc.eu-central-1.chenxuewen.me"
				},
				{
					"region": "cn-zhangjiakou",
					"endpoint": "vpc.cn-zhangjiakou.chenxuewen.me"
				},
				{
					"region": "eu-west-1",
					"endpoint": "vpc.eu-west-1.chenxuewen.me"
				},
				{
					"region": "us-west-1",
					"endpoint": "vpc.chenxuewen.me"
				},
				{
					"region": "us-east-1",
					"endpoint": "vpc.chenxuewen.me"
				}
			],
			"global_endpoint": "vpc.chenxuewen.me",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "gpdb",
			"document_id": "",
			"location_service_code": "gpdb",
			"regional_endpoints": [
				{
					"region": "ap-southeast-3",
					"endpoint": "gpdb.ap-southeast-3.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "gpdb.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "gpdb.chenxuewen.me"
				},
				{
					"region": "us-west-1",
					"endpoint": "gpdb.chenxuewen.me"
				},
				{
					"region": "cn-huhehaote",
					"endpoint": "gpdb.cn-huhehaote.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "gpdb.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "gpdb.chenxuewen.me"
				},
				{
					"region": "us-east-1",
					"endpoint": "gpdb.chenxuewen.me"
				},
				{
					"region": "ap-southeast-5",
					"endpoint": "gpdb.ap-southeast-5.chenxuewen.me"
				},
				{
					"region": "ap-southeast-2",
					"endpoint": "gpdb.ap-southeast-2.chenxuewen.me"
				},
				{
					"region": "eu-central-1",
					"endpoint": "gpdb.eu-central-1.chenxuewen.me"
				},
				{
					"region": "cn-zhangjiakou",
					"endpoint": "gpdb.cn-zhangjiakou.chenxuewen.me"
				},
				{
					"region": "cn-beijing",
					"endpoint": "gpdb.chenxuewen.me"
				},
				{
					"region": "ap-northeast-1",
					"endpoint": "gpdb.ap-northeast-1.chenxuewen.me"
				},
				{
					"region": "eu-west-1",
					"endpoint": "gpdb.eu-west-1.chenxuewen.me"
				},
				{
					"region": "ap-south-1",
					"endpoint": "gpdb.ap-south-1.chenxuewen.me"
				}
			],
			"global_endpoint": "gpdb.chenxuewen.me",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "yunmarket",
			"document_id": "",
			"location_service_code": "yunmarket",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "market.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "pvtz",
			"document_id": "",
			"location_service_code": "pvtz",
			"regional_endpoints": [
				{
					"region": "ap-southeast-1",
					"endpoint": "pvtz.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "pvtz.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "oss",
			"document_id": "",
			"location_service_code": "oss",
			"regional_endpoints": [
				{
					"region": "cn-beijing",
					"endpoint": "oss-cn-beijing.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "oss-cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "oss-cn-shanghai.chenxuewen.me"
				},
				{
					"region": "cn-hongkong",
					"endpoint": "oss-cn-hongkong.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "oss-cn-shenzhen.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "oss-ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "us-west-1",
					"endpoint": "oss-us-west-1.chenxuewen.me"
				},
				{
					"region": "cn-qingdao",
					"endpoint": "oss-cn-qingdao.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "foas",
			"document_id": "",
			"location_service_code": "foas",
			"regional_endpoints": [
				{
					"region": "cn-qingdao",
					"endpoint": "foas.cn-qingdao.chenxuewen.me"
				},
				{
					"region": "cn-beijing",
					"endpoint": "foas.cn-beijing.chenxuewen.me"
				},
				{
					"region": "cn-zhangjiakou",
					"endpoint": "foas.cn-zhangjiakou.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "foas.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "foas.cn-shanghai.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "foas.cn-shenzhen.chenxuewen.me"
				},
				{
					"region": "ap-northeast-1",
					"endpoint": "foas.ap-northeast-1.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "ddos",
			"document_id": "",
			"location_service_code": "ddos",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "ddospro.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "cn-hongkong",
					"endpoint": "ddospro.cn-hongkong.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "cbn",
			"document_id": "",
			"location_service_code": "cbn",
			"regional_endpoints": [
				{
					"region": "ap-southeast-1",
					"endpoint": "cbn.chenxuewen.me"
				},
				{
					"region": "ap-northeast-1",
					"endpoint": "cbn.chenxuewen.me"
				},
				{
					"region": "cn-qingdao",
					"endpoint": "cbn.chenxuewen.me"
				},
				{
					"region": "cn-hongkong",
					"endpoint": "cbn.chenxuewen.me"
				},
				{
					"region": "ap-southeast-2",
					"endpoint": "cbn.chenxuewen.me"
				},
				{
					"region": "us-west-1",
					"endpoint": "cbn.chenxuewen.me"
				},
				{
					"region": "cn-beijing",
					"endpoint": "cbn.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "cbn.chenxuewen.me"
				},
				{
					"region": "ap-southeast-3",
					"endpoint": "cbn.chenxuewen.me"
				},
				{
					"region": "ap-southeast-5",
					"endpoint": "cbn.chenxuewen.me"
				},
				{
					"region": "eu-west-1",
					"endpoint": "cbn.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "cbn.chenxuewen.me"
				},
				{
					"region": "cn-huhehaote",
					"endpoint": "cbn.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "cbn.chenxuewen.me"
				},
				{
					"region": "us-east-1",
					"endpoint": "cbn.chenxuewen.me"
				},
				{
					"region": "eu-central-1",
					"endpoint": "cbn.chenxuewen.me"
				},
				{
					"region": "me-east-1",
					"endpoint": "cbn.chenxuewen.me"
				},
				{
					"region": "ap-south-1",
					"endpoint": "cbn.chenxuewen.me"
				},
				{
					"region": "cn-zhangjiakou",
					"endpoint": "cbn.chenxuewen.me"
				}
			],
			"global_endpoint": "cbn.chenxuewen.me",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "nlp",
			"document_id": "",
			"location_service_code": "nlp",
			"regional_endpoints": [
				{
					"region": "cn-shanghai",
					"endpoint": "nlp.cn-shanghai.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "hsm",
			"document_id": "",
			"location_service_code": "hsm",
			"regional_endpoints": [
				{
					"region": "cn-beijing",
					"endpoint": "hsm.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "hsm.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "hsm.chenxuewen.me"
				},
				{
					"region": "cn-hongkong",
					"endpoint": "hsm.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "hsm.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "hsm.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "ons",
			"document_id": "44416",
			"location_service_code": "ons",
			"regional_endpoints": [
				{
					"region": "ap-southeast-1",
					"endpoint": "ons.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "cn-huhehaote",
					"endpoint": "ons.cn-huhehaote.chenxuewen.me"
				},
				{
					"region": "us-east-1",
					"endpoint": "ons.us-east-1.chenxuewen.me"
				},
				{
					"region": "cn-hongkong",
					"endpoint": "ons.cn-hongkong.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "ons.cn-shenzhen.chenxuewen.me"
				},
				{
					"region": "ap-southeast-3",
					"endpoint": "ons.ap-southeast-3.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "ons.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "cn-qingdao",
					"endpoint": "ons.cn-qingdao.chenxuewen.me"
				},
				{
					"region": "cn-zhangjiakou",
					"endpoint": "ons.cn-zhangjiakou.chenxuewen.me"
				},
				{
					"region": "me-east-1",
					"endpoint": "ons.me-east-1.chenxuewen.me"
				},
				{
					"region": "ap-northeast-1",
					"endpoint": "ons.ap-northeast-1.chenxuewen.me"
				},
				{
					"region": "ap-southeast-2",
					"endpoint": "ons.ap-southeast-2.chenxuewen.me"
				},
				{
					"region": "ap-south-1",
					"endpoint": "ons.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "eu-central-1",
					"endpoint": "ons.eu-central-1.chenxuewen.me"
				},
				{
					"region": "eu-west-1",
					"endpoint": "ons.eu-west-1.chenxuewen.me"
				},
				{
					"region": "us-west-1",
					"endpoint": "ons.us-west-1.chenxuewen.me"
				},
				{
					"region": "cn-beijing",
					"endpoint": "ons.cn-beijing.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "ons.cn-shanghai.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "kms",
			"document_id": "",
			"location_service_code": "kms",
			"regional_endpoints": [
				{
					"region": "cn-hongkong",
					"endpoint": "kms.cn-hongkong.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "kms.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "ap-south-1",
					"endpoint": "kms.ap-south-1.chenxuewen.me"
				},
				{
					"region": "cn-qingdao",
					"endpoint": "kms.cn-qingdao.chenxuewen.me"
				},
				{
					"region": "eu-west-1",
					"endpoint": "kms.eu-west-1.chenxuewen.me"
				},
				{
					"region": "us-east-1",
					"endpoint": "kms.us-east-1.chenxuewen.me"
				},
				{
					"region": "cn-beijing",
					"endpoint": "kms.cn-beijing.chenxuewen.me"
				},
				{
					"region": "cn-zhangjiakou",
					"endpoint": "kms.cn-zhangjiakou.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "kms.cn-shanghai.chenxuewen.me"
				},
				{
					"region": "ap-southeast-5",
					"endpoint": "kms.ap-southeast-5.chenxuewen.me"
				},
				{
					"region": "cn-huhehaote",
					"endpoint": "kms.cn-huhehaote.chenxuewen.me"
				},
				{
					"region": "me-east-1",
					"endpoint": "kms.me-east-1.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "kms.cn-shenzhen.chenxuewen.me"
				},
				{
					"region": "ap-southeast-3",
					"endpoint": "kms.ap-southeast-3.chenxuewen.me"
				},
				{
					"region": "us-west-1",
					"endpoint": "kms.us-west-1.chenxuewen.me"
				},
				{
					"region": "ap-southeast-2",
					"endpoint": "kms.ap-southeast-2.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "kms.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "eu-central-1",
					"endpoint": "kms.eu-central-1.chenxuewen.me"
				},
				{
					"region": "ap-northeast-1",
					"endpoint": "kms.ap-northeast-1.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "cps",
			"document_id": "",
			"location_service_code": "cps",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "cloudpush.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "ensdisk",
			"document_id": "",
			"location_service_code": "ensdisk",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "ens.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "cloudapi",
			"document_id": "43590",
			"location_service_code": "apigateway",
			"regional_endpoints": [
				{
					"region": "ap-southeast-2",
					"endpoint": "apigateway.ap-southeast-2.chenxuewen.me"
				},
				{
					"region": "ap-south-1",
					"endpoint": "apigateway.ap-south-1.chenxuewen.me"
				},
				{
					"region": "us-east-1",
					"endpoint": "apigateway.us-east-1.chenxuewen.me"
				},
				{
					"region": "me-east-1",
					"endpoint": "apigateway.me-east-1.chenxuewen.me"
				},
				{
					"region": "cn-qingdao",
					"endpoint": "apigateway.cn-qingdao.chenxuewen.me"
				},
				{
					"region": "cn-beijing",
					"endpoint": "apigateway.cn-beijing.chenxuewen.me"
				},
				{
					"region": "ap-southeast-5",
					"endpoint": "apigateway.ap-southeast-5.chenxuewen.me"
				},
				{
					"region": "ap-southeast-3",
					"endpoint": "apigateway.ap-southeast-3.chenxuewen.me"
				},
				{
					"region": "cn-zhangjiakou",
					"endpoint": "apigateway.cn-zhangjiakou.chenxuewen.me"
				},
				{
					"region": "cn-huhehaote",
					"endpoint": "apigateway.cn-huhehaote.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "apigateway.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "eu-central-1",
					"endpoint": "apigateway.eu-central-1.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "apigateway.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "us-west-1",
					"endpoint": "apigateway.us-west-1.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "apigateway.cn-shenzhen.chenxuewen.me"
				},
				{
					"region": "eu-west-1",
					"endpoint": "apigateway.eu-west-1.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "apigateway.cn-shanghai.chenxuewen.me"
				},
				{
					"region": "ap-northeast-1",
					"endpoint": "apigateway.ap-northeast-1.chenxuewen.me"
				},
				{
					"region": "cn-hongkong",
					"endpoint": "apigateway.cn-hongkong.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": "apigateway.[RegionId].chenxuewen.me"
		},
		{
			"code": "eci",
			"document_id": "",
			"location_service_code": "eci",
			"regional_endpoints": [
				{
					"region": "cn-shanghai",
					"endpoint": "eci.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "eci.chenxuewen.me"
				},
				{
					"region": "us-west-1",
					"endpoint": "eci.chenxuewen.me"
				},
				{
					"region": "cn-beijing",
					"endpoint": "eci.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "eci.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "eci.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "onsvip",
			"document_id": "",
			"location_service_code": "onsvip",
			"regional_endpoints": [
				{
					"region": "cn-beijing",
					"endpoint": "ons.cn-beijing.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "ons.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "ons.cn-shanghai.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "ons.cn-shenzhen.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "ons.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "cn-qingdao",
					"endpoint": "ons.cn-qingdao.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "linkwan",
			"document_id": "",
			"location_service_code": "linkwan",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "linkwan.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "linkwan.cn-shanghai.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "ddosdip",
			"document_id": "",
			"location_service_code": "ddosdip",
			"regional_endpoints": [
				{
					"region": "ap-southeast-1",
					"endpoint": "ddosdip.ap-southeast-1.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "batchcompute",
			"document_id": "44717",
			"location_service_code": "batchcompute",
			"regional_endpoints": [
				{
					"region": "us-west-1",
					"endpoint": "batchcompute.us-west-1.chenxuewen.me"
				},
				{
					"region": "cn-qingdao",
					"endpoint": "batchcompute.cn-qingdao.chenxuewen.me"
				},
				{
					"region": "cn-zhangjiakou",
					"endpoint": "batchcompute.cn-zhangjiakou.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "batchcompute.cn-shenzhen.chenxuewen.me"
				},
				{
					"region": "cn-huhehaote",
					"endpoint": "batchcompute.cn-huhehaote.chenxuewen.me"
				},
				{
					"region": "cn-beijing",
					"endpoint": "batchcompute.cn-beijing.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "batchcompute.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "batchcompute.cn-shanghai.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "batchcompute.ap-southeast-1.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": "batchcompute.[RegionId].chenxuewen.me"
		},
		{
			"code": "aegis",
			"document_id": "28449",
			"location_service_code": "vipaegis",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "aegis.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "ap-southeast-3",
					"endpoint": "aegis.ap-southeast-3.chenxuewen.me"
				}
			],
			"global_endpoint": "aegis.cn-hangzhou.chenxuewen.me",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "arms",
			"document_id": "42924",
			"location_service_code": "arms",
			"regional_endpoints": [
				{
					"region": "cn-hongkong",
					"endpoint": "arms.cn-hongkong.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "arms.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "arms.cn-shenzhen.chenxuewen.me"
				},
				{
					"region": "cn-qingdao",
					"endpoint": "arms.cn-qingdao.chenxuewen.me"
				},
				{
					"region": "cn-beijing",
					"endpoint": "arms.cn-beijing.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "arms.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "arms.cn-shanghai.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": "arms.[RegionId].chenxuewen.me"
		},
		{
			"code": "live",
			"document_id": "48207",
			"location_service_code": "live",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "live.chenxuewen.me"
				},
				{
					"region": "ap-northeast-1",
					"endpoint": "live.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "live.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "live.chenxuewen.me"
				},
				{
					"region": "eu-central-1",
					"endpoint": "live.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "live.chenxuewen.me"
				},
				{
					"region": "cn-beijing",
					"endpoint": "live.chenxuewen.me"
				}
			],
			"global_endpoint": "live.chenxuewen.me",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "alimt",
			"document_id": "",
			"location_service_code": "alimt",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "mt.cn-hangzhou.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "actiontrail",
			"document_id": "",
			"location_service_code": "actiontrail",
			"regional_endpoints": [
				{
					"region": "cn-shanghai",
					"endpoint": "actiontrail.cn-shanghai.chenxuewen.me"
				},
				{
					"region": "cn-qingdao",
					"endpoint": "actiontrail.cn-qingdao.chenxuewen.me"
				},
				{
					"region": "us-east-1",
					"endpoint": "actiontrail.us-east-1.chenxuewen.me"
				},
				{
					"region": "ap-southeast-2",
					"endpoint": "actiontrail.ap-southeast-2.chenxuewen.me"
				},
				{
					"region": "ap-southeast-3",
					"endpoint": "actiontrail.ap-southeast-3.chenxuewen.me"
				},
				{
					"region": "ap-southeast-5",
					"endpoint": "actiontrail.ap-southeast-5.chenxuewen.me"
				},
				{
					"region": "ap-south-1",
					"endpoint": "actiontrail.ap-south-1.chenxuewen.me"
				},
				{
					"region": "me-east-1",
					"endpoint": "actiontrail.me-east-1.chenxuewen.me"
				},
				{
					"region": "cn-hongkong",
					"endpoint": "actiontrail.cn-hongkong.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "actiontrail.cn-shenzhen.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "actiontrail.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "eu-west-1",
					"endpoint": "actiontrail.eu-west-1.chenxuewen.me"
				},
				{
					"region": "cn-huhehaote",
					"endpoint": "actiontrail.cn-huhehaote.chenxuewen.me"
				},
				{
					"region": "ap-northeast-1",
					"endpoint": "actiontrail.ap-northeast-1.chenxuewen.me"
				},
				{
					"region": "us-west-1",
					"endpoint": "actiontrail.us-west-1.chenxuewen.me"
				},
				{
					"region": "eu-central-1",
					"endpoint": "actiontrail.eu-central-1.chenxuewen.me"
				},
				{
					"region": "cn-zhangjiakou",
					"endpoint": "actiontrail.cn-zhangjiakou.chenxuewen.me"
				},
				{
					"region": "cn-beijing",
					"endpoint": "actiontrail.cn-beijing.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "actiontrail.ap-southeast-1.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "smartag",
			"document_id": "",
			"location_service_code": "smartag",
			"regional_endpoints": [
				{
					"region": "ap-southeast-3",
					"endpoint": "smartag.ap-southeast-3.chenxuewen.me"
				},
				{
					"region": "ap-southeast-5",
					"endpoint": "smartag.ap-southeast-5.chenxuewen.me"
				},
				{
					"region": "eu-central-1",
					"endpoint": "smartag.eu-central-1.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "smartag.cn-shanghai.chenxuewen.me"
				},
				{
					"region": "cn-hongkong",
					"endpoint": "smartag.cn-hongkong.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "smartag.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "ap-southeast-2",
					"endpoint": "smartag.ap-southeast-2.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "vod",
			"document_id": "60574",
			"location_service_code": "vod",
			"regional_endpoints": [
				{
					"region": "cn-shanghai",
					"endpoint": "vod.cn-shanghai.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "vod.cn-shanghai.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "vod.cn-shanghai.chenxuewen.me"
				},
				{
					"region": "cn-beijing",
					"endpoint": "vod.cn-shanghai.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "vod.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "eu-central-1",
					"endpoint": "vod.eu-central-1.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "domain",
			"document_id": "42875",
			"location_service_code": "domain",
			"regional_endpoints": [
				{
					"region": "ap-southeast-1",
					"endpoint": "domain-intl.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "domain.chenxuewen.me"
				}
			],
			"global_endpoint": "domain.chenxuewen.me",
			"regional_endpoint_pattern": "domain.chenxuewen.me"
		},
		{
			"code": "ros",
			"document_id": "28899",
			"location_service_code": "ros",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "ros.chenxuewen.me"
				}
			],
			"global_endpoint": "ros.chenxuewen.me",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "cloudphoto",
			"document_id": "59902",
			"location_service_code": "cloudphoto",
			"regional_endpoints": [
				{
					"region": "cn-shanghai",
					"endpoint": "cloudphoto.cn-shanghai.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": "cloudphoto.[RegionId].chenxuewen.me"
		},
		{
			"code": "rtc",
			"document_id": "",
			"location_service_code": "rtc",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "rtc.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "odpsmayi",
			"document_id": "",
			"location_service_code": "odpsmayi",
			"regional_endpoints": [
				{
					"region": "cn-shanghai",
					"endpoint": "bsb.cloud.alipay.com"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "bsb.cloud.alipay.com"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "ims",
			"document_id": "",
			"location_service_code": "ims",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "ims.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "csb",
			"document_id": "64837",
			"location_service_code": "csb",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "csb.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "cn-beijing",
					"endpoint": "csb.cn-beijing.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": "csb.[RegionId].chenxuewen.me"
		},
		{
			"code": "cds",
			"document_id": "62887",
			"location_service_code": "codepipeline",
			"regional_endpoints": [
				{
					"region": "cn-beijing",
					"endpoint": "cds.cn-beijing.chenxuewen.me"
				}
			],
			"global_endpoint": "cds.cn-beijing.chenxuewen.me",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "ddosbgp",
			"document_id": "",
			"location_service_code": "ddosbgp",
			"regional_endpoints": [
				{
					"region": "cn-huhehaote",
					"endpoint": "ddosbgp.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "cn-beijing",
					"endpoint": "ddosbgp.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "cn-zhangjiakou",
					"endpoint": "ddosbgp.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "cn-hongkong",
					"endpoint": "ddosbgp.cn-hongkong.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "ddosbgp.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "us-west-1",
					"endpoint": "ddosbgp.us-west-1.chenxuewen.me"
				},
				{
					"region": "cn-qingdao",
					"endpoint": "ddosbgp.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "ddosbgp.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "ddosbgp.cn-hangzhou.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "dybaseapi",
			"document_id": "",
			"location_service_code": "dybaseapi",
			"regional_endpoints": [
				{
					"region": "cn-beijing",
					"endpoint": "dybaseapi.chenxuewen.me"
				},
				{
					"region": "cn-chengdu",
					"endpoint": "dybaseapi.chenxuewen.me"
				},
				{
					"region": "ap-southeast-3",
					"endpoint": "dybaseapi.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "dybaseapi.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "us-west-1",
					"endpoint": "dybaseapi.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "dybaseapi.chenxuewen.me"
				},
				{
					"region": "cn-zhangjiakou",
					"endpoint": "dybaseapi.chenxuewen.me"
				},
				{
					"region": "me-east-1",
					"endpoint": "dybaseapi.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "us-east-1",
					"endpoint": "dybaseapi.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "ap-northeast-1",
					"endpoint": "dybaseapi.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "ap-southeast-5",
					"endpoint": "dybaseapi.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "dybaseapi.chenxuewen.me"
				},
				{
					"region": "cn-huhehaote",
					"endpoint": "dybaseapi.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "dybaseapi.chenxuewen.me"
				},
				{
					"region": "eu-central-1",
					"endpoint": "dybaseapi.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "cn-qingdao",
					"endpoint": "dybaseapi.chenxuewen.me"
				},
				{
					"region": "cn-hongkong",
					"endpoint": "dybaseapi.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "ap-southeast-2",
					"endpoint": "dybaseapi.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "ap-south-1",
					"endpoint": "dybaseapi.ap-southeast-1.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "ecs",
			"document_id": "25484",
			"location_service_code": "ecs",
			"regional_endpoints": [
				{
					"region": "cn-huhehaote",
					"endpoint": "ecs.cn-huhehaote.chenxuewen.me"
				},
				{
					"region": "ap-northeast-1",
					"endpoint": "ecs.ap-northeast-1.chenxuewen.me"
				},
				{
					"region": "ap-southeast-2",
					"endpoint": "ecs.ap-southeast-2.chenxuewen.me"
				},
				{
					"region": "ap-south-1",
					"endpoint": "ecs.ap-south-1.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "ecs-cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "ecs-cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "ap-southeast-3",
					"endpoint": "ecs.ap-southeast-3.chenxuewen.me"
				},
				{
					"region": "eu-central-1",
					"endpoint": "ecs.eu-central-1.chenxuewen.me"
				},
				{
					"region": "cn-zhangjiakou",
					"endpoint": "ecs.cn-zhangjiakou.chenxuewen.me"
				},
				{
					"region": "cn-hongkong",
					"endpoint": "ecs-cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "eu-west-1",
					"endpoint": "ecs.eu-west-1.chenxuewen.me"
				},
				{
					"region": "me-east-1",
					"endpoint": "ecs.me-east-1.chenxuewen.me"
				},
				{
					"region": "cn-qingdao",
					"endpoint": "ecs-cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "ecs-cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "us-west-1",
					"endpoint": "ecs-cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "us-east-1",
					"endpoint": "ecs-cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "ap-southeast-5",
					"endpoint": "ecs.ap-southeast-5.chenxuewen.me"
				},
				{
					"region": "cn-beijing",
					"endpoint": "ecs-cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "ecs-cn-hangzhou.chenxuewen.me"
				}
			],
			"global_endpoint": "ecs-cn-hangzhou.chenxuewen.me",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "ccc",
			"document_id": "63027",
			"location_service_code": "ccc",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "ccc.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "ccc.cn-shanghai.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": "ccc.[RegionId].chenxuewen.me"
		},
		{
			"code": "cs",
			"document_id": "26043",
			"location_service_code": "cs",
			"regional_endpoints": null,
			"global_endpoint": "cs.chenxuewen.me",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "drdspre",
			"document_id": "",
			"location_service_code": "drdspre",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "drds.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "drds.cn-shanghai.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "drds.cn-shenzhen.chenxuewen.me"
				},
				{
					"region": "cn-hongkong",
					"endpoint": "drds.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "drds.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "cn-qingdao",
					"endpoint": "drds.cn-qingdao.chenxuewen.me"
				},
				{
					"region": "cn-beijing",
					"endpoint": "drds.cn-beijing.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "dcdn",
			"document_id": "",
			"location_service_code": "dcdn",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "dcdn.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "dcdn.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "linkedmall",
			"document_id": "",
			"location_service_code": "linkedmall",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "linkedmall.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "linkedmall.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "trademark",
			"document_id": "",
			"location_service_code": "trademark",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "trademark.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "openanalytics",
			"document_id": "",
			"location_service_code": "openanalytics",
			"regional_endpoints": [
				{
					"region": "cn-shenzhen",
					"endpoint": "openanalytics.cn-shenzhen.chenxuewen.me"
				},
				{
					"region": "eu-west-1",
					"endpoint": "openanalytics.eu-west-1.chenxuewen.me"
				},
				{
					"region": "cn-beijing",
					"endpoint": "openanalytics.cn-beijing.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "openanalytics.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "openanalytics.cn-shanghai.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "openanalytics.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "ap-southeast-3",
					"endpoint": "openanalytics.ap-southeast-3.chenxuewen.me"
				},
				{
					"region": "cn-zhangjiakou",
					"endpoint": "openanalytics.cn-zhangjiakou.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "sts",
			"document_id": "28756",
			"location_service_code": "sts",
			"regional_endpoints": null,
			"global_endpoint": "sts.chenxuewen.me",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "waf",
			"document_id": "62847",
			"location_service_code": "waf",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "wafopenapi.cn-hangzhou.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "ots",
			"document_id": "",
			"location_service_code": "ots",
			"regional_endpoints": [
				{
					"region": "me-east-1",
					"endpoint": "ots.me-east-1.chenxuewen.me"
				},
				{
					"region": "ap-southeast-5",
					"endpoint": "ots.ap-southeast-5.chenxuewen.me"
				},
				{
					"region": "eu-west-1",
					"endpoint": "ots.eu-west-1.chenxuewen.me"
				},
				{
					"region": "cn-huhehaote",
					"endpoint": "ots.cn-huhehaote.chenxuewen.me"
				},
				{
					"region": "cn-beijing",
					"endpoint": "ots.cn-beijing.chenxuewen.me"
				},
				{
					"region": "ap-southeast-2",
					"endpoint": "ots.ap-southeast-2.chenxuewen.me"
				},
				{
					"region": "us-west-1",
					"endpoint": "ots.us-west-1.chenxuewen.me"
				},
				{
					"region": "us-east-1",
					"endpoint": "ots.us-east-1.chenxuewen.me"
				},
				{
					"region": "ap-south-1",
					"endpoint": "ots.ap-south-1.chenxuewen.me"
				},
				{
					"region": "eu-central-1",
					"endpoint": "ots.eu-central-1.chenxuewen.me"
				},
				{
					"region": "cn-zhangjiakou",
					"endpoint": "ots.cn-zhangjiakou.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "ots.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "ap-northeast-1",
					"endpoint": "ots.ap-northeast-1.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "ots.cn-shanghai.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "ots.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "ap-southeast-3",
					"endpoint": "ots.ap-southeast-3.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "cloudfirewall",
			"document_id": "",
			"location_service_code": "cloudfirewall",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "cloudfw.cn-hangzhou.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "dm",
			"document_id": "29434",
			"location_service_code": "dm",
			"regional_endpoints": [
				{
					"region": "ap-southeast-2",
					"endpoint": "dm.ap-southeast-2.chenxuewen.me"
				}
			],
			"global_endpoint": "dm.chenxuewen.me",
			"regional_endpoint_pattern": "dm.[RegionId].chenxuewen.me"
		},
		{
			"code": "oas",
			"document_id": "",
			"location_service_code": "oas",
			"regional_endpoints": [
				{
					"region": "cn-shenzhen",
					"endpoint": "cn-shenzhen.oas.chenxuewen.me"
				},
				{
					"region": "cn-beijing",
					"endpoint": "cn-beijing.oas.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "cn-hangzhou.oas.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "ddoscoo",
			"document_id": "",
			"location_service_code": "ddoscoo",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "ddoscoo.cn-hangzhou.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "jaq",
			"document_id": "35037",
			"location_service_code": "jaq",
			"regional_endpoints": null,
			"global_endpoint": "jaq.chenxuewen.me",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "iovcc",
			"document_id": "",
			"location_service_code": "iovcc",
			"regional_endpoints": [
				{
					"region": "cn-shanghai",
					"endpoint": "iovcc.cn-shanghai.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "sas-api",
			"document_id": "28498",
			"location_service_code": "sas",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "sas.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "chatbot",
			"document_id": "60760",
			"location_service_code": "beebot",
			"regional_endpoints": [
				{
					"region": "cn-shanghai",
					"endpoint": "chatbot.cn-shanghai.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "chatbot.cn-hangzhou.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": "chatbot.[RegionId].chenxuewen.me"
		},
		{
			"code": "airec",
			"document_id": "",
			"location_service_code": "airec",
			"regional_endpoints": [
				{
					"region": "cn-beijing",
					"endpoint": "airec.cn-beijing.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "airec.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "airec.cn-shanghai.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "dmsenterprise",
			"document_id": "",
			"location_service_code": "dmsenterprise",
			"regional_endpoints": [
				{
					"region": "cn-shanghai",
					"endpoint": "dms-enterprise.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "dms-enterprise.chenxuewen.me"
				},
				{
					"region": "ap-northeast-1",
					"endpoint": "dms-enterprise.chenxuewen.me"
				},
				{
					"region": "cn-qingdao",
					"endpoint": "dms-enterprise.chenxuewen.me"
				},
				{
					"region": "cn-beijing",
					"endpoint": "dms-enterprise.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "dms-enterprise.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "ivision",
			"document_id": "",
			"location_service_code": "ivision",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "ivision.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "cn-beijing",
					"endpoint": "ivision.cn-beijing.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "odpsplusmayi",
			"document_id": "",
			"location_service_code": "odpsplusmayi",
			"regional_endpoints": [
				{
					"region": "cn-shanghai",
					"endpoint": "bsb.cloud.alipay.com"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "bsb.cloud.alipay.com"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "bsb.cloud.alipay.com"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "gameshield",
			"document_id": "",
			"location_service_code": "gameshield",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "gameshield.chenxuewen.me"
				},
				{
					"region": "cn-zhangjiakou",
					"endpoint": "gameshield.cn-zhangjiakou.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "scdn",
			"document_id": "",
			"location_service_code": "scdn",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "scdn.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "hitsdb",
			"document_id": "",
			"location_service_code": "hitsdb",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "hitsdb.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "hdm",
			"document_id": "",
			"location_service_code": "hdm",
			"regional_endpoints": [
				{
					"region": "cn-shanghai",
					"endpoint": "hdm-api.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "slb",
			"document_id": "27565",
			"location_service_code": "slb",
			"regional_endpoints": [
				{
					"region": "cn-shenzhen",
					"endpoint": "slb.chenxuewen.me"
				},
				{
					"region": "eu-west-1",
					"endpoint": "slb.eu-west-1.chenxuewen.me"
				},
				{
					"region": "ap-northeast-1",
					"endpoint": "slb.ap-northeast-1.chenxuewen.me"
				},
				{
					"region": "ap-southeast-2",
					"endpoint": "slb.ap-southeast-2.chenxuewen.me"
				},
				{
					"region": "cn-qingdao",
					"endpoint": "slb.chenxuewen.me"
				},
				{
					"region": "cn-hongkong",
					"endpoint": "slb.chenxuewen.me"
				},
				{
					"region": "cn-huhehaote",
					"endpoint": "slb.cn-huhehaote.chenxuewen.me"
				},
				{
					"region": "ap-south-1",
					"endpoint": "slb.ap-south-1.chenxuewen.me"
				},
				{
					"region": "eu-central-1",
					"endpoint": "slb.eu-central-1.chenxuewen.me"
				},
				{
					"region": "cn-beijing",
					"endpoint": "slb.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "slb.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "slb.chenxuewen.me"
				},
				{
					"region": "us-west-1",
					"endpoint": "slb.chenxuewen.me"
				},
				{
					"region": "ap-southeast-5",
					"endpoint": "slb.ap-southeast-5.chenxuewen.me"
				},
				{
					"region": "ap-southeast-3",
					"endpoint": "slb.ap-southeast-3.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "slb.chenxuewen.me"
				},
				{
					"region": "us-east-1",
					"endpoint": "slb.chenxuewen.me"
				},
				{
					"region": "me-east-1",
					"endpoint": "slb.me-east-1.chenxuewen.me"
				},
				{
					"region": "cn-zhangjiakou",
					"endpoint": "slb.cn-zhangjiakou.chenxuewen.me"
				},
				{
					"region": "cn-wulanchabu",
					"endpoint": "slb.cn-wulanchabu.chenxuewen.me"
				},
				{
					"region": "cn-heyuan",
					"endpoint": "slb.cn-heyuan.chenxuewen.me"
				},
				{
					"region": "cn-guangzhou",
					"endpoint": "slb.cn-guangzhou.chenxuewen.me"
				},
				{
					"region": "cn-chengdu",
					"endpoint": "slb.cn-chengdu.chenxuewen.me"
				},
				{
					"region": "ap-southeast-6",
					"endpoint": "slb.ap-southeast-6.chenxuewen.me"
				},
				{
					"region": "eu-west-1",
					"endpoint": "slb.eu-west-1.chenxuewen.me"
				}
			],
			"global_endpoint": "slb.chenxuewen.me",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "green",
			"document_id": "28427",
			"location_service_code": "green",
			"regional_endpoints": [
				{
					"region": "cn-beijing",
					"endpoint": "green.cn-beijing.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "green.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "green.cn-shanghai.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "green.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "us-west-1",
					"endpoint": "green.us-west-1.chenxuewen.me"
				}
			],
			"global_endpoint": "green.chenxuewen.me",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "cccvn",
			"document_id": "",
			"location_service_code": "cccvn",
			"regional_endpoints": [
				{
					"region": "cn-shanghai",
					"endpoint": "voicenavigator.cn-shanghai.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "ddosrewards",
			"document_id": "",
			"location_service_code": "ddosrewards",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "ddosright.cn-hangzhou.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "iot",
			"document_id": "30557",
			"location_service_code": "iot",
			"regional_endpoints": [
				{
					"region": "us-east-1",
					"endpoint": "iot.us-east-1.chenxuewen.me"
				},
				{
					"region": "ap-northeast-1",
					"endpoint": "iot.ap-northeast-1.chenxuewen.me"
				},
				{
					"region": "us-west-1",
					"endpoint": "iot.us-west-1.chenxuewen.me"
				},
				{
					"region": "eu-central-1",
					"endpoint": "iot.eu-central-1.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "iot.cn-shanghai.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "iot.ap-southeast-1.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": "iot.[RegionId].chenxuewen.me"
		},
		{
			"code": "bssopenapi",
			"document_id": "",
			"location_service_code": "bssopenapi",
			"regional_endpoints": [
				{
					"region": "cn-shanghai",
					"endpoint": "business.chenxuewen.me"
				},
				{
					"region": "cn-hongkong",
					"endpoint": "business.chenxuewen.me"
				},
				{
					"region": "ap-southeast-2",
					"endpoint": "business.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "ap-southeast-3",
					"endpoint": "business.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "ap-northeast-1",
					"endpoint": "business.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "eu-central-1",
					"endpoint": "business.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "me-east-1",
					"endpoint": "business.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "cn-zhangjiakou",
					"endpoint": "business.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "business.chenxuewen.me"
				},
				{
					"region": "ap-southeast-5",
					"endpoint": "business.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "us-east-1",
					"endpoint": "business.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "cn-qingdao",
					"endpoint": "business.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "business.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "cn-beijing",
					"endpoint": "business.chenxuewen.me"
				},
				{
					"region": "cn-huhehaote",
					"endpoint": "business.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "business.chenxuewen.me"
				},
				{
					"region": "us-west-1",
					"endpoint": "business.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "eu-west-1",
					"endpoint": "business.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "ap-south-1",
					"endpoint": "business.ap-southeast-1.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "sca",
			"document_id": "",
			"location_service_code": "sca",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "qualitycheck.cn-hangzhou.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "luban",
			"document_id": "",
			"location_service_code": "luban",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "luban.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "luban.cn-shanghai.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "drdspost",
			"document_id": "",
			"location_service_code": "drdspost",
			"regional_endpoints": [
				{
					"region": "cn-shanghai",
					"endpoint": "drds.cn-shanghai.chenxuewen.me"
				},
				{
					"region": "cn-hongkong",
					"endpoint": "drds.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "drds.ap-southeast-1.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "drds",
			"document_id": "51111",
			"location_service_code": "drds",
			"regional_endpoints": [
				{
					"region": "ap-southeast-1",
					"endpoint": "drds.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "drds.cn-hangzhou.chenxuewen.me"
				}
			],
			"global_endpoint": "drds.chenxuewen.me",
			"regional_endpoint_pattern": "drds.chenxuewen.me"
		},
		{
			"code": "httpdns",
			"document_id": "52679",
			"location_service_code": "httpdns",
			"regional_endpoints": null,
			"global_endpoint": "httpdns-api.chenxuewen.me",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "cas",
			"document_id": "",
			"location_service_code": "cas",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "cas.chenxuewen.me"
				},
				{
					"region": "ap-southeast-2",
					"endpoint": "cas.ap-southeast-2.chenxuewen.me"
				},
				{
					"region": "ap-northeast-1",
					"endpoint": "cas.ap-northeast-1.chenxuewen.me"
				},
				{
					"region": "eu-central-1",
					"endpoint": "cas.eu-central-1.chenxuewen.me"
				},
				{
					"region": "me-east-1",
					"endpoint": "cas.me-east-1.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "hpc",
			"document_id": "35201",
			"location_service_code": "hpc",
			"regional_endpoints": [
				{
					"region": "cn-beijing",
					"endpoint": "hpc.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "hpc.chenxuewen.me"
				}
			],
			"global_endpoint": "hpc.chenxuewen.me",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "ddosbasic",
			"document_id": "",
			"location_service_code": "ddosbasic",
			"regional_endpoints": [
				{
					"region": "ap-south-1",
					"endpoint": "antiddos-openapi.ap-south-1.chenxuewen.me"
				},
				{
					"region": "cn-zhangjiakou",
					"endpoint": "antiddos-openapi.cn-zhangjiakou.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "antiddos.chenxuewen.me"
				},
				{
					"region": "us-east-1",
					"endpoint": "antiddos.chenxuewen.me"
				},
				{
					"region": "ap-southeast-3",
					"endpoint": "antiddos-openapi.ap-southeast-3.chenxuewen.me"
				},
				{
					"region": "me-east-1",
					"endpoint": "antiddos-openapi.me-east-1.chenxuewen.me"
				},
				{
					"region": "ap-southeast-5",
					"endpoint": "antiddos-openapi.ap-southeast-5.chenxuewen.me"
				},
				{
					"region": "ap-southeast-2",
					"endpoint": "antiddos-openapi.ap-southeast-2.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "antiddos.chenxuewen.me"
				},
				{
					"region": "cn-hongkong",
					"endpoint": "antiddos.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "antiddos.chenxuewen.me"
				},
				{
					"region": "eu-west-1",
					"endpoint": "antiddos-openapi.eu-west-1.chenxuewen.me"
				},
				{
					"region": "cn-huhehaote",
					"endpoint": "antiddos-openapi.cn-huhehaote.chenxuewen.me"
				},
				{
					"region": "cn-beijing",
					"endpoint": "antiddos.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "antiddos.chenxuewen.me"
				},
				{
					"region": "us-west-1",
					"endpoint": "antiddos.chenxuewen.me"
				},
				{
					"region": "eu-central-1",
					"endpoint": "antiddos-openapi.eu-central-1.chenxuewen.me"
				},
				{
					"region": "cn-qingdao",
					"endpoint": "antiddos.chenxuewen.me"
				},
				{
					"region": "ap-northeast-1",
					"endpoint": "antiddos-openapi.ap-northeast-1.chenxuewen.me"
				}
			],
			"global_endpoint": "antiddos.chenxuewen.me",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "clouddesktop",
			"document_id": "",
			"location_service_code": "clouddesktop",
			"regional_endpoints": [
				{
					"region": "cn-beijing",
					"endpoint": "clouddesktop.cn-beijing.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "clouddesktop.cn-shanghai.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "clouddesktop.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "clouddesktop.cn-shenzhen.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "uis",
			"document_id": "",
			"location_service_code": "uis",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "uis.cn-hangzhou.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "imm",
			"document_id": "",
			"location_service_code": "imm",
			"regional_endpoints": [
				{
					"region": "cn-beijing",
					"endpoint": "imm.cn-beijing.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "imm.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "cn-zhangjiakou",
					"endpoint": "imm.cn-zhangjiakou.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "imm.cn-shanghai.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "imm.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "imm.cn-shenzhen.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "ens",
			"document_id": "",
			"location_service_code": "ens",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "ens.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "ram",
			"document_id": "28672",
			"location_service_code": "ram",
			"regional_endpoints": null,
			"global_endpoint": "ram.chenxuewen.me",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "hcs_mgw",
			"document_id": "",
			"location_service_code": "hcs_mgw",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "mgw.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "mgw.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "mgw.cn-shanghai.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "itaas",
			"document_id": "55759",
			"location_service_code": "itaas",
			"regional_endpoints": null,
			"global_endpoint": "itaas.chenxuewen.me",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "qualitycheck",
			"document_id": "50807",
			"location_service_code": "qualitycheck",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "qualitycheck.cn-hangzhou.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "alikafka",
			"document_id": "",
			"location_service_code": "alikafka",
			"regional_endpoints": [
				{
					"region": "cn-beijing",
					"endpoint": "alikafka.cn-beijing.chenxuewen.me"
				},
				{
					"region": "cn-zhangjiakou",
					"endpoint": "alikafka.cn-zhangjiakou.chenxuewen.me"
				},
				{
					"region": "cn-huhehaote",
					"endpoint": "alikafka.cn-huhehaote.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "alikafka.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "alikafka.cn-shanghai.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "alikafka.cn-shenzhen.chenxuewen.me"
				},
				{
					"region": "cn-hongkong",
					"endpoint": "alikafka.cn-hongkong.chenxuewen.me"
				},
				{
					"region": "cn-qingdao",
					"endpoint": "alikafka.cn-qingdao.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "faas",
			"document_id": "",
			"location_service_code": "faas",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "faas.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "faas.cn-shenzhen.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "faas.cn-shanghai.chenxuewen.me"
				},
				{
					"region": "cn-beijing",
					"endpoint": "faas.cn-beijing.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "alidfs",
			"document_id": "",
			"location_service_code": "alidfs",
			"regional_endpoints": [
				{
					"region": "cn-beijing",
					"endpoint": "dfs.cn-beijing.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "dfs.cn-shanghai.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "cms",
			"document_id": "28615",
			"location_service_code": "cms",
			"regional_endpoints": [
				{
					"region": "ap-southeast-3",
					"endpoint": "metrics.ap-southeast-3.chenxuewen.me"
				},
				{
					"region": "ap-southeast-5",
					"endpoint": "metrics.ap-southeast-5.chenxuewen.me"
				},
				{
					"region": "ap-south-1",
					"endpoint": "metrics.ap-south-1.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "metrics.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "eu-west-1",
					"endpoint": "metrics.eu-west-1.chenxuewen.me"
				},
				{
					"region": "cn-huhehaote",
					"endpoint": "metrics.cn-huhehaote.chenxuewen.me"
				},
				{
					"region": "ap-northeast-1",
					"endpoint": "metrics.ap-northeast-1.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "domain-intl",
			"document_id": "",
			"location_service_code": "domain-intl",
			"regional_endpoints": null,
			"global_endpoint": "domain-intl.chenxuewen.me",
			"regional_endpoint_pattern": "domain-intl.chenxuewen.me"
		},
		{
			"code": "kvstore",
			"document_id": "",
			"location_service_code": "kvstore",
			"regional_endpoints": [
				{
					"region": "ap-northeast-1",
					"endpoint": "r-kvstore.ap-northeast-1.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "ccs",
			"document_id": "",
			"location_service_code": "ccs",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "ccs.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "ess",
			"document_id": "25925",
			"location_service_code": "ess",
			"regional_endpoints": [
				{
					"region": "cn-huhehaote",
					"endpoint": "ess.cn-huhehaote.chenxuewen.me"
				},
				{
					"region": "ap-southeast-2",
					"endpoint": "ess.ap-southeast-2.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "ess.chenxuewen.me"
				},
				{
					"region": "cn-hongkong",
					"endpoint": "ess.chenxuewen.me"
				},
				{
					"region": "eu-west-1",
					"endpoint": "ess.eu-west-1.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "ess.chenxuewen.me"
				},
				{
					"region": "us-west-1",
					"endpoint": "ess.chenxuewen.me"
				},
				{
					"region": "ap-southeast-5",
					"endpoint": "ess.ap-southeast-5.chenxuewen.me"
				},
				{
					"region": "ap-southeast-3",
					"endpoint": "ess.ap-southeast-3.chenxuewen.me"
				},
				{
					"region": "cn-zhangjiakou",
					"endpoint": "ess.cn-zhangjiakou.chenxuewen.me"
				},
				{
					"region": "cn-beijing",
					"endpoint": "ess.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "ess.chenxuewen.me"
				},
				{
					"region": "me-east-1",
					"endpoint": "ess.me-east-1.chenxuewen.me"
				},
				{
					"region": "ap-northeast-1",
					"endpoint": "ess.ap-northeast-1.chenxuewen.me"
				},
				{
					"region": "ap-south-1",
					"endpoint": "ess.ap-south-1.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "ess.chenxuewen.me"
				},
				{
					"region": "eu-central-1",
					"endpoint": "ess.eu-central-1.chenxuewen.me"
				},
				{
					"region": "cn-qingdao",
					"endpoint": "ess.chenxuewen.me"
				},
				{
					"region": "us-east-1",
					"endpoint": "ess.chenxuewen.me"
				}
			],
			"global_endpoint": "ess.chenxuewen.me",
			"regional_endpoint_pattern": "ess.[RegionId].chenxuewen.me"
		},
		{
			"code": "dds",
			"document_id": "61715",
			"location_service_code": "dds",
			"regional_endpoints": [
				{
					"region": "me-east-1",
					"endpoint": "mongodb.me-east-1.chenxuewen.me"
				},
				{
					"region": "ap-southeast-3",
					"endpoint": "mongodb.ap-southeast-3.chenxuewen.me"
				},
				{
					"region": "ap-southeast-7",
					"endpoint": "mongodb.ap-southeast-7.chenxuewen.me"
				},
				{
					"region": "ap-southeast-5",
					"endpoint": "mongodb.ap-southeast-5.chenxuewen.me"
				},
				{
					"region": "cn-hongkong",
					"endpoint": "mongodb.chenxuewen.me"
				},
				{
					"region": "ap-northeast-1",
					"endpoint": "mongodb.ap-northeast-1.chenxuewen.me"
				},
				{
					"region": "eu-central-1",
					"endpoint": "mongodb.eu-central-1.chenxuewen.me"
				},
				{
					"region": "cn-beijing",
					"endpoint": "mongodb.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "mongodb.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "mongodb.chenxuewen.me"
				},
				{
					"region": "eu-west-1",
					"endpoint": "mongodb.eu-west-1.chenxuewen.me"
				},
				{
					"region": "us-east-1",
					"endpoint": "mongodb.chenxuewen.me"
				},
				{
					"region": "cn-huhehaote",
					"endpoint": "mongodb.cn-huhehaote.chenxuewen.me"
				},
				{
					"region": "cn-qingdao",
					"endpoint": "mongodb.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "mongodb.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "mongodb.chenxuewen.me"
				},
				{
					"region": "ap-southeast-2",
					"endpoint": "mongodb.ap-southeast-2.chenxuewen.me"
				},
				{
					"region": "cn-zhangjiakou",
					"endpoint": "mongodb.cn-zhangjiakou.chenxuewen.me"
				},
				{
					"region": "ap-south-1",
					"endpoint": "mongodb.ap-south-1.chenxuewen.me"
				},
				{
					"region": "us-west-1",
					"endpoint": "mongodb.chenxuewen.me"
				}
			],
			"global_endpoint": "mongodb.chenxuewen.me",
			"regional_endpoint_pattern": "mongodb.[RegionId].chenxuewen.me"
		},
		{
			"code": "mts",
			"document_id": "29212",
			"location_service_code": "mts",
			"regional_endpoints": [
				{
					"region": "cn-beijing",
					"endpoint": "mts.cn-beijing.chenxuewen.me"
				},
				{
					"region": "ap-northeast-1",
					"endpoint": "mts.ap-northeast-1.chenxuewen.me"
				},
				{
					"region": "cn-hongkong",
					"endpoint": "mts.cn-hongkong.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "mts.cn-shenzhen.chenxuewen.me"
				},
				{
					"region": "cn-zhangjiakou",
					"endpoint": "mts.cn-zhangjiakou.chenxuewen.me"
				},
				{
					"region": "ap-south-1",
					"endpoint": "mts.ap-south-1.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "mts.cn-shanghai.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "mts.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "us-west-1",
					"endpoint": "mts.us-west-1.chenxuewen.me"
				},
				{
					"region": "eu-central-1",
					"endpoint": "mts.eu-central-1.chenxuewen.me"
				},
				{
					"region": "eu-west-1",
					"endpoint": "mts.eu-west-1.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "mts.cn-hangzhou.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "push",
			"document_id": "30074",
			"location_service_code": "push",
			"regional_endpoints": null,
			"global_endpoint": "cloudpush.chenxuewen.me",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "hcs_sgw",
			"document_id": "",
			"location_service_code": "hcs_sgw",
			"regional_endpoints": [
				{
					"region": "eu-central-1",
					"endpoint": "sgw.eu-central-1.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "sgw.cn-shanghai.chenxuewen.me"
				},
				{
					"region": "cn-zhangjiakou",
					"endpoint": "sgw.cn-shanghai.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "sgw.ap-southeast-1.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "sgw.cn-shanghai.chenxuewen.me"
				},
				{
					"region": "cn-hongkong",
					"endpoint": "sgw.cn-shanghai.chenxuewen.me"
				},
				{
					"region": "ap-southeast-2",
					"endpoint": "sgw.ap-southeast-2.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "sgw.cn-shanghai.chenxuewen.me"
				},
				{
					"region": "cn-qingdao",
					"endpoint": "sgw.cn-shanghai.chenxuewen.me"
				},
				{
					"region": "cn-beijing",
					"endpoint": "sgw.cn-shanghai.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "hbase",
			"document_id": "",
			"location_service_code": "hbase",
			"regional_endpoints": [
				{
					"region": "cn-huhehaote",
					"endpoint": "hbase.cn-huhehaote.chenxuewen.me"
				},
				{
					"region": "ap-south-1",
					"endpoint": "hbase.ap-south-1.chenxuewen.me"
				},
				{
					"region": "us-west-1",
					"endpoint": "hbase.chenxuewen.me"
				},
				{
					"region": "me-east-1",
					"endpoint": "hbase.me-east-1.chenxuewen.me"
				},
				{
					"region": "eu-central-1",
					"endpoint": "hbase.eu-central-1.chenxuewen.me"
				},
				{
					"region": "cn-qingdao",
					"endpoint": "hbase.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "hbase.chenxuewen.me"
				},
				{
					"region": "cn-shenzhen",
					"endpoint": "hbase.chenxuewen.me"
				},
				{
					"region": "ap-southeast-2",
					"endpoint": "hbase.ap-southeast-2.chenxuewen.me"
				},
				{
					"region": "ap-southeast-3",
					"endpoint": "hbase.ap-southeast-3.chenxuewen.me"
				},
				{
					"region": "cn-hangzhou",
					"endpoint": "hbase.chenxuewen.me"
				},
				{
					"region": "us-east-1",
					"endpoint": "hbase.chenxuewen.me"
				},
				{
					"region": "ap-southeast-5",
					"endpoint": "hbase.ap-southeast-5.chenxuewen.me"
				},
				{
					"region": "cn-beijing",
					"endpoint": "hbase.chenxuewen.me"
				},
				{
					"region": "ap-southeast-1",
					"endpoint": "hbase.chenxuewen.me"
				}
			],
			"global_endpoint": "hbase.chenxuewen.me",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "bastionhost",
			"document_id": "",
			"location_service_code": "bastionhost",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "yundun-bastionhost.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		},
		{
			"code": "vs",
			"document_id": "",
			"location_service_code": "vs",
			"regional_endpoints": [
				{
					"region": "cn-hangzhou",
					"endpoint": "vs.cn-hangzhou.chenxuewen.me"
				},
				{
					"region": "cn-shanghai",
					"endpoint": "vs.cn-shanghai.chenxuewen.me"
				}
			],
			"global_endpoint": "",
			"regional_endpoint_pattern": ""
		}
	]
}`

var initOnce sync.Once
var data interface{}

func getEndpointConfigData() interface{} {
	initOnce.Do(func() {
		err := json.Unmarshal([]byte(endpointsJson), &data)
		if err != nil {
			panic(fmt.Sprintf("init endpoint config data failed. %s", err))
		}
	})
	return data
}
