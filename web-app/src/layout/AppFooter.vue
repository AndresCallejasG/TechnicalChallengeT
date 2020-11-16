<template>
   <v-footer
      dark
      padless
    >
      <v-card
        class="flex"
        flat
        tile
      >
        <v-card-title
          class="primary justify-center" 
        >
          <strong class="mr-8">Get connected with me on social networks!  </strong>
  
  
          <v-btn
            v-for="item in social"
            :key="item"
            class="mx-1"
            :href= "item.url"
            dark
            icon
          >
            <v-icon size="24px">
              {{ item.icon }}
            </v-icon>
          </v-btn>
        </v-card-title>
  
        <v-card-text class="py-2 white--text text-center">
          <strong>Andrés Callejas González</strong> — {{ new Date().getFullYear() }}
        </v-card-text>
      </v-card>
    </v-footer>
</template>

<script>

export default {
  name: "AppFooter",
  data: () => ({
      social: [ 
        { icon:'mdi-linkedin', url: "https://www.linkedin.com/in/andrescallejasg"},
        { icon:'mdi-github', url: "https://github.com/AndresCallejasG"},
        { icon:'mdi-twitter', url: "https://twitter.com/AndresCallejasG"}
      ],
    }),
}
</script>

