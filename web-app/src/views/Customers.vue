<template>
  <v-main class="pt-8">
    <v-container
      fluid     
    >
      <v-row>
        <v-col
          cols="2"
        >
          <h3>CUSTOMERS</h3>
          <v-divider></v-divider> 
          <div 
            style="max-height: 75vh; overflow-y: scroll">
            <CustomersList/>
          </div>
          <v-divider></v-divider> 
        </v-col>
        <v-divider vertical></v-divider>
        <v-col cols="7">
          <h3>TRANSACTIONS HISTORY</h3>
          <v-divider></v-divider> 
          <div 
            style="max-height: 75vh; overflow-y: scroll">
            <TransactionsList/>
          </div>
        <v-divider></v-divider> 
        </v-col>
        <v-divider vertical></v-divider>         
        <v-col cols="2.5">
          <h3>PRODUCTS</h3>
          <v-divider></v-divider>  
          <div
            style="max-height: 40vh; overflow-y: scroll">
            <ProductsList/>
          </div>
          <v-divider class="mb-2"></v-divider>
          <h4>CUSTOMERS WITH THE SAME IP</h4>
          <v-divider></v-divider> 
          <div 
            style="max-height: 30vh; overflow-y: scroll">
            <SameIpList/>
          </div>
          <v-divider></v-divider> 
        </v-col>
      </v-row>
    </v-container>
  </v-main>
</template>

<script>

import CustomersList from '@/components/CustomersList.vue'
import TransactionsList from '@/components/TransactionsList.vue'
import ProductsList from '@/components/ProductsList.vue'
import SameIpList from '@/components/SameIpList.vue'

export default {
  name: 'Customers',
  components: {
    CustomersList,
    TransactionsList,
    ProductsList,
    SameIpList
  }
}
</script>

<style scoped>
  ::-webkit-scrollbar {
    display: none;
  }
  h3{
    color: #2487f8;
  }
  h4{
    color: #2487f8;
  }
</style>