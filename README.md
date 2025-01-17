# Concert Ticket System - MVVM Implementation

This project demonstrates the MVVM (Model-View-ViewModel) pattern in a Vue.js application.

## Architecture Overview

### Model Layer
- `Concert.js` and `Ticket.js`: Define the data structures
- `api.js`: Handles data operations (mocked for training)

### ViewModel Layer
- Vuex store (`store/index.js`): Manages application state and business logic
- Components' script sections: Handle component-specific logic

### View Layer
- Vue templates in `ConcertList.vue` and `ConcertDetails.vue`
- Handle user interface and event bindings

## Data Flow

1. **Initial Load**
   - `ConcertList` component is created
   - Calls `fetchConcerts` action
   - Store commits mutation to update state
   - View automatically updates through computed properties

2. **Concert Selection**
   - User clicks "View Details"
   - Triggers `selectConcert` method
   - Fetches concert details through store action
   - Routes to detail view

3. **Ticket Purchase**
   - User fills purchase form
   - Form submission triggers `purchaseTicket` action
   - Store updates concert availability
   - UI reflects the changes

## Testing the Flow

1. Start the application
2. Observe the concert list loading
3. Click on a concert to view details
4. Try purchasing a ticket
5. Notice the available tickets count decrease

Note: This implementation uses mock data for training purposes. In a production environment, replace the mock service with real API calls.
