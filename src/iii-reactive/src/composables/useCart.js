import { ref, computed } from 'vue'

export function useCart(tickets, calculateFinalPrice) {
  const cart = ref([])

  const cartItemCount = computed(() => cart.value.length)

  const cartSubtotal = computed(() => {
    return cart.value.reduce((total, item) => total + item.price, 0)
  })

  const totalSavings = computed(() => {
    return cart.value.reduce((total, item) => {
      const discount = (item.price * (item.discount || 0)) / 100
      return total + discount
    }, 0)
  })

  const cartTotal = computed(() => {
    return cart.value.reduce((total, item) => {
      return total + calculateFinalPrice(item)
    }, 0)
  })

  const isCheckoutDisabled = computed(() => cart.value.length === 0)

  function addToCart(ticket) {
    if (ticket.quantity > 0) {
      cart.value.push({ ...ticket })
      ticket.quantity--
    }
  }

  function removeFromCart(index) {
    const removedTicket = cart.value.splice(index, 1)[0]
    const originalTicket = tickets.value.find(t => t.id === removedTicket.id)
    originalTicket.quantity++
  }

  function checkout() {
    alert(`Thank you for your purchase!\nTotal Savings: $${totalSavings.value}\nFinal Total: $${cartTotal.value}`)
    cart.value = []
  }

  return {
    cart,
    cartItemCount,
    cartSubtotal,
    totalSavings,
    cartTotal,
    isCheckoutDisabled,
    addToCart,
    removeFromCart,
    checkout
  }
} 