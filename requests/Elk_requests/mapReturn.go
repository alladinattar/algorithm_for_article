package Elk_requests

type request struct {
}

func MapReturn(funcName string) string {

	var requestMap = map[string]string{
		"mining1860": `{
"size": 10000,
"_source": [
"pAz",   //ФИО специалиста
"u",   //Название услуги
"tn",   //Тип назначения
"qqc",   //qqc 1860
"dE",   //Дата
"Du",   //код ОКМУ
"pID"   //Отделение
],
"query": {
    "bool": {
        "filter": {
            "bool": {
                "filter": {
                    "bool": {
                        "should": [
                            {"term": {"pID.keyword": "оУЗД"}}
                        ]
                    }
                },
                "must": [
                    {"range":{"dE":{"gte":####}}},
                    {"range":{"dE":{"lte":****}}}
                ],
                "minimum_should_match": 1,
                "should": [
                    {"term": {"Du.keyword": "A04.01.001"}},
                    {"term": {"Du.keyword": "A04.06.001"}},
                    {"term": {"Du.keyword": "A04.06.002"}},
                    {"term": {"Du.keyword": "A04.06.004"}},
                    {"term": {"Du.keyword": "A04.06.005"}},
                    {"term": {"Du.keyword": "A04.10.002"}},
                    {"term": {"Du.keyword": "A04.10.002.010"}},
                    {"term": {"Du.keyword": "A04.10.009"}},
                    {"term": {"Du.keyword": "A04.12.001.002"}},
                    {"term": {"Du.keyword": "A04.12.006.001"}},
                    {"term": {"Du.keyword": "A04.12.011"}},
                    {"term": {"Du.keyword": "A04.12.013"}},
                    {"term": {"Du.keyword": "A04.12.048"}},
                    {"term": {"Du.keyword": "A04.14.001"}},
                    {"term": {"Du.keyword": "A04.14.003"}},
                    {"term": {"Du.keyword": "A04.15.001"}},
                    {"term": {"Du.keyword": "A04.20.001"}},
                    {"term": {"Du.keyword": "A04.20.001.002"}},
                    {"term": {"Du.keyword": "A04.20.001.003"}},
                    {"term": {"Du.keyword": "A04.20.002"}},
                    {"term": {"Du.keyword": "A04.20.003"}},
                    {"term": {"Du.keyword": "A04.20.011"}},
                    {"term": {"Du.keyword": "A04.20.012"}},
                    {"term": {"Du.keyword": "A04.20.013"}},
                    {"term": {"Du.keyword": "A04.21.001"}},
                    {"term": {"Du.keyword": "A04.21.001.002"}},
                    {"term": {"Du.keyword": "A04.21.001.003"}},
                    {"term": {"Du.keyword": "A04.21.002"}},
                    {"term": {"Du.keyword": "A04.21.002.001"}},
                    {"term": {"Du.keyword": "A04.21.003"}},
                    {"term": {"Du.keyword": "A04.21.011"}},
                    {"term": {"Du.keyword": "A04.21.012"}},
                    {"term": {"Du.keyword": "A04.21.013"}},
                    {"term": {"Du.keyword": "A04.21.014"}},
                    {"term": {"Du.keyword": "A04.21.015"}},
                    {"term": {"Du.keyword": "A04.21.016"}},
                    {"term": {"Du.keyword": "A04.22.001"}},
                    {"term": {"Du.keyword": "A04.22.002"}},
                    {"term": {"Du.keyword": "A04.22.004"}},
                    {"term": {"Du.keyword": "A04.22.008"}},
                    {"term": {"Du.keyword": "A04.22.010"}},
                    {"term": {"Du.keyword": "A04.22.011"}},
                    {"term": {"Du.keyword": "A04.24.001"}},
                    {"term": {"Du.keyword": "A04.26.001.001"}},
                    {"term": {"Du.keyword": "A04.26.006"}},
                    {"term": {"Du.keyword": "A04.28.001"}},
                    {"term": {"Du.keyword": "A04.28.002.003"}},
                    {"term": {"Du.keyword": "A04.31.004"}},
                    {"term": {"Du.keyword": "A04.31.005"}},
                    {"term": {"Du.keyword": "A04.31.006"}},
                    {"term": {"Du.keyword": "A04.31.010"}},
                    {"term": {"Du.keyword": "A04.31.012"}},
                    {"term": {"Du.keyword": "A06.26.010.001"}},
                    {"term": {"Du.keyword": "A06.26.010.002"}},
                    {"term": {"Du.keyword": "B03.001.005"}},
                    {"term": {"Du.keyword": "B03.058.34"}},
                    {"term": {"Du.keyword": "B03.058.35"}}
                ]
            }
        }
    }
}
}`,

		"mining174": `{
"size": 10000, 
"_source": [
"puR",   //Должность специалиста
"pvs",   //Вид поступления
"pKDiag",   //Диагноз выписки
"pAG",   //Дата выписки пациента
"pAz",   //ФИО специалиста
"pID",   //Название отделения
"qqc"
],
"query": {
    "bool":{
        "filter":[
           {"range":{"pAG": {"gte":####}}},
           {"range":{"pAG": {"lte":****}}}
        ]
    }
}
}`,
		"get186": `{
  "size": 10000,
  "_source": [
     "qqc",   //id обьекта
    "pANdop",   //Состояние_назначения
    "puR",   //Должность_Ресурс
    "pAE"  
  ],
  "query": {
    "bool": {
      "filter": {
        "bool": {
          "should": [
            ####
          ]
        }
      }
    }
  }
}`,
		"get153": `{
  "size": 10000,
  "_source": [
    "pB",
    "qqc",
    "pJ",
    "pI"
  ],
  "query": {
    "bool": {
        "filter": {
            "bool": {
                "should": [
                    ####
                ]
            }
        }
    }
  }
}
`,

		"get83": `{
  "size": 10000,
  "_source": [
    "pu",
    "pu1",
	"Du"
  ],
  "query": {
    "bool": {
        "filter": {
            "bool": {
                "should": [
                    ####
                ]
            }
        }
    }
  }
}`,

		"consistedPatients": `{
"size": 10000, 
"_source": [
"pID",   //Название отделения
"datqq",   //Дата поступления   
"pKDiag"   //Диагноз
],
"query": {
    "bool":{
        "filter":[
           {"range":{"datqq": {"gte":####}}},
           {"range":{"datqq": {"lte":****}}}
        ],
        "must_not": [
           {"range":{"pAG": {"gte":####}}},
           {"range":{"pAG": {"lte":****}}} 
        ]
    }
}
}`,
		"mammography": `{
"size": 10000,
"_source": [
"pAz",   //ФИО специалиста
"u",   //Название услуги
"tn",   //Тип назначения
"qqc",   //qqc 1860
"dE",   //Дата
"Du",   //код ОКМУ
"pID"   //Отделение
],
"query": {
    "bool": {
        "filter": {
            "bool": {
                "filter": {
                    "bool": {
                        "should": [
                            {"term": {"pID.keyword": "КМам"}}
                        ]
                    }
                },
                "must": [
                    {"range":{"dE":{"gte":####}}},
                    {"range":{"dE":{"lte":****}}}
                ],
                "minimum_should_match": 1,
                "should": [
                    {"term": {"Du.keyword": "A06.20.017"}},
                    {"term": {"Du.keyword": "A06.20.018"}},
                    {"term": {"Du.keyword": "A06.20.019"}},
                    {"term": {"Du.keyword": "A06.31.049"}},
                    {"term": {"Du.keyword": "A06.31.050"}},
                    {"term": {"Du.keyword": "A06.31.051"}},
                    {"term": {"Du.keyword": "A06.31.052"}}
                ]
            }
        }
    }
}
}`,
		"CTScan": `{
"size": 10000,
"_source": [
"pAz",   //ФИО специалиста
"u",   //Название услуги
"tn",   //Тип назначения
"qqc",   //qqc 1860
"dE",   //Дата
"Du",   //код ОКМУ
"pID"   //Отделение
],
"query": {
    "bool": {
        "filter": {
            "bool": {
                "filter": {
                    "bool": {
                        "should": [
                            {"term": {"pID.keyword": "ОКиМРТ"}}
                        ]
                    }
                },
                "must": [
                    {"range":{"dE":{"gte":####}}},
                    {"range":{"dE":{"lte":****}}}
                ],
                "minimum_should_match": 1,
                "should": [
                    {"term": {"Du.keyword": "A06.03.107.006"}},
                    {"term": {"Du.keyword": "A06.03.107.007"}},
                    {"term": {"Du.keyword": "A06.03.107.008"}},
                    {"term": {"Du.keyword": "A06.08.014"}},
                    {"term": {"Du.keyword": "A06.08.015"}},
                    {"term": {"Du.keyword": "A06.08.016"}},
                    {"term": {"Du.keyword": "A06.09.018"}},
                    {"term": {"Du.keyword": "A06.10.018"}},
                    {"term": {"Du.keyword": "A06.11.007"}},
                    {"term": {"Du.keyword": "A06.12.098"}},
                    {"term": {"Du.keyword": "A06.22.011"}},
                    {"term": {"Du.keyword": "A06.23.029"}},
                    {"term": {"Du.keyword": "A06.23.030"}},
                    {"term": {"Du.keyword": "A06.26.316.004"}},
                    {"term": {"Du.keyword": "A06.28.024"}},
                    {"term": {"Du.keyword": "A06.31.033.002"}},
                    {"term": {"Du.keyword": "A06.31.042"}},
                    {"term": {"Du.keyword": "A06.31.043"}},
                    {"term": {"Du.keyword": "A06.31.044"}},
                    {"term": {"Du.keyword": "A06.31.045"}},
                    {"term": {"Du.keyword": "A06.31.046"}},
                    {"term": {"Du.keyword": "A06.31.047"}},
                    {"term": {"Du.keyword": "D02.01.001"}}
                ]
            }
        }
    }
}
}`,
		"homeVisit": `{  //выезды на дом
"size": 10000,
"_source": [
"pAz",   //ФИО специалиста
"u",   //Название услуги
"tn",   //Тип назначения
"qqc",   //qqc 1860
"dE",   //Дата
"Du",   //код ОКМУ
"pID"   //Отделение
],
"query": {
    "bool": {
        "filter": {
            "bool": {
                "filter": {
                    "bool": {
                        "should": [
                            {"term": {"pID.keyword": "ОТЭП"}},
                            {"term": {"pID.keyword": "ОТЭ"}}
                        ]
                    }
                },
                "must": [
                    {"range":{"dE":{"gte":####}}},
                    {"range":{"dE":{"lte":****}}}
                ],
                "minimum_should_match": 1,
                "should": [
                    {"term": {"Du.keyword": "B01.058.015"}},
                    {"term": {"Du.keyword": "B01.058.014"}},
                    {"term": {"Du.keyword": "B01.058.013"}},
                    {"term": {"Du.keyword": "A11.12.009.001"}},
                    {"term": {"Du.keyword": "A15.01.002.002"}},
                    {"term": {"Du.keyword": "A04.22.010"}},
                    {"term": {"Du.keyword": "B01.028.08"}},
                    {"term": {"Du.keyword": "B01.015.17"}},
                    {"term": {"Du.keyword": "B01.025.06"}}
                ]
            }
        }
    }
}
}`,
		"MRT": `{
"size": 10000,
"_source": [
"pAz",   //ФИО специалиста
"u",   //Название услуги
"tn",   //Тип назначения
"qqc",   //qqc 1860
"dE",   //Дата
"Du",   //код ОКМУ
"pID"   //Отделение
],
"query": {
    "bool": {
        "filter": {
            "bool": {
                "filter": {
                    "bool": {
                        "should": [
                            {"term": {"pID.keyword": "ОКиМРТ"}}
                        ]
                    }
                },
                "must": [
                    {"range":{"dE":{"gte":####}}},
                    {"range":{"dE":{"lte":****}}}
                ],
                "minimum_should_match": 1,
                "should": [
                    {"term": {"Du.keyword": "A05.01.004"}},
                    {"term": {"Du.keyword": "A05.01.005"}},
                    {"term": {"Du.keyword": "A05.03.003"}},
                    {"term": {"Du.keyword": "A05.03.003.004"}},
                    {"term": {"Du.keyword": "A05.04.001"}},
                    {"term": {"Du.keyword": "A05.04.011"}},
                    {"term": {"Du.keyword": "A05.10.004002"}},
                    {"term": {"Du.keyword": "A05.14.002"}},
                    {"term": {"Du.keyword": "A05.23.004"}},
                    {"term": {"Du.keyword": "A05.23.004.001"}},
                    {"term": {"Du.keyword": "A05.26.008"}},
                    {"term": {"Du.keyword": "A05.31.001"}},
                    {"term": {"Du.keyword": "A05.31.002"}},
                    {"term": {"Du.keyword": "A05.31.006"}},
                    {"term": {"Du.keyword": "A05.31.011"}},
                    {"term": {"Du.keyword": "A05.31.018"}},
                    {"term": {"Du.keyword": "A05.31.021"}},
                    {"term": {"Du.keyword": "A05.31.022"}},
                    {"term": {"Du.keyword": "A05.31.023"}},
                    {"term": {"Du.keyword": "A05.31.024"}},
                    {"term": {"Du.keyword": "A05.31.025"}},
                    {"term": {"Du.keyword": "A05.31.422"}},
                    {"term": {"Du.keyword": "A06.31.048"}},
                    {"term": {"Du.keyword": "D02.01.001"}}
                ]
            }
        }
    }
}
}`,
		"ultrasoundControl": `{
"size": 10000,
"_source": [
"pAz",   //ФИО специалиста
"u",   //Название услуги
"tn",   //Тип назначения
"qqc",   //qqc 1860
"dE",   //Дата
"Du",   //код ОКМУ
"pID"   //Отделение
],
"query": {
    "bool": {
        "filter": {
            "bool": {
                "filter": {
                    "bool": {
                        "should": [
                            {"term": {"pID.keyword": "Пукц"}},
                            {"term": {"pID.keyword": "КДО ЛРО"}}
                        ]
                    }
                },
                "must": [
                    {"range":{"dE":{"gte":####}}},
                    {"range":{"dE":{"lte":****}}}
                ],
                "minimum_should_match": 1,
                "should": [
                    {"term": {"Du.keyword": "A04.22.012"}},
                    {"term": {"Du.keyword": "A04.22.009"}}
                ]
            }
        }
    }
}
}`,
		"ultrasound": `{
"size": 10000,
"_source": [
"pAz",   //ФИО специалиста
"u",   //Название услуги
"tn",   //Тип назначения
"qqc",   //qqc 1860
"dE",   //Дата
"Du",   //код ОКМУ
"pID"   //Отделение
],
"query": {
    "bool": {
        "filter": {
            "bool": {
                "filter": {
                    "bool": {
                        "should": [
                            {"term": {"pID.keyword": "оУЗД"}}
                        ]
                    }
                },
                "must": [
                    {"range":{"dE":{"gte":####}}},
                    {"range":{"dE":{"lte":****}}}
                ],
                "minimum_should_match": 1,
                "should": [
                    {"term": {"Du.keyword": "A04.01.001"}},
                    {"term": {"Du.keyword": "A04.06.001"}},
                    {"term": {"Du.keyword": "A04.06.002"}},
                    {"term": {"Du.keyword": "A04.06.004"}},
                    {"term": {"Du.keyword": "A04.06.005"}},
                    {"term": {"Du.keyword": "A04.10.002"}},
                    {"term": {"Du.keyword": "A04.10.002.010"}},
                    {"term": {"Du.keyword": "A04.10.009"}},
                    {"term": {"Du.keyword": "A04.12.001.002"}},
                    {"term": {"Du.keyword": "A04.12.006.001"}},
                    {"term": {"Du.keyword": "A04.12.011"}},
                    {"term": {"Du.keyword": "A04.12.013"}},
                    {"term": {"Du.keyword": "A04.12.048"}},
                    {"term": {"Du.keyword": "A04.14.001"}},
                    {"term": {"Du.keyword": "A04.14.003"}},
                    {"term": {"Du.keyword": "A04.15.001"}},
                    {"term": {"Du.keyword": "A04.20.001"}},
                    {"term": {"Du.keyword": "A04.20.001.002"}},
                    {"term": {"Du.keyword": "A04.20.001.003"}},
                    {"term": {"Du.keyword": "A04.20.002"}},
                    {"term": {"Du.keyword": "A04.20.003"}},
                    {"term": {"Du.keyword": "A04.20.011"}},
                    {"term": {"Du.keyword": "A04.20.012"}},
                    {"term": {"Du.keyword": "A04.20.013"}},
                    {"term": {"Du.keyword": "A04.21.001"}},
                    {"term": {"Du.keyword": "A04.21.001.002"}},
                    {"term": {"Du.keyword": "A04.21.001.003"}},
                    {"term": {"Du.keyword": "A04.21.002"}},
                    {"term": {"Du.keyword": "A04.21.002.001"}},
                    {"term": {"Du.keyword": "A04.21.003"}},
                    {"term": {"Du.keyword": "A04.21.011"}},
                    {"term": {"Du.keyword": "A04.21.012"}},
                    {"term": {"Du.keyword": "A04.21.013"}},
                    {"term": {"Du.keyword": "A04.21.014"}},
                    {"term": {"Du.keyword": "A04.21.015"}},
                    {"term": {"Du.keyword": "A04.21.016"}},
                    {"term": {"Du.keyword": "A04.22.001"}},
                    {"term": {"Du.keyword": "A04.22.002"}},
                    {"term": {"Du.keyword": "A04.22.004"}},
                    {"term": {"Du.keyword": "A04.22.008"}},
                    {"term": {"Du.keyword": "A04.22.010"}},
                    {"term": {"Du.keyword": "A05.22.011"}},
                    {"term": {"Du.keyword": "A04.24.001"}},
                    {"term": {"Du.keyword": "A04.26.001.001"}},
                    {"term": {"Du.keyword": "A04.26.006"}},
                    {"term": {"Du.keyword": "A04.28.001"}},
                    {"term": {"Du.keyword": "A04.28.002.003"}},
                    {"term": {"Du.keyword": "A04.31.004"}},
                    {"term": {"Du.keyword": "A04.31.005"}},
                    {"term": {"Du.keyword": "A04.31.006"}},
                    {"term": {"Du.keyword": "A04.31.010"}},
                    {"term": {"Du.keyword": "A04.31.012"}},
                    {"term": {"Du.keyword": "A06.26.010.001"}},
                    {"term": {"Du.keyword": "A06.26.010.002"}},
                    {"term": {"Du.keyword": "B03.001.005"}},
                    {"term": {"Du.keyword": "B03.058.34"}},
                    {"term": {"Du.keyword": "B03.058.35"}}
                ]
            }
        }
    }
}
}`,
		"operations": `{
"size": 10000,
"_source": [
"pAz",   //ФИО специалиста
"u",   //Название услуги
"tn",   //Тип назначения
"qqc",   //qqc 1860
"dE",   //Дата
"Du",   //код ОКМУ
"pID"   //Отделение
],
"query": {
    "bool": {
        "filter": {
            "bool": {
                "filter": {
                    "bool": {
                        "should": [
                            {"term": {"pID.keyword": "ОДС"}},
                            {"term": {"pID.keyword": "ЛДО ДРиОфт"}},
                            {"term": {"pID.keyword": "ОпБлКард"}}
                        ]
                    }
                },
                "must": [
                    {"range":{"dE":{"gte":####}}},
                    {"range":{"dE":{"lte":****}}},
                    {"term": {"tn.keyword": "Операция"}}
                ]
            }
        }
    }
}
}`,
		"colonoscopy": `{
"size": 10000,
"_source": [
"pAz",
"u",
"tn",
"qqc",
"dE",
"Du",
"pID"
],
"query": {
    "bool": {
        "filter": {
            "bool": {
                "must": [
                    {"term": {"pID.keyword": "КДо"}},
                    {"range":{"dE":{"gte":####}}},
                    {"range":{"dE":{"lte":****}}}
                ],
                "minimum_should_match": 1,
                "should": [
                    {"term": {"Du.keyword": "A03.18.002"}},
                    {"term": {"Du.keyword": "B03.018.009"}},
                    {"term": {"Du.keyword": "A03.19.003"}},
                    {"term": {"Du.keyword": "A03.19.005"}},
                    {"term": {"Du.keyword": "B03.018.010"}},
                    {"term": {"Du.keyword": "B03.004.011"}}
                ]
            }
        }
    }
}
}`,
		"gastroduodenoskopiya": `{
"size": 10000,
"_source": [
"pAz",
"u",
"tn",
"qqc",
"dE",
"Du",
"pID"
],
"query": {
    "bool": {
        "filter": {
            "bool": {
                "must": [
                    {"match": {"pID": "КДо"}},
                    {"range":{"dE":{"gte":####}}},
                    {"range":{"dE":{"lte":****}}}
                ],
                "minimum_should_match": 1,
                "should": [
                    {"term": {"Du.keyword": "B03.004.009"}},
                    {"term": {"Du.keyword": "A03.16.001"}},
                    {"term": {"Du.keyword": "A03.16.001.01"}},
                    {"term": {"Du.keyword": "B03.004.010"}},
                    {"term": {"Du.keyword": "B03.004.011"}}
                ]
            }
        }
    }
}
}`,
		"xrayDiagnostics": `{
"size": 10000,
"_source": [
"pAz",
"u",
"tn",
"qqc",
"dE",
"Du",
"pID"
],
"query": {
    "bool": {
        "filter": {
            "bool": {
                "filter": {
                    "bool": {
                        "should": [
                            {"term": {"pID.keyword": "ОРд"}},
                            {"term": {"pID.keyword": "оРДиКТ"}}
                        ]
                    }
                },
                "must": [
                    {"range":{"dE":{"gte":####}}},
                    {"range":{"dE":{"lte":****}}}
                ],
                "minimum_should_match": 1,
                "should": [
                    {"term": {"Du.keyword": "A06.03.092"}},
                    {"term": {"Du.keyword": "A06.03.092.001"}},
                    {"term": {"Du.keyword": "A06.03.025"}},
                    {"term": {"Du.keyword": "A06.31.033"}},
                    {"term": {"Du.keyword": "A06.03.006"}},
                    {"term": {"Du.keyword": "A06.04.033"}},
                    {"term": {"Du.keyword": "A06.03.011"}},
                    {"term": {"Du.keyword": "A06.03.071"}},
                    {"term": {"Du.keyword": "A06.03.016"}},
                    {"term": {"Du.keyword": "A06.03.058"}},
                    {"term": {"Du.keyword": "A06.03.052"}},
                    {"term": {"Du.keyword": "A06.28.004"}},
                    {"term": {"Du.keyword": "A06.28.001"}},
                    {"term": {"Du.keyword": "A06.31.011"}},
                    {"term": {"Du.keyword": "A06.16.001.001"}},
                    {"term": {"Du.keyword": "A06.22.005"}},
                    {"term": {"Du.keyword": "A06.08.003"}},
                    {"term": {"Du.keyword": "A06.03.107.005"}}
                ]
            }
        }
    }
}
}`,
		"xrayDensitometry": `{
"size": 10000,
"_source": [
"pAz",   //ФИО специалиста
"u",   //Название услуги
"tn",   //Тип назначения
"qqc",   //qqc 1860
"dE",   //Дата
"Du",   //код ОКМУ
"pID"   //Отделение
],
"query": {
    "bool": {
        "filter": {
            "bool": {
                "filter": {
                    "bool": {
                        "should": [
                            {"term": {"pID.keyword": "ОРд"}},
                            {"term": {"pID.keyword": "оРДиКТ"}}
                        ]
                    }
                },
                "must": [
                    {"range":{"dE":{"gte":####}}},
                    {"range":{"dE":{"lte":****}}}
                ],
                "minimum_should_match": 1,
                "should": [
                    {"term": {"Du.keyword": "A06.03.080"}},
                    {"term": {"Du.keyword": "A06.03.062"}},
                    {"term": {"Du.keyword": "A06.03.063"}},
                    {"term": {"Du.keyword": "A06.03.064"}},
                    {"term": {"Du.keyword": "A06.03.102.002"}},
                    {"term": {"Du.keyword": "A06.03.079.001"}},
                    {"term": {"Du.keyword": "A06.03.083.003"}},
                    {"term": {"Du.keyword": "A06.03.083"}}
                ]
            }
        }
    }
}
}`,
		"ECHOKG": `{
"size": 10000,
"_source": [
"pAz",   //ФИО специалиста
"u",   //Название услуги
"tn",   //Тип назначения
"qqc",   //qqc 1860
"dE",   //Дата
"Du",   //код ОКМУ
"pID"   //Отделение
],
"query": {
    "bool": {
        "filter": {
            "bool": {
                "filter": {
                    "bool": {
                        "should": [
                            {"term": {"pID.keyword": "ОФД"}},
                            {"term": {"pID.keyword": "ЛДО ФункКард"}},
                            {"term": {"pID.keyword": "КДО ЛРО"}},
                            {"term": {"pID.keyword": "ОКЭиСХ"}},
                            {"term": {"pID.keyword": "БИТ"}},
                            {"term": {"pID.keyword": "ПРИТ"}}
                        ]
                    }
                },
                "must": [
                    {"range":{"dE":{"gte":####}}},
                    {"range":{"dE":{"lte":****}}}
                ],
                "minimum_should_match": 1,
                "should": [
                    {"term": {"Du.keyword": "A04.10.002.010"}},
                    {"term": {"Du.keyword": "A04.10.002"}},
                    {"term": {"Du.keyword": "A04.10.009"}}
                ]
            }
        }
    }
}
}`,
		"SMAD": `{
"size": 10000,
"_source": [
"pAz",   //ФИО специалиста
"u",   //Название услуги
"tn",   //Тип назначения
"qqc",   //qqc 1860
"dE",   //Дата
"Du",   //код ОКМУ
"pID"   //Отделение
],
"query": {
    "bool": {
        "filter": {
            "bool": {
                "filter": {
                    "bool": {
                        "should": [
                            {"term": {"pID.keyword": "ОКЭиСХ"}},
                            {"term": {"pID.keyword": "КДО ЛРО"}},
                            {"term": {"pID.keyword": "ОФД"}},
                            {"term": {"pID.keyword": "ОДнефрГД"}}
                        ]
                    }
                },
                "must": [
                    {"range":{"dE":{"gte":####}}},
                    {"range":{"dE":{"lte":****}}},
                    {"term": {"Du.keyword": "A05.10.008"}}
                ]
            }
        }
    }
}
}`,
		"XMECG": `{
"size": 10000,
"_source": [
"pAz",   //ФИО специалиста
"u",   //Название услуги
"tn",   //Тип назначения
"qqc",   //qqc 1860
"dE",   //Дата
"Du",   //код ОКМУ
"pID"   //Отделение
],
"query": {
    "bool": {
        "filter": {
            "bool": {
                "filter": {
                    "bool": {
                        "should": [
                            {"term": {"pID.keyword": "ОКЭиСХ"}},
                            {"term": {"pID.keyword": "КДО ЛРО"}},
                            {"term": {"pID.keyword": "ОФД"}},
                            {"term": {"pID.keyword": "ЛДО ФункКард"}}
                        ]
                    }
                },
                "must": [
                    {"range":{"dE":{"gte":####}}},
                    {"range":{"dE":{"lte":****}}},
                    {"term": {"Du.keyword": "A05.10.004002.002"}}
                ]
            }
        }
    }
}
}`,
		"ECG": `{
"size": 10000,
"_source": [
"pAz",   //ФИО специалиста
"u",   //Название услуги
"tn",   //Тип назначения
"qqc",   //qqc 1860
"dE",   //Дата
"Du",   //код ОКМУ
"pID"   //Отделение
],
"query": {
    "bool": {
        "filter": {
            "bool": {
                "filter": {
                    "bool": {
                        "should": [
                            {"term": {"pID.keyword": "ОФД"}},
                            {"term": {"pID.keyword": "ЛДО ФункКард"}},
                            {"term": {"pID.keyword": "КДО ЛРО"}},
                            {"term": {"pID.keyword": "ОКЭиСХ"}},
                            {"term": {"pID.keyword": "БИТ"}},
                            {"term": {"pID.keyword": "ПРИТ"}},
                            {"term": {"pID.keyword": "КДО"}},
                            {"term": {"pID.keyword": "ОХЛНРС"}}
                        ]
                    }
                },
                "must": [
                    {"range":{"dE":{"gte":####}}},
                    {"range":{"dE":{"lte":****}}}
                ],
                "minimum_should_match": 1,
                "should": [
                    {"term": {"Du.keyword": "A05.10.006"}},
                    {"term": {"Du.keyword": "A12.10.011"}}
                ]
            }
        }
    }
}
}`,
		"radiopaqueProcedures": `{
"size": 10000,
"_source": [
"pAz",   //ФИО специалиста
"u",   //Название услуги
"tn",   //Тип назначения
"qqc",   //qqc 1860
"dE",   //Дата
"Du",   //код ОКМУ
"pID"   //Отделение
],
"query": {
    "bool": {
        "filter": {
            "bool": {
                "filter": {
                    "bool": {
                        "should": [
                            {"term": {"pID.keyword": "ЛДОАиУ"}}
                        ]
                    }
                },
                "must": [
                    {"range":{"dE":{"gte":####}}},
                    {"range":{"dE":{"lte":****}}},
                    {"term": {"Du.keyword": "A03.28.999"}}
                ]
            }
        }
    }
}
}`,
		"ejaculateResearch": `{
"size": 10000,
"_source": [
"pAz",   //ФИО специалиста
"u",   //Название услуги
"tn",   //Тип назначения
"qqc",   //qqc 1860
"dE",   //Дата
"Du",   //код ОКМУ
"pID"   //Отделение
],
"query": {
    "bool": {
        "filter": {
            "bool": {
                "filter": {
                    "bool": {
                        "should": [
                            {"term": {"pID.keyword": "ЛДОВРТ"}}
                        ]
                    }
                },
                "must": [
                    {"range":{"dE":{"gte":####}}},
                    {"range":{"dE":{"lte":****}}}
                ],
                "minimum_should_match": 1,
                "should": [
                    {"term": {"Du.keyword": "A08.21.015"}},
                    {"term": {"Du.keyword": "B03.053.001.010"}},
                    {"term": {"Du.keyword": "A09.21.001.001"}},
                    {"term": {"Du.keyword": "A09.20.016"}},
                    {"term": {"Du.keyword": "A09.21.003.001"}},
                    {"term": {"Du.keyword": "A09.21.017"}},
                    {"term": {"Du.keyword": "A09.21.021"}},
                    {"term": {"Du.keyword": "A09.21.001.002"}}
                ]
            }
        }
    }
}
}`,
		"thyroidScintigraphy": `{
"size": 10000,
"_source": [
"pAz",   //ФИО специалиста
"u",   //Название услуги
"tn",   //Тип назначения
"qqc",   //qqc 1860
"dE",   //Дата
"Du",   //код ОКМУ
"pID"   //Отделение
],
"query": {
    "bool": {
        "filter": {
            "bool": {
                "filter": {
                    "bool": {
                        "should": [
                            {"term": {"pID.keyword": "ОРНДи"}},
                            {"term": {"pID.keyword": "ОРНд"}}
                        ]
                    }
                },
                "must": [
                    {"range":{"dE":{"gte":####}}},
                    {"range":{"dE":{"lte":****}}}
                ],
                "minimum_should_match": 1,
                "should": [
                    {"term": {"Du.keyword": "A07.22.002.001"}},
                    {"term": {"Du.keyword": "A07.22.002.002"}},
                    {"term": {"Du.keyword": "A07.22.002.003"}}
                ]
            }
        }
    }
}
}`,
		"thyroidScintigraphyTc99technetril": `{
"size": 10000,
"_source": [
"pAz",   //ФИО специалиста
"u",   //Название услуги
"tn",   //Тип назначения
"qqc",   //qqc 1860
"dE",   //Дата
"Du",   //код ОКМУ
"pID"   //Отделение
],
"query": {
    "bool": {
        "filter": {
            "bool": {
                "filter": {
                    "bool": {
                        "should": [
                            {"term": {"pID.keyword": "ОРНд"}}
                        ]
                    }
                },
                "must": [
                    {"range":{"dE":{"gte":####}}},
                    {"range":{"dE":{"lte":****}}},
                    {"term": {"Du.keyword": "A07.22.023"}}
                ]
            }
        }
    }
}
}`,
		"PTGscintigraphy": `{
"size": 10000,
"_source": [
"pAz",   //ФИО специалиста
"u",   //Название услуги
"tn",   //Тип назначения
"qqc",   //qqc 1860
"dE",   //Дата
"Du",   //код ОКМУ
"pID"   //Отделение
],
"query": {
    "bool": {
        "filter": {
            "bool": {
                "filter": {
                    "bool": {
                        "should": [
                            {"term": {"pID.keyword": "ОРНд"}},
                            {"term": {"pID.keyword": "ОРНДи"}}
                        ]
                    }
                },
                "must": [
                    {"range":{"dE":{"gte":####}}},
                    {"range":{"dE":{"lte":****}}},
                    {"term": {"Du.keyword": "A07.22.021"}}
                ]
            }
        }
    }
}
}`,
		"neuroendocrineTumors": `{
"size": 10000,
"_source": [
"pAz",   //ФИО специалиста
"u",   //Название услуги
"tn",   //Тип назначения
"qqc",   //qqc 1860
"dE",   //Дата
"Du",   //код ОКМУ
"pID"   //Отделение
],
"query": {
    "bool": {
        "filter": {
            "bool": {
                "filter": {
                    "bool": {
                        "should": [
                            {"term": {"pID.keyword": "ОРНДи"}},
                            {"term": {"pID.keyword": "ОРНд"}}
                        ]
                    }
                },
                "must": [
                    {"range":{"dE":{"gte":####}}},
                    {"range":{"dE":{"lte":****}}}
                ],
                "minimum_should_match": 1,
                "should": [
                    {"term": {"Du.keyword": "A07.22.002.010"}},
                    {"term": {"Du.keyword": "A07.22.002.011"}}
                ]
            }
        }
    }
}
}
`,
		"osteoscintigraphy": `{
"size": 10000,
"_source": [
"pAz",   //ФИО специалиста
"u",   //Название услуги
"tn",   //Тип назначения
"qqc",   //qqc 1860
"dE",   //Дата
"Du",   //код ОКМУ
"pID"   //Отделение
],
"query": {
    "bool": {
        "filter": {
            "bool": {
                "filter": {
                    "bool": {
                        "should": [
                            {"term": {"pID.keyword": "ОРНд"}},
                            {"term": {"pID.keyword": "ОРНДи"}}
                        ]
                    }
                },
                "must": [
                    {"range":{"dE":{"gte":####}}},
                    {"range":{"dE":{"lte":****}}},
                    {"term": {"Du.keyword": "A07.03.002"}}
                ]
            }
        }
    }
}
}`,
		"nephroscintigraphy": `{
"size": 10000,
"_source": [
"pAz",   //ФИО специалиста
"u",   //Название услуги
"tn",   //Тип назначения
"qqc",   //qqc 1860
"dE",   //Дата
"Du",   //код ОКМУ
"pID"   //Отделение
],
"query": {
    "bool": {
        "filter": {
            "bool": {
                "filter": {
                    "bool": {
                        "should": [
                            {"term": {"pID.keyword": "ОРНд"}},
                            {"term": {"pID.keyword": "ОРНДи"}}
                        ]
                    }
                },
                "must": [
                    {"range":{"dE":{"gte":####}}},
                    {"range":{"dE":{"lte":****}}},
                    {"term": {"Du.keyword": "A07.28.004.001"}}
                ]
            }
        }
    }
}
}`,
		"myocardialScintigraphy": `{
"size": 10000,
"_source": [
"pAz",   //ФИО специалиста
"u",   //Название услуги
"tn",   //Тип назначения
"qqc",   //qqc 1860
"dE",   //Дата
"Du",   //код ОКМУ
"pID"   //Отделение
],
"query": {
    "bool": {
        "filter": {
            "bool": {
                "filter": {
                    "bool": {
                        "should": [
                            {"term": {"pID.keyword": "ОРНд"}}
                        ]
                    }
                },
                "must": [
                    {"range":{"dE":{"gte":####}}},
                    {"range":{"dE":{"lte":****}}},
                    {"term": {"Du.keyword": "A07.10.001.001"}}
                ]
            }
        }
    }
}
}`,
		"myocardialScintigraphyWithExercise": `{
"size": 10000,
"_source": [
"pAz",   //ФИО специалиста
"u",   //Название услуги
"tn",   //Тип назначения
"qqc",   //qqc 1860
"dE",   //Дата
"Du",   //код ОКМУ
"pID"   //Отделение
],
"query": {
    "bool": {
        "filter": {
            "bool": {
                "filter": {
                    "bool": {
                        "should": [
                            {"term": {"pID.keyword": "ОРНд"}}
                        ]
                    }
                },
                "must": [
                    {"range":{"dE":{"gte":####}}},
                    {"range":{"dE":{"lte":****}}},
                    {"term": {"Du.keyword": "A07.10.012.002"}}
                ]
            }
        }
    }
}
}`,
		"scintirationIodine-131wholebody": `{
"size": 10000,
"_source": [
"pAz",   //ФИО специалиста
"u",   //Название услуги
"tn",   //Тип назначения
"qqc",   //qqc 1860
"dE",   //Дата
"Du",   //код ОКМУ
"pID"   //Отделение
],
"query": {
    "bool": {
        "filter": {
            "bool": {
                "filter": {
                    "bool": {
                        "should": [
                            {"term": {"pID.keyword": "ОРНд"}},
                            {"term": {"pID.keyword": "ОРНДи"}}
                        ]
                    }
                },
                "must": [
                    {"range":{"dE":{"gte":####}}},
                    {"range":{"dE":{"lte":****}}},
                    {"term": {"Du.keyword": "A07.22.002.009"}}
                ]
            }
        }
    }
}
}`,
		"scintirationIodine-123wholebody": `{
"size": 10000,
"_source": [
"pAz",   //ФИО специалиста
"u",   //Название услуги
"tn",   //Тип назначения
"qqc",   //qqc 1860
"dE",   //Дата
"Du",   //код ОКМУ
"pID"   //Отделение
],
"query": {
    "bool": {
        "filter": {
            "bool": {
                "filter": {
                    "bool": {
                        "should": [
                            {"term": {"pID.keyword": "ОРНд"}},
                            {"term": {"pID.keyword": "ОРНДи"}}
                        ]
                    }
                },
                "must": [
                    {"range":{"dE":{"gte":####}}},
                    {"range":{"dE":{"lte":****}}},
                    {"term": {"Du.keyword": "A07.22.002.008"}}
                ]
            }
        }
    }
}
}`,
		"neurophoniatricRehabilitation": `{
"size": 10000,
"_source": [
"pAz",   //ФИО специалиста
"u",   //Название услуги
"tn",   //Тип назначения
"qqc",   //qqc 1860
"dE",   //Дата
"Du",   //код ОКМУ
"pID"   //Отделение
],
"query": {
    "bool": {
        "filter": {
            "bool": {
                "filter": {
                    "bool": {
                        "should": [
                            {"term": {"pID.keyword": ""}}
                        ]
                    }
                },
                "must": [
                    {"range":{"dE":{"gte":####}}},
                    {"range":{"dE":{"lte":****}}},
                    {"term": {"Du.keyword": "A07.08.003"}}
                ]
            }
        }
    }
}
}
`,
		"postTherapeuticFullBodyScintigraphy": `{
"size": 10000,
"_source": [
"pAz",   //ФИО специалиста
"u",   //Название услуги
"tn",   //Тип назначения
"qqc",   //qqc 1860
"dE",   //Дата
"Du",   //код ОКМУ
"pID"   //Отделение
],
"query": {
    "bool": {
        "filter": {
            "bool": {
                "filter": {
                    "bool": {
                        "should": [
                            {"term": {"pID.keyword": "ОРНд"}},
                            {"term": {"pID.keyword": "ОРНДи"}}
                        ]
                    }
                },
                "must": [
                    {"range":{"dE":{"gte":####}}},
                    {"range":{"dE":{"lte":****}}},
                    {"term": {"Du.keyword": "A07.31.063"}}
                ]
            }
        }
    }
}
}`,
		"radionuclideTherapyThyroidCancer": `{
"size": 10000,
"_source": [
"pAz",   //ФИО специалиста
"u",   //Название услуги
"tn",   //Тип назначения
"qqc",   //qqc 1860
"dE",   //Дата
"Du",   //код ОКМУ
"pID"   //Отделение
],
"query": {
    "bool": {
        "filter": {
            "bool": {
                "filter": {
                    "bool": {
                        "should": [
                            {"term": {"pID.keyword": ""}}
                        ]
                    }
                },
                "must": [
                    {"range":{"dE":{"gte":####}}},
                    {"range":{"dE":{"lte":****}}}
                ],
                "minimum_should_match": 1,
                "should": [
                    {"term": {"Du.keyword": "A07.31.011.008"}},
                    {"term": {"Du.keyword": "A07.31.011.009"}}
                ]
            }
        }
    }
}
}`,
		"radiometricExaminationPresenceNasalLiquorrhea": `{
"size": 10000,
"_source": [
"pAz",
"u",
"tn",
"qqc",
"dE",
"Du",
"pID"
],
"query": {
"bool": {
"filter": {
"bool": {
"must": [
{"term": {"Du.keyword": "A07.31.015.009"}}
]
}
}
}
}
}`,
		"lungScintigraphy": `{
"size": 10000,
"_source": [
"pAz",   //ФИО специалиста
"u",   //Название услуги
"tn",   //Тип назначения
"qqc",   //qqc 1860
"dE",   //Дата
"Du",   //код ОКМУ
"pID"   //Отделение
],
"query": {
    "bool": {
        "filter": {
            "bool": {
                "filter": {
                    "bool": {
                        "should": [
                            {"term": {"pID.keyword": ""}}
                        ]
                    }
                },
                "must": [
                    {"range":{"dE":{"gte":####}}},
                    {"range":{"dE":{"lte":****}}},
                    {"term": {"Du.keyword": "A07.09.003.001"}}
                ]
            }
        }
    }
}
}`,

		"radioiodineTherapyThyrotoxicosis": `{
"size": 10000,
"_source": [
"pAz",   //ФИО специалиста
"u",   //Название услуги
"tn",   //Тип назначения
"qqc",   //qqc 1860
"dE",   //Дата
"Du",   //код ОКМУ
"pID"   //Отделение
],
"query": {
    "bool": {
        "filter": {
            "bool": {
                "filter": {
                    "bool": {
                        "should": [
                            {"term": {"pID.keyword": ""}}
                        ]
                    }
                },
                "must": [
                    {"range":{"dE":{"gte":####}}},
                    {"range":{"dE":{"lte":****}}},
                    {"term": {"Du.keyword": "A07.31.010.001"}}
                ]
            }
        }
    }
}
}`,
		"ambulatoryRadioiodineTherapyToxicGoiterFunctionalAutonomy": `{
"size": 10000,
"_source": [
"pAz",   //ФИО специалиста
"u",   //Название услуги
"tn",   //Тип назначения
"qqc",   //qqc 1860
"dE",   //Дата
"Du",   //код ОКМУ
"pID"   //Отделение
],
"query": {
    "bool": {
        "filter": {
            "bool": {
                "filter": {
                    "bool": {
                        "should": [
                            {"term": {"pID.keyword": ""}}
                        ]
                    }
                },
                "must": [
                    {"range":{"dE":{"gte":####}}},
                    {"range":{"dE":{"lte":****}}},
                    {"term": {"Du.keyword": "A07.31.062"}}
                ]
            }
        }
    }
}
}`,
	}

	return requestMap[funcName]
}