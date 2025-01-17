import { createStore } from 'vuex'
import { Concert } from '../models/Concert'

// Mock data
const mockConcerts = [
  new Concert(1, 'Spring Musical', '2024-04-15', 'School Auditorium', 100, 15),
  new Concert(2, 'Jazz Band Night', '2024-05-01', 'Music Hall', 75, 10),
  new Concert(3, 'Choir Performance', '2024-05-20', 'School Theater', 150, 12)
]

export default createStore({
  state: {
    concerts: mockConcerts,
    selectedConcert: null
  },
  mutations: {
    SELECT_CONCERT(state, concertId) {
      state.selectedConcert = state.concerts.find(c => c.id === concertId)
    },
    PURCHASE_TICKET(state, ticket) {
      const concert = state.concerts.find(c => c.id === ticket.concertId)
      if (concert && concert.availableTickets > 0) {
        concert.availableTickets--
      }
    }
  },
  actions: {
    selectConcert({ commit }, concertId) {
      commit('SELECT_CONCERT', concertId)
    },
    purchaseTicket({ commit }, ticket) {
      commit('PURCHASE_TICKET', ticket)
      return Promise.resolve({ success: true })
    }
  },
  getters: {
    availableConcerts: (state) => state.concerts.filter(c => c.availableTickets > 0)
  }
})