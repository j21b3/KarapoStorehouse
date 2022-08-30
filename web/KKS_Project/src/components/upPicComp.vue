<template>
	<div :class="'display-area'">
		<el-row :gutter="20">
			<el-col
			v-for="(file, index) in fileList "
			:key="file"
			:span="4"			
			>
			<el-card :body-style="{ padding: '0px' }">
				<div :class="'image_show'">
					<el-image :src="file.url" 
					:fit="scale-down" 
					:loading="lazy"
					:preview-src-list="fileUrlList"
					:initial-index="index"
					/>
				</div>
				<div :class="'detailShow'">
					<!-- <span>Yummy hamburger</span>
					<div class="bottom">
						<time class="time">{{ currentDate }}</time>
						<el-button text class="button">Operating</el-button>
					</div> -->
					<p style="margin-top: 0px; margin-bottom: 0px">
						上传：{{file.form.Uploader}} <br/>
						标题：{{file.form.Title}}	<br/>
						Msg: {{file.form.Title}}	<br/>
					</p>
					
					<TagComp :Tags="file.form.Tags" v-if="file.form.Tags.length > 0"/>
					<td v-if="file.form.Tags.length <= 0">tags:</td>

				</div>
			</el-card>
			</el-col>
		</el-row>
	</div>
	
	<el-button text @click="dialogFormVisible = true">
		信息修改占位
	</el-button>
	<el-dialog 
		:before-close="closeButtonClick" 
		:close-on-click-modal="false"
		v-model="dialogFormVisible" 
		title="图片信息设置" 
		width="44%">
		<el-upload
			v-model:file-list="tmpfileList"
			action="none"
			list-type="picture-card"
			:on-preview="handlePictureCardPreview"
			:on-remove="handleRemove"
			:auto-upload="false"
			:on-change="onChangeTest"
			:multiple="true"
			ref="upload"
			accept=".jpg,.jpeg,.png,.gif"
			
			>
			<el-icon><Plus /></el-icon>
		</el-upload>

		<el-form :model="form" label-width="140px">
				
				<div class="el-upload__tip">
				        下方属性会赋值给所有图片，如需添加不同属性需要多次填写表单
				</div>
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
					  <el-input v-model="form.Tags" placeholder="以空格隔开"/>
				</el-form-item>
			
			
			<el-button type="danger" :icon="Close" @click="closeButtonClick" circle />
			<el-button type="success" :icon="Check" @click="okButtonClick" circle/>
			
		</el-form>
	</el-dialog> 


	<el-dialog v-model="dialogVisible" :class="'dialogView'">
		<el-image :src="dialogImageUrl" 
				:fit="'contain'"
				alt="Preview Image" 
				/>
	</el-dialog>

</template>

<script lang="ts" setup>
	export interface UploadForm{
		Uploader: '',
		Title: '',
		Message: '',
		Tags: [],
		FileName: '',
	}
	
	import { reactive, ref } from 'vue'
	import {Close, Check,Plus} from '@element-plus/icons-vue'
	import { UploadProps, UploadUserFile } from 'element-plus'
	import { ElMessage } from 'element-plus'
	import TagComp from './TagComp.vue'

	const dialogFormVisible = ref(false)
	const dialogImageUrl = ref('')
	const dialogVisible = ref(false)
	
	const tmpfileList = ref<UploadUserFile[]>([])
	const fileList = ref<object[]>([])
	const fileUrlList = ref<string[]>([])
	
	const upload = ref(null)
	
	const form = reactive({
	  Uploader: '',
	  Title: '',
	  Message: '',
	  Tags: [],
	})
	
	const handleRemove: UploadProps['onRemove'] = (uploadFile, uploadFiles) => {
	  console.log(uploadFile, uploadFiles)
	}
	
	const handlePictureCardPreview: UploadProps['onPreview'] = (uploadFile) => {
		console.log("preview")
	  	dialogImageUrl.value = uploadFile.url!
	  	dialogVisible.value = true
	}
	
	const onChangeTest: UploadProps['onChange'] = (uploadFile, uploadFiles) => {
		console.log("onchange")
		console.log(uploadFile, uploadFiles)
	}

	const createUpPic = (url, fileName,formData) => {
		var pic = {
			url: url,
			form: {
				Uploader: formData.Uploader,
				Title: formData.Title,
				Message: formData.Message,
				Data: formData.Data,
				FileName: fileName,
				Tags:Array<string>(),
			}
		}
		if (formData.Tags != "") {
			var tmp = formData.Tags.split(' ')
			if (tmp.length != 0) {
				var i = 0
				for (i = 0; i < tmp.length; i++) {
					if (tmp[i] != "") {
						pic.form.Tags.push(tmp[i])
					}
				}
			}
		}
		fileUrlList.value.push(pic.url)
		return pic
	}

	const okButtonClick = () => {
		console.log(tmpfileList.value.length)
		console.log(tmpfileList)
		if (tmpfileList.value.length == 0){
			dialogFormVisible.value = false
			ElMessage({
			        type: 'warning',
			        message: `没有要上传的东西`,
			      })
			return
		}
		// TODO：這裏要加入表單校驗
		
		var i = 0
		for (i = 0; i < tmpfileList.value.length; i++){
			var tmp = createUpPic(tmpfileList.value[i].url, tmpfileList.value[i].name, form)
			fileList.value.push(tmp)
		}
		form.Uploader = ""
		form.Title = ""
		form.Message = ""
		form.Tags = []
		dialogFormVisible.value = false
		upload.value.clearFiles()
	}
		
	const closeButtonClick = () => {
		dialogFormVisible.value = false
		upload.value.clearFiles()
	}
	

</script>
	
<style scoped>

	.el-select {
	  width: 300px;
	}
	.el-input {
	  width: 300px;
	}
	.display-area{
		padding-left: 72px;
		padding-right: 72px;

	}
	.image_show{
		/* max-width: 186px;
		max-height: 186px;
		min-width: 186px;
		min-height: 186px; */
		position: relative;
	}

	.el-row {
		margin-bottom: 20px;
	}
	.el-row:last-child {
		margin-bottom: 0;	
	}

	.dialogView{
		width: 60%;
	}

	.detailShow{
		padding: 14px;
		text-align: left;
	}

</style>