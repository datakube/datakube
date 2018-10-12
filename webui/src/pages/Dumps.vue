<template>
  <v-layout row>
    <v-flex xs12 sm12>
      <v-toolbar>
        <target-filter v-bind:targets="targets" @filterTriggered="filterDumps"></target-filter>
      </v-toolbar>
      <v-divider class="mt-1"></v-divider>
      <v-data-table
          :headers="headers"
          :items="filteredDumps"
          :rows-per-page-items="[10,25,50,{text:'All','value':-1}]"
          class="elevation-1"
      >
          <template slot="items" slot-scope="dump">
          <td>{{ dump.item.ID }}</td>
          <td>{{ dump.item.Target }}</td>
          <td class="text-xs-left">{{ dump.item.File.name }}</td>
          <td class="text-xs-left">{{ dump.item.CreatedAt }}</td>
          <td class="text-xs-left">
            <v-btn small flat color="primary" v-on:click="downloadFile(dump.item.File.name)" class="ml-0 pl-0">
              <v-icon>cloud_download</v-icon>
            </v-btn>
            <v-btn small flat disabled icon color="grey" class="ml-0 pl-0">
              <v-icon>delete</v-icon>
            </v-btn>
          </td>
          </template>
      </v-data-table>
    </v-flex>
  </v-layout>
</template>

<script>
  import axios from 'axios'
  import filter from '@/components/targets/Filter'

  export default {
    name: 'Dumps',
    components: {
      'target-filter': filter
    },
    data () {
      return {
        headers: [
          {
            text: '#',
            align: 'left',
            sortable: false,
            value: 'ID'
          },
          { text: 'Target', value: 'deadline' },
          { text: 'File', value: 'deadline' },
          { text: 'CreatedAt', value: 'progress' },
          { text: 'Actions', value: 'status' }
        ],
        dumps: this.getDumps(),
        filteredDumps: [],
        targets: []
      }
    },
    methods: {
      getDumps () {
        axios.get('/api/dumps/', {
        })
        .then(response => {
          this.dumps = response.data.dumps
          this.filteredDumps = this.dumps
          this.targets = this.getTargetsFromDumps()
        })
      },
      filterDumps (targets) {
        let dumps = this.dumps.filter(dump => {
          return targets.includes(dump.Target)
        })

        this.filteredDumps = dumps
      },
      getTargetsFromDumps () {
        let targets = []
        this.dumps.forEach(function (dump) {
          if (targets.indexOf(dump.Target) === -1) {
            targets.push(dump.Target)
          }
        })
        return targets
      },
      downloadFile (filename) {
        window.open('/files/download/' + filename, '_blank')
      }
    },
    computed: {
    }
  }
</script>