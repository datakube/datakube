<template>
    <v-select
        :items="states"
        v-model="filteredStates"
        multiple
        small-chips
        item-text="name"
        label="Filter Job By States"
        @change="$emit('filterTriggered', filteredStates)"
    >

    <template
      slot="selection"
      slot-scope="{ item, index }"
    >
      <v-chip label small :color="getStatusColour(item)" v-text="item"></v-chip>
    </template>
    </v-select>
</template>

<script>
  import status from '@/mixins/JobStatus'
  import JobsFilter from '@/components/jobs/BaseFilter'

  export default {
    extends: JobsFilter,
    name: 'StateFilter',
    props: ['states'],
    mixins: [status],
    data () {
      return {
        filteredStates: []
      }
    }
  }
</script>