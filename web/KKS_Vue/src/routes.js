import { createRouter, createWebHistory } from 'vue-router'

const routes=[
	{
	    path: '/',
	    component: () => import('./App.vue')
	},
	{
	    path:'/pic/:id',
	    component: () => import('./components/PicShowPage.vue')
	},
	{
		path:'/upload',
		component: () => import('./components/uploadPage.vue')
	},
]


var router = createRouter({
	history: createWebHistory(),
	routes,
	})

export default router