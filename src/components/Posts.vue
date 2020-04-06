<template>
<div class="container">
    <v-app>
        <v-list three-line>
            <template v-for="(item, index) in content">
                
                <v-divider
               
                :key="index"
                :inset="item.inset"
                ></v-divider>
                <v-list-item
              
                :key="item.Title"
                >
            
                <v-list-item-content>
                     <v-list-item-title v-html="item.Title"></v-list-item-title>
                      <v-list-item-subtitle  v-html="item.Subtitle"></v-list-item-subtitle>
                </v-list-item-content>     
              <v-dialog
              v-model="dialog"
              width="500"
            >
        <template v-slot:activator="{ on }">
          <v-btn
            color="red lighten-2"
            dark
            v-on="on"
          >
            Detaylar
          </v-btn>
        </template>
  
        <v-card>
          <v-card-title
            class="headline grey lighten-2"
            primary-title
            v-html="item.Title"
          >
           </v-card-title>
  
          <v-card-text v-html="item.Subtitle">
           
          </v-card-text>
  
          <v-divider></v-divider>
  
          <v-card-actions>
            <v-spacer></v-spacer>
           <!--- <v-btn
              color="primary"
              text
              @click="dialog = false"
            >
              Kapat
            </v-btn>--->
          </v-card-actions>
        </v-card>
      </v-dialog>
                </v-list-item>
            </template>
        </v-list>

    </v-app>
</div>
</template>


<script>
import axios from 'axios'
export default {
    name: 'Posts',
 data(){

      return{
        content:[
          
        ] , 
          
     items: [
         
        
     
      {
       title: 'Brunch this weekend?',
        subtitle: "<span class='text--primary'>Ali Connors</span> &mdash; I'll be in your neighborhood doing errands this weekend. Do you Lorem Ipsum, dizgi ve baskı endüstrisinde kullanılan mıgır metinlerdir. Lorem Ipsum, adı bilinmeyen bir matbaacının bir hurufat numune kitabı oluşturmak üzere bir yazı galerisini alarak karıştırdığı 1500'lerden beri endüstri standardı sahte metinler olarak kullanılmıştır. Beşyüz yıl boyunca varlığını sürdürmekle kalmamış, aynı zamanda pek değişmeden elektronik dizgiye de sıçramıştır. 1960'larda Lorem Ipsum pasajları da içeren Letraset yapraklarının yayınlanması ile ve yakın zamanda Aldus PageMaker gibi Lorem Ipsum sürümleri içeren masaüstü yayıncılık yazılımları ile popüler olmuştLorem Ipsum, dizgi ve baskı endüstrisinde kullanılan mıgır metinlerdir. Lorem Ipsum, adı bilinmeyen bir matbaacının bir hurufat numune kitabı oluşturmak üzere bir yazı galerisini alarak karıştırdığı 1500'lerden beri endüstri standardı sahte metinler olarak kullanılmıştır. Beşyüz yıl boyunca varlığını sürdürmekle kalmamış, aynı zamanda pek değişmeden elektronik dizgiye de sıçramıştır. 1960'larda Lorem Ipsum pasajları da içeren Letraset yapraklarının yayınlanması ile ve yakın zamanda Aldus PageMaker gibi Lorem Ipsum sürümleri içeren masaüstü yayıncılık yazılımları ile popüler olmuştur.ur.want to hang out?",
        
      },
      
      {
        
        title: 'Summer BBQ <span class="grey--text text--lighten-1">4</span>',
        subtitle: "<span class='text--primary'>to Alex, Scott, Jennifer</span> &mdash; Wish I could come, but I'm out of town this weekend.",
      },
  
      {
      
        title: 'Oui oui',
        subtitle: "<span class='text--primary'>Sandra Adams</span> &mdash; Do you have Paris recommendations? Have you ever been?",
      },
      
      {
        
        title: 'Birthday gift',
        subtitle: "<span class='text--primary'>Trevor Hansen</span> &mdash; Have any ideas about what we should get Heidi for her birthday?",
      },
      
      {
       
        title: 'Recipe to try',
        subtitle: "<span class='text--primary'>Britta Holt</span> &mdash; We should eat this: Grate, Squash, Corn, and tomatillo Tacos.",
      },
    ],
 
        
      }

    },
    created(){
        axios.get("http://localhost:8070/posts")
        .then(response=>{
         /*for(let i in response.data){
           
          }
          */
        
        //console.log(data) 
        
        //this.content = response.data
       this.content.push(response.data)
        console.log(response.data)
        
        }
        
        )


    }
    
}
</script> 
    