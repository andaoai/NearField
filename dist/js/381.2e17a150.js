"use strict";(self["webpackChunkweb_system"]=self["webpackChunkweb_system"]||[]).push([[381],{7381:function(e,t,s){s.r(t),s.d(t,{default:function(){return u}});var a=function(){var e=this,t=e._self._c;return t("div",[t("el-container",[t("el-header",[t("el-input",{attrs:{placeholder:"搜索 用户账号",clearable:""},model:{value:e.inputUserID,callback:function(t){e.inputUserID=t},expression:"inputUserID"}}),t("el-input",{attrs:{placeholder:"搜索 用户名",clearable:""},model:{value:e.inputUserName,callback:function(t){e.inputUserName=t},expression:"inputUserName"}}),t("el-button",{attrs:{type:"success"},on:{click:function(t){e.centerDialogVisible=!0}}},[e._v("新 增")])],1),t("el-main",[t("el-table",{ref:"multipleTable",staticStyle:{width:"100%"},attrs:{data:this.filterUsername(),"tooltip-effect":"dark"},on:{"selection-change":e.handleSelectionChange}},[t("el-table-column",{attrs:{type:"selection",width:"55"}}),t("el-table-column",{attrs:{label:"用户创建时间",align:"center",width:"200"},scopedSlots:e._u([{key:"default",fn:function(s){return[t("i",{staticClass:"el-icon-time"}),t("span",{staticStyle:{"margin-left":"10px"}},[e._v(e._s(s.row.CreatedAt.slice(0,19)))])]}}])}),t("el-table-column",{attrs:{label:"用户账号",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(t.row.ID))]}}])}),t("el-table-column",{attrs:{label:"用户名",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(t.row.Name))]}}])}),t("el-table-column",{attrs:{label:"用户角色",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(t.row.Role))]}}])}),t("el-table-column",{attrs:{label:"操作",align:"center",width:"400"},scopedSlots:e._u([{key:"default",fn:function(s){return[t("el-switch",{attrs:{"active-color":"#13ce66","inactive-color":"#ff4949","active-text":s.row.Available?"启用":"禁用"},on:{change:function(t){return e.turnOnOff(s.$index,s.row)}},model:{value:s.row.Available,callback:function(t){e.$set(s.row,"Available",t)},expression:"scope.row.Available"}}),t("el-button",{attrs:{size:"mini",type:"primary",plain:""},on:{click:function(t){return e.handleResetPassward(s.$index,s.row)}}},[e._v("重置密码")]),t("el-button",{attrs:{size:"mini",type:"danger",plain:""},on:{click:function(t){return e.handleDeleteUser(s.$index,s.row)}}},[e._v("删除")])]}}])})],1)],1)],1),t("el-dialog",{attrs:{title:"添加新用户",visible:e.centerDialogVisible,width:"35%",center:""},on:{"update:visible":function(t){e.centerDialogVisible=t}}},[t("el-form",{attrs:{"label-position":e.labelPosition,"label-width":"100px",model:e.formLabelAlign}},[t("el-form-item",{attrs:{label:"用户名称"}},[t("el-input",{model:{value:e.formLabelAlign.userName,callback:function(t){e.$set(e.formLabelAlign,"userName",t)},expression:"formLabelAlign.userName"}})],1),t("el-form-item",{directives:[{name:"show",rawName:"v-show",value:e.userDefinedPassword,expression:"userDefinedPassword"}],attrs:{label:"用户密码"}},[t("el-input",{attrs:{type:"password"},model:{value:e.formLabelAlign.userPassword,callback:function(t){e.$set(e.formLabelAlign,"userPassword",t)},expression:"formLabelAlign.userPassword"}})],1),t("el-form-item",{directives:[{name:"show",rawName:"v-show",value:e.userDefinedPassword,expression:"userDefinedPassword"}],attrs:{label:"再输入密码"}},[t("el-input",{attrs:{type:"password"},model:{value:e.formLabelAlign.userPasswordAgain,callback:function(t){e.$set(e.formLabelAlign,"userPasswordAgain",t)},expression:"formLabelAlign.userPasswordAgain"}})],1),t("el-form-item",[t("el-switch",{attrs:{"active-text":"自定义密码","inactive-text":"默认密码"},model:{value:e.userDefinedPassword,callback:function(t){e.userDefinedPassword=t},expression:"userDefinedPassword"}})],1),t("el-form-item",[t("el-button",{on:{click:function(t){e.centerDialogVisible=!1}}},[e._v("取 消")]),t("el-button",{attrs:{type:"primary"},on:{click:function(t){return e.handleAddUser()}}},[e._v("确 定")])],1)],1)],1)],1)},l=[],n=s(586),i={data(){return{userDefinedPassword:!1,centerDialogVisible:!1,inputUserID:"",inputUserName:"",labelPosition:"right",formLabelAlign:{userName:"",userPassword:"123456",userPasswordAgain:"123456"},tableData:[]}},mounted(){this.init()},methods:{init(){(0,n.zL)().then((e=>{this.tableData=e.data}))},handleResetPassward(e,t){console.log("handleResetPassward"),console.log(e),console.log(t.userID),this.$confirm("此操作将该用户设置为默认密码123456, 是否继续?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning",center:!0}).then((()=>{(0,n.ES)({name:this.$store.state.userName,password:"123456",id:t.ID}).then((()=>{console.log("设置为默认密码:123456成功!"),this.$message({type:"success",message:"设置为默认密码:123456成功!"})}))})).catch((()=>{this.$message({type:"info",message:"已取消删除"})}))},turnOnOff(e,t){(0,n.mT)({id:t.ID,name:this.$store.state.userName,available:t.Available}).then((e=>{""===e.data?this.$message({type:"success",message:"修改成功!"}):this.$message("权限不足,修改失败!")}))},handleSelectionChange(){console.log("handleSelectionChange")},handleSearch(){console.log("handleSearch"),console.log(this.inputUserID),console.log(this.inputUserName)},handleAddUser(){(0,n.Qe)({rootName:this.$store.state.userName,name:this.formLabelAlign.userName,password:this.formLabelAlign.userPassword,role:"user"}).then((()=>{this.$message({type:"success",message:"添加成功!"}),this.init()})),this.centerDialogVisible=!1},handleDeleteUser(e,t){console.log("handleDeleteUser"),console.log(e),console.log(t.userID),this.$confirm("此操作将永久删除该用户, 是否继续?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning",center:!0}).then((()=>{console.log("删除成功!"),(0,n._4)({id:t.ID,name:this.$store.state.userName}).then((e=>{console.log(e),this.$message({type:"success",message:"删除成功!"}),this.init()}))})).catch((()=>{this.$message({type:"info",message:"已取消删除"}),console.log("取消删除!")}))},filterUsername(){return this.tableData.filter((e=>(!this.inputUserName||e.Name.includes(this.inputUserName))&&(!this.inputUserID||e.ID==this.inputUserID)))}},watch:{inputUserName(e,t){console.log(e),console.log(t)}}},r=i,o=s(1001),c=(0,o.Z)(r,a,l,!1,null,"1b32f898",null),u=c.exports}}]);
//# sourceMappingURL=381.2e17a150.js.map