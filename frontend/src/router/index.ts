import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import auth from '../auth'

Vue.use(VueRouter)

const ifAuthenticated = (to: any, from: any, next: any) => {
  if(auth.logged) {
    next();
    return
  }
  next('/login')
};

const routes = [
  {
    path: '/',
    name: 'home',
    meta: {
      title: 'Home',
    },
    component: Home
  },
  {
    path: '/login',
    name: 'login',
    meta: {
      title: 'Login',
    },
    component: () => import(/* webpackChunkName: "login" */ '../views/Login.vue'),
  },
  {
    path: '/auction',
    name: 'auction',
    meta: {
      title: 'Auction',
    },
    beforeEnter: ifAuthenticated,
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "auction" */ '../views/Auction.vue')
  }
]

const router = new VueRouter({
  routes
})

export default router
