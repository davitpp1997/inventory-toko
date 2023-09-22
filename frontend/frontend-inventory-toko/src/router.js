import { createWebHistory, createRouter } from "vue-router";
import Home from './components/HelloWorld.vue'
import Barang from './components/Barang.vue'

const routes = [
    {
      path: "/",
      name: "Home",
      component: Home,
    }
    ,{
      path: "/barang",
      name: "List Barang",
      component: Barang,
    },
  ];
  
  const router = createRouter({
    history: createWebHistory(),
    routes,
  });
  
  export default router;