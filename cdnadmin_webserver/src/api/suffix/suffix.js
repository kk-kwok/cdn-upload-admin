import {BaseURL, axios} from '../config'
import qs from 'qs'

// get Suffix list
export const getSuffixList = params => { return axios.get(`${BaseURL}/api/v1/suffix/list`, { params: params }) }

// delete Suffix
export const deleteSuffix = params => { return axios.get(`${BaseURL}/api/v1/suffix/delete`, { params: params }) }

// add Suffix
export const addSuffix = params => { return axios.post(`${BaseURL}/api/v1/suffix/add`, qs.stringify(params)) }

// update Suffix
export const updateSuffix = params => { return axios.post(`${BaseURL}/api/v1/suffix/update`, qs.stringify(params)) }
