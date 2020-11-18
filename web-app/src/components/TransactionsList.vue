<template>
    <v-component>
        TransactionList
        <v-card
            v-for="(tran, i) in transactions"
            :key="i"
            elevation="2"
            outlined
        >
                <v-card-title>Transaction {{ tran.transactionID }}</v-card-title>
                <v-card-subtitle>{{tran.ip}} - {{ tran.device }}</v-card-subtitle>
                <v-card-actions>
                    <v-btn
                        color="orange lighten-2"
                        text
                    >
                        Products
                    </v-btn>

                    <v-spacer></v-spacer>

                    <v-btn
                        icon
                        @click="show = !show"
                    >
                        <v-icon>{{ show ? 'mdi-chevron-up' : 'mdi-chevron-down' }}</v-icon>
                    </v-btn>
                </v-card-actions>
                <v-expand-transition>
                    <div v-show="show">
                        <v-divider></v-divider>
                        <v-list>
                            <v-list-item
                                v-for="(prod, j) in tran.products"
                                :key="j"
                            >
                            {{prod.name}} - {{prod.price}}
                            </v-list-item>
                        </v-list>
                    </div>
                </v-expand-transition>
        </v-card>
    </v-component>    
</template>

<script>


export default {
  name: 'TransactionsList',
  data: () => ({
      show: false,
      transactions : [
          {transactionID: "121212", ip: "129.123.44.234", device: "Linux", 
          products: [{name:"CocaCola", price: 5}, {name:"Juice", price: 10}, {name:"Best Item ever", price: 200}]},
          {transactionID: "121212", ip: "129.123.44.234", device: "Linux",
          products: [{name:"CocaCola", price: 5}, {name:"Juice", price: 10}]},
          {transactionID: "121212", ip: "129.123.44.234", device: "Linux", 
          products: [{name:"CocaCola", price: 5}, {name:"Juice", price: 10}, {name:"Best Item ever", price: 200}]},
          {transactionID: "121212", ip: "129.123.44.234", device: "Linux", 
          products: [{name:"CocaCola", price: 5}, {name:"Juice", price: 10}, {name:"Best Item ever", price: 200}]},
          {transactionID: "121212", ip: "129.123.44.234", device: "Linux", 
          products: [{name:"CocaCola", price: 5}, {name:"Juice", price: 10}, {name:"Best Item ever", price: 200}]},
          
      ]
  }),
  methods: {
      fetchCustomers(){
          
      }
  }
  
}
</script>
