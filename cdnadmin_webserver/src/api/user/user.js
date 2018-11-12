import {BaseURL, axios} from '../config'
import qs from 'qs'

// get User list
export const getUserList = params => { return axios.get(`${BaseURL}/api/v1/user/list`, { params: params }) }

// delete User
export const deleteUser = params => { return axios.get(`${BaseURL}/api/v1/user/delete`, { params: params }) }

// add User
export const addUser = params => { return axios.post(`${BaseURL}/api/v1/user/add`, qs.stringify(params)) }

// update User
export const updateUser = params => { return axios.post(`${BaseURL}/api/v1/user/update`, qs.stringify(params)) }

// update User Pwd
export const updateUserPwd = params => { return axios.post(`${BaseURL}/api/v1/user/updatePwd`, qs.stringify(params)) }

// reset User Pwd
export const resetUserPwd = params => { return axios.post(`${BaseURL}/api/v1/user/reset`, qs.stringify(params)) }

// set user admin
export const setUserAdmin = params => { return axios.post(`${BaseURL}/api/v1/user/setAdmin`, qs.stringify(params)) }
