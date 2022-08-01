import Vue from 'vue'
import axios from 'axios'
import api from 'api_config.js'
import obj from './api_config'

new Vue({
	  el: '#testPic',
	  data () {
	    return {
	      info: null
	    }
	  },
	  mounted () {
	    axios
	      .get(obj.raw.raw+"62dbd0d283d5641480925801")
	      .then(response => (console.log(response.data)))
	      .catch(function (error) { // 请求失败处理
	        console.log(error);
	      });
	  }
})

export default