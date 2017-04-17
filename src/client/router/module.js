export default [{
  path: '/dashboard',
  component: resolve => {
    import('../view/Dashboard.vue').then(resolve)
  }
}, {
  path: '/create',
  component: resolve => {
    import('../view/Create.vue').then(resolve)
  }
}, {
  path: '/app/:name',
  component: resolve => {
    import('../view/app/App.vue').then(resolve)
  },
  children: [{
    path: '/overview',
    component: resolve => {
      import('../view/app/Overview.vue').then(resolve)
    }
  }, {
    path: '/resources',
    component: resolve => {
      import('../view/app/Resources.vue').then(resolve)
    }
  }, {
    path: '/deploy',
    component: resolve => {
      import('../view/app/Deploy.vue').then(resolve)
    }
  }, {
    path: '/activities',
    component: resolve => {
      import('../view/app/Activities.vue').then(resolve)
    }
  }, {
    path: '/settings',
    component: resolve => {
      import('../view/app/Settings.vue').then(resolve)
    }
  }]
}]
