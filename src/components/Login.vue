<template>
<div class="container">

   <br><br><br>
<v-app>
  <v-card>
    
  <v-card-text>
    <v-form>
      <v-text-field
        v-model="userName"
        label="User Name"
        name="UserName"
        prepend-icon="mdi-emoticon"
        type="text"
      ></v-text-field>
      <v-text-field
        v-model="password"
        id="password"
        label="Password"
        name="password"
        prepend-icon="mdi-lock"
        type="password"
      ></v-text-field>
    </v-form>
  </v-card-text>
  <v-card-actions>
    <v-spacer></v-spacer>
    <v-btn   @click="login()"   color="red lighten-2" dark>login</v-btn>
  </v-card-actions>
  </v-card>
<!---<Header :control= "control" ></Header>---->
</v-app>

</div>
</template>



<script>
import axios from  "axios"
///import Header from './Header';


export default {
 /*components:{
   Header,
 },*/

      name: 'Login',
    data(){
      return{
        
        control:"",
        userName : "",
        password : ""
        
      }

    },
    methods : {
      
      login(){  
      
        axios.post("http://localhost:8070/login",{
    userName:this.userName,
    password:this.password
  })
        .then(response => {
          console.log(response.data)
         

          if (response.data == false) {
            alert("kullanıcı adı veya parola yanlış")
           // alert(this.control)
          }else{
            //const control1 = this.control
            //this.$router.push('/');
            this.control=response.data
          }
          
        })
        .catch(e =>{
          console.log(e)
        })
          //this.$emit('login',this.response.data)

      }

    }
  
    
}
</script>