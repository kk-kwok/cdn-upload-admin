import {BaseURL, axios} from '../config'
import qs from 'qs'

// login auth get token
export const doLogin = params => { return axios.post(`${BaseURL}/api/v1/auth`, qs.stringify(params)) }
