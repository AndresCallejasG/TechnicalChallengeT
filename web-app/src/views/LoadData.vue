<template>
  <v-main>
    <v-container>
      <v-row>
      <v-col cols=7>
        <v-img
          src="https://i.imgur.com/wIoy27u.png"
          class="my-3"
          contain
          height="500"
        />
      </v-col>
       <v-spacer></v-spacer>
      <v-col cols=4 align="center">
          <div>
          <v-date-picker class="mt-1"
            block
            v-model="picker"
            elevation="15"
          ></v-date-picker>        
          <v-btn
            width="100%"
            color="primary"
            class="mt-5 white--text"
          >
          Upload
            <v-icon
              right
              dark
            >
              mdi-cloud-upload
            </v-icon>
          </v-btn>
          </div>
      </v-col>
      </v-row>
    </v-container>
  </v-main>
</template>

<script>
// @ is an alias to /src

export default {
  name: 'LoadData',
  data () {
      return {
        picker: null,
      }
    },
}
</script>
