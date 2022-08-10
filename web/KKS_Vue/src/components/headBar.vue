<template>
	<el-header>
		<el-row :glutter="20" :class="'myheader'">
			<el-col :span="1">
				<router-link to="/">
					<!-- <img :src="klpHead" alt="头像" :class="klphead" /> -->
					<el-avatar :src="klpHead" :class="hlphead"></el-avatar>
				</router-link>
			</el-col>
			<el-col :span="1">
				<router-link to="/upload/">
					上传
				</router-link>
			</el-col>
			<el-col :span="7">
				<searchBar></searchBar>
			</el-col>
		</el-row>
	</el-header>
</template>

<script>
	import {Download} from '@element-plus/icons-vue'
	import searchBar from './searchBar.vue'
	
	export default{
		name:'headBar',
		data(){
			return {
				searchText:String,
				klpHead:"/klphead.jpg",
			}
		},
		components:{
			searchBar
		},
	}
	
	
		
	
</script>

<style scoped="scoped">
	header {
		display: flex;	
		align-items: center;
		height: 58px;
		justify-content: start;
		
		background-color: #D8E3E7;
		
		position: fixed;
		top: 0%;left: 0%;right: 0%;
		width: 100%;
		z-index: 2000;
		margin-bottom: 3px;
		
		border-bottom-left-radius: 10px;
		border-bottom-right-radius: 10px;
		
		padding-left: 15px;
		padding-right: 15px;
	}
	
	.myheader{
		display: flex;
		align-items: center;
		height: 58px;
		justify-content: start;
		
		background-color: #D8E3E7;
		
		position: fixed;
		top: 0%;left: 0%;right: 0%;
		width: 100%;
		z-index: 999999;
		margin-bottom: 3px;
		
		border-bottom-left-radius: 10px;
		border-bottom-right-radius: 10px;
		
		padding-left: 15px;
		padding-right: 15px;
	}
	
	img{
		width: 40px;
		background-color: transparent;
		border-style: none;
		visibility: visible;
		border-radius: 20px;
	}
	
</style>