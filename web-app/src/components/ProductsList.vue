<template>
   <v-component>        
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
  methods: {
      fetchProducts(){
          
      }
  }
  
}
</script>
