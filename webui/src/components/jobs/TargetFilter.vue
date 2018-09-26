<template>
    <v-select
        :items="targets"
        v-model="filteredTargets"
        multiple
        small-chips
        item-text="name"
        label="Filter Jobs by Target"
        @change="$emit('filterTriggered', filteredTargets)"
    >
    </v-select>
</template>

<script>
  import JobsFilter from '@/components/jobs/BaseFilter'

  export default {
    extends: JobsFilter,
    name: 'TargetFilter',
    props: ['targets'],
    data () {
      return {
        filteredTargets: []
      }
    }
  }
</script>