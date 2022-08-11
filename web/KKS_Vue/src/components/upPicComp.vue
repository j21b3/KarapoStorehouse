<template>
	<el-upload
	    v-model:file-list="fileList"
	    action="none"
	    list-type="picture-card"
	    :on-preview="handlePictureCardPreview"
	    :on-remove="handleRemove"
		auto-upload=false
		:on-change="onChangeTest"
	  >
	    <el-icon><Plus /></el-icon>
	  </el-upload>
	  
	<el-dialog v-model="dialogVisible">
	    <img w-full :src="dialogImageUrl" alt="Preview Image" />
	</el-dialog>
	
	<el-button text @click="dialogFormVisible = true">
		信息修改占位
	</el-button>
	<el-dialog v-model="dialogFormVisible" title="图片信息设置" width="40%">
		<el-form :model="form" label-width="140px">
			
				<el-form-item label="上传者"  required >
					  <el-input v-model="form.Uploader" placeholder="人人都可以是Karapo"/>
				</el-form-item>
				<el-form-item label="标题" >
					  <el-input v-model="form.Title" />
				</el-form-item>
				<el-form-item label="留言" >
					  <el-input v-model="form.Message" />
				</el-form-item>
				<el-form-item label="标签">
					  <el-input v-model="form.Tags" />
				</el-form-item>
			
			
			<el-button type="danger" :icon="Close" circle />
			<el-button type="success" :icon="Check" circle/>
			
		</el-form>
	</el-dialog> 
		
</template>

<script lang="ts" setup>
	import { reactive, ref } from 'vue'
	import {Close, Check,Plus} from '@element-plus/icons-vue'
	import { UploadProps, UploadUserFile } from 'element-plus'

	const dialogFormVisible = ref(false)
	const dialogImageUrl = ref('')
	const dialogVisible = ref(false)
	
	const fileList = ref<UploadUserFile[]>([])
	
	const form = reactive({
	  Uploader: '',
	  Title: '',
	  Message: '',
	  Tags: [],
	  Data: '',
	  FileName: '',
	})
	
	const handleRemove: UploadProps['onRemove'] = (uploadFile, uploadFiles) => {
	  console.log(uploadFile, uploadFiles)
	}
	
	const handlePictureCardPreview: UploadProps['onPreview'] = (uploadFile) => {
	  dialogImageUrl.value = uploadFile.url!
	  dialogVisible.value = true
	}
	
	const onChangeTest: UploadProps['onChange'] = (uploadFile, uploadFiles) => {
		console.log("onchange")
		console.log(uploadFile, uploadFiles)
	}
	
</script>
	
<style scoped>

	.el-select {
	  width: 300px;
	}
	.el-input {
	  width: 300px;
	}


</style>