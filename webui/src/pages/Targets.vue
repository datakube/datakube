<template>
  <v-layout row>
    <v-flex xs12 sm6>
        <v-container fluid grid-list-sm>
          <v-layout row wrap>
            <v-flex
                v-for="target in targets"
                v-bind:data="target"
                v-bind:key="target.Name"
                sm6
            >
                <v-card>         
                  <v-card-text>
                      <div slot="header"> {{target.Name}} <v-btn flat small :color="getScheduleColour(target.Schedule.Interval)">{{getCleanScheduleString(target.Schedule.Interval)}}</v-btn></div>
                      <v-btn flat small color="grey"> Host: {{ target.DBConfig.DatabaseHost }}</v-btn>
                      <v-btn flat small color="grey"> Database: {{ target.DBConfig.DatabaseName }}</v-btn>
                      <v-btn flat small color="grey"> Port: {{ target.DBConfig.DatabasePort }}</v-btn>
                      <v-btn flat small color="grey"> User: {{ target.DBConfig.DatabaseUserName }}</v-btn>
                      <v-btn flat small color="grey">Password: {{ target.DBConfig.DatabasePassword }}</v-btn>
                  </v-card-text>
                </v-card>
            </v-flex>
          </v-layout>
        </v-container>
    </v-flex>
  </v-layout>
</template>

<script>
  import axios from 'axios'

  export default {
    name: 'Targets',
    data () {
      return {
        show: false,
        targets: this.getTargets()
      }
    },
    methods: {
      getTargets () {
        axios.get('/api/targets/', {
        })
        .then(response => {
          console.log(response)
          this.targets = response.data.targets
        })
      },
      getScheduleColour: function (interval) {
        switch (interval) {
          case 'every_minute':
            return 'blue'
          case 'hourly':
            return 'green'
          case 'daily':
            return 'yelow'
          case 'weekly':
            return 'grey'
          case 'monthly':
            return 'orange'
          default:
            return 'accent'
        }
      },
      getCleanScheduleString: function (interval) {
        switch (interval) {
          case 'every_minute':
            return 'minutely'
          default:
            return 'interval'
        }
      }
    },
    computed: {
    }
  }
</script>