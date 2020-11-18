<template>
  <v-main class="pt-5">
    <v-container
      fluid     
    >
      <v-row >
        <v-col
          cols="2" 
        >
          <div 
            style="max-height: 80vh; overflow-y: scroll">
            <CustomersList/>
          </div>
          
        </v-col>
        <v-divider vertical></v-divider>
        <v-col cols="7">
          <div 
            style="max-height: 80vh; overflow-y: scroll">
            <TransactionsList/>
          </div>
        </v-col>
        <v-divider vertical></v-divider>         
        <v-col cols="2.5">
          <div 
            style="max-height: 80vh; overflow-y: scroll">
            <ProductsList/>
          </div>
        </v-col>
      </v-row>
    </v-container>
  </v-main>
</template>

<script>

import CustomersList from '@/components/CustomersList.vue'
import TransactionsList from '@/components/TransactionsList.vue'
import ProductsList from '@/components/ProductsList.vue'

export default {
  name: 'Customers',
  components: {
    CustomersList,
    TransactionsList,
    ProductsList
  }
}
</script>

<style scoped>
  ::-webkit-scrollbar {
    display: none;
  }
</style>