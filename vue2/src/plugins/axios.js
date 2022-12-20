"use strict";

import axios from "axios";

// Full config:  https://github.com/axios/axios#request-config
// axios.defaults.baseURL = process.env.baseURL || process.env.apiUrl || '';
// axios.defaults.headers.common['Authorization'] = AUTH_TOKEN;
// axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded';


const axiosInstance = axios.create({
    //baseURL: "http://localhost:7077"
    baseURL: process.env.VUE_APP_BASE_URL,
    timeout: 60 * 1000,
});

axiosInstance.interceptors.request.use(config => {
        config.headers = {
            'Content-Type': 'application/json'
        };
        return config;
    },
    error => {
        return Promise.reject(error);
    }
);


axiosInstance.interceptors.response.use(
    response => {
        return response;
    },
    error => {
        return Promise.reject(error);
    }
);

// Plugin.install = function (Vue, options) {
//     console.log(options)
//     Vue.axios = _axios;
//     window.axios = _axios;
//     Object.defineProperties(Vue.prototype, {
//         axios: {
//             get() {
//                 return _axios;
//             }
//         },
//         $axios: {
//             get() {
//                 return _axios;
//             }
//         },
//     });
// };
// Vue.use(Plugin)
// export default Plugin;

export default {
    Get(url, params, headers) {
        return axiosInstance({
            method: "get",
            url: url,
            headers: headers,
            params: params
        });
    },
    Delete(url, params, headers) {
        return axiosInstance({
            method: "delete",
            url: url,
            headers: headers,
            params: params
        });
    },
    Post(url, data, headers) {
        return axiosInstance({
            method: "post",
            url: url,
            headers: headers,
            data: data
        });
    },
    Put(url, data, headers) {
        return axiosInstance({
            method: "put",
            url: url,
            headers: headers,
            data: data
        });
    }
};


