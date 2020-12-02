<template>
    <v-component>
        <v-skeleton-loader
          v-show="loadingCustomers"
          type="card-heading, list-item-three-line, card-heading, list-item-three-line"
        >
        </v-skeleton-loader>
        <v-card
            v-show="emptyArray">
            <v-card-subtitle >You may load some data</v-card-subtitle>
        </v-card>  
        <v-card
            v-for="(card, i) in customers"
            :key="i"
            elevation="2"
            outlined
            @click="setSelectedCustomer(card.customerID)"
        >
            <v-card-title >{{ card.name }}</v-card-title>
            <v-card-subtitle>ID {{card.customerID}} - {{ card.age }} years</v-card-subtitle>

        </v-card>        
    </v-component>
    
</template>

<script>


export default {
  name: 'CustomersList',
  data: () => ({
      customers : [],
      loadingCustomers: false,
      emptyArray : false
  }),
  created(){
      this.fetchCustomers()
  },
  methods: {
        async fetchCustomers(){
        
        this.loadingCustomers = true

        var axios = require('axios');
        var data = JSON.stringify({
        query: `query {
        queryCustomer {
            customerID
            name
            age
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

        this.customers = await axios(config)
        .then(function (response) {
            //var x = JSON.stringify(response.data)
            return response.data.data.queryCustomer;
        })
        .catch(function (error) {
        console.log(error);
        });

        if (this.customers.length) {
            this.emptyArray = false
        }else{
            this.emptyArray = true
        }
         
        this.loadingCustomers = false
      },
      setSelectedCustomer(customerID){
          this.$store.dispatch('SetSelectedCustomer', customerID)
          this.$store.dispatch('FetchSelectedTransactions')
      }
  }  
}
</script>
