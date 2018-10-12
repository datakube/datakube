<script>
  export default {
    methods: {
      getIcon: function (status) {
        switch (status) {
          case 'queued':
            return 'watch_later'
          case 'in_progress':
            return 'yellow'
          case 'error':
            return 'error'
          case 'success':
            return 'done'
          default:
            return 'primary'
        }
      },
      getStatusColour: function (status) {
        switch (status) {
          case 'queued':
            return 'grey'
          case 'in_progress':
            return 'yellow'
          case 'error':
            return 'red'
          case 'success':
            return 'green'
          default:
            return 'primary'
        }
      }
    }
}
</script>
