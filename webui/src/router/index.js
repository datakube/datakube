import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'
import Target from '@/pages/Targets'
import Jobs from '@/pages/Jobs'
import Dumps from '@/pages/Dumps'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'HelloWorld',
      component: HelloWorld
    },
    {
      path: '/targets',
      name: 'Targets',
      component: Target
    },
    {
      path: '/jobs',
      name: 'Jobs',
      component: Jobs
    },
    {
      path: '/dumps',
      name: 'Dumps',
      component: Dumps
    }
  ]
})
