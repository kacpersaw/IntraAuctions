import Vue from 'vue'
import VueRouter from 'vue-router'
import Auction from '../views/Auction.vue'
import auth from '../auth'

Vue.use(VueRouter)

const ifAuthenticated = (to: any, from: any, next: any) => {
  if(auth.logged) {
    next();
    return
  }
  next('/login')
};

const ifAdmin = (to: any, from: any, next: any) => {
  if(auth.logged && auth.admin) {
    next();
    return
  }
  next('/')
};

const routes = [
  {
    path: '/',
    name: 'auctoin',
    meta: {
      title: 'Auction',
    },
    component: Auction,
    beforeEnter: ifAuthenticated
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
    path: '/admin/auctions',
    name: 'auctions',
    meta: {
      title: 'Auctions',
    },
    beforeEnter: ifAdmin,
    component: () => import(/* webpackChunkName: "auction" */ '../views/admin/Auctions.vue')
  }
];

const router = new VueRouter({
  routes
});

export default router
