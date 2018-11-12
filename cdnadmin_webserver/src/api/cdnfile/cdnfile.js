import {BaseURL, axios} from '../config'
import qs from 'qs'

// get cdnfile list
export const getCDNFileList = params => { return axios.get(`${BaseURL}/api/v1/cdnfile/list`, { params: params }) }

// delete cdnfile
export const deleteCDNFile = params => { return axios.get(`${BaseURL}/api/v1/cdnfile/delete`, { params: params }) }

// add cdnfile
let ax = axios.create()
export const addCDNFile = (params, config) => { return ax.post(`${BaseURL}/api/v1/cdnfile/add`, params, config) }

// update cdnfile
export const updateCDNFile = params => { return axios.post(`${BaseURL}/api/v1/cdnfile/update`, qs.stringify(params)) }

// refresh cdnfile
export const pushCDNFile = params => { return axios.post(`${BaseURL}/api/v1/cdnfile/push`, qs.stringify(params)) }
