import Vue from 'vue'
import Router from 'vue-router'

import HelloWorld from '@/components/HelloWorld'
import Login from '@/components/Login'
import Index from '@/components/pages/Index'
import Team from '@/components/pages/Team'
import Bet from '@/components/pages/Bet'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Index',
      component: Index
    },
    {
      path: '/login',
      name: 'Login',
      component: Login
    },
    {
      path: '/hello',
      name: 'HelloWorld',
      component: HelloWorld
    },
    {
      path: '/bet',
      name: 'Bet',
      component: Bet
    },
    {
      path: '/teams/:id',
      name: 'Team',
      component: Team
    }
  ]
})
