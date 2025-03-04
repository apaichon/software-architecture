<!-- App.vue -->
<template>
    <div id="app">
      <h1>{{ concert.name }} Ticket Booking</h1>
      
      <!-- Ticket Status Section -->
      <div class="status-bar">
        <span :class="inventoryStatus.class">
          {{ inventoryStatus.message }}
        </span>
      </div>

      <!-- Tickets Section -->
      <div v-if="tickets.length > 0">
        <h2>Available Tickets: {{ availableTickets }}</h2>
        <TicketItem
          v-for="ticket in tickets"
          :key="ticket.id"
          :ticket="ticket"
          :calculate-final-price="calculateFinalPrice"
          :get-ticket-status="getTicketStatus"
          :get-ticket-status-class="getTicketStatusClass"
          @add-to-cart="addToCart"
        />
      </div>
      <div v-else>
        <p>Sorry, no tickets available.</p>
      </div>

      <!-- Cart Section -->
      <div v-if="cart.length > 0" class="cart-section">
        <h2>Your Cart ({{ cartItemCount }} items)</h2>
        <div v-for="(item, index) in cart" :key="index" class="cart-item">
          <p>{{ item.type }} - ${{ calculateFinalPrice(item) }}</p>
          <button @click="removeFromCart(index)">Remove</button>
        </div>
        <div class="cart-summary">
          <h3>Summary:</h3>
          <p>Subtotal: ${{ cartSubtotal }}</p>
          <p>Savings: ${{ totalSavings }}</p>
          <h3>Total: ${{ cartTotal }}</h3>
        </div>
        <button 
          @click="checkout"
          :disabled="isCheckoutDisabled"
          class="checkout-button"
        >
          Checkout
        </button>
      </div>
    </div>
  </template>
  
  <script>
  import { ref, watch, onMounted } from 'vue'
  import { useTickets } from './composables/useTickets'
  import { useCart } from './composables/useCart'
  import TicketItem from './components/TicketItem.vue'
  
  export default {
    name: 'App',
    components: {
      TicketItem
    },
    setup() {
      const concert = ref({
        id: 1,
        name: 'Rock Festival 2024',
        date: '2024-07-15',
      })
  
      const {
        tickets,
        availableTickets,
        inventoryStatus,
        calculateFinalPrice,
        getTicketStatus,
        getTicketStatusClass
      } = useTickets()
  
      const {
        cart,
        cartItemCount,
        cartSubtotal,
        totalSavings,
        cartTotal,
        isCheckoutDisabled,
        addToCart,
        removeFromCart,
        checkout
      } = useCart(tickets, calculateFinalPrice)
  
      // Watchers
      watch(availableTickets, (newValue, oldValue) => {
        if (newValue === 0) {
          alert('All tickets have been sold out!')
        } else if (newValue <= 10 && oldValue > 10) {
          alert('Only a few tickets left! Hurry up!')
        }
      })
  
      // Lifecycle Hooks
      onMounted(() => {
        // Simulate dynamic discount updates
        setInterval(() => {
          const randomTicket = tickets.value[Math.floor(Math.random() * tickets.value.length)]
          randomTicket.discount = Math.floor(Math.random() * 15)
        }, 30000)
      })
  
      return {
        concert,
        tickets,
        availableTickets,
        inventoryStatus,
        cart,
        cartItemCount,
        cartSubtotal,
        cartTotal,
        totalSavings,
        isCheckoutDisabled,
        calculateFinalPrice,
        getTicketStatus,
        getTicketStatusClass,
        addToCart,
        removeFromCart,
        checkout
      }
    }
  }
  </script>
  
  <style scoped>
  #app {
    font-family: Arial, sans-serif;
    max-width: 800px;
    margin: 0 auto;
    padding: 20px;
  }
  
  .status-bar {
    margin: 20px 0;
    padding: 10px;
    text-align: center;
  }
  
  .status-available { color: green; }
  .status-low { color: orange; }
  .status-sold-out { color: red; }
  
  .cart-section {
    margin-top: 30px;
    border-top: 2px solid #ddd;
    padding-top: 20px;
  }
  
  .cart-item {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 10px;
    border-bottom: 1px solid #eee;
  }
  
  .cart-summary {
    margin: 20px 0;
    padding: 15px;
    background-color: #f9f9f9;
    border-radius: 5px;
  }
  
  .checkout-button {
    width: 100%;
    padding: 15px;
    background-color: #2196F3;
    color: white;
    font-size: 1.1em;
    margin-top: 20px;
  }
  
  button {
    margin: 5px;
    padding: 8px 15px;
    cursor: pointer;
    border-radius: 4px;
    border: 1px solid #ddd;
    background-color: #4CAF50;
    color: white;
  }
  
  button:disabled {
    cursor: not-allowed;
    opacity: 0.5;
    background-color: #cccccc;
  }
  </style>