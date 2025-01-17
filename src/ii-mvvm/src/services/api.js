import { Concert } from '../models/Concert';

// Mock data for development
const mockConcerts = [
  new Concert(1, 'Spring Musical', '2024-04-15', 'School Auditorium', 100, 15),
  new Concert(2, 'Jazz Band Night', '2024-05-01', 'Music Hall', 75, 10),
  new Concert(3, 'Choir Performance', '2024-05-20', 'School Theater', 150, 12)
];

export default {
  async getConcerts() {
    return Promise.resolve(mockConcerts);
  },

  async getConcert(id) {
    const concert = mockConcerts.find(c => c.id === parseInt(id));
    return Promise.resolve(concert);
  },

  async purchaseTicket(ticket) {
    console.log('Purchased ticket:', ticket);
    return Promise.resolve({ success: true });
  }
};