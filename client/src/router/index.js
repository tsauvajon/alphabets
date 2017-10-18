import Vue from 'vue'
import Router from 'vue-router'

import HelloWorld from '@/components/HelloWorld'
import Login from '@/components/Login'
import Index from '@/components/pages/Index'
import ResultsPage from '@/components/pages/ResultsPage'
import Team from '@/components/pages/Team'
import Bet from '@/components/pages/Bet'
import FixtureDetail from '@/components/pages/FixtureDetail'
import Standing from '@/components/pages/Standing'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Index',
      component: Index
    },
    {
      path: '/results',
      name: 'Results',
      component: ResultsPage
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
    },
    {
      path: '/fixtures/:id',
      name: 'Fixture',
      component: FixtureDetail
    },
    {
      path: '/standing/:saisonid',
      name: 'Standing',
      component: Standing
    }
  ]
})
