import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import Index from '../pages/Index.vue'
import Login from '../pages/Login.vue'
import Register from '../pages/Register.vue'

const routes: Array<RouteRecordRaw> = [
    {
        path: '/',
        name: 'index',
        component: Index,
    },
    {
        path: '/login',
        name: 'login',
        component: Login,
    },
    {
        path: '/register',
        name: 'register',
        component: Register,
    },
]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
})

export default router

