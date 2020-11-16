<template>
  <v-app-bar
      app
      color="primary"
      dark
    >
      <div class="d-flex align-center">
        <v-img
          alt="TheShop Logo"
          class="shrink mr-2"
          contain
          src="https://i.imgur.com/RRelP7u.png"
          transition="scale-transition"
          width="40"
        />

        <v-img
          alt="TheShop Name"
          class="shrink mt-1 hidden-sm-and-down"
          contain
          min-width="100"
          src="https://i.imgur.com/CZTWePa.png"
          width="100"
        />
      </div>

      <v-spacer></v-spacer>
      <router-link to="/">
        <v-btn
          target="_blank"
          text
        >
          
          <span>Home</span>
        </v-btn>
      </router-link>
      <router-link to="/load-data">
        <v-btn
          target="_blank"
          text
        >
          <span>Load Data</span>
        </v-btn>
      </router-link>
      <router-link to="/customers">
        <v-btn
          target="_blank"
          text
        >
          <span>Dashboard</span>
        </v-btn>
      </router-link>
    </v-app-bar>
</template>

<script>

export default {
  name: "AppHeader"
}
</script>