webpackJsonp([6],{CUl6:function(e,r,t){"use strict";Object.defineProperty(r,"__esModule",{value:!0});var a=t("Dd8w"),o=t.n(a),u=t("NYxO"),n=t("nv77"),l=t("phQZ"),i={name:"AddDrugRequest",data:function(){return{ruleForm:{proprietor:"",drugName:"",tradeName:"",drugNumber:0},accountList:[],rules:{proprietor:[{required:!0,message:"请选择用户",trigger:"change"}],drugName:[{required:!0,message:"请输入药品名",trigger:"change"}],tradeName:[{required:!0,message:"请输入商品名",trigger:"change"}],drugNumber:[{validator:function(e,r,t){r<=0?t(new Error("必须大于0")):t()},trigger:"blur"}]},loading:!1}},computed:o.a({},u.b(["accountId","userName"])),created:function(){var e=this;n.c().then(function(r){null!==r&&(e.accountList=r.filter(function(e){return"管理员"!==e.userName}))})},methods:{submitForm:function(e){var r=this;this.$refs[e].validate(function(e){if(!e)return!1;r.$confirm("是否立即新增?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"success"}).then(function(){r.loading=!0,l.a({accountId:r.accountId,proprietor:r.ruleForm.proprietor,drugName:r.ruleForm.drugName,tradeName:r.ruleForm.tradeName,drugNumber:r.ruleForm.drugNumber}).then(function(e){r.loading=!1,null!==e?r.$message({type:"success",message:"新增成功!"}):r.$message({type:"error",message:"新增失败!"})}).catch(function(e){r.loading=!1})}).catch(function(){r.loading=!1,r.$message({type:"info",message:"已取消新增"})})})},resetForm:function(e){this.$refs[e].resetFields()},selectGet:function(e){this.ruleForm.proprietor=e}}},s={render:function(){var e=this,r=e.$createElement,t=e._self._c||r;return t("div",{staticClass:"app-container"},[t("el-form",{directives:[{name:"loading",rawName:"v-loading",value:e.loading,expression:"loading"}],ref:"ruleForm",attrs:{model:e.ruleForm,rules:e.rules,"label-width":"100px"}},[t("el-form-item",{attrs:{label:"用户",prop:"proprietor"}},[t("el-select",{attrs:{placeholder:"请选择用户"},on:{change:e.selectGet},model:{value:e.ruleForm.proprietor,callback:function(r){e.$set(e.ruleForm,"proprietor",r)},expression:"ruleForm.proprietor"}},[t("el-option",{key:e.accountId,attrs:{label:e.userName,value:e.accountId}},[t("span",{staticStyle:{float:"left",color:"#8492a6","font-size":"13px"}},[e._v(e._s(e.userName))]),e._v(" "),t("span",{staticStyle:{float:"right",color:"#8492a6","font-size":"13px"}},[e._v(e._s(e.accountId))])])],1)],1),e._v(" "),t("el-form-item",{attrs:{label:"药品名",prop:"drugName"}},[t("el-input",{staticStyle:{width:"200px"},attrs:{placeholder:"请输入药品名"},model:{value:e.ruleForm.drugName,callback:function(r){e.$set(e.ruleForm,"drugName",r)},expression:"ruleForm.drugName"}})],1),e._v(" "),t("el-form-item",{attrs:{label:"商品名",prop:"tradeName"}},[t("el-input",{staticStyle:{width:"200px"},attrs:{placeholder:"请输入商品名"},model:{value:e.ruleForm.tradeName,callback:function(r){e.$set(e.ruleForm,"tradeName",r)},expression:"ruleForm.tradeName"}})],1),e._v(" "),t("el-form-item",{attrs:{label:"数量",prop:"drugNumber"}},[t("el-input-number",{attrs:{precision:2,step:1,min:0},model:{value:e.ruleForm.drugNumber,callback:function(r){e.$set(e.ruleForm,"drugNumber",r)},expression:"ruleForm.drugNumber"}})],1),e._v(" "),t("el-form-item",[t("el-button",{attrs:{type:"primary"},on:{click:function(r){return e.submitForm("ruleForm")}}},[e._v("立即发布")]),e._v(" "),t("el-button",{on:{click:function(r){return e.resetForm("ruleForm")}}},[e._v("重置")])],1)],1)],1)},staticRenderFns:[]};var c=t("VU/8")(i,s,!1,function(e){t("PkTh")},"data-v-dc32cf9a",null);r.default=c.exports},PkTh:function(e,r){},phQZ:function(e,r,t){"use strict";r.a=function(e){return a.a({url:"/createDrugRequest",method:"post",data:e})},r.b=function(e){return a.a({url:"/queryDrugRequestList",method:"post",data:e})};var a=t("vLgD")}});
//# sourceMappingURL=6.9773c408dd2244861d96.js.map