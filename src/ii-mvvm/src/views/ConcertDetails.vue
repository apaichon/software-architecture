<template>
  <div class="concert-details" v-if="selectedConcert">
    <h2>{{ selectedConcert.name }}</h2>
    <div class="concert-info">
      <p>Date: {{ selectedConcert.date }}</p>
      <p>Venue: {{ selectedConcert.venue }}</p>
      <p>Available Tickets: {{ selectedConcert.availableTickets }}</p>
      <p>Price: ${{ selectedConcert.ticketPrice }}</p>
    </div>

    <button v-if="!showForm" @click="showForm = true" class="purchase-btn">
      Purchase Ticket
    </button>

    <form v-else @submit.prevent="submitPurchase" class="purchase-form">
      <div class="form-group">
        <label>Student Name:</label>
        <input v-model="studentName" required>
      </div>
      <div class="form-group">
        <label>Class:</label>
        <input v-model="studentClass" required>
      </div>
      <button type="submit">Confirm Purchase</button>
    </form>
  </div>
</template>

<script>
import { mapState, mapActions } from 'vuex'
import { Ticket } from '../models/Ticket'

export default {
  name: 'ConcertDetails',
  props: {
    id: {
      type: String,
      required: true
    }
  },
  data() {
    return {
      showForm: false,
      studentName: '',
      studentClass: ''
    }
  },
  computed: {
    ...mapState(['selectedConcert'])
  },
  methods: {
    ...mapActions(['selectConcert', 'purchaseTicket']),
    async submitPurchase() {
      const ticket = new Ticket(
        parseInt(this.id),
        this.studentName,
        this.studentClass
      )
      await this.purchaseTicket(ticket)
      this.$router.push('/')
    }
  },
  created() {
    this.selectConcert(parseInt(this.id))
  }
}
</script>

<style scoped>
.concert-details {
  padding: 20px;
}
.purchase-form {
  margin-top: 20px;
}
.form-group {
  margin: 10px 0;
}
input {
  width: 100%;
  padding: 8px;
  margin: 5px 0;
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