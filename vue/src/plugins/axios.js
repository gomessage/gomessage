"use strict";

import axios from "axios";
import store from '@/store';
import router from "@/router";

// Full config:  https://github.com/axios/axios#request-config
// axios.defaults.baseURL = process.env.baseURL || process.env.apiUrl || '';
// axios.defaults.headers.common['Authorization'] = AUTH_TOKEN;
// axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded';


const axiosInstance = axios.create({
    //baseURL: "http://localhost:7077"
    baseURL: process.env.VUE_APP_BASE_URL,
    timeout: 60 * 1000,
});

axiosInstance.interceptors.request.use(request => {
        let token = store.getters.getToken

        //如果token为空字符串，则跳转到登录页面
        if (token === "") {
            router.push("/login")
        } else {
            request.headers = {
                'Content-Type': 'application/json',
                'Authorization': token
            };
        }
        return request;
    },
    error => {
        return Promise.reject(error);
    }
);


axiosInstance.interceptors.response.use(response => {
        return response;
    },
    error => {
        switch (error.response.status) {
            case 403:
                console.log("没有权限")
                router.push("/login")
                break;
            case 401:
                console.log("没有权限")
                router.push("/login")
                break;
            default:
                break;
        }
        return Promise.reject(error);
    }
);


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


