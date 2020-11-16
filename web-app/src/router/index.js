import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Customers from '../views/Customers.vue'
import CustomerDetail from '../views/CustomerDetail.vue'
import Page404 from '../views/Page404.vue'
import AppHeader from '../layout/AppHeader.vue'
import AppFooter from '../layout/AppFooter.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    components: {
      header: AppHeader,
      default: Home,
      footer: AppFooter
    }
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  },
  {
    path: '/customers',
    name: 'Customers',
    components: {
      header: AppHeader,
      default: Customers,
      footer: AppFooter
    }
  },
  {
    path: '/customer-detail/:customer_id',
    name: 'CustomerDetail',
    components: {
      header: AppHeader,
      default: CustomerDetail,
      footer: AppFooter
    }
  },
  {
    path: "*",
    name: "404",
    components: {
      header: AppHeader,
      default: Page404,
      footer: AppFooter
    }
  }

]

const router = new VueRouter({
  routes
})

export default router
