import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import VueRouter from 'vue-router';

Vue.use(VueRouter);


import Login from './components/Login.vue'
import NewPost from './components/NewPost.vue'
import Posts from './components/Posts.vue'

Vue.config.productionTip = false
const router = new VueRouter({
  routes: [
    {path: '/', component:Posts, meta:'home'},
    {path: '/login', component:Login},
    {path: '/newposts', component:NewPost}

  ],
  mode: 'history'
})
new Vue({
  vuetify,
  render: h => h(App),
  router: router
}).$mount('#app')
