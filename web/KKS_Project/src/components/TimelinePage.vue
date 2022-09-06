<template>
    <div :class="'display-area'">
		<el-row :gutter="20">
			<el-col
			v-for="(file, index) in fileList "
			:key="file"
			:span="4"			
			>
                <el-card 
                    :body-style="{ padding: '0px'}"
                    shadow="hover">
                        <div :class="'image_show'">
                            <el-image :src="file.url" 
                            fit="scale-down" 
                            :preview-src-list="fileUrlList"
                            :initial-index="index"
                            />
                        </div>
                        <div :class="'detailShow'">
                            <p style="margin-top: 0px; margin-bottom: 0px">
                                上传：{{file.form.Uploader}} <br/>
                                标题：{{file.form.Title}}	<br/>
                                Msg: {{file.form.Message}}	<br/>
                            </p>
                            <td v-if="file.form.Tags.length <= 0">tags:</td>
                            <TagComp :Tags="file.form.Tags" v-if="file.form.Tags.length > 0"/>
                            

                        </div>
                </el-card>
			</el-col>
		</el-row>
	</div>
</template>

<script>

</script>

<style>

</style>