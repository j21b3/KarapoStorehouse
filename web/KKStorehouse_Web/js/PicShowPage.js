import Vue from 'vue'
import axios from 'axios'
import obj from './api_config'


new Vue({
	  el: '#testPic',
	  data :{    
		
		title:"",
		filename:"",
		message:"",
		tags:[],
		uploader:"",
		data:null,
	  },
	  mounted () {
		axios
		  .get(obj.raw.raw+"62e8e95446b76ab6631787b3")
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
					this.data = "data:image/jpg;base64," + response.data.data.data
					//console.log(base64String)
					//console.log(this.data)
			
				}else{
					console.log("获取失败"+response.data.err)
				}
			})
		  .catch(function (error) { // 请求失败处理
			console.log(error);
		  });
	  }
})


export default