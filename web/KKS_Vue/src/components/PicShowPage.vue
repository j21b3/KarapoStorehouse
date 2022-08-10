<template>
	<headBar />
	<!-- 布局需要优化 -->
	<div :class="'ShowBox'">
	  
		<div v-if="success" :class="'imageShowBox'">
			<img :src="imglink" alt="this is a test" :class="'image'">
		</div>
		<div v-else :class="'imageShowBox'">
			<img :src="failimg" alt="find image fail" :class="'image'">
		</div>
		
	  
		<el-aside :class="'SideShowBox'">
			<div :class="'floatBox'">
				
				<div :class="'UserBox'">
					<div :class="'TextBox'">
						<h1 :class="'PicUploader'">上传者：{{ uploader }}</h1>
						<h4>{{ title }}</h4>
						
						<p :class="'PicMessage'">{{ message }}</p>
						<ul :class="'TagsBox'">
							<li v-for="value in tags" :class="'Tag1'">
								<span style="color: rgb(61, 118, 153);">
									<span>
										<!-- TODO：以后做好标签查询之后把href="url"填进去 -->
										<a  style="color: inherit;text-decoration: none;" v-on:click="notDone">
											#{{value}}
										</a>
									</span>
								</span>
							</li>
						</ul>

					
					</div>
				</div>
				
				<div :class="'buttonBox'">
					<el-container :class="'TextBox'">
						<el-button-group  >
							<el-button type="primary" :icon="Download" round v-on:click="DownloadButtonClick">
								下载原图
							</el-button>
							
						</el-button-group>
					</el-container>
				</div>
			</div>
		</el-aside>
		
	  
	</div>
	
</template>

<script >
	import axios from 'axios'
	import api from './api_config.js'
	import {Download} from '@element-plus/icons-vue'
	import download from 'downloadjs'
	
	import headBar from './headBar.vue'
	
	import { ElMessage } from 'element-plus'
	
	export default{
		name:'PicShowPage',
		
		data(){
			return {
				title:"人是会思考的芦苇",
				filename:"",
				message:"你舍得打破这份宁静吗",
				tags:[],
				uploader:"404",
				success:false,
				pid:this.$route.params.id,
				// pid:"62dbd0d283d5641480925801",
				imglink:"",
				failimg:"/failpic.jpeg",
				Download,
			}
		},
		
		components:{
			headBar,
		},
	
		mounted () {
			// 设置页面背景颜色
			document.querySelector("body").setAttribute('style',"background-color: #132C33;	min-width: 1160px; ")
			
			axios.get(api.raw.detail ,{
				params:{
					pid:this.pid
				}
			})
			  // .get("http://127.0.0.1:25790/raw/62dbd0d283d5641480925801")
			  .then(response => {
					if (response.data.status){
						console.log("获取图片成功")
						//由于有filename，就根据filename的文件类型来选择解码方式，demo中固定用jpg
						this.title=response.data.data.title
						this.filename=response.data.data.filename
						this.message=response.data.data.message
						this.tags=response.data.data.tags
						this.uploader=response.data.data.uploader
						//this.data=response.data.data.data
						
						/*
						let base64String = btoa(
									String.fromCharCode.apply(null, new Uint8Array(response.data.data.data))
								  );
								  */
						// this.data = "data:image/jpg;base64," + response.data.data.data
						//console.log(base64String)
						//console.log(this.data)
						this.success=true
						this.imglink=api.raw.raw+this.pid
					}else{
						console.log("获取失败"+response.data.err)
					}
				})
			  .catch(function (error) { // 请求失败处理
				console.log(error)
			  })
		},
		
		beforeDestroy() {
			document.querySelector("body").removeAttribute('style')
		},
		
		methods:{
			DownloadButtonClick(){
				if (this.success){
					ElMessage({
					        type: 'success',
					        message: `已开始下载`,
							
					      })
					download(this.imglink)	
				}else{
					ElMessage({
					        type: 'warning',
					        message: `没有能下载的东西`,
					      })
				}
			},
			
			notDone(){
				ElMessage({
				        type: 'warning',
				        message: `未完待续`,
						
				      })
			}
		}
	}
</script>

<style>
	
	.ShowBox{
		padding: 0px 72px;
		margin: 8px 8px 8px 8px;
		
	
		display: flex;
		padding-bottom: 12px;
		padding-left: 12px;
		padding-right: 12px;
		padding-top: 12px;
		max-width: 100%;
		max-height: 100%;
		
		/* background-color:  rgba(216, 227, 231, 0.4); */
		/* #D8E3E7 */
		border-radius: 30px;
		
		/* position: absolute;
		top: 0;
		left: 0;
		bottom: 0;
		right: 0; */
		
		
	}
	
	.image{
		width: 100%;
		height: 100%;
	
	
		justify-content: center;
		object-fit: contain;
		border-radius: 20px
		
	}
	
	.imageShowBox{
	
		display: flex;
		justify-content: center;
		
		width: 100%;
		height: 100%;
		min-width: 200px;
		/* min-height: 480px; */
		margin-right: 6px;
		max-width: calc(100% - 12px - 288px);
		max-height: calc(100%);
		
		padding-bottom: 0px;
		position: relative;
		border-radius: 20px;
	
	}
	
	.SideShowBox{
		margin-left: 6px;
		width: 288px;
		
		
		position: relative;
		
		
		margin-bottom: 2px;
		
		border-radius: 20px;
	}
	
	.UserBox{
		
		width: 288px;
		height: calc(62.5% - 6px);
		flex: 0 0 auto;	
		
		margin-bottom: 6px;
		/* background-color: rgba(81, 196, 211, 0.6) ; */
		background-color: rgba(255, 255, 255, 0.5);
		/* #51C4D3 */
		
		border-radius: 20px;
	}
	
	.buttonBox{
		
		width: 288px;
		
		flex: 0 0 auto;	
		margin-top: 2%;
		margin-bottom: 0px;
		
		/* background-color: rgba(81, 196, 211, 0.6) ; */
		/* #126E82 */
		
		background-color: rgba(255, 255, 255, 0.5);
		
		position: relative;
		
		border-radius: 20px;
		
		
	}
	
	.floatBox{
		position: fixed;
		right: 28px;
		/* top: 20px; */
	}
	
	.TextBox{
	
		padding-top: 5px;
		padding-bottom: 5px;
		padding-left: 15px;
		padding-right: 15px;
		text-align: left;
	}
	
	.PicUploader{
		
		
		color: rgb(31, 31, 31);
		font-size: 20px;
		line-height: 24px;
		font-weight: bold;
	}
	
	.PicMessage{
		
		color: rgb(92, 92, 92);
		line-height: 1.33;
	}
	
	.TagsBox{
		word-break: break-all;
		padding-inline-start: 0px;	
	}
	
	.Tag1{
		display: inline;
		margin: 0px 12px 0px 0px;
	}
	
</style>