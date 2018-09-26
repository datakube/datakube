<template>
    <v-select
        :items="targets"
        v-model="value"
        multiple
        small-chips
        label="Select Target"
        @change="$emit('filterTriggered', value)"
    ></v-select>
</template>

<script>
  export default {
    name: 'TargetsFilter',
    mixins: [status],
    props: ['targets'],
    data () {
      return {
        value: []
      }
    }
  }
</script>