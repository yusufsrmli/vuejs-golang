<template>
<v-app>
  
 <v-navigation-drawer
    v-model="drawer"
    :color="color"
    :expand-on-hover="expandOnHover"
    :mini-variant="miniVariant"
    :right="right"
    stateless
    dark
    app>    
    <v-list
    dense
    nav
    class="py-0"
    >
    <br>
    <br>
    <v-divider></v-divider>
    <v-list-item
      v-for="item in items"
      :key="item.title"
      :to="item.link"
      link
      >
      <v-list-item-icon>
        <v-icon>{{item.icon}}</v-icon>
      </v-list-item-icon>
      <v-list-item-content>
        <v-list-item-title >{{item.title }}</v-list-item-title>
      </v-list-item-content>
      <br>
      <br>
      <br>
    </v-list-item>
    </v-list>
  </v-navigation-drawer>
  <v-app-bar
    app
    dark
    :color="color"
  >
  
  <v-spacer/>
  <v-app-bar-title>{{title}}</v-app-bar-title >
   <v-spacer/>
   <div class="logout" v-if="control"  @click="login()" ><v-btn light>Logout<v-icon>{{"fas fa-sign-out-alt"}}</v-icon></v-btn></div>
  </v-app-bar>
 <v-content>
      <v-container
        fluid
        fill-height
      >
        <v-layout
          align-center
          justify-center
        >
         <router-view></router-view>
        
        </v-layout>
      </v-container>
 </v-content>

  </v-app>
</template>

<script>
//import axios from  "axios"

  export default {
    props: [ 'control'],
    data(){
      
      return{
       // control1:this.control,
        title:'Blog',
        drawer:true,
        items:[
          {title: 'Ana Sayfa', icon: 'fas fa-home', link:"/"},
          {title: 'Giriş Yap', icon: 'fas fa-lock', link:"/login"},
          {title:  'yazı ekle', icon: 'fas fa-edit',link:"/newposts"}
          
        ],
        color: 'blue-grey darken-2',
        right: false,
        miniVariant: true,
        expandOnHover: false,
        background: false,
        
      }
},
    methods : {
      login (){
        
          
         
       
        alert(this.control)
      
      }

    }
    
}
</script>
