import axios from "axios";

// 1.创建axios的实例
const request = axios.create({
  baseURL: process.env.VUE_APP_URL,
  timeout: 10000,
});

// 2.请求拦截的作用
request.interceptors.request.use(
  (config) => {
    return config;
  },
  (err) => {
    return Promise.reject(err);
  }
);

// 3.响应拦截
request.interceptors.response.use(
  (res) => {
    return res.data;
  },
  (err) => {
    return Promise.reject(err);
  }
);

export default request;
