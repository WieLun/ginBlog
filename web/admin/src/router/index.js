import Vue from 'vue'
import VueRouter from 'vue-router'
const Login = () => import("views/login/Login");

Vue.use(VueRouter)

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
]

const router = new VueRouter({
  base: "/",
  mode: "history",
  routes,
})

export default router
