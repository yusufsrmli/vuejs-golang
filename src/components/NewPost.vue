<template>
    <div class="container">
        <v-app>
            <v-card>
                <v-card-text>
        <v-form>
            <v-divider></v-divider>
            <v-text-field
            v-model="subject"
            label="Subject"
           
            single-line
            hide-details 
            ></v-text-field>
            <v-divider></v-divider>
            <v-textarea
            v-model="message"
                label="Message"
                counter
                
                single-line
            ></v-textarea>
            


        </v-form>
                </v-card-text>
                <v-card-actions>
                    <v-spacer></v-spacer>
    <v-btn color="red lighten-2" dark  @click="post()">Post</v-btn>
                </v-card-actions>
            </v-card>
        </v-app>
    </div>
</template>


<script> 
import axios from  "axios"
export default {
   
    name: 'NewPosts',
     data(){
      return{
        subject:"",
        message:""
        
      }

    },
     methods : {
      
     post(){  
    axios.post("http://localhost:8070/addpost",{
    subject:this.subject,
    message:this.message,
    
  })
    this.$router.push('/')
}
     }}
     
</script>