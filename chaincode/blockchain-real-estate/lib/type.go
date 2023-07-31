/**
 * @Author: 夜央 Oh oh oh oh oh oh (https://github.com/togettoyou1)
 * @Email: zoujh99@qq.com
 * @Date: 2020/3/4 2:00 下午
 * @Description: 定义的数据结构体、常量
 */
package lib

//账户，虚拟管理员和若干业主账号
type Account struct {
	AccountId string `json:"accountId"` //账号ID
	UserName  string `json:"userName"`  //账号名
	Identity  string `json:"identity"`  //身份
	PassWd    string `json:"passwd"`    //密码

}

//是否允许登录
type Access struct {
	Access bool `json:access`
}

//药品信息
type Drug struct {
	DrugID             string  `json:"drugId"` //药品采购ID
	YPID               string  `json:"ypid"`
	Downer             string  `json:"downer"`                 //数据提供者(用户AccountId)
	DName              string  `json:"dname"`                  //产品通用名
	DtradeName         string  `json:"dtradename"`             //商品名
	DosageForm         string  `json:"dosageform"`             //剂型
	DosageClass        string  `json:"dosageclass"`            //剂型分类名称
	DosageCode         string  `json:"dosagecode"`             //剂型分类码
	Specifications     string  `json:"specifications"`         //规格
	SpecificationsCode string  `json:"specificationscode"`     //规格分类码
	PharmaClass        string  `json:"pharmaclass`             //药理、药效分类
	PharmaCode         string  `json:"pharmacode`              //药理、药效分类编码
	ConvertRate        int64   `json:"convertrate`             //转换比
	PackageMaterial    string  `json:"packagematerial"`        // 包装材质
	MiniPackageUnit    string  `json:"minipackageunit`         //最小包装单位
	Conpany            string  `json:"conpany"`                //生产企业
	USCC               string  `json:"uscc"`                   //统一社会信用代码
	City               string  `json:"city"`                   //地市
	ConpanyCode        string  `json:"conpanycode"`            //企业编码
	ApprovalNum        string  `json:"approvalnum"`            //批准文号
	Hospital           string  `json:"hospital"`               //医疗机构
	Rank               string  `json:"rank"`                   //医院等级
	Number             int64   `json:"number"`                 //采购数量
	Amount             float64 `json:"amount"`                 //采购金额
	InsureClass        string  `json:"insureclass"`            //医保甲乙类
	BasicDrug          string  `json:"basicdrug"`              //基药
	Year               int64   `json:"year"`                   //采购年份
	DrugHash           uint64  `json:"unique_index:hash_idx;"` //该药品采购的唯一hash值
}

//定义一个医用耗材详细信息
type Material struct {
	MaterialID       string  `json:"materialid"`             //采购耗材编码
	Mcode            string  `json:"mcode"`                  //耗材编码
	Mowner           string  `json:"mowner"`                 //数据提供者(用户AccountId)
	MName            string  `json:"mname"`                  //耗材名称
	Mclass           string  `json:"mclass"`                 //产品类别
	CatalogNum       string  `json:"catalognum"`             //目录序号
	CountryCode      string  `json:"countrycode"`            //国家对应耗材编码
	CatalogName      string  `json:"catalogname"`            //目录名称
	MSpecifications  string  `json:"mspecifications"`        //产品规格
	Model            string  `json:"model"`                  //型号
	Unit             string  `json:"unit`                    //单位
	BiddingCompany   string  `json:"biddingcompany`          //投标企业
	ProduceCompany   string  `json:"producecompany`          //生产企业
	DomesticImports  string  `json:"domesticimports"`        //国产进口
	RegistrantNumber string  `json:"registrantnumber"`       //注册证编号
	RegistrantName   string  `json:"registrantname"`         //注册证名称
	MNumber          int     `json:"mnumber"`                //数量
	MAmount          float64 `json:"mamount"`                //金额
	MYear            int     `json:"myear"`                  //年份
	MaterialHash     uint64  `json:"unique_index:hash_idx;"` //该耗材的唯一hash值
}

