<template>	
	<upPicComp></upPicComp>
	
</template>

<script>
	// import type { UploadProps, UploadUserFile } from 'element-plus'
	import api from './api_config.js'
	import upPicComp from './upPicComp.vue'
	
	export default{
		name:'uploadPage',
		components:{
			upPicComp,
		},

	}

</script>

<style>
</style>