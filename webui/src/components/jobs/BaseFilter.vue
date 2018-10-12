<template>
    <v-toolbar>
        <v-flex sm12 md4>
            <target-filter
                title="Filter by Target"
                v-bind:targets="targets"
                v-bind:jobs="jobs"
                @filterTriggered="filteredTargets = $event"
            ></target-filter> 
        </v-flex>
        <v-flex sm12 md4>
            <state-filter
                title="Filter by State"
                v-bind:states="states"
                v-bind:jobs="jobs"
                @filterTriggered="filteredStates = $event"
            ></state-filter>
        </v-flex>
        <v-flex sm12 md2>
            <date-filter
                label="From Date"
                v-bind:date="fromDate"
                @dateSelected="fromDate = $event"
            ></date-filter>
        </v-flex>
        <v-flex sm12 md2>
            <date-filter
                label="To Date"
                v-bind:date="toDate"
                @dateSelected="toDate = $event"
            ></date-filter>
        </v-flex>
    </v-toolbar>
</template>

<script>
  import stateFilter from '@/components/jobs/StateFilter'
  import targetFilter from '@/components/jobs/TargetFilter'
  import dateFilter from '@/components/jobs/DateFilter'

  export default {
    name: 'JobsFilter',
    props: ['jobs'],
    components: {
      'state-filter': stateFilter,
      'target-filter': targetFilter,
      'date-filter': dateFilter
    },
    data () {
      return {
        targets: [],
        filteredTargets: [],
        states: [],
        filteredStates: [],
        filters: [stateFilter, targetFilter],
        fromDate: null,
        toDate: null
      }
    },
    watch: {
      jobs: function () {
        this.states = this.filterJobsForStates()
        this.targets = this.filterJobsForTargets()
      },
      filteredTargets: function () {
        this.filterJobs()
      },
      filteredStates: function () {
        this.filterJobs()
      },
      fromDate: function () {
        this.filterJobs(true)
      },
      toDate: function () {
        this.filterJobs(true)
      }
    },
    methods: {
      filterJobs: function (isDateFilter) {
        let filteredJobs = this.jobs

        if (this.filteredTargets.length > 0) {
          filteredJobs = filteredJobs.filter(job => {
            return this.filteredTargets.includes(job.provider)
          })
        }

        if (this.filteredStates.length > 0) {
          filteredJobs = filteredJobs.filter(job => {
            return this.filteredStates.includes(job.state)
          })
        }

        if (isDateFilter) {
          filteredJobs = filteredJobs.filter(job => {
            let date = job.runAt.split('T')

            let jobDate = new Date(date[0])
            let fromDate = new Date(this.fromDate)
            let toDate = new Date(this.toDate)

            if ((this.fromDate != null && jobDate < fromDate) || (this.toDate != null && jobDate > toDate)) {
              return false
            }
            return true
          })
        }

        this.$emit('jobsFiltered', filteredJobs)
      },
      filterJobsForTargets: function () {
        let targets = []
        this.jobs.forEach(function (job) {
          if (targets.indexOf(job.provider) === -1) {
            targets.push(job.provider)
          }
        })
        return targets
      },
      filterJobsForStates: function () {
        let states = []
        this.jobs.forEach(function (job) {
          if (states.indexOf(job.state) === -1) {
            states.push(job.state)
          }
        })
        return states
      }
    }
  }
</script>