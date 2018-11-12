import {BaseURL, axios} from '../config'
import qs from 'qs'

// get platform list
export const getPlatformList = params => { return axios.get(`${BaseURL}/api/v1/platform/list`, { params: params }) }

// delete platform
export const deletePlatform = params => { return axios.get(`${BaseURL}/api/v1/platform/delete`, { params: params }) }

// add platform
export const addPlatform = params => { return axios.post(`${BaseURL}/api/v1/platform/add`, qs.stringify(params)) }

// update platform
export const updatePlatform = params => { return axios.post(`${BaseURL}/api/v1/platform/update`, qs.stringify(params)) }
