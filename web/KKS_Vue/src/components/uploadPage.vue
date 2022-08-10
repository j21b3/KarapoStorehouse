<template>
	<headBar />
	
	<el-upload
	    v-model:file-list="fileList"
	    class="upload-demo"
	    :action="uploadUrl"
	    :on-preview="handlePreview"
	    :on-remove="handleRemove"
	    list-type="picture"
	  >
	    <el-button type="primary">Click to upload</el-button>
	    <template #tip>
	      <div class="el-upload__tip">
	        jpg/png files with a size less than 500kb
	      </div>
	    </template>
	</el-upload>
	
</template>

<script>
	import headBar from './headBar.vue'
	// import type { UploadProps, UploadUserFile } from 'element-plus'
	import api from './api_config.js'
	
	
	export default{
		name:'uploadPage',
		
		components:{
			headBar
		},
		
		data(){
			return{
				uploadUrl:api.upload,
			}
		}
	}

</script>

<style>
</style>