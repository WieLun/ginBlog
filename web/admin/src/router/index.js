import Vue from 'vue'
import VueRouter from 'vue-router'
const Login = () => import("views/login/Login");
const Admin = () => import("views/admin/Admin");

Vue.use(VueRouter)

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/admin',
    name: 'Admin',
    component: Admin
  },
]

const router = new VueRouter({
  base: "/",
  mode: "history",
  routes,
})

export default router
