<template>
	<div>
	    <el-input placeholder="" round v-model="searchOption.input" class="searchClass">
			<template #prepend>
				<el-select v-model="searchOption.select" placeholder="搜索模式" style="width: 80px;">
					<el-option v-for="(text,index) in searchMode" :label="text" :value="text" />
				</el-select>
			</template>
			<template #append>
				<el-button round :icon="Search" @click="searchButtonClick"/>
			</template>
	    </el-input>
	</div>
</template>

<script>
	import {Search} from '@element-plus/icons-vue'
	import { ref } from 'vue'
	import api from './api_config.js'
	import { ElMessage } from 'element-plus'
	
	export default{
		
		name:"searchBar",
		
		data(){
			return{
				searchOption:{
					input:"",
					select:"ID",
				},
				Search,
				searchMode:[
					"ID",
					"标签(未实现)",
				],
			}
		},
		
		methods:{
			searchButtonClick(){
				
				if(this.searchOption.select=="ID"){
					window.open(api.domain.pic+this.searchOption.input,'_blank')
				}else{
					ElMessage({
					        type: 'warning',
					        message: `未完待续`,
					      })
				}
			}
		}
		
	}
</script>

<style scoped="scoped">
	.searchClass{
		border: 1px solid #c5c5c5;
		/* border-radius: 20px; */
		background:#f4f4f4;
		width:400px;
	}
	.searchClass .el-input-group__prepend {
	  border: none;
	  background-color: transparent;
	  padding: 0 10px 0 30px;
	}
	.searchClass .el-input-group__append {
	  border: none;
	  background-color: transparent;
	}
	.searchClass .el-input__inner {
	  height: 36px;
	  line-height: 36px;
	  border: none;
	  background-color: transparent;
	}
	.searchClass .el-icon-search{
	  font-size: 16px;
	}
	.searchClass .centerClass {
	  height: 100%;
	  line-height: 100%;
	  display: inline-block;
	  vertical-align: middle;
	  text-align: right;
	}
	.searchClass .line {
	  width: 1px;
	  height: 26px;
	  background-color: #c5c5c5;
	  margin-left: 14px;
	}
	.searchClass:hover{
	  border: 1px solid #D5E3E8;
	  background: #fff;
	}
	.searchClass:hover .line {
	  background-color: #D5E3E8;
	}
	.searchClass:hover .el-icon-search{
	  color: #409eff;
	  font-size: 16px;
	}
	
</style>