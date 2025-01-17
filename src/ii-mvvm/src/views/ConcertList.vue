<template>
  <div class="concert-list">
    <h2>School Concert Events</h2>
    <div v-for="concert in availableConcerts" :key="concert.id" class="concert-card">
      <h3>{{ concert.name }}</h3>
      <p>{{ concert.date }} at {{ concert.venue }}</p>
      <p>Available Tickets: {{ concert.availableTickets }}</p>
      <p>Price: ${{ concert.ticketPrice }}</p>
      <button @click="viewDetails(concert.id)">View Details</button>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
  name: 'ConcertList',
  computed: {
    ...mapGetters(['availableConcerts'])
  },
  methods: {
    viewDetails(id) {
      this.$router.push({ name: 'ConcertDetails', params: { id: id.toString() } })
    }
  }
}
</script>

<style scoped>
.concert-list {
  padding: 20px;
}
.concert-card {
  border: 1px solid #ddd;
  border-radius: 8px;
  padding: 15px;
  margin-bottom: 15px;
}
button {
  background-color: #4CAF50;
  color: white;
  padding: 8px 16px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}
</style>