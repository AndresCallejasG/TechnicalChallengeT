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
  mounted() {
      this.fetchProducts()
  },
  methods: {
      async fetchProducts(){
        var axios = require('axios');
        var data = JSON.stringify({
        query: `query {
        queryProduct {
            productID
            name
            price
        }
        }`,
        variables: {}
        });

        var config = {
        method: 'post',
        url: 'http://localhost:8080/graphql',
        headers: { 
            'Content-Type': 'application/json'
        },
        data : data
        };

        this.products = await axios(config)
        .then(function (response) {
            //var x = JSON.stringify(response.data)
            return response.data.data.queryProduct;
        })
        .catch(function (error) {
        console.log(error);
        });
          
      }
  }
  
}
</script>
