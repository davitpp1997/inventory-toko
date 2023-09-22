import axios from 'axios'

class BarangService {

    listDataBarang() {
      // let barangs
      axios.defaults.baseURL = 'http://localhost:8001/v1';

      return axios.get("/list_barang");
      // // axios.defaults.headers.post['Content-Type'] ='application/json;charset=utf-8';
      // // axios.defaults.headers.post['Access-Control-Allow-Origin'] = '*';

      // axios.get("/list_barang")
      //   .then((res)=>{
      //     barangs = res.data;
      //     // console.log(barangs)
          
      //   }).catch(err=> console.log(err))

      // return barangs;
    }
}

export default new BarangService()