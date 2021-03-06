import Vue from 'vue'
import Router from 'vue-router'

import HelloWorld from '@/components/HelloWorld'
import Index from '@/components/pages/Index'
import ResultsPage from '@/components/pages/ResultsPage'
import Team from '@/components/pages/Team'
import Bet from '@/components/pages/Bet'
import Bets from '@/components/pages/Bets'
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
      path: '/hello',
      name: 'HelloWorld',
      component: HelloWorld
    },
    {
      path: '/bets/:id',
      name: 'Bet',
      component: Bet
    },
    {
      path: '/mybets',
      name: 'Bets',
      component: Bets
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
