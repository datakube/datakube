<template>
     <v-dialog
        ref="dateDialog"
        persistent
        v-model="modal"
        lazy
        full-width
        width="290px"
        :return-value.sync="date"
    >
        <v-text-field
        slot="activator"
        :label="label"
        v-model="date"
        prepend-icon="event"
        readonly
        ></v-text-field>
        <v-date-picker v-model="date" scrollable>
            <v-spacer></v-spacer>
            <v-btn flat color="primary" @click="modal = false">Cancel</v-btn>
            <v-btn flat color="primary" @click="saveDate(date)">OK</v-btn>
        </v-date-picker>
    </v-dialog>
</template>

<script>
  import JobsFilter from '@/components/jobs/BaseFilter'

  export default {
    extends: JobsFilter,
    name: 'DateFilter',
    props: ['label', 'selectedDate'],
    data () {
      return {
        date: null,
        modal: false
      }
    },
    methods: {
      saveDate: function (date) {
        this.$refs.dateDialog.save(date)
        this.$emit('dateSelected', date)
      }
    }
  }
</script>