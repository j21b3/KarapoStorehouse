import { createApp } from 'vue'
import App from './App.vue'
import { createRouter,createWebHashHistory } from 'vue-router'

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

import PicShowPage from './components/PicShowPage.vue'



const routes=[
	{
	    path: '/hello',
	    component: App
	},
	{
	    path:'/pic/:id',
		name:'picshow',
	    component:PicShowPage
	},
]

const Router=createRouter(
	{
		history:createWebHashHistory(),
		routes,
	}
)



const app=createApp(PicShowPage)
app.use(Router)
app.use(ElementPlus)
app.mount('#app')