import { ref, computed } from 'vue'

export function useTickets() {
  const tickets = ref([
    { 
      id: 1, 
      type: 'General Admission', 
      price: 50, 
      quantity: 100,
      discount: 0 
    },
    { 
      id: 2, 
      type: 'VIP', 
      price: 150, 
      quantity: 20,
      discount: 10 
    },
    { 
      id: 3, 
      type: 'Backstage Pass', 
      price: 250, 
      quantity: 5,
      discount: 0 
    },
  ])

  const availableTickets = computed(() => {
    return tickets.value.reduce((total, ticket) => total + ticket.quantity, 0)
  })

  const inventoryStatus = computed(() => {
    if (availableTickets.value === 0) {
      return { message: 'SOLD OUT', class: 'status-sold-out' }
    }
    if (availableTickets.value <= 10) {
      return { message: 'SELLING FAST', class: 'status-low' }
    }
    return { message: 'TICKETS AVAILABLE', class: 'status-available' }
  })

  function calculateFinalPrice(ticket) {
    const discount = (ticket.price * (ticket.discount || 0)) / 100
    return ticket.price - discount
  }

  function getTicketStatus(ticket) {
    if (ticket.quantity === 0) return 'Sold Out'
    if (ticket.quantity <= 5) return 'Almost Gone'
    return 'Available'
  }

  function getTicketStatusClass(ticket) {
    if (ticket.quantity === 0) return 'sold-out'
    if (ticket.quantity <= 5) return 'low-stock'
    return 'in-stock'
  }

  return {
    tickets,
    availableTickets,
    inventoryStatus,
    calculateFinalPrice,
    getTicketStatus,
    getTicketStatusClass
  }
} 