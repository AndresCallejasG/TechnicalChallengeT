<template>
    <v-component>
        SameIp
    </v-component>
    
</template>

<script>


export default {
  name: 'SameIPList',
  data: () => ({
      sameip : [
          {customerID: "121212", name: "Julian Franco", age: "28"},
          {customerID: "288923", name: "Andres Calejas", age: "25"},    
      ]
  }),
  methods: {
      fetchCustomersIP(){
          
      }
  }
  
}
</script>
