export default [{
  path: '/dashboard',
  component: resolve => {
    import('../view/Dashboard.vue').then(resolve)
  }
}, {
  path: '/app-create',
  component: resolve => {
    import('../view/NewApp.vue').then(resolve)
  }
}, {
  path: '/apps/:name',
  component: resolve => {
    import('../view/app/App.vue').then(resolve)
  },
  children: [{
    path: 'overview',
    component: resolve => {
      import('../view/app/Overview.vue').then(resolve)
    }
  }, {
    path: 'resources',
    component: resolve => {
      import('../view/app/Resources.vue').then(resolve)
    }
  }, {
    path: 'deploy',
    component: resolve => {
      import('../view/app/Deploy.vue').then(resolve)
    }
  }, {
    path: 'activity',
    component: resolve => {
      import('../view/app/Activity.vue').then(resolve)
    }
  }, {
    path: 'logs',
    component: resolve => {
      import('../view/app/Logs.vue').then(resolve)
    }
  }, {
    path: 'settings',
    component: resolve => {
      import('../view/app/Settings.vue').then(resolve)
    }
  }, {
    path: '/',
    redirect: 'overview'
  }]
}]
