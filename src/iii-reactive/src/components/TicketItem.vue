<template>
  <div class="ticket-item">
    <h3>{{ ticket.type }}</h3>
    <div class="ticket-details">
      <p>Base Price: ${{ ticket.price }}</p>
      <p v-if="ticket.discount" class="discount">
        Discount: {{ ticket.discount }}% off
      </p>
      <p class="final-price">
        Final Price: ${{ calculateFinalPrice(ticket) }}
      </p>
    </div>
    <div class="inventory-status">
      <span :class="getTicketStatusClass(ticket)">
        {{ getTicketStatus(ticket) }}
      </span>
    </div>
    <button 
      @click="$emit('add-to-cart', ticket)" 
      :disabled="ticket.quantity === 0"
      :class="{ 'low-stock': ticket.quantity <= 5 }"
    >
      Add to Cart ({{ ticket.quantity }} left)
    </button>
  </div>
</template>

<script>
export default {
  name: 'TicketItem',
  props: {
    ticket: {
      type: Object,
      required: true
    },
    calculateFinalPrice: {
      type: Function,
      required: true
    },
    getTicketStatus: {
      type: Function,
      required: true
    },
    getTicketStatusClass: {
      type: Function,
      required: true
    }
  },
  emits: ['add-to-cart']
}
</script>

<style scoped>
.ticket-item {
  border: 1px solid #ddd;
  margin: 10px 0;
  padding: 15px;
  border-radius: 5px;
}

.ticket-details {
  margin: 10px 0;
}

.discount {
  color: #e44d26;
  font-weight: bold;
}

.final-price {
  font-size: 1.2em;
  font-weight: bold;
}

.inventory-status {
  margin: 10px 0;
}
</style> 