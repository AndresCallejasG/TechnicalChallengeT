<template>
   <v-component>
       <v-skeleton-loader
          v-show="loading"
          type="card-heading, list-item-three-line, card-heading, list-item-three-line"
        >
        </v-skeleton-loader>        
        <v-card
            v-for="(prod, i) in products"
            :key="i"
            elevation="2"
            outlined
        >
            <v-card-title>{{ prod.name }}</v-card-title>
            <v-card-subtitle>ID {{prod.productID}} - {{ prod.price }} usd </v-card-subtitle>

        </v-card>
    </v-component>    
</template>

<script>


export default {
  name: 'ProductsList',
  data: () => ({
      products : [
          {productID: "25eaefff", name: "Organic broth", price: 444},
          {productID: "25eaefff", name: "Organic broth", price: 444},
          {productID: "25eaefff", name: "Organic broth", price: 444},
          {productID: "25eaefff", name: "Organic broth", price: 444},
          {productID: "25eaefff", name: "Organic broth", price: 444},
          {productID: "25eaefff", name: "Organic broth", price: 444},
          {productID: "25eaefff", name: "Organic broth", price: 444},        
      ]
  }),
  created() {
      this.$store.dispatch('FetchProductRecommendations')
  },
  computed: {
      products(){
          return this.$store.state.productRecommendations
      },
      loading(){
          var tr = this.$store.state.productRecommendations.length

          if (tr === 0) {
              return true
          }
          return false
      }
  },  
  methods: {
      
  }
  
}
</script>
