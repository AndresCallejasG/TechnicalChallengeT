<template>
    <v-component>        
        <v-card
            v-for="(card, i) in customers"
            :key="i"
            elevation="2"
            outlined
        >
            <v-card-title>{{ card.name }}</v-card-title>
            <v-card-subtitle>ID {{card.customerID}} - {{ card.age }} years</v-card-subtitle>

        </v-card>
    </v-component>
    
</template>

<script>


export default {
  name: 'CustomersList',
  data: () => ({
      customers : [
          {customerID: "121212", name: "Julian Franco", age: 28},
          {customerID: "288923", name: "Andres Callejas", age: 25},
          {customerID: "929833", name: "Alejandro Ramirez", age: 18},
          {customerID: "209370", name: "Santiago Castillo", age: 25},
          {customerID: "121212", name: "Julian Franco", age: 28},
          {customerID: "288923", name: "Andres Callejas", age: 25},
          {customerID: "929833", name: "Alejandro Ramirez", age: 18},
          {customerID: "209370", name: "Santiago Castillo", age: 25},
          {customerID: "121212", name: "Julian Franco", age: 28},
          {customerID: "288923", name: "Andres Callejas", age: 25},
          {customerID: "929833", name: "Alejandro Ramirez", age: 18},
          {customerID: "209370", name: "Santiago Castillo", age: 25},
          {customerID: "121212", name: "Julian Franco", age: 28},
          {customerID: "288923", name: "Andres Callejas", age: 25},
          {customerID: "929833", name: "Alejandro Ramirez", age: 18},
          {customerID: "209370", name: "Santiago Castillo", age: 25},
          {customerID: "121212", name: "Julian Franco", age: 28},
          {customerID: "288923", name: "Andres Callejas", age: 25},
          {customerID: "929833", name: "Alejandro Ramirez", age: 18},
          {customerID: "209370", name: "Santiago Castillo", age: 25},
      ]
  }),
    mounted() {
      this.fetchCustomers()
  },
  methods: {
      async fetchCustomers(){
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

      }
  }
  
}
</script>
