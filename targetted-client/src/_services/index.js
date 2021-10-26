import axios from "axios";
//import { throttleAdapterEnhancer } from "axios-extensions";

let headers = {
    "Content-Type": "application/json"
  };

const instance = axios.create({
    //adapter: throttleAdapterEnhancer(axios.defaults.adapter, CACHE_TIME),
    headers
})
  
axios.interceptors.response.use((response) => {
    if(response.status === 401) {
         console.log("You are not authorized");
         window.location.href = "/login";
    }
    return response;
}, (error) => {
    if (error.response && error.response.data) {
        return Promise.reject(error.response.data);
    }
    return Promise.reject(error.message);
});

export {instance};
