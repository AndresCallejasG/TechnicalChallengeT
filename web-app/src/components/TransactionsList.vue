<template>
    <v-component>
        <v-skeleton-loader
          v-show="loading"
          type="card-heading, list-item-three-line, card-heading, list-item-three-line"
        >
        </v-skeleton-loader>
        <v-skeleton-loader
          v-show="loading"
          type="card-heading, list-item-three-line, card-heading, list-item-three-line"
        >
        </v-skeleton-loader>  
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

//import { mapState } from 'vuex';

export default {
  name: 'TransactionsList',
  data: () => ({
      show: false
  }),
  computed: {
      transactions(){
          return this.$store.state.transactions
      },
      loading(){
          var tr = this.$store.state.transactions.length

          if (tr === 0) {
              return true
          }

          return false
      }
  }    
}
</script>
