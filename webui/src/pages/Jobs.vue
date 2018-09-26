<template>
  <v-container container fluid grid-list-md>
  <v-layout row>
    <v-flex xs12 sm12>
      <jobs-filter
        v-bind:jobs="jobs"
        @jobsFiltered="filteredJobs = $event"
      ></jobs-filter>
      <v-divider class="mt-1"></v-divider>
      <jobs-list
        v-bind:jobs="filteredJobs"
      ></jobs-list>
    </v-flex>
  </v-layout>
  </v-container>
</template>

<script>
  import axios from 'axios'
  import status from '@/mixins/JobStatus'
  import jobList from '@/components/jobs/List'
  import filter from '@/components/jobs/BaseFilter'

  export default {
    name: 'Jobs',
    mixins: [status],
    components: {
      'jobs-filter': filter,
      'jobs-list': jobList
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
          { text: 'Status', value: 'progress' },
          { text: 'Run', value: 'status' }
        ],
        jobs: [],
        filteredJobs: [],
        fromDate: null,
        toDate: null,
        modal: null,
        modal2: null
      }
    },
    created () {
      this.getJobs().then(response => {
        this.jobs = response.data.jobs
        this.filteredJobs = this.jobs
      })
    },
    methods: {
      getJobs () {
        return axios.get('/api/jobs/', {
        })
      },
      updateJobs (jobs) {
        this.filteredJobs = jobs
      }
    }
  }
</script>