//定义一个不良反应详细信息
type Adverse struct {
	AdverseID         string `json:"adverseid"`              //不良反应id
	ACode             string `json:"acode"`                  //报告表编码
	Aowner            string `json:"aowner"`                 //数据提供者（用户ID）
	FirstRecord       string `json:"firstrecord"`            //首次\跟踪报告
	RecordTypeNew     string `json:"recordtypenew"`          //报告类型-新的
	RecordTypeLevel   string `json:"recordtypelevel"`        //报告类型-严重程度
	AdverseLevel      string `json:"adverselevel"`           //严重药品不良反应
	RecordUnit        string `json:"recordunit"`             //报告单位类别
	PName             string `json:"pname"`                  //患者姓名
	PSex              string `json:"psex"`                   //性别
	PBirth            string `json:"pbirth`                  //出生日期
	PAge              string `json:"page`                    //年龄
	PAgeUnit          int    `json:"pageunit`                //年龄单位
	PNation           string `json:"pnation"`                // 民族
	PWeight           string `json:"pweight"`                //体重
	Pphone            string `json:"pphone"`                 //患者联系方式
	PDisease          string `json:"pdisease"`               //原患疾病
	HName             string `json:"hname"`                  //医院名称
	HRecordCode       string `json:"hrecordcode"`            //病历号
	PreAdverse        string `json:"preadverse"`             //既往药品不良反应事件
	FamilyAdverse     string `json:"familyadverse"`          //家族药品不良反应事件
	RelatedInfo       string `json:"relatedinfo"`            //相关重要信息
	DoubtUse          string `json:"doubtuse"`               //怀疑\并用
	DrugOrder         string `json:"drugorder"`              //药品序号
	ApprovalNum       string `json:"approvalnum"`            //批准文号
	DtradeName        string `json:"dtradename"`             //药物商品名称
	DName             string `json:"dname"`                  //药物通用名称
	DosageForm        string `json:"dosageform"`             //剂型
	Conpany           string `json:"conpany"`                //生产企业
	ProductBatchnum   string `json:"productbatchnum"`        //生产批号
	Usage             string `json:"usage"`                  //用量
	UsageUnit         string `json:"usageunit"`              //用量单位
	UsageDay          string `json:"usageday"`               //用法-天
	UsageTimes        string `json:"usagetimes"`             //用法-次
	UseRoute          string `json:"useroute"`               //给药途径
	UseStartTime      string `json:"usestarttime"`           //用药开始时间
	UseEndTime        int    `json:"useendtime"`             //用药结束时间
	UseReason         string `json:"usereason"`              // 用药原因
	AdeverseName      string `json:"adeversename"`           //不良反应名称
	AdverseStartTime  string `json:"adversestarttime"`       //不良反应开始时间
	AdeverseDescribe  string `json:"adeversedescribe"`       //不良反应过程描述
	AdverseResult     string `json:"adverseresult"`          //不良反应结果
	DeadTime          string `json:"deadtime"`               //死亡时间
	DirectCause       string `json:"directcause"`            //直接原因
	StopUseAffect     string `json:"stopuseaffect"`          //停药\减药后症状是否减轻
	AgainUseAffect    string `json:"againuseaffect"`         //再次使用可疑药物是否出现同样反应
	ImpactToDisease   string `json:"impacttodisease"`        //对原患疾病影响
	RecorderJudge     string `json:"recorderjudge"`          //报告人评价
	RecorderJudgeSign string `json:"recorderjudgesign"`      //报告人评价签名
	RecordUnitJudge   string `json:"recordunitjudge"`        //报告单位评价
	RecordUnitSign    string `json:"recordunitsign"`         //报告单位签名
	RecorderPhone     string `json:"recorderphone"`          //报告人联系方式
	RecorderCareer    string `json:"recordercareer"`         //报告人职业
	RecorderMail      string `json:"recordermail"`           //报告人邮箱
	RecorderSign      string `json:"recordersign"`           //报告人签名
	RecordUnitName    string `json:"recordunitname"`         //报告单位名称
	RecordUnitPerson  string `json:"recordunitperson"`       //报告单位联系人
	RecordUnitPhone   string `json:"recordunitphone"`        //报告单位电话
	RecordTime        string `json:"recordtime"`             //报告日期
	RecordSource      string `json:"recordsource"`           //报告单位信息来源
	Notes             string `json:"notes"`                  //备注
	CountryRecTime    string `json:"countryrectime"`         //国家中心接收
	Entered           string `json:"entered"`                //录入人
	JudgeState        string `json:"judgestate"`             //评价状态
	City              string `json:"city"`                   //报告地区名称
	EndTime           string `json:"endtime"`                //最后修改时间
	XJudge            string `json:"xjudge"`                 //县评价
	SJudge            string `json:"sjudge"`                 //市评价
	ShJudge           string `json:"shjudge"`                //省评价
	GJudge            string `json:"gjudge"`                 //国家评价
	FalseReport       string `json:"falsereport"`            //假报告
	RepeatReport      string `json:"repeatreport"`           //重复报告
	DelFlag           string `json:"delflag"`                //删除标识
	XJudgerName       string `json:"xjudgename"`             //县评价人姓名
	SJudgerName       string `json:"sjudgername"`            //市评价人姓名
	ShJudgerName      string `json:"shjudgername"`           //省评价人姓名
	GJudgerName       string `json:"gjudgername"`            //国家评价人姓名
	XJudgerTime       string `json:"xjudgertime"`            //县评价时间
	SJudgerTime       string `json:"sjudgertime"`            //市评价时间
	ShJudgerTime      string `json:"shjudgertime"`           //省评价时间
	GJudgerTime       string `json:"gjudgertime"`            //国家评价时间
	AdverseHash       uint64 `json:"unique_index:hash_idx;"` //该不良反应的唯一hash值

}

const (
	AccountKey  = "account-key"
	DrugKey     = "drug-key"
	MaterialKey = "material-key"
	AdverseKey  = "adverse-key"
)
