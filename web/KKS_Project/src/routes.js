import { createRouter, createWebHistory ,} from 'vue-router'

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
		path:'/pic',
		component: () => import('./components/PicShowPage.vue')
	},
	{
		path:'/uploadpic',
		component: () => import('./components/uploadPicPage.vue')
	},
	{
		path: '/timeline',
		component: () => import('./components/TimelinePage.vue')
	}
]


var router = createRouter({
	history: createWebHistory(),
	routes,
	})

export default router