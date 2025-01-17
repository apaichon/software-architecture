import { createRouter, createWebHistory } from 'vue-router';
import ConcertList from '../views/ConcertList.vue';
import ConcertDetails from '../views/ConcertDetails.vue';

const routes = [
  {
    path: '/',
    name: 'ConcertList',
    component: ConcertList
  },
  {
    path: '/concert/:id',
    name: 'ConcertDetails',
    component: ConcertDetails,
    props: true
  }
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
});

export default router;