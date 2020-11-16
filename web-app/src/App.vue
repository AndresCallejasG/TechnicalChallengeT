<template>
  <v-app>
    <v-main>
      <router-view vcol name="header"/>
      <router-view></router-view>
      <router-view name="footer"/>
    </v-main>
  </v-app>
</template>

<script>
//import HelloWorld from './components/HelloWorld';

export default {
  name: 'App',

  components: {
    //HelloWorld,
  },

  data: () => ({
    //
  }),
};
</script>
