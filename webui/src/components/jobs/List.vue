<template>
    <v-data-table
        :headers="headers"
        :items="jobs"
        :rows-per-page-items="[10,25,50,{text:'All','value':-1}]"
        class="elevation-1"
    >
        <template slot="items" slot-scope="job">
        <td>{{ job.item.ID }}</td>
        <td>{{ job.item.provider }}</td>
        <td class="text-xs-left"><v-icon :color="getStatusColour(job.item.state)">{{ getIcon(job.item.state) }}</v-icon></td>
        <td class="text-xs-left">{{ job.item.runAt }}</td>
        <td class="text-xs-right">
            <v-btn flat disabled icon color="grey">
            <v-icon>delete</v-icon>
            </v-btn>
        </td>
        </template>
    </v-data-table>
</template>

<script>
  import status from '@/mixins/JobStatus'

  export default {
    name: 'JobsList',
    mixins: [status],
    props: ['jobs'],
    data () {
      return {
        headers: [
          {
            text: '#',
            align: 'left',
            sortable: false,
            value: 'ID'
          },
          { text: 'Target', value: 'provider' },
          { text: 'Status', value: 'state' },
          { text: 'Run', value: 'runAt' }
        ]
      }
    }
  }
</script>