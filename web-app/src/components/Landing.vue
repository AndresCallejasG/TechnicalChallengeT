<template>
  <v-container>
    <v-row class="text-center">
      <v-col cols="12">
        <v-img
          src="https://i.imgur.com/RRelP7u.png"
          class="my-3"
          contain
          height="200"
        />
      </v-col>

      <v-col class="mb-4">
        <h1 class="display-2 font-weight-bold mb-3">
          Welcome to TheShop
        </h1>

        <p class="subheading font-weight-regular">
          This website let to visualize data about customers, products and transactions,
          <br>using a GraphQL API. Feel free explore.
        </p>
        <p class="subheading font-weight-regular">
          You can also load data from previous days to the Dgraph database using a
          <br>REST API made in Golang that works as an ETL.
        </p>
      </v-col>

      <v-col
        class="mb-5"
        cols="12"
      >
        <h2 class="headline font-weight-bold mb-3">
          What's Dgraph
        </h2>

        <v-row justify="center">
          <a
            v-for="(elem, i) in Dgraph"
            :key="i"
            :href="elem.href"
            class="subheading mx-3"
            target="_blank"
          >
            {{ elem.text }}
          </a>
        </v-row>
      </v-col>

      <v-col
        class="mb-5"
        cols="12"
      >
        <h2 class="headline font-weight-bold mb-3">
         GraphQL
        </h2>

        <v-row justify="center">
          <a
            v-for="(elem, i) in GraphQL"
            :key="i"
            :href="elem.href"
            class="subheading mx-3"
            target="_blank"
          >
            {{ elem.text }}
          </a>
        </v-row>
      </v-col>

      <v-col
        class="mb-5"
        cols="12"
      >
        <h2 class="headline font-weight-bold mb-3">
          TheShop Project
        </h2>

        <v-row justify="center">
          <a
            v-for="(elem, i) in project"
            :key="i"
            :href="elem.href"
            class="subheading mx-3"
            target="_blank"
          >
            {{ elem.text }}
          </a>
        </v-row>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
  export default {
    name: 'HelloWorld',

    data: () => ({
      project: [
        {
          text: 'Repository',
          href: 'https://github.com/AndresCallejasG/TechnicalChallengeT',
        },
        {
          text: 'Architecture',
          href: 'https://i.imgur.com/tR9ybwU.png',
        },        
        {
          text: 'Wireframe',
          href: 'https://i.imgur.com/4jTJEV9.png',
        },
        {
          text: 'DGraph Schema',
          href: 'https://i.imgur.com/oHoXg4R.png',
        },
      ],
      GraphQL: [
        {
          text: 'GraphQL Docs',
          href: 'https://graphql.org/',
        },
        {
          text: 'How to GraphQL',
          href: 'https://www.howtographql.com/',
        },
      ],
      Dgraph: [
        {
          text: 'Dgraph Docs',
          href: 'https://dgraph.io/docs',
        },
        {
          text: 'GraphQL API',
          href: 'https://dgraph.io/docs/graphql/',
        }, 
        {
          text: 'Dgraph Tour',
          href: 'https://dgraph.io/tour/',
        },
      ],
    }),
  }
</script>
