<template>
    <v-component>
        <v-card
            v-for="(ipCust, i) in ipCustomers"
            :key="i"
            elevation="2"
            outlined
        >
            <v-card-title>{{ ipCust.name }}</v-card-title>
            <v-card-subtitle>ID {{ipCust.customerID}} - {{ ipCust.age }} usd </v-card-subtitle>

        </v-card>
    </v-component>
    
</template>

<script>


export default {
  name: 'SameIPList',
  data: () => ({
      ipCustomers : [
          {customerID: "121212", name: "Julian Franco", age: "28"},
          {customerID: "288923", name: "Andres Calejas", age: "25"},
          {customerID: "288923", name: "Andres Calejas", age: "25"},  
          {customerID: "288923", name: "Andres Calejas", age: "25"},  
          {customerID: "288923", name: "Andres Calejas", age: "25"},      
      ]
  }),
  computed: {
      ipCustomers(){
          return this.$store.state.sameIpCustomers
      },
      loading(){
          var tr = this.$store.state.sameIpCustomers.length

          if (tr === 0) {
              return true
          }

          return false
      }
  } 
  
}
</script>
