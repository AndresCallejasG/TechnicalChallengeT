import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    selectedCustomerID: "",
    transactions: [],
    productRecommendations: [],
    sameIpCustomers: []
  },
  mutations: {
    SetSelectedCustomer(state, customerID){
      state.selectedCustomerID = customerID
    },
    SetTransactions(state, tranList){
      state.transactions = tranList
    },
    SetProductRecommendations(state, prodList){
      state.productRecommendations = prodList
    },
    SetSameIpCustomers(state, sameIpList){
      state.sameIpCustomers = sameIpList
    }
  },
  actions: {
    SetSelectedCustomer(context, customerID){
      context.commit('SetSelectedCustomer', customerID)
    },
    async FetchSelectedTransactions(context){
      var axios = require('axios');
      var data = JSON.stringify({
      query: `query ($custID: String!) {
      getCustomer(customerID: $custID){
        name
        transactions{
            ip
            transactionID
            device
            products{
                productID
                name
                price
            }
        }
      }
    }`,
      variables: {"custID": context.state.selectedCustomerID}
    });

    var config = {
      method: 'post',
      url: 'http://localhost:8080/graphql',
      headers: { 
        'Content-Type': 'application/json'
      },
      data : data
    };

      const trans = await axios(config)
      .then(function (response) {
        return response.data.data.getCustomer.transactions
      })
      .catch(function (error) {
        console.log(error);
      });
      
      context.commit('SetTransactions', trans)

    },
    async FetchProductRecommendations(context){
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

      const prods = await axios(config)
      .then(function (response) {
          //var x = JSON.stringify(response.data)
          return response.data.data.queryProduct;
      })
      .catch(function (error) {
      console.log(error);
      });

      context.commit('SetProductRecommendations', prods)
    },
    /* async FetchSameIp(context, ip){
      var axios = require('axios');
      
      context.commit('SetProductRecommendations', sameip)
    }, */

  },
  modules: {
  }
})
