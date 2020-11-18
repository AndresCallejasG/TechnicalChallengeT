import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import LoadData from '../views/LoadData.vue'
import Customers from '../views/Customers.vue'
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
    path: '/customers',
    name: 'Customers',
    components: {
      header: AppHeader,
      default: Customers,
      footer: AppFooter
    }
  },
  {
    path: '/load-data',
    name: 'LoadData',
    components: {
      header: AppHeader,
      default: LoadData,
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
