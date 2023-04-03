"use strict";(self["webpackChunkweb_system"]=self["webpackChunkweb_system"]||[]).push([[384],{4384:function(e,t,i){i.r(t),i.d(t,{default:function(){return h}});var o=function(){var e=this,t=e._self._c;return t("div",{staticClass:"video-css"},[t("el-container",[t("el-main",{staticStyle:{padding:"0px",height:"calc(100vh - 60px)"}},[t("el-header",{staticStyle:{"text-align":"-webkit-center",height:"30px",margin:"10px"}},[t("el-button",{attrs:{type:e.ViewSatus?"danger":"primary",size:"small",round:"",icon:"el-icon-video-play",autofocus:!1},on:{click:function(t){return e.handleViewSatus()}}},[e._v(" 预览 ")]),t("el-button",{directives:[{name:"show",rawName:"v-show",value:e.checkTimerVideoTaskStatus(),expression:"checkTimerVideoTaskStatus()"}],attrs:{type:"primary",size:"small",round:"",icon:"el-icon-video-camera",autofocus:!0},on:{click:function(t){return e.handleTimingTecordingTimeOpen()}}},[e._v("定时录制")]),t("el-button",{directives:[{name:"show",rawName:"v-show",value:!e.checkTimerVideoTaskStatus(),expression:"!checkTimerVideoTaskStatus()"}],attrs:{type:"danger",size:"small",round:"",icon:"el-icon-video-camera",autofocus:!0},on:{click:function(t){return e.handleStopTimedVideoTask()}}},[e._v("等待录制结束，点击停止")]),t("el-button",{attrs:{type:"primary",size:"small",round:"",icon:"el-icon-download",loading:!1},on:{click:function(t){e.LocalFilesDialogVisible=!0}}},[e._v(" 批量回传 ")]),t("el-button",{attrs:{type:e.VideoRecordingStatus?"danger":"primary",size:"small",round:"",icon:"el-icon-bell",loading:!1},on:{click:function(t){return e.handleVideoRecordingStatus()}}},[e._v("录制")]),t("el-button",{attrs:{type:"primary",size:"small",round:"",icon:"el-icon-folder"},on:{click:function(t){return e.handleOpenLocalFlie()}}},[e._v("本地存储")]),t("el-button",{attrs:{type:"primary",size:"small",round:"",icon:"el-icon-time"},on:{click:function(t){return e.handleTimeSyn()}}},[e._v("时间同步 ")]),t("el-button",{attrs:{type:"primary",size:"small",round:"",icon:"el-icon-time"},on:{click:function(t){return e.DeviceList()}}},[e._v(" "+e._s(this.showDevice?"收缩相机列表":"展开相机列表")+" ")])],1),e.showBigVideo?e._e():t("el-main",{staticStyle:{padding:"0px",overflow:"hidden",height:"calc(100vh - 170px)",width:"100%"}},[0===e.VideoView.length?t("el-empty",{attrs:{description:"无设备上线，请稍后。。。"}}):e._e(),e._l(e.VideoView,(function(i,o){return t("div",{key:""+o,staticClass:"video-main"},[o+1===e.nowpage?t("div",{staticStyle:{height:"calc(100vh - 170px)"}},e._l(i,(function(i,a){return t("div",{key:a,staticClass:"video-main",staticStyle:{display:"flex",height:"33.33333333333333333%"}},e._l(i,(function(i,s){return t("div",{key:s,staticStyle:{width:"33.33333333333333333%"}},[t("div",{staticStyle:{width:"100%",height:"100%"}},[t("object",{directives:[{name:"show",rawName:"v-show",value:!e.timingDialogVisible,expression:"!timingDialogVisible"}],attrs:{type:"application/x-vlc-plugin",id:"vlc-"+o+"-"+a+"-"+s,events:"True",width:"100%",height:"100%"}},[t("param",{attrs:{name:"volume",value:"50"}}),t("param",{attrs:{name:"autoplay",value:"true"}}),t("param",{attrs:{name:"loop",value:"false"}}),t("param",{attrs:{name:"fullscreen",value:"false"}}),t("param",{attrs:{name:"controls",value:"false"}})])])])})),0)})),0):e._e()])}))],2),e.showBigVideo?t("el-main",{staticStyle:{padding:"0px",overflow:"hidden",height:"calc(100vh - 120px)"}},[t("object",{directives:[{name:"show",rawName:"v-show",value:!e.timingDialogVisible,expression:"!timingDialogVisible"}],attrs:{type:"application/x-vlc-plugin",id:"FullScreen",events:"True",width:"100%",height:"100%"}},[t("param",{attrs:{name:"volume",value:"50"}}),t("param",{attrs:{name:"autoplay",value:"true"}}),t("param",{attrs:{name:"loop",value:"false"}}),t("param",{attrs:{name:"fullscreen",value:"false"}}),t("param",{attrs:{name:"controls",value:"false"}})])]):e._e(),e.showBigVideo?e._e():t("el-footer",[t("el-pagination",{staticStyle:{margin:"0px 0px",height:"4%","text-align-last":"center",padding:"18px 5px"},attrs:{background:"",layout:"prev, pager, next","page-size":9,total:this.$store.state.deviceList.length},on:{"current-change":e.handleCurrentChange}})],1)],1),t("el-aside",{directives:[{name:"show",rawName:"v-show",value:e.showDevice,expression:"showDevice"}],staticStyle:{padding:"0px",height:"calc(100vh - 60px)"},attrs:{width:"400px"}},[e._l(this.$store.state.downLoadStatus,(function(i,o){return t("div",{key:o,staticStyle:{width:"100%"}},[i.progress!=i.total?t("div",[t("el-progress",{staticClass:"my-progress-bar",attrs:{"text-inside":!0,"stroke-width":26,percentage:Math.ceil(i.progress/i.total*100),status:"exception",format:t=>e.formatText(t,i.file_name,i.total),color:"#c03838"}})],1):e._e()])})),t("el-main",{staticClass:"test",staticStyle:{padding:"0px",height:"95%"}},[t("el-table",{staticStyle:{width:"100%"},attrs:{data:this.$store.state.deviceList}},[t("el-table-column",{attrs:{label:"相机列表",align:"center",width:"150"},scopedSlots:e._u([{key:"default",fn:function(i){return[t("el-popover",{attrs:{trigger:"hover",placement:"top"}},[t("p",[e._v("GoProIP: "+e._s(i.row.gopro_ip))]),t("div",{staticClass:"name-wrapper",attrs:{slot:"reference"},slot:"reference"},[t("el-tag",{attrs:{size:"medium"}},[e._v(e._s(i.row.device_name))])],1)])]}}])}),t("el-table-column",{attrs:{label:"操作",align:"center"},scopedSlots:e._u([{key:"default",fn:function(i){return[t("el-tooltip",{staticClass:"item",attrs:{effect:"dark",content:"下载最新文件",placement:"right-end"}},[t("el-button",{attrs:{size:"mini",type:"warning",round:"",plain:""},on:{click:function(t){return e.handleReturnVideo(i.$index,i.row)}}},[e._v("回传")])],1),t("el-button",{attrs:{size:"mini",round:""},on:{click:function(t){return e.clickFullscreen(i.$index,i.row)}}},[e._v("全屏")]),t("el-button",{attrs:{size:"mini",type:"success",round:"",plain:""},on:{click:function(t){return e.handleReturnVideoFileList(i.$index,i.row)}}},[e._v("文件")])]}}])})],1)],1)],2)],1),t("el-dialog",{attrs:{title:"设置定时录制视频",visible:e.timingDialogVisible,width:"30%"},on:{"update:visible":function(t){e.timingDialogVisible=t}}},[t("div",{staticClass:"block"},[t("el-date-picker",{attrs:{type:"datetimerange","range-separator":"至","start-placeholder":"开始日期","end-placeholder":"结束日期"},model:{value:e.timingTecordingTimeValue1,callback:function(t){e.timingTecordingTimeValue1=t},expression:"timingTecordingTimeValue1"}})],1),t("span",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[t("el-button",{on:{click:function(t){e.timingDialogVisible=!1}}},[e._v("取 消")]),t("el-button",{attrs:{type:"primary"},on:{click:function(t){return e.handletimingTecordingTime()}}},[e._v("保 存")])],1)]),t("el-dialog",{attrs:{title:"本地文件存储",visible:e.localFileStorageDialogTableVisible,width:"60%"},on:{"update:visible":function(t){e.localFileStorageDialogTableVisible=t}}},[t("el-table",{staticStyle:{width:"100%"},attrs:{data:e.VideoFilesTableData}},[t("el-table-column",{attrs:{prop:"Name",label:"文件名称",width:"180"}}),t("el-table-column",{attrs:{prop:"ModTime",label:"创建日期",width:"180"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(t.row.ModTime.split(".")[0])+" ")]}}])}),t("el-table-column",{attrs:{prop:"Size",label:"文件大小"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(e.bytesToSize(parseInt(t.row.Size)))+" ")]}}])}),t("el-table-column",{attrs:{prop:"address",label:"操作",width:"300"},scopedSlots:e._u([{key:"default",fn:function(i){return[t("el-button",{attrs:{size:"mini",type:"primary"},on:{click:function(t){return e.openRenameLacolVideoFile(i.row.Name)}}},[e._v(" 重命名 ")]),t("el-button",{attrs:{size:"mini",type:"danger"},on:{click:function(t){return e.handleRemoveLocalVideoFile(i.row.Name)}}},[e._v("删除 ")])]}}])})],1)],1),t("el-dialog",{attrs:{title:this.GoProFileTable.GoProIP+"-视频文件",visible:e.ReturnVideoFileListdialogTableVisible,width:"60%"},on:{"update:visible":function(t){e.ReturnVideoFileListdialogTableVisible=t}}},[t("el-table",{attrs:{data:e.GoProFileTable.FileList}},[t("el-table-column",{attrs:{property:"n",label:"文件名称",width:"auto"}}),t("el-table-column",{attrs:{property:"cre",label:"创建日期",width:"auto"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(e._s(new Date(1e3*t.row.cre).toISOString()))]}}])}),t("el-table-column",{attrs:{property:"s",label:"文件大小",width:"auto"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(e._s(e.bytesToSize(parseInt(t.row.s))))]}}])}),t("el-table-column",{attrs:{label:"操作",width:"auto"},scopedSlots:e._u([{key:"default",fn:function(i){return[t("el-button",{attrs:{size:"mini",type:"primary"},on:{click:function(t){return e.handleReturnOneVideoFile(i.row)}}},[e._v("下载 ")])]}}])})],1)],1),t("el-dialog",{attrs:{title:"批量回传",visible:e.LocalFilesDialogVisible,width:"30%"},on:{"update:visible":function(t){e.LocalFilesDialogVisible=t}}},[t("span",[t("el-checkbox",{attrs:{indeterminate:e.isIndeterminate},on:{change:e.handleCheckAllChange},model:{value:e.checkAll,callback:function(t){e.checkAll=t},expression:"checkAll"}},[e._v("全选 ")]),t("div",{staticStyle:{margin:"15px 0"}}),t("el-checkbox-group",{on:{change:e.handleCheckedCitiesChange},model:{value:e.DownloadLatestFileFromSelectedDevice,callback:function(t){e.DownloadLatestFileFromSelectedDevice=t},expression:"DownloadLatestFileFromSelectedDevice"}},e._l(e.cities,(function(i){return t("el-checkbox",{key:i,attrs:{label:i}},[e._v(e._s(i))])})),1)],1),t("span",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[t("el-button",{on:{click:function(t){e.LocalFilesDialogVisible=!1}}},[e._v("取 消")]),t("el-button",{attrs:{type:"primary"},on:{click:function(t){return e.handleBatchVideoBack()}}},[e._v("确 定")])],1)])],1)},a=[],s=(i(7658),i(586)),l=i(6028),n={data(){return{intervalID:null,ViewSatus:!1,VideoRecordingStatus:!1,timingTecordingTimeValue1:[new Date,new Date+3600],timingDialogVisible:!1,dialogTableVisible:!1,LocalFilesDialogVisible:!1,ReturnVideoFileListdialogTableVisible:!1,localFileStorageDialogTableVisible:!1,checkAll:!1,DownloadLatestFileFromSelectedDevice:[],cities:[],isIndeterminate:!0,showBigVideo:!1,BigVideoID:null,showDevice:!0,VideoFilesTableData:null,DevTableData:this.$store.state.deviceList,GoProFileTable:{GoProIP:null,DeviceName:null,FileList:null},VideoView:[],flvVideoList:[],nowpage:1}},methods:{initVideo(){var e=this.$store.state.deviceList;let t=[],i=[],o=3;for(let a=0,s=e.length,l=0;a<s;a+=o)i.push(e.slice(a,a+o)),l>=2?(t.push(i),console.log(l),i=[],l=0):l+=1;return i.length>0&&t.push(i),t},clickVideo(e){this.showBigVideo=!this.showBigVideo,this.BigVideoID=e,console.log(e)},handleViewLocalFiles(e){window.open("http://"+l.c.ServerIP+":8081/VideoDownload/"+e,"_blank")},handleTimingTecordingTimeOpen(){!1===this.ViewSatus&&!1===this.VideoRecordingStatus&&0==this.showBigVideo?this.timingDialogVisible=!0:(this.ViewSatus&&this.$notify.error({title:"注意！",message:"必须先关闭所有设备视频预览"}),this.VideoRecordingStatus&&this.$notify.error({title:"注意！",message:"必须先关闭所有设备视频录制"}),this.showBigVideo&&this.$notify.error({title:"注意！",message:"必须先关闭全屏"}))},handletimingTecordingTime(){console.log(this.timingTecordingTimeValue1[0].getTime()/1e3),console.log(this.timingTecordingTimeValue1[1].getTime()/1e3),this.$store.state.TimedVideoTaskStateWebSocket=[this.timingTecordingTimeValue1[0].getTime(),this.timingTecordingTimeValue1[1].getTime()],(0,s._z)(parseInt(this.timingTecordingTimeValue1[0].getTime()/1e3),parseInt(this.timingTecordingTimeValue1[1].getTime()/1e3)),this.timingDialogVisible=!1},handleStopTimedVideoTask(){(0,s.ci)()},handleBatchVideoBack(){this.LocalFilesDialogVisible=!1;const e=this.DownloadLatestFileFromSelectedDevice;e.forEach((e=>{const t=this.$store.state.deviceList.filter((t=>t.device_name===e)),i=t.find((t=>t.device_name===e));(0,s.aP)(i.gopro_ip).then((e=>{const t=e.data.media[0].fs[e.data.media[0].fs.length-1].n,o=this.formatDate(e.data.media[0].fs[e.data.media[0].fs.length-1].cre);console.log(o),(0,s.xn)(i.gopro_ip,t,`${i.device_name}_${o}_${t}`)}))}))},handleOpenLocalFlie(){(0,s.WQ)(),console.log("handleOpenLocalFlie")},bytesToSize(e){if(0===e)return"0 B";var t=1e3,i=["B","KB","MB","GB","TB","PB","EB","ZB","YB"],o=Math.floor(Math.log(e)/Math.log(t));return(e/Math.pow(t,o)).toPrecision(3)+" "+i[o]},handleCurrentChange(e){return console.log(`当前页: ${e}`),console.log(e),this.nowpage=e,e},handleEdit(e,t){console.log(e,t)},handleReturnVideo(e,t){console.log("handleReturnVideo"),(0,s.aP)(t.gopro_ip).then((e=>{console.log(e.data.media[0].fs[e.data.media[0].fs.length-1].n);var i=e.data.media[0].fs[e.data.media[0].fs.length-1].n;const o=this.formatDate(e.data.media[0].fs[e.data.media[0].fs.length-1].cre);console.log(o),(0,s.xn)(t.gopro_ip,i,t.device_name+"_"+o+"_"+i)}))},handleWatchVideo(e,t){window.open("http://"+this.GoProFileTable.GoProIP+"/videos/DCIM/100GOPRO/"+t.n,"_blank"),console.log(this.GoProFileTable.GoProIP),console.log(t.n)},handleReturnVideoFileList(e,t){this.GoProFileTable.FileList=null,this.GoProFileTable.GoProIP=null,this.GoProFileTable.DeviceName=null,this.ReturnVideoFileListdialogTableVisible=!0,console.log("handleReturnVideoFileList"),console.log(e,t),(0,s.aP)(t.gopro_ip).then((e=>{console.log(e.data.media[0].fs),this.GoProFileTable.FileList=e.data.media[0].fs,this.GoProFileTable.GoProIP=t.gopro_ip,this.GoProFileTable.DeviceName=t.device_name}))},openRenameLacolVideoFile(e){this.$prompt("请输入新文件名","提示",{confirmButtonText:"确定",cancelButtonText:"取消"}).then((({value:t})=>{(0,s.wq)({oldName:e,newName:t}).then((e=>{console.log(e),666===e.data.code?(this.$message({type:"success",message:"新文件名: "+t}),this.handleOpenLocalFlie()):this.$message({type:"info",message:e.data.message})})),console.log(t,e)})).catch((()=>{this.$message({type:"info",message:"取消输入"})}))},handleRemoveLocalVideoFile(e){this.$confirm("此操作将永久删除该文件, 是否继续?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then((()=>{(0,s.db)({name:e}).then((e=>{666===e.data.code?(this.$message({type:"success",message:"删除成功!"}),this.handleOpenLocalFlie()):this.$message({type:"info",message:"删除失败,错误情况:"+e.data.message})}))})).catch((()=>{this.$message({type:"info",message:"已取消删除"})})),console.log(e)},DeviceList(){this.showDevice=!this.showDevice},handleViewSatus(){if(!1===this.VideoRecordingStatus&&0==this.showBigVideo&&!0===this.checkTimerVideoTaskStatus())return this.ViewSatus=!this.ViewSatus;this.VideoRecordingStatus&&this.$notify.error({title:"注意！",message:"必须先关闭所有设备视频录制"}),this.showBigVideo&&this.$notify.error({title:"注意！",message:"必须先关闭全屏"}),this.checkTimerVideoTaskStatus()||this.$notify.error({title:"注意！",message:"必须先等待定时录像任务结束"})},handleVideoRecordingStatus(){if(!1===this.ViewSatus&&0==this.showBigVideo&&1==this.checkTimerVideoTaskStatus())return this.VideoRecordingStatus=!this.VideoRecordingStatus;this.ViewSatus&&this.$notify.error({title:"注意！",message:"必须先关闭所有设备视频预览"}),this.showBigVideo&&this.$notify.error({title:"注意！",message:"必须先关闭全屏"}),this.checkTimerVideoTaskStatus()||this.$notify.error({title:"注意！",message:"必须等待定时录制完成"})},handleTimeSyn(){this.$store.state.deviceList.forEach((function(e){console.log(e.device_ip.split(":")[0]),(0,s.iG)(e.device_ip.split(":")[0]).then((e=>{console.log(e)})).then((e=>{console.log(e)}))}))},handleCheckAllChange(e){this.DownloadLatestFileFromSelectedDevice=e?this.cities:[],this.isIndeterminate=!1},handleCheckedCitiesChange(e){let t=e.length;this.checkAll=t===this.cities.length,this.isIndeterminate=t>0&&t<this.cities.length},handleAddVideoTitle(e,t){let i={},o={},a=[],s=t;e[s-1].forEach((function(e,t){e.forEach((function(e,l){i["vlc-"+(s-1)+"-"+t+"-"+l]=e.device_name,o["vlc-"+(s-1)+"-"+t+"-"+l]=e.utc_time,a.push("vlc-"+(s-1)+"-"+t+"-"+l)}))})),a.forEach((function(e,t){let a=document.getElementById(e);try{a.video.marquee.text=i[e]+"-"+o[e][0].toString()+"-"+o[e][1].toString().padStart(2,"0")+"-"+o[e][2].toString().padStart(2,"0")+" "+o[e][3].toString().padStart(2,"0")+":"+o[e][4].toString().padStart(2,"0")+":"+o[e][5].toString().padStart(2,"0"),a.video.marquee.color=16776960,a.video.marquee.size=100,a.video.marquee.x=50,a.video.marquee.y=50,a.video.marquee.position="TOP",a.video.marquee.enable()}catch(s){console.log("no video stream,so add failed!")}}))},handleAddVideoUrl(){let e=this.nowpage;this.VideoView[e-1].forEach((function(t,i){t.forEach((function(t,o){console.log(t),(0,s.AZ)(t.device_ip.split(":")[0]).then((t=>{console.log(t.data.video_port);let a=document.getElementById("vlc-"+(e-1)+"-"+i+"-"+o);a.playlist.add("udp://@0.0.0.0:"+t.data.video_port),a.playlist.play()}))}))}))},handleClearVideoUrl(){let e=this.nowpage;this.VideoView[e-1].forEach((function(t,i){t.forEach((function(t,o){let a=document.getElementById("vlc-"+(e-1)+"-"+i+"-"+o);a.playlist.clear()}))}))},clickFullscreen(e,t){if(!1===this.ViewSatus&&1==this.checkTimerVideoTaskStatus()&&!1===this.VideoRecordingStatus)if(this.showBigVideo=!this.showBigVideo,this.showBigVideo)console.log("clickFullscreen"),(0,s.AZ)(t.device_ip.split(":")[0]).then((e=>{console.log(e.data.video_port);let i=document.getElementById("FullScreen");i.playlist.add("udp://@0.0.0.0:"+e.data.video_port),i.playlist.play(),(0,s.eb)(t.device_ip.split(":")[0],t.gopro_ip),this.$notify({title:"成功",message:"开始全屏播放"+t.device_name+"设备视频",type:"success"})}));else{let e=document.getElementById("FullScreen");e.playlist.clear(),(0,s.Z_)(t.device_ip.split(":")[0],t.gopro_ip).then((()=>{this.$notify.error({title:"注意！",message:"关闭全屏播放"})}))}else this.ViewSatus&&this.$notify.error({title:"注意！",message:"必须先关闭所有设备视频预览"}),this.VideoRecordingStatus&&this.$notify.error({title:"注意！",message:"必须先关闭所有设备视频录制"}),this.checkTimerVideoTaskStatus()||this.$notify.error({title:"注意！",message:"必须等待定时录制完成"})},formatDate(e){const t=new Date(1e3*e),i=new Date(t.getTime()-288e5),o=i.getFullYear(),a=String(i.getMonth()+1).padStart(2,"0"),s=String(i.getDate()).padStart(2,"0"),l=String(i.getHours()).padStart(2,"0"),n=String(i.getMinutes()).padStart(2,"0"),r=String(i.getSeconds()).padStart(2,"0");return`${o}-${a}-${s}_${l}-${n}-${r}`},formatText(e,t,i){return`名称:${t},进度:${e}%,${this.bytesToSize(i)}`},handleReturnOneVideoFile(e){console.log("handleReturnVideo");const t=this.formatDate(e.cre);(0,s.xn)(this.GoProFileTable.GoProIP,e.n,this.GoProFileTable.DeviceName+"_"+t+"_"+e.n)},checkTimerVideoTaskStatus(){return null===this.$store.state.TimedVideoTaskStateWebSocket},showTimerVideoTaskStatusS(){return this.$store.state.TimedVideoTaskStateWebSocket.end_time>Date.now()>this.$store.state.TimedVideoTaskStateWebSocket.start_time?"距离结束录像时间剩余：":this.$store.state.TimedVideoTaskStateWebSocket.end_time>this.$store.state.TimedVideoTaskStateWebSocket.start_time>Date.now()?"距离开始录像时间剩余：":void 0},test(e){console.log(e)}},watch:{"$store.state.err":{deep:!0,handler:function(){this.$notify.error({title:"错误",message:this.$store.state.err,duration:0})}},"$store.state.deviceList":{deep:!0,handler:function(){this.VideoView=this.initVideo(),this.DevTableData=this.$store.state.deviceList,this.cities=[],this.$store.state.deviceList.forEach((e=>{this.cities.push(e.device_name)}))}},ViewSatus:{deep:!0,handler:function(e,t){console.log("ViewSatus数据发生变化啦"),console.log(e,t),e?(this.handleAddVideoUrl(),this.intervalID=setInterval((()=>{this.handleAddVideoTitle(this.VideoView,this.nowpage)}),500),console.log(this.DevTableData.slice(9*(this.nowpage-1),9*this.nowpage)),this.DevTableData.slice(9*(this.nowpage-1),9*this.nowpage).forEach((function(e){console.log(e.device_ip.split(":")[0]),(0,s.eb)(e.device_ip.split(":")[0],e.gopro_ip)})),this.$notify({title:"成功",message:"已成功打开所有设备视频流",type:"success"})):(this.handleClearVideoUrl(),clearInterval(this.intervalID),this.DevTableData.slice(9*(this.nowpage-1),9*this.nowpage).forEach((function(e){console.log(e.device_ip.split(":")[0]),(0,s.Z_)(e.device_ip.split(":")[0],e.gopro_ip)})),this.$notify.error({title:"关闭",message:"已成功关闭所有设备视频流"}))}},VideoRecordingStatus:{deep:!0,handler:function(e,t){console.log("VideoRecordingStatus数据发生变化啦"),console.log(e,t),e?(this.$store.state.deviceList.forEach((function(e){console.log(e.gopro_ip),(0,s.RR)(e.gopro_ip,"start")})),this.$notify({title:"成功",message:"已开始所有设备录制视频",type:"success"})):(this.$store.state.deviceList.forEach((function(e){console.log(e.gopro_ip),(0,s.RR)(e.gopro_ip,"stop")})),this.$notify.error({title:"关闭",message:"已停止所有设备录制视频"}))}}},mounted(){this.$store.state.deviceList.forEach((e=>{this.cities.push(e.device_name)})),this.VideoView=this.initVideo()},beforeDestroy(){clearInterval(this.intervalID)}},r=n,c=i(1001),d=(0,c.Z)(r,o,a,!1,null,null,null),h=d.exports}}]);
//# sourceMappingURL=384.118d5fb1.js.map