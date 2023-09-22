<template>
    <div class="container">
      <h3>All Users</h3>
      <div v-if="message" class="alert alert-success">
        {{ this.message }}</div>
      <div class="container">
        <table class="table">
          <thead>
            <tr> 
              <th>Kode</th>
              <th>Nama</th>
              <th>Jumlah</th>
              <th>Deskripsi</th>
              <th>Status</th>
              <th>Aksi</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="barang in barangs" v-bind:key="barang.id">
            
              <td>{{ barang.kode }}</td>
              <td>{{ barang.nama }}</td>
              <td>{{ barang.jumlah }}</td>
              <td>{{ barang.deskripsi }}</td>
              <td>{{ barang.status_aktif = true? "Aktive":"Inactive" }}</td>
              <td>
                <button class="btn btn-warning" 
                v-on:click="updateUser(barang.id)">
                  Update
                </button>
              </td>
              <td>
                <button class="btn btn-danger" 
                v-on:click="deleteUser(barang.id)">
                  Delete
                </button>
              </td>
            </tr>
          </tbody>
        </table>
        <div class="row">
          <button class="btn btn-success" 
          v-on:click="addUser()">Add</button>
        </div>
      </div>
    </div>
  </template>

  <script>
  import BarangService from "../service/BarangService";

  
  export default {
    name: "BarangList",
    data() {
      return {
        barangs: [],
        message: "",
      };
    },
    methods: {
      dataBarang() {
        // databarangs = BarangService.listDataBarang();
        BarangService.listDataBarang().then((res) => {
          this.barangs = res.data.data;
        });

        // console.log(databarangs)
      },
    },
    created() {
      this.dataBarang();
    },
  };
  </script>